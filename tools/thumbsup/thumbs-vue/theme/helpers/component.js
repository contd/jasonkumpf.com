module.exports = component => {
  return component.replace(/[-_]/g, "")
}