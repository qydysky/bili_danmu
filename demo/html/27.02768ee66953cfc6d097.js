(window.webpackJsonp = window.webpackJsonp || []).push([
  [27],
  {
    1439: function (e, r, t) {
      t(431)("WeakMap");
    },
    1440: function (e, r, t) {
      t(432)("WeakMap");
    },
    1441: function (e, r, t) {
      "use strict";
      var n = t(361),
        _ = t(228).getWeak,
        a = t(94),
        o = t(73),
        c = t(362),
        E = t(197),
        i = t(545),
        u = t(134),
        f = t(356),
        T = i(5),
        s = i(6),
        A = 0,
        P = function (e) {
          return e._l || (e._l = new l());
        },
        l = function () {
          this.a = [];
        },
        d = function (e, r) {
          return T(e.a, function (e) {
            return e[0] === r;
          });
        };
      (l.prototype = {
        get: function (e) {
          var r = d(this, e);
          if (r) return r[1];
        },
        has: function (e) {
          return !!d(this, e);
        },
        set: function (e, r) {
          var t = d(this, e);
          t ? (t[1] = r) : this.a.push([e, r]);
        },
        delete: function (e) {
          var r = s(this.a, function (r) {
            return r[0] === e;
          });
          return ~r && this.a.splice(r, 1), !!~r;
        },
      }),
        (e.exports = {
          getConstructor: function (e, r, t, a) {
            var i = e(function (e, n) {
              c(e, i, r, "_i"),
                (e._t = r),
                (e._i = A++),
                (e._l = void 0),
                void 0 != n && E(n, t, e[a], e);
            });
            return (
              n(i.prototype, {
                delete: function (e) {
                  if (!o(e)) return !1;
                  var t = _(e);
                  return !0 === t
                    ? P(f(this, r)).delete(e)
                    : t && u(t, this._i) && delete t[this._i];
                },
                has: function (e) {
                  if (!o(e)) return !1;
                  var t = _(e);
                  return !0 === t ? P(f(this, r)).has(e) : t && u(t, this._i);
                },
              }),
              i
            );
          },
          def: function (e, r, t) {
            var n = _(a(r), !0);
            return !0 === n ? P(e).set(r, t) : (n[e._i] = t), e;
          },
          ufstore: P,
        });
    },
    1442: function (e, r, t) {
      "use strict";
      var n,
        _ = t(59),
        a = t(545)(0),
        o = t(440),
        c = t(228),
        E = t(547),
        i = t(1441),
        u = t(73),
        f = t(356),
        T = t(356),
        s = !_.ActiveXObject && "ActiveXObject" in _,
        A = c.getWeak,
        P = Object.isExtensible,
        l = i.ufstore,
        d = function (e) {
          return function () {
            return e(this, arguments.length > 0 ? arguments[0] : void 0);
          };
        },
        p = {
          get: function (e) {
            if (u(e)) {
              var r = A(e);
              return !0 === r
                ? l(f(this, "WeakMap")).get(e)
                : r
                ? r[this._i]
                : void 0;
            }
          },
          set: function (e, r) {
            return i.def(f(this, "WeakMap"), e, r);
          },
        },
        W = (e.exports = t(433)("WeakMap", d, p, i, !0, !0));
      T &&
        s &&
        (E((n = i.getConstructor(d, "WeakMap")).prototype, p),
        (c.NEED = !0),
        a(["delete", "has", "get", "set"], function (e) {
          var r = W.prototype,
            t = r[e];
          o(r, e, function (r, _) {
            if (u(r) && !P(r)) {
              this._f || (this._f = new n());
              var a = this._f[e](r, _);
              return "set" == e ? this : a;
            }
            return t.call(this, r, _);
          });
        }));
    },
    1443: function (e, r, t) {
      t(233), t(198), t(1442), t(1440), t(1439), (e.exports = t(43).WeakMap);
    },
    1444: function (e, r, t) {
      e.exports = { default: t(1443), __esModule: !0 };
    },
    2222: function (e, r, t) {
      "use strict";
      t.r(r);
      var n = t(200),
        _ = t.n(n), //new _.a() = new Map()
        a = t(1444),
        o = t.n(a), //new o.a() = new WeakMap()
        c = t(12),
        E = t.n(c), //E()(util) = typeof util
        i = t(15),
        u = t.n(i); //u()(t) = Object.keys(t)
      r.default = function () {
        var e = { STDWEB_PRIVATE: {} };
        (e.STDWEB_PRIVATE.to_utf8 = function (r, t) {
          for (var n = e.HEAPU8, _ = 0; _ < r.length; ++_) {
            var a = r.charCodeAt(_);
            a >= 55296 &&
              a <= 57343 &&
              (a = (65536 + ((1023 & a) << 10)) | (1023 & r.charCodeAt(++_))),
              a <= 127
                ? (n[t++] = a)
                : a <= 2047
                ? ((n[t++] = 192 | (a >> 6)), (n[t++] = 128 | (63 & a)))
                : a <= 65535
                ? ((n[t++] = 224 | (a >> 12)),
                  (n[t++] = 128 | ((a >> 6) & 63)),
                  (n[t++] = 128 | (63 & a)))
                : a <= 2097151
                ? ((n[t++] = 240 | (a >> 18)),
                  (n[t++] = 128 | ((a >> 12) & 63)),
                  (n[t++] = 128 | ((a >> 6) & 63)),
                  (n[t++] = 128 | (63 & a)))
                : a <= 67108863
                ? ((n[t++] = 248 | (a >> 24)),
                  (n[t++] = 128 | ((a >> 18) & 63)),
                  (n[t++] = 128 | ((a >> 12) & 63)),
                  (n[t++] = 128 | ((a >> 6) & 63)),
                  (n[t++] = 128 | (63 & a)))
                : ((n[t++] = 252 | (a >> 30)),
                  (n[t++] = 128 | ((a >> 24) & 63)),
                  (n[t++] = 128 | ((a >> 18) & 63)),
                  (n[t++] = 128 | ((a >> 12) & 63)),
                  (n[t++] = 128 | ((a >> 6) & 63)),
                  (n[t++] = 128 | (63 & a)));
          }
        }),
          (e.STDWEB_PRIVATE.noop = function () {}),
          (e.STDWEB_PRIVATE.to_js = function (r) {
            var t = e.HEAPU8[r + 12];
            if (0 !== t) {
              if (1 === t) return null;
              if (2 === t) return e.HEAP32[r / 4];
              if (3 === t) return e.HEAPF64[r / 8];
              if (4 === t) {
                var n = e.HEAPU32[r / 4],
                  _ = e.HEAPU32[(r + 4) / 4];
                return e.STDWEB_PRIVATE.to_js_string(n, _);
              }
              if (5 === t) return !1;
              if (6 === t) return !0;
              if (7 === t) {
                (n = e.STDWEB_PRIVATE.arena + e.HEAPU32[r / 4]),
                  (_ = e.HEAPU32[(r + 4) / 4]);
                for (var a = [], o = 0; o < _; ++o)
                  a.push(e.STDWEB_PRIVATE.to_js(n + 16 * o));
                return a;
              }
              if (8 === t) {
                var c = e.STDWEB_PRIVATE.arena,
                  E = c + e.HEAPU32[r / 4],
                  i = ((_ = e.HEAPU32[(r + 4) / 4]), c + e.HEAPU32[(r + 8) / 4]);
                for (a = {}, o = 0; o < _; ++o) {
                  var u = e.HEAPU32[(i + 8 * o) / 4],
                    f = e.HEAPU32[(i + 4 + 8 * o) / 4],
                    T = e.STDWEB_PRIVATE.to_js_string(u, f),
                    s = e.STDWEB_PRIVATE.to_js(E + 16 * o);
                  a[T] = s;
                }
                return a;
              }
              if (9 === t)
                return e.STDWEB_PRIVATE.acquire_js_reference(e.HEAP32[r / 4]);
              if (10 === t || 12 === t || 13 === t) {
                var A = e.HEAPU32[r / 4],
                  P = ((n = e.HEAPU32[(r + 4) / 4]), e.HEAPU32[(r + 8) / 4]),
                  l = 0,
                  d = !1;
                return (
                  ((a = function r() {
                    if (0 === n || !0 === d)
                      throw 10 === t
                        ? new ReferenceError("Already dropped Rust function called!")
                        : 12 === t
                        ? new ReferenceError("Already dropped FnMut function called!")
                        : new ReferenceError(
                            "Already called or dropped FnOnce function called!"
                          );
                    var _ = n;
                    if (
                      (13 === t && ((r.drop = e.STDWEB_PRIVATE.noop), (n = 0)),
                      0 !== l && (12 === t || 13 === t))
                    )
                      throw new ReferenceError(
                        "FnMut function called multiple times concurrently!"
                      );
                    var a = e.STDWEB_PRIVATE.alloc(16);
                    e.STDWEB_PRIVATE.serialize_array(a, arguments);
                    try {
                      (l += 1), e.STDWEB_PRIVATE.dyncall("vii", A, [_, a]);
                      var o = e.STDWEB_PRIVATE.tmp;
                      e.STDWEB_PRIVATE.tmp = null;
                    } finally {
                      l -= 1;
                    }
                    return !0 === d && 0 === l && r.drop(), o;
                  }).drop = function () {
                    if (0 === l) {
                      a.drop = e.STDWEB_PRIVATE.noop;
                      var r = n;
                      (n = 0), 0 != r && e.STDWEB_PRIVATE.dyncall("vi", P, [r]);
                    } else d = !0;
                  }),
                  a
                );
              }
              if (14 === t) {
                (n = e.HEAPU32[r / 4]), (_ = e.HEAPU32[(r + 4) / 4]);
                var p = e.HEAPU32[(r + 8) / 4],
                  W = n + _;
                switch (p) {
                  case 0:
                    return e.HEAPU8.subarray(n, W);
                  case 1:
                    return e.HEAP8.subarray(n, W);
                  case 2:
                    return e.HEAPU16.subarray(n, W);
                  case 3:
                    return e.HEAP16.subarray(n, W);
                  case 4:
                    return e.HEAPU32.subarray(n, W);
                  case 5:
                    return e.HEAP32.subarray(n, W);
                  case 6:
                    return e.HEAPF32.subarray(n, W);
                  case 7:
                    return e.HEAPF64.subarray(n, W);
                }
              } else if (15 === t)
                return e.STDWEB_PRIVATE.get_raw_value(e.HEAPU32[r / 4]);
            }
          }),
          (e.STDWEB_PRIVATE.serialize_object = function (r, t) {
            var n = u()(t),
              _ = n.length,
              a = e.STDWEB_PRIVATE.alloc(8 * _),
              o = e.STDWEB_PRIVATE.alloc(16 * _);
            (e.HEAPU8[r + 12] = 8),
              (e.HEAPU32[r / 4] = o),
              (e.HEAPU32[(r + 4) / 4] = _),
              (e.HEAPU32[(r + 8) / 4] = a);
            for (var c = 0; c < _; ++c) {
              var E = n[c],
                i = a + 8 * c;
              e.STDWEB_PRIVATE.to_utf8_string(i, E),
                e.STDWEB_PRIVATE.from_js(o + 16 * c, t[E]);
            }
          }),
          (e.STDWEB_PRIVATE.serialize_array = function (r, t) {
            var n = t.length,
              _ = e.STDWEB_PRIVATE.alloc(16 * n);
            (e.HEAPU8[r + 12] = 7),
              (e.HEAPU32[r / 4] = _),
              (e.HEAPU32[(r + 4) / 4] = n);
            for (var a = 0; a < n; ++a) e.STDWEB_PRIVATE.from_js(_ + 16 * a, t[a]);
          });
        var r =
          "function" == typeof TextEncoder
            ? new TextEncoder("utf-8")
            : "object" === ("undefined" == typeof util ? "undefined" : E()(util)) &&
              util &&
              "function" == typeof util.TextEncoder
            ? new util.TextEncoder("utf-8")
            : null;
        (e.STDWEB_PRIVATE.to_utf8_string =
          null != r
            ? function (t, n) {
                var _ = r.encode(n),
                  a = _.length,
                  o = 0;
                a > 0 && ((o = e.STDWEB_PRIVATE.alloc(a)), e.HEAPU8.set(_, o)),
                  (e.HEAPU32[t / 4] = o),
                  (e.HEAPU32[(t + 4) / 4] = a);
              }
            : function (r, t) {
                var n = e.STDWEB_PRIVATE.utf8_len(t),
                  _ = 0;
                n > 0 &&
                  ((_ = e.STDWEB_PRIVATE.alloc(n)), e.STDWEB_PRIVATE.to_utf8(t, _)),
                  (e.HEAPU32[r / 4] = _),
                  (e.HEAPU32[(r + 4) / 4] = n);
              }),
          (e.STDWEB_PRIVATE.from_js = function (r, t) {
            var n = Object.prototype.toString.call(t);
            if ("[object String]" === n)
              (e.HEAPU8[r + 12] = 4), e.STDWEB_PRIVATE.to_utf8_string(r, t);
            else if ("[object Number]" === n)
              t === (0 | t)
                ? ((e.HEAPU8[r + 12] = 2), (e.HEAP32[r / 4] = t))
                : ((e.HEAPU8[r + 12] = 3), (e.HEAPF64[r / 8] = t));
            else if (null === t) e.HEAPU8[r + 12] = 1;
            else if (void 0 === t) e.HEAPU8[r + 12] = 0;
            else if (!1 === t) e.HEAPU8[r + 12] = 5;
            else if (!0 === t) e.HEAPU8[r + 12] = 6;
            else if ("[object Symbol]" === n) {
              var _ = e.STDWEB_PRIVATE.register_raw_value(t);
              (e.HEAPU8[r + 12] = 15), (e.HEAP32[r / 4] = _);
            } else {
              var a = e.STDWEB_PRIVATE.acquire_rust_reference(t);
              (e.HEAPU8[r + 12] = 9), (e.HEAP32[r / 4] = a);
            }
          });
        var t =
          "function" == typeof TextDecoder
            ? new TextDecoder("utf-8")
            : "object" === ("undefined" == typeof util ? "undefined" : E()(util)) &&
              util &&
              "function" == typeof util.TextDecoder
            ? new util.TextDecoder("utf-8")
            : null;
        (e.STDWEB_PRIVATE.to_js_string =
          null != t
            ? function (r, n) {
                return t.decode(e.HEAPU8.subarray(r, r + n));
              }
            : function (r, t) {
                for (
                  var n = e.HEAPU8, _ = (0 | (r |= 0)) + (0 | (t |= 0)), a = "";
                  r < _;
        
                ) {
                  var o = n[r++];
                  if (o < 128) a += String.fromCharCode(o);
                  else {
                    var c = 31 & o,
                      E = 0;
                    r < _ && (E = n[r++]);
                    var i = (c << 6) | (63 & E);
                    if (o >= 224) {
                      var u = 0;
                      r < _ && (u = n[r++]);
                      var f = ((63 & E) << 6) | (63 & u);
                      if (((i = (c << 12) | f), o >= 240)) {
                        var T = 0;
                        r < _ && (T = n[r++]),
                          (i = ((7 & c) << 18) | (f << 6) | (63 & T)),
                          (a += String.fromCharCode(55232 + (i >> 10))),
                          (i = 56320 + (1023 & i));
                      }
                    }
                    a += String.fromCharCode(i);
                  }
                }
                return a;
              }),
          (e.STDWEB_PRIVATE.id_to_ref_map = {}),
          (e.STDWEB_PRIVATE.id_to_refcount_map = {}),
          (e.STDWEB_PRIVATE.ref_to_id_map = new o.a()),
          (e.STDWEB_PRIVATE.ref_to_id_map_fallback = new _.a()),
          (e.STDWEB_PRIVATE.last_refid = 1),
          (e.STDWEB_PRIVATE.id_to_raw_value_map = {}),
          (e.STDWEB_PRIVATE.last_raw_value_id = 1),
          (e.STDWEB_PRIVATE.acquire_rust_reference = function (r) {
            if (void 0 === r || null === r) return 0;
            var t = e.STDWEB_PRIVATE.id_to_refcount_map,
              n = e.STDWEB_PRIVATE.id_to_ref_map,
              _ = e.STDWEB_PRIVATE.ref_to_id_map,
              a = e.STDWEB_PRIVATE.ref_to_id_map_fallback,
              o = _.get(r);
            if ((void 0 === o && (o = a.get(r)), void 0 === o)) {
              o = e.STDWEB_PRIVATE.last_refid++;
              try {
                _.set(r, o);
              } catch (e) {
                a.set(r, o);
              }
            }
            return o in n ? t[o]++ : ((n[o] = r), (t[o] = 1)), o;
          }),
          (e.STDWEB_PRIVATE.acquire_js_reference = function (r) {
            return e.STDWEB_PRIVATE.id_to_ref_map[r];
          }),
          (e.STDWEB_PRIVATE.increment_refcount = function (r) {
            e.STDWEB_PRIVATE.id_to_refcount_map[r]++;
          }),
          (e.STDWEB_PRIVATE.decrement_refcount = function (r) {
            var t = e.STDWEB_PRIVATE.id_to_refcount_map;
            if (0 == --t[r]) {
              var n = e.STDWEB_PRIVATE.id_to_ref_map,
                _ = e.STDWEB_PRIVATE.ref_to_id_map_fallback,
                a = n[r];
              delete n[r], delete t[r], _.delete(a);
            }
          }),
          (e.STDWEB_PRIVATE.register_raw_value = function (r) {
            var t = e.STDWEB_PRIVATE.last_raw_value_id++;
            return (e.STDWEB_PRIVATE.id_to_raw_value_map[t] = r), t;
          }),
          (e.STDWEB_PRIVATE.unregister_raw_value = function (r) {
            delete e.STDWEB_PRIVATE.id_to_raw_value_map[r];
          }),
          (e.STDWEB_PRIVATE.get_raw_value = function (r) {
            return e.STDWEB_PRIVATE.id_to_raw_value_map[r];
          }),
          (e.STDWEB_PRIVATE.alloc = function (r) {
            return e.web_malloc(r);
          }),
          (e.STDWEB_PRIVATE.dyncall = function (r, t, n) {
            return e.web_table.get(t).apply(null, n);
          }),
          (e.STDWEB_PRIVATE.utf8_len = function (e) {
            for (var r = 0, t = 0; t < e.length; ++t) {
              var n = e.charCodeAt(t);
              n >= 55296 &&
                n <= 57343 &&
                (n = (65536 + ((1023 & n) << 10)) | (1023 & e.charCodeAt(++t))),
                n <= 127
                  ? ++r
                  : (r +=
                      n <= 2047
                        ? 2
                        : n <= 65535
                        ? 3
                        : n <= 2097151
                        ? 4
                        : n <= 67108863
                        ? 5
                        : 6);
            }
            return r;
          }),
          (e.STDWEB_PRIVATE.prepare_any_arg = function (r) {
            var t = e.STDWEB_PRIVATE.alloc(16);
            return e.STDWEB_PRIVATE.from_js(t, r), t;
          }),
          (e.STDWEB_PRIVATE.acquire_tmp = function (r) {
            var t = e.STDWEB_PRIVATE.tmp;
            return (e.STDWEB_PRIVATE.tmp = null), t;
          });
        var n = null,
          a = null,
          c = null,
          i = null,
          f = null,
          T = null,
          s = null,
          A = null;
        function P() {
          var r = e.instance.exports.memory.buffer;
          (n = new Int8Array(r)),
            (a = new Int16Array(r)),
            (c = new Int32Array(r)),
            (i = new Uint8Array(r)),
            (f = new Uint16Array(r)),
            (T = new Uint32Array(r)),
            (s = new Float32Array(r)),
            (A = new Float64Array(r)),
            (e.HEAP8 = n),
            (e.HEAP16 = a),
            (e.HEAP32 = c),
            (e.HEAPU8 = i),
            (e.HEAPU16 = f),
            (e.HEAPU32 = T),
            (e.HEAPF32 = s),
            (e.HEAPF64 = A);
        }
        return (
          Object.defineProperty(e, "exports", { value: {} }),
          {
        imports: {
            env: {
            __cargo_web_snippet_0d39c013e2144171d64e2fac849140a7e54c939a: function (
                r,
                t
            ) {
                (t = e.STDWEB_PRIVATE.to_js(t)),
                e.STDWEB_PRIVATE.from_js(r, t.location);
            },
                __cargo_web_snippet_0f503de1d61309643e0e13a7871406891e3691c9: function (
                  r
                ) {
                  e.STDWEB_PRIVATE.from_js(r, window);
                },
                __cargo_web_snippet_10f5aa3985855124ab83b21d4e9f7297eb496508: function (
                  r
                ) {
                  return (
                    (e.STDWEB_PRIVATE.acquire_js_reference(r) instanceof Array) | 0
                  );
                },
                __cargo_web_snippet_2b0b92aee0d0de6a955f8e5540d7923636d951ae: function (
                  r,
                  t
                ) {
                  (t = e.STDWEB_PRIVATE.to_js(t)),
                    e.STDWEB_PRIVATE.from_js(
                      r,
                      (function () {
                        try {
                          return { value: t.origin, success: !0 };
                        } catch (e) {
                          return { error: e, success: !1 };
                        }
                      })()
                    );
                },
                __cargo_web_snippet_461d4581925d5b0bf583a3b445ed676af8701ca6: function (
                  r,
                  t
                ) {
                  (t = e.STDWEB_PRIVATE.to_js(t)),
                    e.STDWEB_PRIVATE.from_js(
                      r,
                      (function () {
                        try {
                          return { value: t.host, success: !0 };
                        } catch (e) {
                          return { error: e, success: !1 };
                        }
                      })()
                    );
                },
                __cargo_web_snippet_4c895ac2b754e5559c1415b6546d672c58e29da6: function (
                  r,
                  t
                ) {
                  (t = e.STDWEB_PRIVATE.to_js(t)),
                    e.STDWEB_PRIVATE.from_js(
                      r,
                      (function () {
                        try {
                          return { value: t.protocol, success: !0 };
                        } catch (e) {
                          return { error: e, success: !1 };
                        }
                      })()
                    );
                },
                __cargo_web_snippet_614a3dd2adb7e9eac4a0ec6e59d37f87e0521c3b: function (
                  r,
                  t
                ) {
                  (t = e.STDWEB_PRIVATE.to_js(t)),
                    e.STDWEB_PRIVATE.from_js(r, t.error);
                },
                __cargo_web_snippet_62ef43cf95b12a9b5cdec1639439c972d6373280: function (
                  r,
                  t
                ) {
                  (t = e.STDWEB_PRIVATE.to_js(t)),
                    e.STDWEB_PRIVATE.from_js(r, t.childNodes);
                },
                __cargo_web_snippet_6fcce0aae651e2d748e085ff1f800f87625ff8c8: function (
                  r
                ) {
                  e.STDWEB_PRIVATE.from_js(r, document);
                },
                __cargo_web_snippet_7ba9f102925446c90affc984f921f414615e07dd: function (
                  r,
                  t
                ) {
                  (t = e.STDWEB_PRIVATE.to_js(t)),
                    e.STDWEB_PRIVATE.from_js(r, t.body);
                },
                __cargo_web_snippet_80d6d56760c65e49b7be8b6b01c1ea861b046bf0: function (
                  r
                ) {
                  e.STDWEB_PRIVATE.decrement_refcount(r);
                },
                __cargo_web_snippet_897ff2d0160606ea98961935acb125d1ddbf4688: function (
                  r
                ) {
                  var t = e.STDWEB_PRIVATE.acquire_js_reference(r);
                  return t instanceof DOMException && "SecurityError" === t.name;
                },
                __cargo_web_snippet_8c32019649bb581b1b742eeedfc410e2bedd56a6: function (
                  r,
                  t
                ) {
                  var n = e.STDWEB_PRIVATE.acquire_js_reference(r);
                  e.STDWEB_PRIVATE.serialize_array(t, n);
                },
                __cargo_web_snippet_a466a2ab96cd77e1a77dcdb39f4f031701c195fc: function (
                  r,
                  t
                ) {
                  (t = e.STDWEB_PRIVATE.to_js(t)),
                    e.STDWEB_PRIVATE.from_js(
                      r,
                      (function () {
                        try {
                          return { value: t.pathname, success: !0 };
                        } catch (e) {
                          return { error: e, success: !1 };
                        }
                      })()
                    );
                },
                __cargo_web_snippet_ab05f53189dacccf2d365ad26daa407d4f7abea9: function (
                  r,
                  t
                ) {
                  (t = e.STDWEB_PRIVATE.to_js(t)),
                    e.STDWEB_PRIVATE.from_js(r, t.value);
                },
                __cargo_web_snippet_b06dde4acf09433b5190a4b001259fe5d4abcbc2: function (
                  r,
                  t
                ) {
                  (t = e.STDWEB_PRIVATE.to_js(t)),
                    e.STDWEB_PRIVATE.from_js(r, t.success);
                },
                __cargo_web_snippet_b33a39de4ca954888e26fe9caa277138e808eeba: function (
                  r,
                  t
                ) {
                  (t = e.STDWEB_PRIVATE.to_js(t)),
                    e.STDWEB_PRIVATE.from_js(r, t.length);
                },
                __cargo_web_snippet_cdf2859151791ce4cad80688b200564fb08a8613: function (
                  r,
                  t
                ) {
                  (t = e.STDWEB_PRIVATE.to_js(t)),
                    e.STDWEB_PRIVATE.from_js(
                      r,
                      (function () {
                        try {
                          return { value: t.href, success: !0 };
                        } catch (e) {
                          return { error: e, success: !1 };
                        }
                      })()
                    );
                },
                __cargo_web_snippet_e8ef87c41ded1c10f8de3c70dea31a053e19747c: function (
                  r,
                  t
                ) {
                  (t = e.STDWEB_PRIVATE.to_js(t)),
                    e.STDWEB_PRIVATE.from_js(
                      r,
                      (function () {
                        try {
                          return { value: t.hostname, success: !0 };
                        } catch (e) {
                          return { error: e, success: !1 };
                        }
                      })()
                    );
                },
                __cargo_web_snippet_e9638d6405ab65f78daf4a5af9c9de14ecf1e2ec: function (
                  r
                ) {
                  (r = e.STDWEB_PRIVATE.to_js(r)),
                    console.log(t),
                    e.STDWEB_PRIVATE.unregister_raw_value(r);
                },
                __cargo_web_snippet_ff5103e6cc179d13b4c7a785bdce2708fd559fc0: function (
                  r
                ) {
                  e.STDWEB_PRIVATE.tmp = e.STDWEB_PRIVATE.to_js(r);
                },
                __web_on_grow: P,
              },
            },
            initialize: function (r) {
              return (
                Object.defineProperty(e, "instance", { value: r }),
                Object.defineProperty(e, "web_malloc", {
                  value: e.instance.exports.__web_malloc,
                }),
                Object.defineProperty(e, "web_free", {
                  value: e.instance.exports.__web_free,
                }),
                Object.defineProperty(e, "web_table", {
                  value: e.instance.exports.__indirect_function_table,
                }),
                (e.exports.spyder = function (r, t) {
                  return e.STDWEB_PRIVATE.acquire_tmp(
                    e.instance.exports.spyder(
                      e.STDWEB_PRIVATE.prepare_any_arg(r),
                      e.STDWEB_PRIVATE.prepare_any_arg(t)
                    )
                  );
                }),
                (e.exports.from_js = (r, t) => {//export fromjs
                    return e.STDWEB_PRIVATE.from_js(r, t)
                }),
                P(),
                e.exports
              );
            },
          }
        );
      };
    },
  },
]);
