"use strict";
exports.__esModule = true;
var tablet_1 = require("./device/tablet");
var Tablet1 = new tablet_1["default"]('class1-tablet2');
Tablet1.getData()
    .then(function (res) { return res.json(); })
    .then(console.log)["catch"](console.error);
