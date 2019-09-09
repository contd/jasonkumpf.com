module.exports = crumbs => {
  return crumbs.map(item => {
    return JSON.stringify({
      text: item.title,
      disabled: false,
      href: item.url
    }, null, 2)
  })
}
