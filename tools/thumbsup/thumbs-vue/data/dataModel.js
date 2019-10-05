{
  album: Album,
  gallery: Gallery,
  breadcrumbs: [String],

  Gallery: {
    home: Album,
    title: String,
    footer: String,
    thumbSize: Number,
    largeSize: Number,
    googleAnalytics: String
  },
  Album:
  {
    title: String,
    home: Boolean,
    depth: Number,
    zip: String,
    albums: [Album],
    files: [File],
    stats: {
      albums: Number,
      photos: Number,
      videos: Number,
      fromDate: Date,
      toDate: Date
    }
  },
  File:
  {
    id: Number,
    path: String,
    filename: String,
    date: Date,
    type: 'Video|Image',
    isVideo: Boolean,
    meta: Metadata,
    urls: {
      thumbnail: String,
      small: String,
      large: String,
      video: String,
      download: String
    }
  },
  Metadata:
  {
    date: Date,
    caption: String,
    keywords: [String],
    video: Boolean,
    animated: Boolean,
    rating: Number,
    favourite: Boolean,
    width: Number,
    height: Number,
    exif: Object
  }
}
