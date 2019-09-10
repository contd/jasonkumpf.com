module.exports = fixlink => {
  return fixlink.replace(/^media\//, "/travels/").replace("../../../../../../..", "/travels")
}
