/**
 * 获取wasm并返回instance
 * @param {string} url wasm文件位置
 * @param {{}} importObject new t_func()得到的对象
 */
function fetchAndInstantiate(url, importObject) {
  return fetch(url)
    .then((response) => response.arrayBuffer())
    .then((bytes) => WebAssembly.instantiate(bytes, importObject))
    .then((results) => results.instance);
}

/**
 * 根据
 * https://github.com/lkeme/bilibili-pcheartbeat 及
 * 27.02768ee66953cfc6d097.js(改自bilibili)
 * 编写
 */
function t_func() {
  var tmp = {};
  tmp = (arg) => {
    if (arg == 200) {
      return { a: Map };
    }
    if (arg == 1444) {
      return { a: WeakMap };
    }
    if (arg == 12) {
      return () => {
        return (argc) => {
          return typeof argc;
        };
      };
    }
    if (arg == 15) {
      return () => {
        return (argc) => {
          return Object.keys(argc);
        };
      };
    }
  };
  tmp.n = (arg) => {
    return arg;
  };
  tmp.r = (arg) => {
    arg.default = {};
  };
  return tmp;
}

/**
 * 全局操作对象
 */
var wasm =
  /**
   * 初始化
   */
  (() => {
    let r = {};
    webpackJsonp[0][1][2222](null, r, new t_func());
    return r.default();
  })();

/**
 * 替换27.02768ee66953cfc6d097.js(改自bilibili)中的导入函数
 */
(() => {
  /**
   * 返回host
   */
  wasm.imports.env.__cargo_web_snippet_461d4581925d5b0bf583a3b445ed676af8701ca6 = (
    r,
    t
  ) => {
    wasm.all.from_js(r, { value: "live.bilibili.com", success: !0 }); //hs
  };
  /**
   * 返回protocol
   */
  wasm.imports.env.__cargo_web_snippet_4c895ac2b754e5559c1415b6546d672c58e29da6 = (
    r,
    t
  ) => {
    wasm.all.from_js(r, { value: "https:", success: !0 }); //pe
  };
  /**
   * 返回hostname
   */
  wasm.imports.env.__cargo_web_snippet_e8ef87c41ded1c10f8de3c70dea31a053e19747c = (
    r,
    t
  ) => {
    wasm.all.from_js(r, { value: "live.bilibili.com", success: !0 }); //hn
  };
  /**
   * 返回href
   */
  wasm.imports.env.__cargo_web_snippet_cdf2859151791ce4cad80688b200564fb08a8613 = (
    r,
    t
  ) => {
    wasm.all.from_js(r, {
      value: "https://live.bilibili.com/0",
      success: !0,
    }); //hq
  };
})();

/**
 * 加密
 */
wasm.spyder = (r, t) => {
  try {
    return wasm.all.spyder(JSON.stringify(r), t);
  } catch (err) {
    return err;
  }
};

/**
 * 测试加密是否正常
 */
wasm.test = () => {
  let test = {
    r: {
      id: "[9,371,1,22613059]",
      device: '["AUTO8216117272375373","77bee604-b591-4664-845b-b69603f8c71c"]',
      ets: 1611836581,
      benchmark: "seacasdgyijfhofiuxoannn",
      time: 60,
      ts: 1611836642190,
    },
    t: [2, 5, 1, 4],
  };
  `e4249b7657c2d4a44955548eb814797d41ddd99bfdfa5974462b8c387d701b8c83898f6d7dde1772c67fad6a113d20c20e454be1d1627e7ea99617a8a1f99bd0` ==
  wasm.spyder(test.r, test.t)
    ? (() => {
        console.log(`Test Pass`);
        console.log(`现在可以使用 wasm.spyder(r,t)进行加密`);
      })()
    : console.error(`Test No Pass`);
};

/**
 * 获取wasm并测试加密
 */
(() => {
  fetchAndInstantiate("e791556706f88d88b4846a61a583b31db007f83d.wasm", {
    env: wasm.imports.env,
  })
    .then((instance) => {
      wasm.all = wasm.initialize(instance);
      wasm.test();
    })
    .catch((reason) => console.error(reason));
})();

/**
 * ws 收发
 */
(() => {
  if (window["WebSocket"]) {
    conn = new WebSocket("ws://" + document.location.host + "/ws");
    conn.onclose = function () {
      location.reload();
    };
    conn.onmessage = function (evt) {
      deal(evt.data)
    };

    function deal(data) {
      //或许还没初始化
      if (wasm.all == null) {
        setTimeout(()=>{
          deal(data)
        },100)
        return
      }

      let rt = JSON.parse(data),
          s = wasm.spyder(rt.r, rt.t);

      conn.send(JSON.stringify({
        id:rt.r.id,
        s:s
      }));
    }
  }
})();
