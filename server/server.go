package main

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	s "strings"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
)

// Directory struct is for directory objects
type Directory struct {
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	ModTime time.Time `json:"modified"`
	Path    string    `json:"path"`
	Thumb   string    `json:"thumb"`
	Items   int       `json:"items"`
}

// Exif picture data struct
type Exif struct {
	Lat   float64   `json:"lat"`
	Long  float64   `json:"long"`
	Taken time.Time `json:"taken"`
}

// Picture struct is for picture objects
type Picture struct {
	Name    string    `json:"name"`
	Size    int64     `json:"size"`
	Type    string    `json:"type"`
	ModTime time.Time `json:"modified"`
	Path    string    `json:"path"`
	Thumb   string    `json:"thumb"`
	Width   int       `json:"width"`
	Height  int       `json:"height"`
	Exif    Exif      `json:"exif"`
}

// Files struct stores the current directory's contents in separate arrays of directores or pictures
type Files struct {
	Directories []Directory
	Pictures    []Picture
}

// RootPath is the starting path of directories or pictures
var RootPath string

// Directories is an array of Directory objects in current path
var Directories []Directory

// Pictures is an array of Picture objects in current path
var Pictures []Picture

// OriginsAllowed is for CORS and should leave the localhost there
var OriginsAllowed = []string{"https://jasonkumpf.com", "http://localhost:8088", "http://localhost:3000"}

// ServerPort is the port Echo is running
var ServerPort = os.Getenv("SERVER_PORT")

// ServerHost is the hostname Echo is running
var ServerHost = os.Getenv("SERVER_HOST")

func main() {
	if os.Getenv("PHOTOS_PATH") == "" {
		RootPath = "photos"
	} else {
		RootPath = os.Getenv("PHOTOS_PATH")
	}
	Directories, Pictures = readPath(RootPath, "")

	router := echo.New()
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())
	// CORS restricted
	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: OriginsAllowed,
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	router.Static("/photos", "./photos")
	router.File("/favicon.ico", "photos/favicon.ico")
	router.GET("/", listRoot)
	router.GET("/:path", listRoot)
	router.GET("/:path/*", listRoot)
	log.Printf("SERVER_PORT: %s", ServerPort)
	log.Printf("SERVER_HOST: %s", ServerHost)
	log.Printf("PHOTOS_PATH: %s", RootPath)
	router.Logger.Fatal(router.Start(ServerPort))
}

func listRoot(c echo.Context) error {
	ctxPath := s.Join(c.ParamValues(), "/")
	files := getFiles(ctxPath)
	return c.JSON(http.StatusOK, files)
}

func getFiles(ctxPath string) Files {
	Directories, Pictures = readPath(RootPath, ctxPath)

	sort.SliceStable(Directories, func(i, j int) bool {
		return Directories[i].Name < Directories[j].Name
	})
	sort.SliceStable(Pictures, func(i, j int) bool {
		return Pictures[i].Name < Pictures[j].Name
	})

	files := Files{
		Directories: Directories,
		Pictures:    Pictures,
	}
	return files
}

func readPath(fullPath string, ctxPath string) ([]Directory, []Picture) {
	dir, err := os.Open(path.Join(fullPath, ctxPath))
	if err != nil {
		log.Printf("readDir os.Open FAILED [%s] %s : %v", fullPath, ctxPath, err)
		return nil, nil
	}
	defer dir.Close()

	items, err := dir.Readdir(0)
	if err != nil {
		log.Printf("readDir dir.Readdir FAILED [%s] %s : %v", fullPath, ctxPath, err)
		return nil, nil
	}

	var dirs []Directory
	var pics []Picture

	for _, item := range items {
		itemPath := path.Join(fullPath, ctxPath, item.Name())

		if item.IsDir() {
			if item.Name() != "." && item.Name() != ".." {
				firstThumb, itemCnt := getFirstThumb(path.Join(fullPath, ctxPath, item.Name()))

				directory := Directory{
					Name:    item.Name(),
					Size:    item.Size(),
					ModTime: item.ModTime(),
					Path:    fmt.Sprintf("%s", path.Join(ctxPath, item.Name())),
					Thumb:   fmt.Sprintf("%s", path.Join(ctxPath, item.Name(), "thumbs", firstThumb)),
					Items:   itemCnt,
				}
				dirs = append(dirs, directory)
			}
		} else {
			if checkExtension(item.Name()) {
				pic, err := os.Open(itemPath)
				if err != nil {
					log.Printf("failed to open pic: %v", item.Name())
					return dirs, pics
				}

				mimeType := getMimeType(item.Name())
				img, _, err := image.DecodeConfig(pic)
				if err != nil {
					log.Printf("%s: %v", item.Name(), err)
					return dirs, pics
				}

				exifInfo, exifErr := getExif(itemPath)
				if exifErr != nil {
					//t, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
					t := time.Now()
					exifInfo = Exif{Lat: 0, Long: 0, Taken: t}
				}
				picture := Picture{
					Name:    item.Name(),
					Size:    item.Size(),
					Type:    mimeType,
					ModTime: item.ModTime(),
					Path:    fmt.Sprintf("/%s", path.Join("photos", ctxPath, item.Name())),
					Thumb:   fmt.Sprintf("/%s", path.Join("photos", ctxPath, "thumbs", item.Name())),
					Width:   img.Width,
					Height:  img.Height,
					Exif:    exifInfo,
				}
				pics = append(pics, picture)
			}
		}
	}

	return dirs, pics
}

func getFirstThumb(albumPath string) (string, int) {
	dir, err := os.Open(albumPath)
	if err != nil {
		log.Printf("getFirstThumb: failed to open dir: %s", err)
	}
	defer dir.Close()

	items, err := dir.Readdir(0)
	if err != nil {
		log.Printf("failed to read dir: %s", err)
	}
	item := items[0].Name()
	if item == "thumbs" {
		item = items[1].Name()
	}

	return item, len(items)
}

func checkExtension(fname string) bool {
	re := regexp.MustCompile(`.*\.(?:jpg|gif|png|tif)$`)
	if re.MatchString(fname) {
		return true
	}
	return false
}

func getMimeType(fname string) string {
	extn := filepath.Ext(fname)
	mtype := mime.TypeByExtension(extn)
	return mtype
}

func getExif(fname string) (Exif, error) {
	exif.RegisterParsers(mknote.All...)

	f, err := os.Open(fname)
	if err != nil {
		//log.Printf("Failed to open for exif [%s]: %v", fname, err)
		return Exif{}, err
	}

	x, err := exif.Decode(f)
	if err != nil {
		//log.Printf("Failed to decode exif [%s]: %v", fname, err)
		return Exif{}, err
	}

	lat, long, _ := x.LatLong()
	tm, _ := x.DateTime()

	return Exif{
		Lat:   lat,
		Long:  long,
		Taken: tm,
	}, nil
}
