import $ from 'jquery'
window.$ = $
window.jQuery = $

function operate(url, type, dataType, data, callback) {
    let config = {
        url: url,
        type: type,
        dataType: dataType,
        data: JSON.stringify(data),
        beforeSend: function (XMLHttpRequest) {
            XMLHttpRequest.setRequestHeader("Access-Token", sessionStorage.getItem("myAccessToken"));
            XMLHttpRequest.setRequestHeader("Access-Control-Allow-Origin", "*");
            XMLHttpRequest.setRequestHeader("Content-Type", "application/x-www-form-urlencoded; charset=utf-8");
        },
        success: function(result) {
            callback(result);
        }
    };
    $.ajax(config)
}

export function AjaxPost(url, data, callback) {
    operate(serverUrl+url, "post", "json", data, callback);
}

export function AjaxGet(url, callback) {
    operate(serverUrl+url, "get", "json", "", callback);
}

const serverUrl='http://127.0.0.1:8888/'
