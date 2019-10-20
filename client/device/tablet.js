"use strict";
var __extends = (this && this.__extends) || (function () {
    var extendStatics = function (d, b) {
        extendStatics = Object.setPrototypeOf ||
            ({ __proto__: [] } instanceof Array && function (d, b) { d.__proto__ = b; }) ||
            function (d, b) { for (var p in b) if (b.hasOwnProperty(p)) d[p] = b[p]; };
        return extendStatics(d, b);
    };
    return function (d, b) {
        extendStatics(d, b);
        function __() { this.constructor = d; }
        d.prototype = b === null ? Object.create(b) : (__.prototype = b.prototype, new __());
    };
})();
exports.__esModule = true;
var device_1 = require("./device");
var http_1 = require("../common/http");
var Tablet = /** @class */ (function (_super) {
    __extends(Tablet, _super);
    function Tablet(name) {
        var _this = _super.call(this, name) || this;
        _this.sendData = function (battery, currentVideo, date) {
            return http_1["default"].sendData(_this.name, battery, currentVideo, date);
        };
        return _this;
    }
    ;
    Tablet.prototype.getData = function () {
        return http_1["default"].getData(this.name);
    };
    return Tablet;
}(device_1["default"]));
exports["default"] = Tablet;
