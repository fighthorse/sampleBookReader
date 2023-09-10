
// 获取当前页面的URL
var currentUrl = window.location.href;
// 解析URL参数
function parseUrlParam(url, paramName) {
    paramName = paramName.replace(/[\[\]]/g, "\\$&");
    var regex = new RegExp("[?&]" + paramName + "(=([^&#]*)|&|#|$)"),
        results = regex.exec(url);

    if (!results || !results[2]) return '';
    return decodeURIComponent(results[2].replace(/\+/g, " "));
}