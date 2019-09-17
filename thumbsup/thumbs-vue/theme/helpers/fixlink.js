module.exports = fixlink => {
  //return fixlink.replace(/^media\//, "/travels/").replace("../../../../../../..", "/travels")
  return fixlink.replace(/^media\//, "/media/")
}
