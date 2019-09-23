package main

import (
	"flag"
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"mime"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"sort"
	"time"

	"github.com/rwcarlsen/goexif/exif"
	"github.com/rwcarlsen/goexif/mknote"
	"github.com/ttacon/chalk"
	"gopkg.in/h2non/bimg.v1"
)

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

// Directory struct is for directory objects
type Directory struct {
	Name     string    `json:"name"`
	Size     int64     `json:"size"`
	ModTime  time.Time `json:"modified"`
	Path     string    `json:"path"`
	Thumb    string    `json:"thumb"`
	Items    int       `json:"items"`
	Pictures []Picture `json:"pictures"`
}

// Files struct stores the current directory's contents in separate arrays of directores or pictures
type Files struct {
	Directories []Directory
	Pictures    []Picture
}

var lime = chalk.Green.NewStyle().
	WithBackground(chalk.Black).
	WithTextStyle(chalk.Bold).
	Style

func printMsg(prefix string, msg string) {
	fmt.Println(chalk.Red, prefix, ": ", chalk.Cyan, msg, chalk.Reset)
}

func printErr(prefix string, err error) {
	fmt.Println(chalk.Red, prefix, ": ", chalk.Cyan, err, chalk.Reset)
}

func printInfo(msg string) {
	fmt.Println(chalk.Cyan, msg, chalk.Reset)
}

func getNames(pictures []Picture) []string {
	var names []string

	for _, picture := range pictures {
		names = append(names, picture.Name)
	}
	return names
}

func contains(thumbs []string, search string) bool {
	for _, item := range thumbs {
		if item == search {
			return true
		}
	}
	return false
}

func main() {
	// Pass in the base dir to start with (i.e. ./photos/travels)
	// Then get the list of albums (Directories) then for each album get Pictures to process
	fullPath := flag.String("path", "./photos/travels", "Path to the albums of pictures to process")
	flag.Parse()
	RootPath := *fullPath
	items := getFiles(RootPath)

	for _, album := range items.Directories {
		// See if dir original exists and if not create it
		printMsg("ALBUM", album.Path)

		err := album.originalCheck()
		if err != nil {
			printErr("ERROR originalsCheck", err)
			os.Exit(1)
		}
		// For each picture:
		// Move to originals dir
		err = album.moveToOriginal()
		if err != nil {
			printErr("ERROR moveToOriginal", err)
			os.Exit(1)
		}
		// Check that thumbs dir exists
		thumbs, err := album.thumbDirCheck()
		if err != nil {
			printErr("ERROR thumbDirCheck", err)
		}
		// Check that a thumbnail exists in thumbs dir with same name
		// Create thumbnail if one doesn't exist
		for _, picture := range album.getPhotos() {
			err := picture.thumbCheck(thumbs, album.Path)
			if err != nil {
				printErr("ERROR thumbCheck", err)
			}
			// Resize
			err = picture.resizeFile(album.Path)
			if err != nil {
				printErr("ERROR resizeFile", err)
			}
			// // Check if EXIF data still exists and if not then add from original
			// exifInfo, err := getExif(picture.Path)
			// if err != nil {
			// 	printErr("ERROR", err)
			// }
			// fmt.Println(chalk.Blue, "EXIF: ", exifInfo.Lat, " x ", exifInfo.Long, chalk.Reset)
			// if exifInfo.Lat == 0 && exifInfo.Long == 0 {
			// 	//
			// }
		}
	}
}

func (d *Directory) originalCheck() error {
	originals := path.Join(d.Path, "originals")

	_, err := os.Open(originals)
	if err != nil {
		printInfo("No originals dir exists, going to create it")

		e := os.Mkdir(originals, 0777)
		if e != nil {
			printErr("ERROR os.Mkdir", e)
			return e
		}
		printInfo("originals directory created")
	}
	return nil
}

func (d *Directory) moveToOriginal() error {
	files := getFiles(d.Path)
	photos := files.Pictures
	originals := path.Join(d.Path, "originals")

	if len(photos) <= 0 {
		printMsg("ERROR", fmt.Sprintf("no photos exist in album path: %s", d.Path))
		os.Exit(1)
	}
	printInfo("Moving photos to originals dir..")

	for _, photo := range photos {
		oldPath := path.Join(d.Path, photo.Name)
		newPath := path.Join(originals, photo.Name)
		//printInfo(fmt.Sprintf("Moving %s to %s dir.", oldPath, newPath))

		err := os.Rename(oldPath, newPath)
		if err != nil {
			printInfo(fmt.Sprintf("ERROR moving %s to %s", oldPath, newPath))
			return err
		}
	}
	return nil
}

func (d *Directory) thumbDirCheck() ([]string, error) {
	thumbDir := path.Join(d.Path, "thumbs")

	_, err := os.Open(thumbDir)
	if err != nil {
		printInfo("No thumbs dir exists, goint to create it")

		e := os.Mkdir(thumbDir, 0777)
		if e != nil {
			fmt.Fprintln(os.Stderr, e)
			return nil, e
		}
		printInfo("thumbs directory created")
	} else {
		printInfo("thumbs directory exists")
	}
	files := getFiles(thumbDir)
	thumbs := getNames(files.Pictures)

	return thumbs, nil
}

func (p *Picture) thumbCheck(thumbs []string, albumPath string) error {
	if !contains(thumbs, p.Name) {
		// At this point no thumbnail exists, so going to create one
		err := p.thumbCreate(albumPath)
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *Picture) thumbCreate(albumPath string) error {
	original := path.Join(albumPath, "originals", p.Name)
	buffer, err := bimg.Read(original)
	if err != nil {
		printErr("ERROR bimg.Read", err)
		return err
	}
	thumbnail, err := bimg.NewImage(buffer).Thumbnail(200)
	if err != nil {
		printErr("ERROR bimg.Thumbnail", err)
		return err
	}
	bimg.Write(path.Join(albumPath, "thumbs", p.Name), thumbnail)
	printInfo(fmt.Sprintf("CREATED thumbnail: %s", p.Name))
	return nil
}

func (p *Picture) resizeFile(albumPath string) error {
	original := path.Join(albumPath, "originals", p.Name)
	aspect := float64(p.Width) / float64(p.Height)
	newWidth := 1280.0
	newHeight := newWidth / aspect
	// Use the width and height to determine aspect ratio and portait or landscape
	if aspect < 1.0 {
		newWidth = 1280.0
		newHeight = newWidth / aspect
	}
	buffer, err := bimg.Read(original)
	if err != nil {
		printErr("ERROR bimg.Read", err)
		return err
	}
	// Resize the original to either 1024x768 or 768x1024 and save
	newImage, err := bimg.NewImage(buffer).Resize(int(newWidth), int(newHeight))
	if err != nil {
		printErr("ERROR bimg.Resize", err)
		return err
	}
	// size, err := bimg.NewImage(newImage).Size()
	// if float64(size.Width) == newWidth && float64(size.Height) == newHeight {
	// 	printInfo(fmt.Sprintf("The image size is valid: %s", p.Name))
	// 	return err
	// }
	bimg.Write(path.Join(albumPath, p.Name), newImage)
	return nil
}

func (d *Directory) getPhotos() []Picture {
	files := getFiles(d.Path)
	if len(files.Pictures) <= 0 {
		files = getFiles(path.Join(d.Path, "originals"))
	}

	return files.Pictures
}

func getFiles(fullPath string) Files {
	Directories, Pictures := readPath(fullPath)

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

func readPath(fullPath string) ([]Directory, []Picture) {
	dir, err := os.Open(path.Join(fullPath))
	if err != nil {
		log.Printf("readDir os.Open FAILED [%s] %v", fullPath, err)
		return nil, nil
	}
	defer dir.Close()

	items, err := dir.Readdir(0)
	if err != nil {
		log.Printf("readDir dir.Readdir FAILED [%s] %v", fullPath, err)
		return nil, nil
	}

	var dirs []Directory
	var pics []Picture

	for _, item := range items {
		itemPath := path.Join(fullPath, item.Name())

		if item.IsDir() {
			if item.Name() != "." && item.Name() != ".." {
				itemCnt, pictures := getAlbumItems(path.Join(fullPath, item.Name()))

				directory := Directory{
					Name:     item.Name(),
					Size:     item.Size(),
					ModTime:  item.ModTime(),
					Path:     fmt.Sprintf("%s", path.Join(fullPath, item.Name())),
					Items:    itemCnt,
					Pictures: pictures,
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
					t := time.Now()
					exifInfo = Exif{Lat: 0, Long: 0, Taken: t}
				}
				picture := Picture{
					Name:    item.Name(),
					Size:    item.Size(),
					Type:    mimeType,
					ModTime: item.ModTime(),
					Path:    fmt.Sprintf("%s", path.Join(fullPath, item.Name())),
					Thumb:   fmt.Sprintf("%s", path.Join(fullPath, "thumbs", item.Name())),
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

func getAlbumItems(albumPath string) (int, []Picture) {
	files := getFiles(albumPath)
	pictures := files.Pictures
	itemcnt := len(pictures)
	return itemcnt, pictures
}

func checkExtension(fname string) bool {
	re := regexp.MustCompile(`.*\.(?:jpg|jpeg|gif|png|tif|tiff|webp)$`)
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
		return Exif{}, err
	}

	x, err := exif.Decode(f)
	if err != nil {
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
