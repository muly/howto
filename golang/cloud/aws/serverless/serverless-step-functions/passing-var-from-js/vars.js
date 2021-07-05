const today = new Date()
const yesterday = new Date(today)

yesterday.setDate(yesterday.getDate() - 1)


module.exports.year = yesterday.getFullYear()
module.exports.month = ('0' + (yesterday.getMonth()+1)).slice(-2)
module.exports.date = ('0' + yesterday.getDate()).slice(-2)

// console.log(yesterday.getDate())


// function getYear() {
//    return yesterday.getFullYear()
// }