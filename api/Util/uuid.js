const { v4: uuidv4 } = require('uuid');
function uuid() {
    let myuuid = uuidv4();
    return myuuid
}

module.exports = {uuid:uuid}
