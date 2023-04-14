var md5 = require("md5");
var n = {"payload": {"sort": 1, "start": 180, "limit": 20}};
var _p = "W5D80NFZHAYB8EUI2T649RT2MNRMVE2O";
var _keyStr = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=";

var s = JSON.stringify(n), l = JSON.parse(s)

function _u_e(e) {
    if (null == e)
        return null;
    e = e.replace(/\r\n/g, "\n");
    for (var t = "", n = 0; n < e.length; n++) {
        var r = e.charCodeAt(n);
        r < 128 ? t += String.fromCharCode(r) : r > 127 && r < 2048 ? (t += String.fromCharCode(r >> 6 | 192),
            t += String.fromCharCode(63 & r | 128)) : (t += String.fromCharCode(r >> 12 | 224),
            t += String.fromCharCode(r >> 6 & 63 | 128),
            t += String.fromCharCode(63 & r | 128))
    }
    return t
}

function e1(e) {
    if (null == e)
        return null;
    for (var t, n, r, o, i, a, c, u = "", s = 0; s < e.length;)
        o = (t = e.charCodeAt(s++)) >> 2,
            i = (3 & t) << 4 | (n = e.charCodeAt(s++)) >> 4,
            a = (15 & n) << 2 | (r = e.charCodeAt(s++)) >> 6,
            c = 63 & r,
            isNaN(n) ? a = c = 64 : isNaN(r) && (c = 64),
            u = u + _keyStr.charAt(o) + _keyStr.charAt(i) + _keyStr.charAt(a) + _keyStr.charAt(c);
    return u
}

function e2(e) {
    if (null == (e = _u_e(e)))
        return null;
    for (var t = "", n = 0; n < e.length; n++) {
        var r = _p.charCodeAt(n % _p.length);
        t += String.fromCharCode(e.charCodeAt(n) ^ r)
    }
    return t
}

function sig(e) {
    //  标准的md5 加密，可以直接导包使用
    return md5(e + _p).toUpperCase()
}

function result() {
    var h = {};
    h['payload'] = e1(e2(JSON.stringify(l.payload)));
    h['sig'] = sig(h['payload']);
    return h
}


// 返回结果解密

function _u_d(e) {
    for (var t = "", n = 0, r = 0, o = 0, i = 0; n < e.length;)
        (r = e.charCodeAt(n)) < 128 ? (t += String.fromCharCode(r),
            n++) : r > 191 && r < 224 ? (o = e.charCodeAt(n + 1),
            t += String.fromCharCode((31 & r) << 6 | 63 & o),
            n += 2) : (o = e.charCodeAt(n + 1),
            i = e.charCodeAt(n + 2),
            t += String.fromCharCode((15 & r) << 12 | (63 & o) << 6 | 63 & i),
            n += 3);
    return t
}

function d1(e) {
    var t, n, r, o, i, a, c = "", u = 0;
    for (e = e.replace(/[^A-Za-z0-9\+\/\=]/g, ""); u < e.length;)
        t = _keyStr.indexOf(e.charAt(u++)) << 2 | (o = _keyStr.indexOf(e.charAt(u++))) >> 4,
            n = (15 & o) << 4 | (i = _keyStr.indexOf(e.charAt(u++))) >> 2,
            r = (3 & i) << 6 | (a = _keyStr.indexOf(e.charAt(u++))),
            c += String.fromCharCode(t),
        64 != i && (c += String.fromCharCode(n)),
        64 != a && (c += String.fromCharCode(r));
    return c
}

function d2(e) {
    for (var t = "", n = 0; n < e.length; n++) {
        var r = _p.charCodeAt(n % _p.length);
        t += String.fromCharCode(e.charCodeAt(n) ^ r)
    }
    return t = _u_d(t)
}

function result_decode(l) {
    return d2(d1(l));
}

// console.log(result())
// result_decode("LBcnV1QrZGB4bXsuUTYhawgPTRZaPTBXb3RwLjN0AC5jUXZcAC0gbn11PXoBJ2dxC2UGUAtrNVF1K2BvemdcLjpQZgISqcjhrfHmq739d2UQPUV6XCV2CDk8Jyh6Z1s8H1owGgooJzY7JHVgWzcwKEYxYl1UN3YIfHhgfGJ1CnduDXQIAGJkLzglODZdETwkV3YMBQ9lbQF8eWZ7Y3UCf3sXIU5VIDJ4cmO8zoOi68fU8rTRhuezs")