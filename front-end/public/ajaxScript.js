function operate(url, type, dataType, data, callback) {
    let config = {
        url: url,
        type: type,
        dataType: dataType,
        data: data,
        beforeSend: function () {
        },
        success: function(result) {
            callback(result);
        }
    };
    $.ajax(config)
}

function jPost(url, data, callback) {
    operate(serverUrl+url, "post", "json", data, callback);
}


const serverUrl='http://127.0.0.1:8888/'