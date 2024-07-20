/******/ (() => { // webpackBootstrap
/******/ 	var __webpack_modules__ = ({

/***/ "./node_modules/artplayer-plugin-danmuku/dist/artplayer-plugin-danmuku.js":
/*!********************************************************************************!*\
  !*** ./node_modules/artplayer-plugin-danmuku/dist/artplayer-plugin-danmuku.js ***!
  \********************************************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

/* module decorator */ module = __webpack_require__.nmd(module);
/*!
 * artplayer-plugin-danmuku.js v5.1.0
 * Github: https://github.com/zhw2590582/ArtPlayer
 * (c) 2017-2024 Harvey Zack
 * Released under the MIT License.
 */
!function(e,t,i,a,n){var o="undefined"!=typeof globalThis?globalThis:"undefined"!=typeof self?self:"undefined"!=typeof window?window:"undefined"!=typeof __webpack_require__.g?__webpack_require__.g:{},s="function"==typeof o[a]&&o[a],l=s.cache||{},r= true&&"function"==typeof module.require&&module.require.bind(module);function d(t,i){if(!l[t]){if(!e[t]){var n="function"==typeof o[a]&&o[a];if(!i&&n)return n(t,!0);if(s)return s(t,!0);if(r&&"string"==typeof t)return r(t);var p=Error("Cannot find module '"+t+"'");throw p.code="MODULE_NOT_FOUND",p}h.resolve=function(i){var a=e[t][1][i];return null!=a?a:i},h.cache={};var u=l[t]=new d.Module(t);e[t][0].call(u.exports,h,u,u.exports,this)}return l[t].exports;function h(e){var t=h.resolve(e);return!1===t?{}:d(t)}}d.isParcelRequire=!0,d.Module=function(e){this.id=e,this.bundle=d,this.exports={}},d.modules=e,d.cache=l,d.parent=s,d.register=function(t,i){e[t]=[function(e,t){t.exports=i},{}]},Object.defineProperty(d,"root",{get:function(){return o[a]}}),o[a]=d;for(var p=0;p<t.length;p++)d(t[p]);if(i){var u=d(i); true?module.exports=u:0}}({bgm6t:[function(e,t,i){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(i),a.export(i,"default",()=>p);var n=e("./danmuku"),o=a.interopDefault(n),s=e("./setting"),l=a.interopDefault(s),r=e("./heatmap"),d=a.interopDefault(r);function p(e){return t=>{let i=new o.default(t,e),a=new l.default(t,i);return i.option.heatmap&&(0,d.default)(t,i,i.option.heatmap),{name:"artplayerPluginDanmuku",emit:i.emit.bind(i),load:i.load.bind(i),config:i.config.bind(i),hide:i.hide.bind(i),show:i.show.bind(i),reset:i.reset.bind(i),mount:a.mount.bind(a),get option(){return i.option},get isHide(){return i.isHide},get isStop(){return i.isStop}}}}p.icons=l.default.icons,"undefined"!=typeof window&&(window.artplayerPluginDanmuku=p)},{"./danmuku":"4ns48","./setting":"lO8OT","./heatmap":"8AxLD","@parcel/transformer-js/src/esmodule-helpers.js":"9pCYc"}],"4ns48":[function(e,t,i){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(i);var n=e("./bilibili"),o=e("bundle-text:./worker"),s=a.interopDefault(o);class l{constructor(e,t){let{constructor:i,template:a}=e;this.utils=i.utils,this.validator=i.validator,this.$danmuku=a.$danmuku,this.$player=a.$player,this.art=e,this.danmus=[],this.queue=[],this.option={},this.$refs=[],this.isStop=!1,this.isHide=!1,this.timer=null,this.index=0,this.states={wait:[],ready:[],emit:[],stop:[]},this.config(t),this.worker=new Worker(URL.createObjectURL(new Blob([s.default]))),this.start=this.start.bind(this),this.stop=this.stop.bind(this),this.reset=this.reset.bind(this),this.resize=this.resize.bind(this),this.destroy=this.destroy.bind(this),e.on("video:play",this.start),e.on("video:playing",this.start),e.on("video:pause",this.stop),e.on("video:waiting",this.stop),e.on("destroy",this.destroy),e.on("resize",this.resize),this.load()}static get option(){return{danmuku:[],speed:5,margin:[10,"25%"],opacity:1,color:"#FFFFFF",mode:0,modes:[0,1,2],fontSize:25,antiOverlap:!0,synchronousPlayback:!1,mount:void 0,heatmap:!1,points:[],filter:()=>!0,beforeEmit:()=>!0,beforeVisible:()=>!0,visible:!0,emitter:!0,maxLength:200,lockTime:5,theme:"dark",OPACITY:{},FONT_SIZE:{},MARGIN:{},SPEED:{},COLOR:[]}}static get scheme(){return{danmuku:"array|function|string",speed:"number",margin:"array",opacity:"number",color:"string",mode:"number",modes:"array",fontSize:"number|string",antiOverlap:"boolean",synchronousPlayback:"boolean",mount:"?htmldivelement|string",heatmap:"object|boolean",points:"array",filter:"function",beforeEmit:"function",beforeVisible:"function",visible:"boolean",emitter:"boolean",maxLength:"number",lockTime:"number",theme:"string",OPACITY:"object",FONT_SIZE:"object",MARGIN:"object",SPEED:"object",COLOR:"array"}}static get cssText(){return` user-select: none; position: absolute; white-space: pre; pointer-events: none; perspective: 500px; display: inline-block; will-change: transform; font-weight: normal; line-height: 1.125; visibility: hidden; font-family: SimHei, "Microsoft JhengHei", Arial, Helvetica, sans-serif; text-shadow: rgb(0, 0, 0) 1px 0px 1px, rgb(0, 0, 0) 0px 1px 1px, rgb(0, 0, 0) 0px -1px 1px, rgb(0, 0, 0) -1px 0px 1px; `}get isRotate(){return this.art.plugins?.autoOrientation?.state}get marginTop(){let{clamp:e}=this.utils,t=this.option.margin[0],{clientHeight:i}=this.$player;return"number"==typeof t?e(t,0,i):"string"==typeof t&&t.endsWith("%")?e(i*(parseFloat(t)/100),0,i):l.option.margin[0]}get marginBottom(){let{clamp:e}=this.utils,t=this.option.margin[1],{clientHeight:i}=this.$player;return"number"==typeof t?e(t,0,i):"string"==typeof t&&t.endsWith("%")?e(i*(parseFloat(t)/100),0,i):l.option.margin[1]}get fontSize(){let{clamp:e}=this.utils,{clientHeight:t}=this.$player,i=this.option.fontSize;return"number"==typeof i?e(i,12,t):"string"==typeof i&&i.endsWith("%")?e(t*(parseFloat(i)/100),12,t):l.option.fontSize}get $ref(){let e=this.$refs.pop()||document.createElement("div");return e.style.cssText=l.cssText,e.dataset.mode="",e.className="",e}get readys(){let{currentTime:e}=this.art,t=[];return this.filter("ready",e=>t.push(e)),this.filter("wait",i=>{e+.1>=i.time&&i.time>=e-.1&&t.push(i)}),t}get visibles(){let e=[],{clientWidth:t}=this.$player,i=this.getLeft(this.$player);return this.filter("emit",a=>{let n=a.$ref.offsetTop,o=this.getLeft(a.$ref)-i,s=a.$ref.clientHeight,l=a.$ref.clientWidth,r=o+l,d=r/a.$restTime,p={};p.top=n,p.left=o,p.height=s,p.width=l,p.right=t-r,p.speed=d,p.distance=r,p.time=a.$restTime,p.mode=a.mode,e.push(p)}),e}get speed(){return this.option.synchronousPlayback&&this.art.playbackRate?this.option.speed/Number(this.art.playbackRate):this.option.speed}async load(){let{errorHandle:e}=this.utils;try{"function"==typeof this.option.danmuku?this.danmus=await this.option.danmuku():"function"==typeof this.option.danmuku.then?this.danmus=await this.option.danmuku:"string"==typeof this.option.danmuku?this.danmus=await (0,n.bilibiliDanmuParseFromUrl)(this.option.danmuku):this.danmus=this.option.danmuku,e(Array.isArray(this.danmus),"Danmuku need return an array as result"),this.queue=[],this.$danmuku.innerText="",this.danmus.forEach(e=>this.emit(e)),this.art.emit("artplayerPluginDanmuku:loaded",this.queue)}catch(e){throw this.art.emit("artplayerPluginDanmuku:error",e),e}return this}emit(e){let{clamp:t}=this.utils;if(this.validator(e,{text:"string",mode:"?number",color:"?string",time:"?number",border:"?boolean",style:"?object"}),!e.text.trim()||(e.time?e.time=t(e.time,0,1/0):e.time=this.art.currentTime+.5,void 0===e.mode&&(e.mode=this.option.mode),void 0===e.style&&(e.style={}),void 0===e.color&&(e.color=this.option.color),![0,1,2].includes(e.mode)||!this.option.filter(e)))return this;let i={...e,$state:"wait",$id:this.index++,$ref:null,$restTime:0,$lastStartTime:0};return this.setState(i,"wait"),this.queue.push(i),this}config(e){let{clamp:t}=this.utils,{$controlsCenter:i}=this.art.template;return Object.keys(e).some(t=>JSON.stringify(this.option[t])!==JSON.stringify(e[t]))&&(this.option=Object.assign({},l.option,this.option,e),this.validator(this.option,l.scheme),this.option.mode=t(this.option.mode,0,2),this.option.speed=t(this.option.speed,1,10),this.option.opacity=t(this.option.opacity,0,1),this.option.lockTime=t(this.option.lockTime,1,60),this.option.maxLength=t(this.option.maxLength,1,1e3),this.option.mount=this.option.mount||i,e.fontSize&&this.reset(),this.option.visible?this.show():this.hide(),this.art.emit("artplayerPluginDanmuku:config",this.option)),this}getLeft(e){let t=e.getBoundingClientRect();return this.isRotate?t.top:t.left}postMessage(e={}){return new Promise(t=>{e.id=Date.now(),this.worker.postMessage(e),this.worker.onmessage=i=>{let{data:a}=i;a.id===e.id&&t(a)}})}filter(e,t){let i=this.states[e]||[];for(let e=0;e<i.length;e++)t(i[e]);return i}setState(e,t){this.states[e.$state]=this.states[e.$state].filter(t=>t!==e),e.$state=t,e.$ref&&(e.$ref.dataset.state=t),this.states[t].push(e)}makeWait(e){this.setState(e,"wait"),e.$ref&&(e.$ref.style.cssText=l.cssText,e.$ref.style.visibility="hidden",e.$ref.style.marginLeft="0px",e.$ref.style.transform="translateX(0px)",e.$ref.style.transition="transform 0s linear 0s",this.$refs.push(e.$ref),e.$ref=null)}update(){let{setStyles:e}=this.utils;return this.timer=window.requestAnimationFrame(async()=>{if(this.art.playing&&!this.isHide){this.filter("emit",e=>{let t=(Date.now()-e.$lastStartTime)/1e3;e.$restTime-=t,e.$lastStartTime=Date.now(),e.$restTime<=0&&this.makeWait(e)});let t=this.readys;for(let i=0;i<t.length;i++){let a=t[i];if(await this.option.beforeVisible(a)){let{clientWidth:t,clientHeight:i}=this.$player;a.$ref=this.$ref,a.$ref.innerText=a.text,this.$danmuku.appendChild(a.$ref),a.$ref.style.opacity=this.option.opacity,a.$ref.style.fontSize=`${this.fontSize}px`,a.$ref.style.color=a.color,a.$ref.style.border=a.border?`1px solid ${a.color}`:null,a.$ref.style.backgroundColor=a.border?"rgb(0 0 0 / 50%)":null,e(a.$ref,a.style),a.$lastStartTime=Date.now(),a.$restTime=this.speed;let n=t+a.$ref.clientWidth,{top:o}=await this.postMessage({target:{mode:a.mode,height:a.$ref.clientHeight,speed:n/a.$restTime},visibles:this.visibles,antiOverlap:this.option.antiOverlap,clientWidth:t,clientHeight:i,marginBottom:this.marginBottom,marginTop:this.marginTop});if(a.$ref){if(this.isStop||void 0===o)this.setState(a,"ready"),this.$refs.push(a.$ref),a.$ref=null;else{switch(this.setState(a,"emit"),a.$ref.style.top=`${o}px`,a.$ref.style.visibility="visible",a.$ref.dataset.mode=a.mode,a.mode){case 0:a.$ref.style.left=`${t}px`,a.$ref.style.marginLeft="0px",a.$ref.style.transform=`translateX(${-n}px)`,a.$ref.style.transition=`transform ${a.$restTime}s linear 0s`;break;case 1:case 2:a.$ref.style.left="50%",a.$ref.style.marginLeft=`-${a.$ref.clientWidth/2}px`}this.art.emit("artplayerPluginDanmuku:visible",a)}}}}}this.isStop||this.update()}),this}resize(){let{clientWidth:e}=this.$player;this.filter("stop",t=>{0===t.mode&&(t.$ref.style.left=`${e}px`)}),this.filter("emit",t=>{if(t.$lastStartTime=Date.now(),0===t.mode){let i=e+t.$ref.clientWidth;t.$ref.style.left=`${e}px`,t.$ref.style.transform=`translateX(${-i}px)`,t.$ref.style.transition=`transform ${t.$restTime}s linear 0s`}})}continue(){let{clientWidth:e}=this.$player;return this.filter("stop",t=>{if(this.setState(t,"emit"),t.$lastStartTime=Date.now(),0===t.mode){let i=e+t.$ref.clientWidth;t.$ref.style.transform=`translateX(${-i}px)`,t.$ref.style.transition=`transform ${t.$restTime}s linear 0s`}}),this}suspend(){let{clientWidth:e}=this.$player;return this.filter("emit",t=>{if(this.setState(t,"stop"),0===t.mode){let i=e-(this.getLeft(t.$ref)-this.getLeft(this.$player));t.$ref.style.transform=`translateX(${-i}px)`,t.$ref.style.transition="transform 0s linear 0s"}}),this}stop(){return this.isStop=!0,this.suspend(),window.cancelAnimationFrame(this.timer),this.art.emit("artplayerPluginDanmuku:stop"),this}start(){return this.isStop=!1,this.continue(),this.update(),this.art.emit("artplayerPluginDanmuku:start"),this}reset(){return this.queue.forEach(e=>this.makeWait(e)),this.art.emit("artplayerPluginDanmuku:reset"),this}show(){return this.isHide=!1,this.$danmuku.style.opacity=1,this.option.visible=!0,this.art.emit("artplayerPluginDanmuku:show"),this}hide(){return this.isHide=!0,this.$danmuku.style.opacity=0,this.option.visible=!1,this.art.emit("artplayerPluginDanmuku:hide"),this}destroy(){this.stop(),this.worker.terminate(),this.art.off("video:play",this.start),this.art.off("video:playing",this.start),this.art.off("video:pause",this.stop),this.art.off("video:waiting",this.stop),this.art.off("resize",this.reset),this.art.off("destroy",this.destroy),this.art.emit("artplayerPluginDanmuku:destroy")}}i.default=l},{"./bilibili":"f83sx","bundle-text:./worker":"lfIAi","@parcel/transformer-js/src/esmodule-helpers.js":"9pCYc"}],f83sx:[function(e,t,i){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");function n(e){switch(e){case 1:case 2:case 3:default:return 0;case 4:return 2;case 5:return 1}}function o(e){return"string"!=typeof e?[]:Array.from(e.matchAll(/<d (?:.*? )??p="(?<p>.+?)"(?: .*?)?>(?<text>.+?)<\/d>/gs)).map(e=>{let t=e.groups.p.split(",");return t.length>=8?{text:e.groups.text.trim().replaceAll("&quot;",'"').replaceAll("&apos;","'").replaceAll("&lt;","<").replaceAll("&gt;",">").replaceAll("&amp;","&"),time:Number(t[0]),mode:n(Number(t[1])),fontSize:Number(t[2]),color:`#${Number(t[3]).toString(16)}`,timestamp:Number(t[4]),pool:Number(t[5]),userID:t[6],rowID:Number(t[7])}:null}).filter(Boolean)}async function s(e){let t=await fetch(e);return o(await t.text())}a.defineInteropFlag(i),a.export(i,"getMode",()=>n),a.export(i,"bilibiliDanmuParseFromXml",()=>o),a.export(i,"bilibiliDanmuParseFromUrl",()=>s)},{"@parcel/transformer-js/src/esmodule-helpers.js":"9pCYc"}],"9pCYc":[function(e,t,i){i.interopDefault=function(e){return e&&e.__esModule?e:{default:e}},i.defineInteropFlag=function(e){Object.defineProperty(e,"__esModule",{value:!0})},i.exportAll=function(e,t){return Object.keys(e).forEach(function(i){"default"===i||"__esModule"===i||Object.prototype.hasOwnProperty.call(t,i)||Object.defineProperty(t,i,{enumerable:!0,get:function(){return e[i]}})}),t},i.export=function(e,t,i){Object.defineProperty(e,t,{enumerable:!0,get:i})}},{}],lfIAi:[function(e,t,i){t.exports='!function(e,t,n,o,i){var r="undefined"!=typeof globalThis?globalThis:"undefined"!=typeof self?self:"undefined"!=typeof window?window:"undefined"!=typeof global?global:{},f="function"==typeof r[o]&&r[o],l=f.cache||{},d="undefined"!=typeof module&&"function"==typeof module.require&&module.require.bind(module);function u(t,n){if(!l[t]){if(!e[t]){var i="function"==typeof r[o]&&r[o];if(!n&&i)return i(t,!0);if(f)return f(t,!0);if(d&&"string"==typeof t)return d(t);var h=Error("Cannot find module \'"+t+"\'");throw h.code="MODULE_NOT_FOUND",h}p.resolve=function(n){var o=e[t][1][n];return null!=o?o:n},p.cache={};var s=l[t]=new u.Module(t);e[t][0].call(s.exports,p,s,s.exports,this)}return l[t].exports;function p(e){var t=p.resolve(e);return!1===t?{}:u(t)}}u.isParcelRequire=!0,u.Module=function(e){this.id=e,this.bundle=u,this.exports={}},u.modules=e,u.cache=l,u.parent=f,u.register=function(t,n){e[t]=[function(e,t){t.exports=n},{}]},Object.defineProperty(u,"root",{get:function(){return r[o]}}),r[o]=u;for(var h=0;h<t.length;h++)u(t[h]);if(n){var s=u(n);"object"==typeof exports&&"undefined"!=typeof module?module.exports=s:"function"==typeof define&&define.amd&&define(function(){return s})}}({"59OZS":[function(e,t,n){onmessage=e=>{let{data:t}=e;if(!t.id)return;let n=function({target:e,visibles:t,clientWidth:n,clientHeight:o,marginBottom:i,marginTop:r,antiOverlap:f}){let l=o-i,d=t.filter(t=>t.mode===e.mode&&t.top<=l).sort((e,t)=>e.top-t.top);if(0===d.length)return 2===e.mode?l-e.height:r;if(d.unshift({type:"top",top:0,left:0,right:0,height:r,width:n,speed:0,distance:n}),d.push({type:"bottom",top:l,left:0,right:0,height:i,width:n,speed:0,distance:n}),2===e.mode)for(let t=d.length-2;t>=0;t-=1){let n=d[t],o=d[t+1],i=n.top+n.height;if(o.top-i>=e.height)return o.top-e.height}else for(let t=1;t<d.length;t+=1){let n=d[t],o=d[t-1],i=o.top+o.height;if(n.top-i>=e.height)return i}let u=[];for(let e=1;e<d.length-1;e+=1){let t=d[e];if(u.length){let e=u[u.length-1];e[0].top===t.top?e.push(t):u.push([t])}else u.push([t])}if(f)switch(e.mode){case 0:{let t=u.find(t=>t.every(t=>!(n<t.distance)&&(e.speed<t.speed||t.right/(e.speed-t.speed)>t.time)));return t&&t[0]?t[0].top:void 0}case 1:case 2:return}else{switch(e.mode){case 0:u.sort((e,t)=>{let n=Math.min(...t.map(e=>e.right)),o=Math.min(...e.map(e=>e.right));return n*t.length-o*e.length});break;case 1:case 2:u.sort((e,t)=>{let n=Math.max(...t.map(e=>e.width));return Math.max(...e.map(e=>e.width))*e.length-n*t.length})}return u[0][0].top}}(t);self.postMessage({top:n,id:t.id})}},{}]},["59OZS"],"59OZS","parcelRequire4dc0");'},{}],lO8OT:[function(e,t,i){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(i);var n=e("bundle-text:./style.less"),o=a.interopDefault(n),s=e("bundle-text:./img/on.svg"),l=a.interopDefault(s),r=e("bundle-text:./img/off.svg"),d=a.interopDefault(r),p=e("bundle-text:./img/config.svg"),u=a.interopDefault(p),h=e("bundle-text:./img/style.svg"),m=a.interopDefault(h),c=e("bundle-text:./img/mode_0_off.svg"),f=a.interopDefault(c),g=e("bundle-text:./img/mode_0_on.svg"),v=a.interopDefault(g),y=e("bundle-text:./img/mode_1_off.svg"),x=a.interopDefault(y),k=e("bundle-text:./img/mode_1_on.svg"),b=a.interopDefault(k),$=e("bundle-text:./img/mode_2_off.svg"),w=a.interopDefault($),M=e("bundle-text:./img/mode_2_on.svg"),E=a.interopDefault(M),D=e("bundle-text:./img/check_on.svg"),S=a.interopDefault(D),z=e("bundle-text:./img/check_off.svg"),A=a.interopDefault(z);if(i.default=class{constructor(e,t){this.art=e,this.danmuku=t,this.utils=e.constructor.utils;let{setStyle:i}=this.utils,{$controlsCenter:a}=e.template;i(a,"display","flex"),this.template={$controlsCenter:a,$mount:a,$danmuku:null,$toggle:null,$configModes:null,$styleModes:null,$colors:null,$opacitySlider:null,$opacityValue:null,$marginSlider:null,$marginValue:null,$fontSizeSlider:null,$fontSizeValue:null,$speedSlider:null,$speedValue:null,$input:null,$send:null},this.slider={opacity:null,margin:null,fontSize:null,speed:null},this.emitting=!1,this.isLock=!1,this.timer=null,this.createTemplate(),this.createSliders(),this.createEvents(),this.mount(this.option.mount),e.on("fullscreen",e=>this.onFullscreen(e)),e.on("fullscreenWeb",e=>this.onFullscreen(e))}static get icons(){return{$on:l.default,$off:d.default,$config:u.default,$style:m.default,$mode_0_off:f.default,$mode_0_on:v.default,$mode_1_off:x.default,$mode_1_on:b.default,$mode_2_off:w.default,$mode_2_on:E.default,$check_on:S.default,$check_off:A.default}}get option(){return this.danmuku.option}get outside(){return this.template.$mount!==this.template.$controlsCenter}get TEMPLATE(){let{option:e}=this;return`<div class="apd-toggle">${l.default}${d.default}</div><div class="apd-config">${u.default}<div class="apd-config-panel"><div class="apd-config-panel-inner"><div class="apd-config-mode">\u{6309}\u{7C7B}\u{578B}\u{5C4F}\u{853D}<div class="apd-modes"><div data-mode="0" class="apd-mode">${f.default}${v.default}<div>\u{6EDA}\u{52A8}</div></div><div data-mode="1" class="apd-mode">${x.default}${b.default}<div>\u{9876}\u{90E8}</div></div><div data-mode="2" class="apd-mode">${w.default}${E.default}<div>\u{5E95}\u{90E8}</div></div></div></div><div class="apd-config-other"><div class="apd-other apd-anti-overlap">${S.default}${A.default} \u{9632}\u{6B62}\u{5F39}\u{5E55}\u{91CD}\u{53E0}</div><div class="apd-other apd-sync-video">${S.default}${A.default} \u{540C}\u{6B65}\u{89C6}\u{9891}\u{901F}\u{5EA6}</div></div><div class="apd-config-slider apd-config-opacity">\u{4E0D}\u{900F}\u{660E}\u{5EA6}<div class="apd-slider"></div><div class="apd-value">\u{672A}\u{77E5}</div></div><div class="apd-config-slider apd-config-margin">\u{663E}\u{793A}\u{533A}\u{57DF}<div class="apd-slider"></div><div class="apd-value">\u{672A}\u{77E5}</div></div><div class="apd-config-slider apd-config-fontSize">\u{5F39}\u{5E55}\u{5B57}\u{53F7}<div class="apd-slider"></div><div class="apd-value">\u{672A}\u{77E5}</div></div><div class="apd-config-slider apd-config-speed">\u{5F39}\u{5E55}\u{901F}\u{5EA6}<div class="apd-slider"></div><div class="apd-value">\u{672A}\u{77E5}</div></div></div></div></div><div class="apd-emitter"><div class="apd-style">${m.default}<div class="apd-style-panel"><div class="apd-style-panel-inner"><div class="apd-style-mode">\u{6A21}\u{5F0F}<div class="apd-modes"><div data-mode="0" class="apd-mode">${v.default}<div>\u{6EDA}\u{52A8}</div></div><div data-mode="1" class="apd-mode">${b.default}<div>\u{9876}\u{90E8}</div></div><div data-mode="2" class="apd-mode">${E.default}<div>\u{5E95}\u{90E8}</div></div></div></div><div class="apd-style-color">\u{989C}\u{8272}<div class="apd-colors">${this.COLOR.map(e=>`<div data-color="${e}" class="apd-color" style="background-color: ${e}"></div>`).join("")}</div></div></div></div></div><input class="apd-input" placeholder="\u{53D1}\u{4E2A}\u{53CB}\u{5584}\u{7684}\u{5F39}\u{5E55}\u{89C1}\u{8BC1}\u{5F53}\u{4E0B}" autocomplete="off" maxLength="${e.maxLength}" /><div class="apd-send">\u{53D1}\u{9001}</div></div>`}get OPACITY(){return{min:0,max:100,steps:[],...this.option.OPACITY}}get FONT_SIZE(){return{min:12,max:120,steps:[],...this.option.FONT_SIZE}}get MARGIN(){return{min:0,max:3,steps:[{name:"1/4",value:[10,"75%"]},{name:"半屏",value:[10,"50%"]},{name:"3/4",value:[10,"25%"]},{name:"满屏",value:[10,10]}],...this.option.MARGIN}}get SPEED(){return{min:0,max:4,steps:[{name:"极慢",value:10},{name:"较慢",value:7.5,hide:!0},{name:"适中",value:5},{name:"较快",value:2.5,hide:!0},{name:"极快",value:1}],...this.option.SPEED}}get COLOR(){return this.option.COLOR.length?this.option.COLOR:["#FE0302","#FF7204","#FFAA02","#FFD302","#FFFF00","#A0EE00","#00CD00","#019899","#4266BE","#89D5FF","#CC0273","#222222","#9B9B9B","#FFFFFF"]}query(e){let{query:t}=this.utils,{$danmuku:i}=this.template;return t(e,i)}setData(e,t){let{$player:i}=this.art.template,{$mount:a}=this.template;i.dataset[e]=t,this.outside&&(a.dataset[e]=t)}createTemplate(){let{createElement:e,tooltip:t}=this.utils,i=e("div");i.className="artplayer-plugin-danmuku",i.innerHTML=this.TEMPLATE,this.template.$danmuku=i,this.template.$toggle=this.query(".apd-toggle"),this.template.$configModes=this.query(".apd-config-mode .apd-modes"),this.template.$styleModes=this.query(".apd-style-mode .apd-modes"),this.template.$colors=this.query(".apd-colors"),this.template.$antiOverlap=this.query(".apd-anti-overlap"),this.template.$syncVideo=this.query(".apd-sync-video"),this.template.$opacitySlider=this.query(".apd-config-opacity .apd-slider"),this.template.$opacityValue=this.query(".apd-config-opacity .apd-value"),this.template.$marginSlider=this.query(".apd-config-margin .apd-slider"),this.template.$marginValue=this.query(".apd-config-margin .apd-value"),this.template.$fontSizeSlider=this.query(".apd-config-fontSize .apd-slider"),this.template.$fontSizeValue=this.query(".apd-config-fontSize .apd-value"),this.template.$speedSlider=this.query(".apd-config-speed .apd-slider"),this.template.$speedValue=this.query(".apd-config-speed .apd-value"),this.template.$input=this.query(".apd-input"),this.template.$send=this.query(".apd-send");let{$toggle:a}=this.template;this.art.on("artplayerPluginDanmuku:show",()=>{t(a,"关闭弹幕")}),this.art.on("artplayerPluginDanmuku:hide",()=>{t(a,"打开弹幕")})}createEvents(){let{$toggle:e,$configModes:t,$styleModes:i,$colors:a,$antiOverlap:n,$syncVideo:o,$send:s,$input:l}=this.template;this.art.proxy(e,"click",()=>{this.danmuku.config({visible:!this.option.visible}),this.reset()}),this.art.proxy(t,"click",e=>{let t=e.target.closest(".apd-mode");if(!t)return;let i=Number(t.dataset.mode);this.option.modes.includes(i)?this.danmuku.config({modes:this.option.modes.filter(e=>e!==i)}):this.danmuku.config({modes:[...this.option.modes,i]}),this.reset()}),this.art.proxy(n,"click",()=>{this.danmuku.config({antiOverlap:!this.option.antiOverlap}),this.reset()}),this.art.proxy(o,"click",()=>{this.danmuku.config({synchronousPlayback:!this.option.synchronousPlayback}),this.reset()}),this.art.proxy(i,"click",e=>{let t=e.target.closest(".apd-mode");if(!t)return;let i=Number(t.dataset.mode);this.danmuku.config({mode:i}),this.reset()}),this.art.proxy(a,"click",e=>{let t=e.target.closest(".apd-color");t&&(this.danmuku.config({color:t.dataset.color}),this.reset())}),this.art.proxy(s,"click",()=>this.emit()),this.art.proxy(l,"keypress",e=>{"Enter"===e.key&&(e.preventDefault(),this.emit())})}createSliders(){this.slider.opacity=this.createSlider({...this.OPACITY,container:this.template.$opacitySlider,findIndex:()=>Math.round(100*this.option.opacity),onChange:e=>{let{$opacityValue:t}=this.template;t.textContent=`${e}%`,this.danmuku.config({opacity:e/100})}}),this.slider.margin=this.createSlider({...this.MARGIN,container:this.template.$marginSlider,findIndex:()=>this.MARGIN.steps.findIndex(e=>e.value[0]===this.option.margin[0]&&e.value[1]===this.option.margin[1]),onChange:e=>{let t=this.MARGIN.steps[e];if(!t)return;let{$marginValue:i}=this.template;i.textContent=t.name,this.danmuku.config({margin:t.value})}}),this.slider.fontSize=this.createSlider({...this.FONT_SIZE,container:this.template.$fontSizeSlider,findIndex:()=>Math.round(this.danmuku.fontSize),onChange:e=>{let{$fontSizeValue:t}=this.template;t.textContent=`${e}px`,this.danmuku.config({fontSize:e})}}),this.slider.speed=this.createSlider({...this.SPEED,container:this.template.$speedSlider,findIndex:()=>this.SPEED.steps.findIndex(e=>e.value===this.option.speed),onChange:e=>{let t=this.SPEED.steps[e];if(!t)return;let{$speedValue:i}=this.template;i.textContent=t.name,this.danmuku.config({speed:t.value})}})}createSlider({min:e,max:t,container:i,findIndex:a,onChange:n,steps:o=[]}){let{query:s,clamp:l}=this.utils;i.innerHTML=`<div class="apd-slider-line"><div class="apd-slider-points">${o.map(()=>'<div class="apd-slider-point"></div>').join("")}</div><div class="apd-slider-progress"></div></div><div class="apd-slider-dot"></div><div class="apd-slider-steps">${o.map(e=>e.hide?"":`<div class="apd-slider-step">${e.name}</div>`).join("")}</div>`;let r=s(".apd-slider-dot",i),d=s(".apd-slider-progress",i),p=!1;function u(i=a()){if(i<e||i>t)return;let s=(i-e)/(t-e);r.style.left=`${100*s}%`,0===o.length&&(d.style.width=r.style.left),n(i)}function h(a){let{left:n,width:o}=i.getBoundingClientRect();u(Math.round(l(a.clientX-n,0,o)/o*(t-e)+e))}return this.art.proxy(i,"click",e=>{h(e)}),this.art.proxy(i,"mousedown",e=>{p=0===e.button}),this.art.on("document:mousemove",e=>{p&&h(e)}),this.art.on("document:mouseup",e=>{p&&(p=!1,h(e))}),{reset:u}}onFullscreen(e){if(this.outside){let{$danmuku:t,$controlsCenter:i,$mount:a}=this.template;e?i.appendChild(t):a.appendChild(t)}}async emit(){let{$input:e}=this.template,t=e.value.trim();if(!t.length||this.isLock||this.emitting)return;let i={text:t,mode:this.option.mode,color:this.option.color,time:this.art.currentTime};try{this.emitting=!0;let t=await this.option.beforeEmit(i);if(this.emitting=!1,!0!==t)return;i.border=!0,delete i.time,this.danmuku.emit(i),e.value="",this.lock()}catch(e){this.emitting=!1}}lock(){let{addClass:e}=this.utils,{$send:t}=this.template;this.isLock=!0;let i=this.option.lockTime;t.innerText=i,e(t,"apd-lock");let a=()=>{this.timer=setTimeout(()=>{0===i?this.unlock():(i-=1,t.innerText=i,a())},1e3)};a()}unlock(){let{removeClass:e}=this.utils,{$send:t}=this.template;clearTimeout(this.timer),this.isLock=!1,t.innerText="发送",e(t,"apd-lock")}reset(){let{inverseClass:e,tooltip:t}=this.utils,{$toggle:i,$colors:a}=this.template;this.slider.opacity.reset(),this.slider.margin.reset(),this.slider.fontSize.reset(),this.slider.speed.reset(),this.setData("danmukuVisible",this.option.visible),this.setData("danmukuMode",this.option.mode),this.setData("danmukuColor",this.option.color),this.setData("danmukuMode0",this.option.modes.includes(0)),this.setData("danmukuMode1",this.option.modes.includes(1)),this.setData("danmukuMode2",this.option.modes.includes(2)),this.setData("danmukuAntiOverlap",this.option.antiOverlap),this.setData("danmukuSyncVideo",this.option.synchronousPlayback),this.setData("danmukuTheme",this.option.theme),this.setData("danmukuEmitter",this.option.emitter);let n=Array.from(a.children).find(e=>e.dataset.color===this.option.color.toUpperCase());n&&e(n,"apd-active"),t(i,this.option.visible?"关闭弹幕":"打开弹幕")}mount(e){let t="string"==typeof e?document.querySelector(e):e;t.appendChild(this.template.$danmuku),this.template.$mount=t,this.reset()}},"undefined"!=typeof document){let e="artplayer-plugin-danmuku",t=document.getElementById(e);if(t)t.textContent=o.default;else{let t=document.createElement("style");t.id=e,t.textContent=o.default,document.head.appendChild(t)}}},{"bundle-text:./style.less":"hViDo","bundle-text:./img/on.svg":"9pjcf","bundle-text:./img/off.svg":"b2dkP","bundle-text:./img/config.svg":"l8tyy","bundle-text:./img/style.svg":"5iZC3","bundle-text:./img/mode_0_off.svg":"i0Vut","bundle-text:./img/mode_0_on.svg":"hOSvZ","bundle-text:./img/mode_1_off.svg":"bOXC3","bundle-text:./img/mode_1_on.svg":"lKuh0","bundle-text:./img/mode_2_off.svg":"eB8W6","bundle-text:./img/mode_2_on.svg":"bpe2E","bundle-text:./img/check_on.svg":"kL9zy","bundle-text:./img/check_off.svg":"22xpM","@parcel/transformer-js/src/esmodule-helpers.js":"9pCYc"}],hViDo:[function(e,t,i){t.exports='.artplayer-plugin-danmuku{z-index:99;color:#fff;flex-shrink:0;justify-content:center;align-items:center;gap:10px;width:100%;height:32px;font-size:12px;font-weight:300;display:flex;position:relative}.artplayer-plugin-danmuku .apd-icon{cursor:pointer;opacity:.75;fill:#fff;transition:all .2s}.artplayer-plugin-danmuku .apd-icon:hover{opacity:1}.artplayer-plugin-danmuku .apd-config{display:flex;position:relative}.artplayer-plugin-danmuku .apd-config .apd-config-panel{width:320px;padding:10px;display:none;position:absolute;bottom:24px;left:-148px}.artplayer-plugin-danmuku .apd-config .apd-config-panel .apd-config-panel-inner{background-color:#000000d9;border-radius:3px;width:100%;padding:10px}.artplayer-plugin-danmuku .apd-config:hover .apd-config-panel{display:flex}.artplayer-plugin-danmuku .apd-config-mode,.artplayer-plugin-danmuku .apd-config-slider,.artplayer-plugin-danmuku .apd-config-other,.artplayer-plugin-danmuku .apd-style-mode{margin-bottom:15px}.artplayer-plugin-danmuku .apd-modes{align-items:center;gap:20px;margin-top:5px;display:flex}.artplayer-plugin-danmuku .apd-modes .apd-mode{cursor:pointer;text-align:center}.artplayer-plugin-danmuku .apd-modes .apd-mode:hover{color:#00a1d6}.artplayer-plugin-danmuku .apd-config-slider{align-items:center;gap:12px;display:flex}.artplayer-plugin-danmuku .apd-config-slider .apd-value{text-align:right;width:32px}.artplayer-plugin-danmuku .apd-slider{cursor:pointer;flex:1;justify-content:center;align-items:center;height:20px;display:flex;position:relative}.artplayer-plugin-danmuku .apd-slider .apd-slider-line{background-color:#ffffff40;border-radius:3px;width:100%;height:2px;position:relative;overflow:hidden}.artplayer-plugin-danmuku .apd-slider .apd-slider-points{justify-content:space-between;align-items:center;display:flex;position:absolute;inset:0}.artplayer-plugin-danmuku .apd-slider .apd-slider-points .apd-slider-point{background-color:#ffffff80;border-radius:50%;width:2px;height:2px}.artplayer-plugin-danmuku .apd-slider .apd-slider-progress{background-color:#00a1d6;width:0%;height:100%}.artplayer-plugin-danmuku .apd-slider .apd-slider-dot{background-color:#00a1d6;border-radius:50%;width:12px;height:12px;position:absolute;left:0%;transform:translate(-6px)}.artplayer-plugin-danmuku .apd-slider .apd-slider-steps{color:#777;justify-content:space-between;align-items:center;width:calc(100% + 32px);display:flex;position:absolute;bottom:-12px}.artplayer-plugin-danmuku .apd-slider .apd-slider-steps .apd-slider-step{text-align:center;flex-shrink:0;width:36px;scale:.95}.artplayer-plugin-danmuku .apd-config-other{align-items:center;gap:20px;display:flex}.artplayer-plugin-danmuku .apd-config-other .apd-check-off,.artplayer-plugin-danmuku .apd-config-other .apd-check-on{width:16px;height:16px}.artplayer-plugin-danmuku .apd-config-other .apd-other{cursor:pointer;align-items:center;gap:2px;display:flex}.artplayer-plugin-danmuku .apd-config-other .apd-other:hover{color:#00a1d6}.artplayer-plugin-danmuku .apd-emitter{background-color:#ffffff40;border-radius:5px;flex:1;align-items:center;height:100%;display:flex}.artplayer-plugin-danmuku .apd-style{justify-content:center;align-items:center;display:flex;position:relative}.artplayer-plugin-danmuku .apd-style .apd-style-panel{width:200px;padding:10px;display:none;position:absolute;bottom:24px;left:-88px}.artplayer-plugin-danmuku .apd-style .apd-style-panel .apd-style-panel-inner{background-color:#000000d9;border-radius:3px;width:100%;padding:10px}.artplayer-plugin-danmuku .apd-style:hover .apd-style-panel{display:flex}.artplayer-plugin-danmuku .apd-colors{flex-wrap:wrap;gap:8px;margin-top:5px;display:flex}.artplayer-plugin-danmuku .apd-colors .apd-color{cursor:pointer;border-radius:2px;width:16px;height:16px}.artplayer-plugin-danmuku .apd-colors .apd-color.apd-active{border:1px solid #000;box-shadow:0 0 0 1px #fff}.artplayer-plugin-danmuku .apd-input{color:#fff;background-color:#0000;border:none;outline:none;flex:1;width:auto;min-width:0;height:100%;line-height:1}.artplayer-plugin-danmuku .apd-input::placeholder{color:#ffffff80}.artplayer-plugin-danmuku .apd-send{cursor:pointer;text-shadow:none;background-color:#00a1d6;border-top-right-radius:5px;border-bottom-right-radius:5px;flex-shrink:0;justify-content:center;align-items:center;width:60px;height:100%;display:flex}.artplayer-plugin-danmuku .apd-send.apd-lock{cursor:not-allowed;color:#666;background-color:#e7e7e7}.art-controls-center .apd-emitter{flex:none;width:260px}.art-fullscreen .artplayer-plugin-danmuku,.art-fullscreen-web .artplayer-plugin-danmuku{gap:16px;height:38px}.art-fullscreen .artplayer-plugin-danmuku .apd-config-icon,.art-fullscreen-web .artplayer-plugin-danmuku .apd-config-icon,.art-fullscreen .artplayer-plugin-danmuku .apd-toggle-off,.art-fullscreen-web .artplayer-plugin-danmuku .apd-toggle-off,.art-fullscreen .artplayer-plugin-danmuku .apd-toggle-on,.art-fullscreen-web .artplayer-plugin-danmuku .apd-toggle-on{width:28px;height:28px}.art-fullscreen .artplayer-plugin-danmuku .apd-emitter,.art-fullscreen-web .artplayer-plugin-danmuku .apd-emitter{flex:none;width:400px}[data-danmuku-emitter=false] .apd-emitter{display:none!important}[data-danmuku-emitter=false] .art-controls-center .artplayer-plugin-danmuku{justify-content:flex-end;gap:18px}[data-danmuku-emitter=false].art-fullscreen .art-controls-center .artplayer-plugin-danmuku,[data-danmuku-emitter=false].art-fullscreen-web .art-controls-center .artplayer-plugin-danmuku{gap:24px}[data-danmuku-theme=light]>.artplayer-plugin-danmuku .apd-icon{fill:#333}[data-danmuku-theme=light]>.artplayer-plugin-danmuku .apd-emitter{background-color:#f1f2f3}[data-danmuku-theme=light]>.artplayer-plugin-danmuku .apd-input{color:#000}[data-danmuku-theme=light]>.artplayer-plugin-danmuku .apd-input::placeholder{color:#0000004d}[data-danmuku-visible=false] .apd-toggle-off{display:block}[data-danmuku-visible=false] .apd-toggle-on,[data-danmuku-visible=true] .apd-toggle-off{display:none}[data-danmuku-visible=true] .apd-toggle-on{display:block}[data-danmuku-anti-overlap=false] .apd-anti-overlap .apd-check-on{display:none}[data-danmuku-anti-overlap=false] .apd-anti-overlap .apd-check-off,[data-danmuku-anti-overlap=true] .apd-anti-overlap .apd-check-on{display:block}[data-danmuku-anti-overlap=true] .apd-anti-overlap .apd-check-off,[data-danmuku-sync-video=false] .apd-sync-video .apd-check-on{display:none}[data-danmuku-sync-video=false] .apd-sync-video .apd-check-off,[data-danmuku-sync-video=true] .apd-sync-video .apd-check-on{display:block}[data-danmuku-sync-video=true] .apd-sync-video .apd-check-off{display:none}[data-danmuku-mode0=false] .apd-config-mode .apd-mode-0-off{display:block}[data-danmuku-mode0=false] .apd-config-mode .apd-mode-0-on{display:none}[data-danmuku-mode0=false] .art-danmuku [data-mode="0"]{opacity:0!important}[data-danmuku-mode0=true] .apd-config-mode .apd-mode-0-off{display:none}[data-danmuku-mode0=true] .apd-config-mode .apd-mode-0-on{display:block}[data-danmuku-mode="0"] .apd-style-mode [data-mode="0"]{color:#00a1d6}[data-danmuku-mode="0"] .apd-style-mode [data-mode="0"] path{fill:#00a1d6}[data-danmuku-mode1=false] .apd-config-mode .apd-mode-1-off{display:block}[data-danmuku-mode1=false] .apd-config-mode .apd-mode-1-on{display:none}[data-danmuku-mode1=false] .art-danmuku [data-mode="1"]{opacity:0!important}[data-danmuku-mode1=true] .apd-config-mode .apd-mode-1-off{display:none}[data-danmuku-mode1=true] .apd-config-mode .apd-mode-1-on{display:block}[data-danmuku-mode="1"] .apd-style-mode [data-mode="1"]{color:#00a1d6}[data-danmuku-mode="1"] .apd-style-mode [data-mode="1"] path{fill:#00a1d6}[data-danmuku-mode2=false] .apd-config-mode .apd-mode-2-off{display:block}[data-danmuku-mode2=false] .apd-config-mode .apd-mode-2-on{display:none}[data-danmuku-mode2=false] .art-danmuku [data-mode="2"]{opacity:0!important}[data-danmuku-mode2=true] .apd-config-mode .apd-mode-2-off{display:none}[data-danmuku-mode2=true] .apd-config-mode .apd-mode-2-on{display:block}[data-danmuku-mode="2"] .apd-style-mode [data-mode="2"]{color:#00a1d6}[data-danmuku-mode="2"] .apd-style-mode [data-mode="2"] path{fill:#00a1d6}'},{}],"9pjcf":[function(e,t,i){t.exports='<svg class="apd-icon apd-toggle-on" xmlns="http://www.w3.org/2000/svg" data-pointer="none" viewBox="0 0 24 24" width="24" height="24"><path fill-rule="evenodd" d="M11.989 4.828c-.47 0-.975.004-1.515.012l-1.71-2.566a1.008 1.008 0 0 0-1.678 1.118l.999 1.5c-.681.018-1.403.04-2.164.068a4.013 4.013 0 0 0-3.83 3.44c-.165 1.15-.245 2.545-.245 4.185 0 1.965.115 3.67.35 5.116a4.012 4.012 0 0 0 3.763 3.363l.906.046c1.205.063 1.808.095 3.607.095a.988.988 0 0 0 0-1.975c-1.758 0-2.339-.03-3.501-.092l-.915-.047a2.037 2.037 0 0 1-1.91-1.708c-.216-1.324-.325-2.924-.325-4.798 0-1.563.076-2.864.225-3.904.14-.977.96-1.713 1.945-1.747 2.444-.087 4.465-.13 6.063-.131 1.598 0 3.62.044 6.064.13.96.034 1.71.81 1.855 1.814.075.524.113 1.962.141 3.065v.002c.01.342.017.65.025.88a.987.987 0 1 0 1.974-.068c-.008-.226-.016-.523-.025-.856v-.027c-.03-1.118-.073-2.663-.16-3.276-.273-1.906-1.783-3.438-3.74-3.507-.9-.032-1.743-.058-2.531-.078l1.05-1.46a1.008 1.008 0 0 0-1.638-1.177l-1.862 2.59c-.38-.004-.744-.007-1.088-.007h-.13Zm.521 4.775h-1.32v4.631h2.222v.847h-2.618v1.078h2.618l.003.678c.36.026.714.163 1.01.407h.11v-1.085h2.694v-1.078h-2.695v-.847H16.8v-4.63h-1.276a8.59 8.59 0 0 0 .748-1.42L15.183 7.8a14.232 14.232 0 0 1-.814 1.804h-1.518l.693-.308a8.862 8.862 0 0 0-.814-1.408l-1.045.352c.297.396.572.847.825 1.364Zm-4.18 3.564.154-1.485h1.98V8.294h-3.2v.98H9.33v1.43H7.472l-.308 3.453h2.277c0 1.166-.044 1.925-.12 2.277-.078.352-.386.528-.936.528-.308 0-.616-.022-.902-.055l.297 1.067.062.005c.285.02.551.04.818.04 1.001-.067 1.562-.419 1.694-1.057.11-.638.176-1.903.176-3.795h-2.2Zm7.458.11v-.858h-1.254v.858h1.254Zm-2.376-.858v.858h-1.199v-.858h1.2Zm-1.199-.946h1.2v-.902h-1.2v.902Zm2.321 0v-.902h1.254v.902h-1.254Z" clip-rule="evenodd"/><path fill="#00AEEC" fill-rule="evenodd" d="M22.846 14.627a1 1 0 0 0-1.412.075l-5.091 5.703-2.216-2.275-.097-.086-.008-.005a1 1 0 0 0-1.322 1.493l2.963 3.041.093.083.007.005a1 1 0 0 0 1.354-.124l5.81-6.505.08-.102.005-.008a1 1 0 0 0-.166-1.295Z" clip-rule="evenodd"/></svg>'},{}],b2dkP:[function(e,t,i){t.exports='<svg class="apd-icon apd-toggle-off" xmlns="http://www.w3.org/2000/svg" data-pointer="none" viewBox="0 0 24 24" width="24" height="24"><path fill-rule="evenodd" d="m8.085 4.891-.999-1.499a1.008 1.008 0 0 1 1.679-1.118l1.709 2.566c.54-.008 1.045-.012 1.515-.012h.13c.345 0 .707.003 1.088.007l1.862-2.59a1.008 1.008 0 0 1 1.637 1.177l-1.049 1.46c.788.02 1.631.046 2.53.078 1.958.069 3.468 1.6 3.74 3.507.088.613.13 2.158.16 3.276l.001.027c.01.333.017.63.025.856a.987.987 0 0 1-1.974.069c-.008-.23-.016-.539-.025-.881v-.002c-.028-1.103-.066-2.541-.142-3.065-.143-1.004-.895-1.78-1.854-1.813a179.14 179.14 0 0 0-6.064-.131c-1.598 0-3.619.044-6.063.13a2.037 2.037 0 0 0-1.945 1.748c-.15 1.04-.225 2.341-.225 3.904 0 1.874.11 3.474.325 4.798.154.949.95 1.66 1.91 1.708a97.58 97.58 0 0 0 5.416.139.988.988 0 0 1 0 1.975 99.78 99.78 0 0 1-5.513-.141A4.012 4.012 0 0 1 2.197 17.7c-.236-1.446-.351-3.151-.351-5.116 0-1.64.08-3.035.245-4.184A4.013 4.013 0 0 1 5.92 4.96c.761-.027 1.483-.05 2.164-.069Zm4.436 4.707h-1.32v4.63h2.222v.848h-2.618v1.078h2.431a5.01 5.01 0 0 1 3.575-3.115V9.598h-1.276a8.59 8.59 0 0 0 .748-1.42l-1.089-.384a14.232 14.232 0 0 1-.814 1.804h-1.518l.693-.308a8.862 8.862 0 0 0-.814-1.408l-1.045.352c.297.396.572.847.825 1.364Zm-4.18 3.564.154-1.485h1.98V8.289h-3.2v.979h2.067v1.43H7.483l-.308 3.454h2.277c0 1.166-.044 1.925-.12 2.277-.078.352-.386.528-.936.528-.308 0-.616-.022-.902-.055l.297 1.067.062.004c.285.02.551.04.818.04 1.001-.066 1.562-.418 1.694-1.056.11-.638.176-1.903.176-3.795h-2.2Zm7.458.11v-.858h-1.254v.858H15.8Zm-2.376-.858v.858h-1.199v-.858h1.2Zm-1.199-.946h1.2v-.902h-1.2v.902Zm2.321 0v-.902H15.8v.902h-1.254Zm3.517 10.594a4 4 0 1 0 0-8 4 4 0 0 0 0 8Zm-.002-1.502a2.5 2.5 0 0 1-2.217-3.657l3.326 3.398a2.49 2.49 0 0 1-1.109.259Zm2.5-2.5c0 .42-.103.815-.286 1.162l-3.328-3.401a2.5 2.5 0 0 1 3.614 2.239Z" clip-rule="evenodd"/></svg>'},{}],l8tyy:[function(e,t,i){t.exports='<svg class="apd-icon apd-config-icon" xmlns="http://www.w3.org/2000/svg" data-pointer="none" viewBox="0 0 24 24" width="24" height="24"><path fill-rule="evenodd" d="m15.645 4.881 1.06-1.473a.998.998 0 1 0-1.622-1.166L13.22 4.835a110.67 110.67 0 0 0-1.1-.007h-.131c-.47 0-.975.004-1.515.012L8.783 2.3A.998.998 0 0 0 7.12 3.408l.988 1.484c-.688.019-1.418.042-2.188.069a4.013 4.013 0 0 0-3.83 3.44c-.165 1.15-.245 2.545-.245 4.185 0 1.965.115 3.67.35 5.116a4.012 4.012 0 0 0 3.763 3.363 99.78 99.78 0 0 0 5.513.141.988.988 0 0 0 0-1.975 97.58 97.58 0 0 1-5.416-.139 2.037 2.037 0 0 1-1.91-1.708c-.216-1.324-.325-2.924-.325-4.798 0-1.563.076-2.864.225-3.904.14-.977.96-1.713 1.945-1.747 2.444-.087 4.465-.13 6.063-.131 1.598 0 3.62.044 6.064.13.96.034 1.71.81 1.855 1.814.075.524.113 1.962.141 3.065v.002c.005.183.01.07.014-.038.004-.096.008-.189.011-.081a.987.987 0 1 0 1.974-.069c-.004-.105-.007-.009-.011.09a2.175 2.175 0 0 1-.007.135l-.002.01a.574.574 0 0 1-.005-.091v-.027c-.03-1.118-.073-2.663-.16-3.276-.273-1.906-1.783-3.438-3.74-3.507-.905-.032-1.752-.058-2.543-.079Zm-3.113 4.703h-1.307v4.643h2.2v.04l.651-1.234c.113-.215.281-.389.482-.509v-.11h.235c.137-.049.283-.074.433-.074h1.553V9.584h-1.264a8.5 8.5 0 0 0 .741-1.405l-1.078-.381c-.24.631-.501 1.23-.806 1.786h-1.503l.686-.305a8.613 8.613 0 0 0-.806-1.394l-1.034.348c.294.392.566.839.817 1.35Zm-1.7 5.502h2.16l-.564 1.068h-1.595v-1.068Zm-2.498-1.863.152-1.561h1.96V8.289H7.277v.969h2.048v1.435h-1.84l-.306 3.51h2.254c0 1.155-.043 1.906-.12 2.255-.076.348-.38.523-.925.523-.305 0-.61-.022-.893-.055l.294 1.056.061.005c.282.02.546.039.81.039.991-.065 1.547-.414 1.677-1.046.11-.631.175-1.883.175-3.757H8.334Zm5.09-.8v.85h-1.188v-.85h1.187Zm-1.188-.955h1.187v-.893h-1.187v.893Zm2.322.007v-.893h1.241v.893h-1.241Zm.528 2.757a1.26 1.26 0 0 1 1.087-.627l4.003-.009a1.26 1.26 0 0 1 1.094.63l1.721 2.982c.226.39.225.872-.001 1.263l-1.743 3a1.26 1.26 0 0 1-1.086.628l-4.003.009a1.26 1.26 0 0 1-1.094-.63l-1.722-2.982a1.26 1.26 0 0 1 .002-1.263l1.742-3Zm1.967.858a1.26 1.26 0 0 0-1.08.614l-.903 1.513a1.26 1.26 0 0 0-.002 1.289l.885 1.492c.227.384.64.62 1.086.618l2.192-.005a1.26 1.26 0 0 0 1.08-.615l.904-1.518a1.26 1.26 0 0 0 .001-1.288l-.884-1.489a1.26 1.26 0 0 0-1.086-.616l-2.193.005Zm2.517 2.76a1.4 1.4 0 1 1-2.8 0 1.4 1.4 0 0 1 2.8 0Z" clip-rule="evenodd"/></svg>'},{}],"5iZC3":[function(e,t,i){t.exports='<svg class="apd-icon apd-style-icon" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" data-pointer="none" style="enable-background:new 0 0 22 22" viewBox="0 0 22 22" width="36" height="24"><path d="M17 16H5c-.55 0-1 .45-1 1s.45 1 1 1h12c.55 0 1-.45 1-1s-.45-1-1-1zM6.96 15c.39 0 .74-.24.89-.6l.65-1.6h5l.66 1.6c.15.36.5.6.89.6.69 0 1.15-.71.88-1.34l-3.88-8.97C11.87 4.27 11.46 4 11 4s-.87.27-1.05.69l-3.88 8.97c-.27.63.2 1.34.89 1.34zM11 5.98 12.87 11H9.13L11 5.98z"/></svg>'},{}],i0Vut:[function(e,t,i){t.exports='<svg class="apd-icon apd-mode-0-off" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" data-pointer="none" style="enable-background:new 0 0 28 28" viewBox="0 0 28 28" width="28" height="28"><path d="M23 15c1.487 0 2.866.464 4 1.255V7a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h11.674A7 7 0 0 1 23 15zM11 9h6a1 1 0 0 1 0 2h-6a1 1 0 0 1 0-2zm-3 2H6V9h2v2zm4 4h-2v-2h2v2zm2-1a1 1 0 0 1 1-1h1a1 1 0 0 1 0 2h-1a1 1 0 0 1-1-1z" fill="#00AEEC"/><path d="M26.536 18.464a5 5 0 0 0-7.071 0 5 5 0 0 0 0 7.071 5 5 0 1 0 7.071-7.071zm-5.657 5.657a3 3 0 0 1-.586-3.415l4.001 4.001a3 3 0 0 1-3.415-.586zm4.829-.827-4.001-4.001a3.002 3.002 0 0 1 4.001 4.001z" fill="#00AEEC"/></svg>'},{}],hOSvZ:[function(e,t,i){t.exports='<svg class="apd-icon apd-mode-0-on" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" data-pointer="none" style="enable-background:new 0 0 28 28" viewBox="0 0 28 28" width="28" height="28"><path d="M23 3H5a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h18a4 4 0 0 0 4-4V7a4 4 0 0 0-4-4zM11 9h6a1 1 0 0 1 0 2h-6a1 1 0 0 1 0-2zm-3 2H6V9h2v2zm4 4h-2v-2h2v2zm9 0h-6a1 1 0 0 1 0-2h6a1 1 0 0 1 0 2z" fill="#FFF"/></svg>'},{}],bOXC3:[function(e,t,i){t.exports='<svg class="apd-icon apd-mode-1-off" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" data-pointer="none" style="enable-background:new 0 0 28 28" viewBox="0 0 28 28" width="28" height="28"><path d="M23 15c1.487 0 2.866.464 4 1.255V7a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h11.674A7 7 0 0 1 23 15zm-4-8h2v2h-2V7zM9 9H7V7h2v2zm4 0h-2V7h2v2zm2-2h2v2h-2V7z" fill="#00AEEC"/><path d="M26.536 18.464a5 5 0 0 0-7.071 0 5 5 0 0 0 0 7.071 5 5 0 1 0 7.071-7.071zm-5.657 5.657a3 3 0 0 1-.586-3.415l4.001 4.001a3 3 0 0 1-3.415-.586zm4.829-.827-4.001-4.001a3.002 3.002 0 0 1 4.001 4.001z" fill="#00AEEC"/></svg>'},{}],lKuh0:[function(e,t,i){t.exports='<svg class="apd-icon apd-mode-1-on" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" data-pointer="none" style="enable-background:new 0 0 28 28" viewBox="0 0 28 28" width="28" height="28"><path d="M23 3H5a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h18a4 4 0 0 0 4-4V7a4 4 0 0 0-4-4zM9 9H7V7h2v2zm4 0h-2V7h2v2zm4 0h-2V7h2v2zm4 0h-2V7h2v2z" fill="#FFF"/></svg>'},{}],eB8W6:[function(e,t,i){t.exports='<svg class="apd-icon apd-mode-2-off" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" data-pointer="none" style="enable-background:new 0 0 28 28" viewBox="0 0 28 28" width="28" height="28"><path d="M23 15c1.487 0 2.866.464 4 1.255V7a4 4 0 0 0-4-4H5a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h11.674A7 7 0 0 1 23 15zM9 21H7v-2h2v2zm4 0h-2v-2h2v2z" fill="#00AEEC"/><path d="M26.536 18.464a5 5 0 0 0-7.071 0 5 5 0 0 0 0 7.071 5 5 0 1 0 7.071-7.071zm-5.657 5.657a3 3 0 0 1-.586-3.415l4.001 4.001a3 3 0 0 1-3.415-.586zm4.829-.827-4.001-4.001a3.002 3.002 0 0 1 4.001 4.001z" fill="#00AEEC"/></svg>'},{}],bpe2E:[function(e,t,i){t.exports='<svg class="apd-icon apd-mode-2-on" xmlns="http://www.w3.org/2000/svg" xml:space="preserve" data-pointer="none" style="enable-background:new 0 0 28 28" viewBox="0 0 28 28" width="28" height="28"><path d="M23 3H5a4 4 0 0 0-4 4v14a4 4 0 0 0 4 4h18a4 4 0 0 0 4-4V7a4 4 0 0 0-4-4zM9 21H7v-2h2v2zm4 0h-2v-2h2v2zm4 0h-2v-2h2v2zm4 0h-2v-2h2v2z" fill="#FFF"/></svg>'},{}],kL9zy:[function(e,t,i){t.exports='<svg class="apd-icon apd-check-on" xmlns="http://www.w3.org/2000/svg" data-pointer="none" viewBox="0 0 32 32" width="32" height="32"><path d="m13 18.25-1.8-1.8c-.6-.6-1.65-.6-2.25 0s-.6 1.5 0 2.25l2.85 2.85c.318.318.762.468 1.2.448.438.02.882-.13 1.2-.448l8.85-8.85c.6-.6.6-1.65 0-2.25s-1.65-.6-2.25 0l-7.8 7.8zM8 4h16c2.21 0 4 1.79 4 4v16c0 2.21-1.79 4-4 4H8c-2.21 0-4-1.79-4-4V8c0-2.21 1.79-4 4-4z" fill="#00AEEC"/></svg>'},{}],"22xpM":[function(e,t,i){t.exports='<svg class="apd-icon apd-check-off" xmlns="http://www.w3.org/2000/svg" data-pointer="none" viewBox="0 0 32 32" width="32" height="32"><path d="M8 6a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h16a2 2 0 0 0 2-2V8a2 2 0 0 0-2-2H8zm0-2h16c2.21 0 4 1.79 4 4v16c0 2.21-1.79 4-4 4H8c-2.21 0-4-1.79-4-4V8c0-2.21 1.79-4 4-4z" fill="#FFF"/></svg>'},{}],"8AxLD":[function(e,t,i){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(i),a.export(i,"default",()=>s);let n={map:(e,t,i,a,n)=>(e-t)*(n-a)/(i-t)+a,range(e,t,i){let a=Math.round(e/i)*i;return Array.from({length:Math.floor((t-e)/i)},(e,t)=>t*i+a)}},o=(e,t)=>{let i=t[0]-e[0],a=t[1]-e[1];return{length:Math.sqrt(Math.pow(i,2)+Math.pow(a,2)),angle:Math.atan2(a,i)}};function s(e,t,i){let{query:a}=e.constructor.utils;e.controls.add({name:"heatmap",position:"top",html:"",style:{position:"absolute",top:"-100px",left:"0px",right:"0px",height:"100px",width:"100%",pointerEvents:"none"},mounted(s){let l=null,r=null;function d(d=[]){if(l=null,r=null,s.innerHTML="",!e.duration||e.option.isLive)return;let p={w:s.offsetWidth,h:s.offsetHeight},u={xMin:0,xMax:p.w,yMin:0,yMax:128,scale:.25,opacity:.2,minHeight:Math.floor(.05*p.h),sampling:Math.floor(p.w/100),smoothing:.2,flattening:.2};"object"==typeof i&&Object.assign(u,i);let h=[];if(Array.isArray(d)&&d.length)h=[...d];else{let i=e.duration/p.w;for(let e=0;e<=p.w;e+=u.sampling){let a=t.danmus.filter(({time:t})=>t>e*i&&t<=(e+u.sampling)*i).length;h.push([e,a])}}if(0===h.length)return;let m=h[h.length-1],c=m[0],f=m[1];c!==p.w&&h.push([p.w,f]);let g=h.map(e=>e[1]),v=(Math.min(...g)+Math.max(...g))/2;for(let e=0;e<h.length;e++){let t=h[e],i=t[1];t[1]=i*(i>v?1+u.scale:1-u.scale)+u.minHeight}let y=(e,t,i,a)=>{let s=o(t||e,i||e),l=n.map(Math.cos(s.angle)*u.flattening,0,1,1,0),r=s.angle*l+(a?Math.PI:0),d=s.length*u.smoothing;return[e[0]+Math.cos(r)*d,e[1]+Math.sin(r)*d]},x=(e,t,i)=>{let a=y(i[t-1],i[t-2],e),n=y(e,i[t-1],i[t+1],!0),o=t===i.length-1?" z":"";return`C ${a[0]},${a[1]} ${n[0]},${n[1]} ${e[0]},${e[1]}${o}`},k=h.map(e=>[n.map(e[0],u.xMin,u.xMax,0,p.w),n.map(e[1],u.yMin,u.yMax,p.h,0)]).reduce((e,t,i,a)=>0===i?`M ${a[a.length-1][0]},${p.h} L ${t[0]},${p.h} L ${t[0]},${t[1]}`:`${e} ${x(t,i,a)}`,"");s.innerHTML=`<svg viewBox="0 0 ${p.w} ${p.h}"><defs><linearGradient id="heatmap-solids" x1="0%" y1="0%" x2="100%" y2="0%"><stop offset="0%" style="stop-color:var(--art-theme);stop-opacity:${u.opacity}" /><stop offset="0%" style="stop-color:var(--art-theme);stop-opacity:${u.opacity}" id="heatmap-start" /><stop offset="0%" style="stop-color:var(--art-progress-color);stop-opacity:1" id="heatmap-stop" /><stop offset="100%" style="stop-color:var(--art-progress-color);stop-opacity:1" /></linearGradient></defs><path fill="url(#heatmap-solids)" d="${k}"></path></svg>`,l=a("#heatmap-start",s),r=a("#heatmap-stop",s),l.setAttribute("offset",`${100*e.played}%`),r.setAttribute("offset",`${100*e.played}%`)}e.on("video:timeupdate",()=>{l&&r&&(l.setAttribute("offset",`${100*e.played}%`),r.setAttribute("offset",`${100*e.played}%`))}),e.on("setBar",(e,t)=>{l&&r&&"played"===e&&(l.setAttribute("offset",`${100*t}%`),r.setAttribute("offset",`${100*t}%`))}),e.on("ready",()=>d()),e.on("resize",()=>d()),e.on("artplayerPluginDanmuku:loaded",()=>d()),e.on("artplayerPluginDanmuku:points",e=>d(e))}})}},{"@parcel/transformer-js/src/esmodule-helpers.js":"9pCYc"}]},["bgm6t"],"bgm6t","parcelRequire4dc0");

/***/ }),

/***/ "./node_modules/artplayer/dist/artplayer.js":
/*!**************************************************!*\
  !*** ./node_modules/artplayer/dist/artplayer.js ***!
  \**************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

/* module decorator */ module = __webpack_require__.nmd(module);

/*!
 * artplayer.js v5.1.6
 * Github: https://github.com/zhw2590582/ArtPlayer
 * (c) 2017-2024 Harvey Zack
 * Released under the MIT License.
 */
!function(e,t,r,a,i){var o="undefined"!=typeof globalThis?globalThis:"undefined"!=typeof self?self:"undefined"!=typeof window?window:"undefined"!=typeof __webpack_require__.g?__webpack_require__.g:{},n="function"==typeof o[a]&&o[a],s=n.cache||{},l= true&&"function"==typeof module.require&&module.require.bind(module);function c(t,r){if(!s[t]){if(!e[t]){var i="function"==typeof o[a]&&o[a];if(!r&&i)return i(t,!0);if(n)return n(t,!0);if(l&&"string"==typeof t)return l(t);var u=Error("Cannot find module '"+t+"'");throw u.code="MODULE_NOT_FOUND",u}d.resolve=function(r){var a=e[t][1][r];return null!=a?a:r},d.cache={};var p=s[t]=new c.Module(t);e[t][0].call(p.exports,d,p,p.exports,this)}return s[t].exports;function d(e){var t=d.resolve(e);return!1===t?{}:c(t)}}c.isParcelRequire=!0,c.Module=function(e){this.id=e,this.bundle=c,this.exports={}},c.modules=e,c.cache=s,c.parent=n,c.register=function(t,r){e[t]=[function(e,t){t.exports=r},{}]},Object.defineProperty(c,"root",{get:function(){return o[a]}}),o[a]=c;for(var u=0;u<t.length;u++)c(t[u]);if(r){var p=c(r); true?module.exports=p:0}}({abjMI:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("bundle-text:./style/index.less"),o=a.interopDefault(i),n=e("option-validator"),s=a.interopDefault(n),l=e("./utils/emitter"),c=a.interopDefault(l),u=e("./utils"),p=e("./scheme"),d=a.interopDefault(p),f=e("./config"),h=a.interopDefault(f),m=e("./template"),g=a.interopDefault(m),v=e("./i18n"),y=a.interopDefault(v),b=e("./player"),x=a.interopDefault(b),w=e("./control"),j=a.interopDefault(w),k=e("./contextmenu"),S=a.interopDefault(k),I=e("./info"),T=a.interopDefault(I),O=e("./subtitle"),E=a.interopDefault(O),M=e("./events"),$=a.interopDefault(M),F=e("./hotkey"),C=a.interopDefault(F),H=e("./layer"),D=a.interopDefault(H),B=e("./loading"),z=a.interopDefault(B),R=e("./notice"),A=a.interopDefault(R),L=e("./mask"),P=a.interopDefault(L),N=e("./icons"),Z=a.interopDefault(N),_=e("./setting"),q=a.interopDefault(_),V=e("./storage"),W=a.interopDefault(V),U=e("./plugins"),Y=a.interopDefault(U);let K=0,X=[];class G extends c.default{constructor(e,t){super(),this.id=++K;let r=u.mergeDeep(G.option,e);if(r.container=e.container,this.option=(0,s.default)(r,d.default),this.isLock=!1,this.isReady=!1,this.isFocus=!1,this.isInput=!1,this.isRotate=!1,this.isDestroy=!1,this.template=new g.default(this),this.events=new $.default(this),this.storage=new W.default(this),this.icons=new Z.default(this),this.i18n=new y.default(this),this.notice=new A.default(this),this.player=new x.default(this),this.layers=new D.default(this),this.controls=new j.default(this),this.contextmenu=new S.default(this),this.subtitle=new E.default(this),this.info=new T.default(this),this.loading=new z.default(this),this.hotkey=new C.default(this),this.mask=new P.default(this),this.setting=new q.default(this),this.plugins=new Y.default(this),"function"==typeof t&&this.on("ready",()=>t.call(this,this)),G.DEBUG){let e=e=>console.log(`[ART.${this.id}] -> ${e}`);e("Version@"+G.version),e("Env@"+G.env),e("Build@"+G.build);for(let t=0;t<h.default.events.length;t++)this.on("video:"+h.default.events[t],t=>e("Event@"+t.type))}X.push(this)}static get instances(){return X}static get version(){return"5.1.6"}static get env(){return"production"}static get build(){return"2024-06-15 19:54:25"}static get config(){return h.default}static get utils(){return u}static get scheme(){return d.default}static get Emitter(){return c.default}static get validator(){return s.default}static get kindOf(){return s.default.kindOf}static get html(){return g.default.html}static get option(){return{id:"",container:"#artplayer",url:"",poster:"",type:"",theme:"#f00",volume:.7,isLive:!1,muted:!1,autoplay:!1,autoSize:!1,autoMini:!1,loop:!1,flip:!1,playbackRate:!1,aspectRatio:!1,screenshot:!1,setting:!1,hotkey:!0,pip:!1,mutex:!0,backdrop:!0,fullscreen:!1,fullscreenWeb:!1,subtitleOffset:!1,miniProgressBar:!1,useSSR:!1,playsInline:!0,lock:!1,fastForward:!1,autoPlayback:!1,autoOrientation:!1,airplay:!1,layers:[],contextmenu:[],controls:[],settings:[],quality:[],highlight:[],plugins:[],thumbnails:{url:"",number:60,column:10,width:0,height:0},subtitle:{url:"",type:"",style:{},name:"",escape:!0,encoding:"utf-8",onVttLoad:e=>e},moreVideoAttr:{controls:!1,preload:u.isSafari?"auto":"metadata"},i18n:{},icons:{},cssVar:{},customType:{},lang:navigator.language.toLowerCase()}}get proxy(){return this.events.proxy}get query(){return this.template.query}get video(){return this.template.$video}destroy(e=!0){this.events.destroy(),this.template.destroy(e),X.splice(X.indexOf(this),1),this.isDestroy=!0,this.emit("destroy")}}r.default=G,G.DEBUG=!1,G.CONTEXTMENU=!0,G.NOTICE_TIME=2e3,G.SETTING_WIDTH=250,G.SETTING_ITEM_WIDTH=200,G.SETTING_ITEM_HEIGHT=35,G.RESIZE_TIME=200,G.SCROLL_TIME=200,G.SCROLL_GAP=50,G.AUTO_PLAYBACK_MAX=10,G.AUTO_PLAYBACK_MIN=5,G.AUTO_PLAYBACK_TIMEOUT=3e3,G.RECONNECT_TIME_MAX=5,G.RECONNECT_SLEEP_TIME=1e3,G.CONTROL_HIDE_TIME=3e3,G.DBCLICK_TIME=300,G.DBCLICK_FULLSCREEN=!0,G.MOBILE_DBCLICK_PLAY=!0,G.MOBILE_CLICK_PLAY=!1,G.AUTO_ORIENTATION_TIME=200,G.INFO_LOOP_TIME=1e3,G.FAST_FORWARD_VALUE=3,G.FAST_FORWARD_TIME=1e3,G.TOUCH_MOVE_RATIO=.5,G.VOLUME_STEP=.1,G.SEEK_STEP=5,G.PLAYBACK_RATE=[.5,.75,1,1.25,1.5,2],G.ASPECT_RATIO=["default","4:3","16:9"],G.FLIP=["normal","horizontal","vertical"],G.FULLSCREEN_WEB_IN_BODY=!1,G.LOG_VERSION=!0,G.USE_RAF=!1,u.isBrowser&&(window.Artplayer=G,u.setStyleText("artplayer-style",o.default),setTimeout(()=>{G.LOG_VERSION&&console.log(`%c ArtPlayer %c ${G.version} %c https://artplayer.org`,"color: #fff; background: #5f5f5f","color: #fff; background: #4bc729","")},100))},{"bundle-text:./style/index.less":"kfOe8","option-validator":"9I0i9","./utils/emitter":"2bGVu","./utils":"h3rH9","./scheme":"AdvwB","./config":"9Xmqu","./template":"2gKYH","./i18n":"1AdeF","./player":"556MW","./control":"14IBq","./contextmenu":"7iUum","./info":"hD2Lg","./subtitle":"lum0D","./events":"1Epl5","./hotkey":"eTow4","./layer":"4fDoD","./loading":"fE0Sp","./notice":"9PuGy","./mask":"2etr0","./icons":"6dYSr","./setting":"bRHiA","./storage":"f2Thp","./plugins":"96ThS","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],kfOe8:[function(e,t,r){t.exports='.art-video-player{--art-theme:red;--art-font-color:#fff;--art-background-color:#000;--art-text-shadow-color:#00000080;--art-transition-duration:.2s;--art-padding:10px;--art-border-radius:3px;--art-progress-height:6px;--art-progress-color:#ffffff40;--art-hover-color:#ffffff40;--art-loaded-color:#ffffff40;--art-state-size:80px;--art-state-opacity:.8;--art-bottom-height:100px;--art-bottom-offset:20px;--art-bottom-gap:5px;--art-highlight-width:8px;--art-highlight-color:#ffffff80;--art-control-height:46px;--art-control-opacity:.75;--art-control-icon-size:36px;--art-control-icon-scale:1.1;--art-volume-height:120px;--art-volume-handle-size:14px;--art-lock-size:36px;--art-indicator-scale:0;--art-indicator-size:16px;--art-fullscreen-web-index:9999;--art-settings-icon-size:24px;--art-settings-max-height:300px;--art-selector-max-height:300px;--art-contextmenus-min-width:250px;--art-subtitle-font-size:20px;--art-subtitle-gap:5px;--art-subtitle-bottom:15px;--art-subtitle-border:#000;--art-widget-background:#000000d9;--art-tip-background:#000000b3;--art-scrollbar-size:4px;--art-scrollbar-background:#ffffff40;--art-scrollbar-background-hover:#ffffff80;--art-mini-progress-height:2px}.art-bg-cover{background-position:50%;background-repeat:no-repeat;background-size:cover}.art-bottom-gradient{background-image:linear-gradient(#0000,#0006,#000);background-position:bottom;background-repeat:repeat-x}.art-backdrop-filter{backdrop-filter:saturate(180%)blur(20px);background-color:#000000bf!important}.art-truncate{text-overflow:ellipsis;white-space:nowrap;overflow:hidden}.art-video-player{zoom:1;text-align:left;user-select:none;box-sizing:border-box;color:var(--art-font-color);background-color:var(--art-background-color);text-shadow:0 0 2px var(--art-text-shadow-color);-webkit-tap-highlight-color:#0000;-ms-touch-action:manipulation;touch-action:manipulation;-ms-high-contrast-adjust:none;direction:ltr;outline:0;width:100%;height:100%;margin:0 auto;padding:0;font-family:PingFang SC,Helvetica Neue,Microsoft YaHei,Roboto,Arial,sans-serif;font-size:14px;line-height:1.3;position:relative}.art-video-player *,.art-video-player :before,.art-video-player :after{box-sizing:border-box}.art-video-player ::-webkit-scrollbar{width:var(--art-scrollbar-size);height:var(--art-scrollbar-size)}.art-video-player ::-webkit-scrollbar-thumb{background-color:var(--art-scrollbar-background)}.art-video-player ::-webkit-scrollbar-thumb:hover{background-color:var(--art-scrollbar-background-hover)}.art-video-player img{vertical-align:top;max-width:100%}.art-video-player svg{fill:var(--art-font-color)}.art-video-player a{color:var(--art-font-color);text-decoration:none}.art-icon{justify-content:center;align-items:center;line-height:1;display:flex}.art-video-player.art-backdrop .art-contextmenus,.art-video-player.art-backdrop .art-info,.art-video-player.art-backdrop .art-settings,.art-video-player.art-backdrop .art-layer-auto-playback,.art-video-player.art-backdrop .art-selector-list,.art-video-player.art-backdrop .art-volume-inner{backdrop-filter:saturate(180%)blur(20px);background-color:#000000bf!important}.art-video{z-index:10;cursor:pointer;width:100%;height:100%;position:absolute;inset:0}.art-poster{z-index:11;pointer-events:none;background-position:50%;background-repeat:no-repeat;background-size:cover;width:100%;height:100%;position:absolute;inset:0}.art-video-player .art-subtitle{z-index:20;text-align:center;pointer-events:none;justify-content:center;align-items:center;gap:var(--art-subtitle-gap);bottom:var(--art-subtitle-bottom);font-size:var(--art-subtitle-font-size);transition:bottom var(--art-transition-duration)ease;text-shadow:var(--art-subtitle-border)1px 0 1px,var(--art-subtitle-border)0 1px 1px,var(--art-subtitle-border)-1px 0 1px,var(--art-subtitle-border)0 -1px 1px,var(--art-subtitle-border)1px 1px 1px,var(--art-subtitle-border)-1px -1px 1px,var(--art-subtitle-border)1px -1px 1px,var(--art-subtitle-border)-1px 1px 1px;flex-direction:column;width:100%;padding:0 5%;display:none;position:absolute}.art-video-player.art-subtitle-show .art-subtitle{display:flex}.art-video-player.art-control-show .art-subtitle{bottom:calc(var(--art-control-height) + var(--art-subtitle-bottom))}.art-danmuku{z-index:30;pointer-events:none;width:100%;height:100%;position:absolute;inset:0;overflow:hidden}.art-video-player .art-layers{z-index:40;pointer-events:none;width:100%;height:100%;display:none;position:absolute;inset:0}.art-video-player .art-layers .art-layer{pointer-events:auto}.art-video-player.art-layer-show .art-layers{display:flex}.art-video-player .art-mask{z-index:50;pointer-events:none;justify-content:center;align-items:center;width:100%;height:100%;display:flex;position:absolute;inset:0}.art-video-player .art-mask .art-state{opacity:0;width:var(--art-state-size);height:var(--art-state-size);transition:all var(--art-transition-duration)ease;justify-content:center;align-items:center;display:flex;transform:scale(2)}.art-video-player.art-mask-show .art-state{cursor:pointer;pointer-events:auto;opacity:var(--art-state-opacity);transform:scale(1)}.art-video-player.art-loading-show .art-state{display:none}.art-video-player .art-loading{z-index:70;pointer-events:none;justify-content:center;align-items:center;width:100%;height:100%;display:none;position:absolute;inset:0}.art-video-player.art-loading-show .art-loading{display:flex}.art-video-player .art-bottom{z-index:60;opacity:0;pointer-events:none;padding:0 var(--art-padding);transition:all var(--art-transition-duration)ease;background-size:100% var(--art-bottom-height);background-image:linear-gradient(#0000,#0006,#000);background-position:bottom;background-repeat:repeat-x;flex-direction:column;justify-content:flex-end;width:100%;height:100%;display:flex;position:absolute;inset:0;overflow:hidden}.art-video-player .art-bottom .art-controls,.art-video-player .art-bottom .art-progress{transform:translateY(var(--art-bottom-offset));transition:transform var(--art-transition-duration)ease}.art-video-player.art-control-show .art-bottom,.art-video-player.art-hover .art-bottom{opacity:1}.art-video-player.art-control-show .art-bottom .art-controls,.art-video-player.art-hover .art-bottom .art-controls,.art-video-player.art-control-show .art-bottom .art-progress,.art-video-player.art-hover .art-bottom .art-progress{transform:translateY(0)}.art-bottom .art-progress{z-index:0;pointer-events:auto;padding-bottom:var(--art-bottom-gap);position:relative}.art-bottom .art-progress .art-control-progress{cursor:pointer;height:var(--art-progress-height);justify-content:center;align-items:center;display:flex;position:relative}.art-bottom .art-progress .art-control-progress .art-control-progress-inner{transition:height var(--art-transition-duration)ease;background-color:var(--art-progress-color);align-items:center;width:100%;height:50%;display:flex;position:relative}.art-bottom .art-progress .art-control-progress .art-control-progress-inner .art-progress-hover{z-index:0;background-color:var(--art-hover-color);width:0%;height:100%;position:absolute;inset:0}.art-bottom .art-progress .art-control-progress .art-control-progress-inner .art-progress-loaded{z-index:10;background-color:var(--art-loaded-color);width:0%;height:100%;position:absolute;inset:0}.art-bottom .art-progress .art-control-progress .art-control-progress-inner .art-progress-played{z-index:20;background-color:var(--art-theme);width:0%;height:100%;position:absolute;inset:0}.art-bottom .art-progress .art-control-progress .art-control-progress-inner .art-progress-highlight{z-index:30;pointer-events:none;width:100%;height:100%;position:absolute;inset:0}.art-bottom .art-progress .art-control-progress .art-control-progress-inner .art-progress-highlight span{z-index:0;pointer-events:auto;transform:translateX(calc(var(--art-highlight-width)/-2));background-color:var(--art-highlight-color);width:100%;height:100%;position:absolute;inset:0 auto 0 0;width:var(--art-highlight-width)!important}.art-bottom .art-progress .art-control-progress .art-control-progress-inner .art-progress-indicator{z-index:40;width:var(--art-indicator-size);height:var(--art-indicator-size);transform:scale(var(--art-indicator-scale));margin-left:calc(var(--art-indicator-size)/-2);transition:transform var(--art-transition-duration)ease;border-radius:50%;justify-content:center;align-items:center;display:flex;position:absolute;left:0}.art-bottom .art-progress .art-control-progress .art-control-progress-inner .art-progress-indicator .art-icon{pointer-events:none;width:100%;height:100%}.art-bottom .art-progress .art-control-progress .art-control-progress-inner .art-progress-indicator:hover{transform:scale(1.2)!important}.art-bottom .art-progress .art-control-progress .art-control-progress-inner .art-progress-indicator:active{transform:scale(1)!important}.art-bottom .art-progress .art-control-progress .art-control-progress-inner .art-progress-tip{z-index:50;border-radius:var(--art-border-radius);white-space:nowrap;background-color:var(--art-tip-background);padding:3px 5px;font-size:12px;line-height:1;display:none;position:absolute;top:-25px;left:0}.art-bottom .art-progress .art-control-progress:hover .art-control-progress-inner{height:100%}.art-bottom .art-progress .art-control-thumbnails{bottom:calc(var(--art-bottom-gap) + 10px);border-radius:var(--art-border-radius);pointer-events:none;background-color:var(--art-widget-background);display:none;position:absolute;left:0;box-shadow:0 1px 3px #0003,0 1px 2px -1px #0003}.art-bottom:hover .art-progress .art-control-progress .art-control-progress-inner .art-progress-indicator{transform:scale(1)}.art-controls{z-index:10;pointer-events:auto;height:var(--art-control-height);justify-content:space-between;align-items:center;display:flex;position:relative}.art-controls .art-controls-left,.art-controls .art-controls-right{height:100%;display:flex}.art-controls .art-controls-center{flex:1;justify-content:center;align-items:center;height:100%;padding:0 10px;display:none}.art-controls .art-controls-right{justify-content:flex-end}.art-controls .art-control{cursor:pointer;white-space:nowrap;opacity:var(--art-control-opacity);min-height:var(--art-control-height);min-width:var(--art-control-height);transition:opacity var(--art-transition-duration)ease;flex-shrink:0;justify-content:center;align-items:center;display:flex}.art-controls .art-control .art-icon{height:var(--art-control-icon-size);width:var(--art-control-icon-size);transform:scale(var(--art-control-icon-scale));transition:transform var(--art-transition-duration)ease}.art-controls .art-control .art-icon:active{transform:scale(calc(var(--art-control-icon-scale)*.8))}.art-controls .art-control:hover{opacity:1}.art-control-volume{position:relative}.art-control-volume .art-volume-panel{text-align:center;cursor:default;opacity:0;pointer-events:none;left:0;right:0;bottom:var(--art-control-height);width:var(--art-control-height);height:var(--art-volume-height);transition:all var(--art-transition-duration)ease;justify-content:center;align-items:center;padding:0 5px;font-size:12px;display:flex;position:absolute;transform:translateY(10px)}.art-control-volume .art-volume-panel .art-volume-inner{border-radius:var(--art-border-radius);background-color:var(--art-widget-background);flex-direction:column;align-items:center;gap:10px;width:100%;height:100%;padding:10px 0 12px;display:flex}.art-control-volume .art-volume-panel .art-volume-inner .art-volume-slider{cursor:pointer;flex:1;justify-content:center;width:100%;display:flex;position:relative}.art-control-volume .art-volume-panel .art-volume-inner .art-volume-slider .art-volume-handle{border-radius:var(--art-border-radius);background-color:#ffffff40;justify-content:center;width:2px;display:flex;position:relative;overflow:hidden}.art-control-volume .art-volume-panel .art-volume-inner .art-volume-slider .art-volume-handle .art-volume-loaded{z-index:0;background-color:var(--art-theme);width:100%;height:100%;position:absolute;inset:0}.art-control-volume .art-volume-panel .art-volume-inner .art-volume-slider .art-volume-indicator{width:var(--art-volume-handle-size);height:var(--art-volume-handle-size);margin-top:calc(var(--art-volume-handle-size)/-2);background-color:var(--art-theme);transition:transform var(--art-transition-duration)ease;border-radius:100%;flex-shrink:0;position:absolute;transform:scale(1)}.art-control-volume .art-volume-panel .art-volume-inner .art-volume-slider:active .art-volume-indicator{transform:scale(.9)}.art-control-volume:hover .art-volume-panel{opacity:1;pointer-events:auto;transform:translateY(0)}.art-video-player .art-notice{z-index:80;padding:var(--art-padding);pointer-events:none;width:100%;height:auto;display:none;position:absolute;inset:0 0 auto}.art-video-player .art-notice .art-notice-inner{border-radius:var(--art-border-radius);background-color:var(--art-tip-background);padding:5px;line-height:1;display:inline-flex}.art-video-player.art-notice-show .art-notice{display:flex}.art-video-player .art-contextmenus{z-index:120;border-radius:var(--art-border-radius);background-color:var(--art-widget-background);min-width:var(--art-contextmenus-min-width);flex-direction:column;padding:5px 0;font-size:12px;display:none;position:absolute}.art-video-player .art-contextmenus .art-contextmenu{cursor:pointer;border-bottom:1px solid #ffffff1a;padding:10px 15px;display:flex}.art-video-player .art-contextmenus .art-contextmenu span{padding:0 8px}.art-video-player .art-contextmenus .art-contextmenu span:hover,.art-video-player .art-contextmenus .art-contextmenu span.art-current{color:var(--art-theme)}.art-video-player .art-contextmenus .art-contextmenu:hover{background-color:#ffffff1a}.art-video-player .art-contextmenus .art-contextmenu:last-child{border-bottom:none}.art-video-player.art-contextmenu-show .art-contextmenus{display:flex}.art-video-player .art-settings{z-index:90;border-radius:var(--art-border-radius);transform-origin:100% 100%;max-height:var(--art-settings-max-height);left:auto;right:var(--art-padding);bottom:var(--art-control-height);transform:scale(var(--art-settings-scale));transition:all var(--art-transition-duration)ease;background-color:var(--art-widget-background);flex-direction:column;display:none;position:absolute;overflow:hidden auto}.art-video-player .art-settings .art-setting-panel{flex-direction:column;display:none}.art-video-player .art-settings .art-setting-panel.art-current{display:flex}.art-video-player .art-settings .art-setting-panel .art-setting-item{cursor:pointer;transition:background-color var(--art-transition-duration)ease;justify-content:space-between;align-items:center;padding:0 5px;display:flex;overflow:hidden}.art-video-player .art-settings .art-setting-panel .art-setting-item:hover{background-color:#ffffff1a}.art-video-player .art-settings .art-setting-panel .art-setting-item.art-current{color:var(--art-theme)}.art-video-player .art-settings .art-setting-panel .art-setting-item .art-icon-check{visibility:hidden;height:15px}.art-video-player .art-settings .art-setting-panel .art-setting-item.art-current .art-icon-check{visibility:visible}.art-video-player .art-settings .art-setting-panel .art-setting-item .art-setting-item-left{justify-content:center;align-items:center;gap:5px;display:flex}.art-video-player .art-settings .art-setting-panel .art-setting-item .art-setting-item-left .art-setting-item-left-icon{height:var(--art-settings-icon-size);width:var(--art-settings-icon-size);justify-content:center;align-items:center;display:flex}.art-video-player .art-settings .art-setting-panel .art-setting-item .art-setting-item-right{justify-content:center;align-items:center;gap:5px;font-size:12px;display:flex}.art-video-player .art-settings .art-setting-panel .art-setting-item .art-setting-item-right .art-setting-item-right-tooltip{white-space:nowrap;color:#ffffff80}.art-video-player .art-settings .art-setting-panel .art-setting-item .art-setting-item-right .art-setting-item-right-icon{justify-content:center;align-items:center;min-width:32px;height:24px;display:flex}.art-video-player .art-settings .art-setting-panel .art-setting-item .art-setting-item-right .art-setting-range{appearance:none;background-color:#fff3;outline:none;width:80px;height:3px}.art-video-player .art-settings .art-setting-panel .art-setting-item-back{border-bottom:1px solid #ffffff1a}.art-video-player.art-setting-show .art-settings{display:flex}.art-video-player .art-info{left:var(--art-padding);top:var(--art-padding);z-index:100;border-radius:var(--art-border-radius);background-color:var(--art-widget-background);padding:10px;font-size:12px;display:none;position:absolute}.art-video-player .art-info .art-info-panel{flex-direction:column;gap:5px;display:flex}.art-video-player .art-info .art-info-panel .art-info-item{align-items:center;gap:5px;display:flex}.art-video-player .art-info .art-info-panel .art-info-item .art-info-title{text-align:right;width:100px}.art-video-player .art-info .art-info-panel .art-info-item .art-info-content{text-overflow:ellipsis;white-space:nowrap;user-select:all;width:250px;overflow:hidden}.art-video-player .art-info .art-info-close{cursor:pointer;position:absolute;top:5px;right:5px}.art-video-player.art-info-show .art-info{display:flex}.art-hide-cursor *{cursor:none!important}.art-video-player[data-aspect-ratio]{overflow:hidden}.art-video-player[data-aspect-ratio] .art-video{object-fit:fill;box-sizing:content-box}.art-fullscreen{--art-progress-height:8px;--art-indicator-size:20px;--art-control-height:60px;--art-control-icon-scale:1.3}.art-fullscreen-web{--art-progress-height:8px;--art-indicator-size:20px;--art-control-height:60px;--art-control-icon-scale:1.3;z-index:var(--art-fullscreen-web-index);width:100%;height:100%;position:fixed;inset:0}.art-mini-popup{z-index:9999;border-radius:var(--art-border-radius);cursor:move;user-select:none;background:#000;width:320px;height:180px;transition:opacity .2s;position:fixed;overflow:hidden;box-shadow:0 0 5px #00000080}.art-mini-popup svg{fill:#fff}.art-mini-popup .art-video{pointer-events:none}.art-mini-popup .art-mini-close{z-index:20;cursor:pointer;opacity:0;transition:opacity .2s;position:absolute;top:10px;right:10px}.art-mini-popup .art-mini-state{z-index:30;pointer-events:none;opacity:0;background-color:#00000040;justify-content:center;align-items:center;width:100%;height:100%;transition:opacity .2s;display:flex;position:absolute;inset:0}.art-mini-popup .art-mini-state .art-icon{opacity:.75;cursor:pointer;pointer-events:auto;transition:transform .2s;transform:scale(3)}.art-mini-popup .art-mini-state .art-icon:active{transform:scale(2.5)}.art-mini-popup.art-mini-droging{opacity:.9}.art-mini-popup:hover .art-mini-close,.art-mini-popup:hover .art-mini-state{opacity:1}.art-video-player[data-flip=horizontal] .art-video{transform:scaleX(-1)}.art-video-player[data-flip=vertical] .art-video{transform:scaleY(-1)}.art-video-player .art-layer-lock{height:var(--art-lock-size);width:var(--art-lock-size);top:50%;left:var(--art-padding);background-color:var(--art-tip-background);border-radius:50%;justify-content:center;align-items:center;display:none;position:absolute;transform:translateY(-50%)}.art-video-player .art-layer-auto-playback{border-radius:var(--art-border-radius);left:var(--art-padding);bottom:calc(var(--art-control-height) + var(--art-bottom-gap) + 10px);background-color:var(--art-widget-background);align-items:center;gap:10px;padding:10px;line-height:1;display:none;position:absolute}.art-video-player .art-layer-auto-playback .art-auto-playback-close{cursor:pointer;justify-content:center;align-items:center;display:flex}.art-video-player .art-layer-auto-playback .art-auto-playback-close svg{fill:var(--art-theme);width:15px;height:15px}.art-video-player .art-layer-auto-playback .art-auto-playback-jump{color:var(--art-theme);cursor:pointer}.art-video-player.art-lock .art-subtitle{bottom:var(--art-subtitle-bottom)!important}.art-video-player.art-mini-progress-bar .art-bottom,.art-video-player.art-lock .art-bottom{opacity:1;background-image:none;padding:0}.art-video-player.art-mini-progress-bar .art-bottom .art-controls,.art-video-player.art-lock .art-bottom .art-controls,.art-video-player.art-mini-progress-bar .art-bottom .art-progress,.art-video-player.art-lock .art-bottom .art-progress{transform:translateY(calc(var(--art-control-height) + var(--art-bottom-gap) + var(--art-progress-height)/4))}.art-video-player.art-mini-progress-bar .art-bottom .art-progress-indicator,.art-video-player.art-lock .art-bottom .art-progress-indicator{display:none!important}.art-video-player.art-control-show .art-layer-lock{display:flex}.art-control-selector{position:relative}.art-control-selector .art-selector-list{text-align:center;border-radius:var(--art-border-radius);opacity:0;pointer-events:none;bottom:var(--art-control-height);max-height:var(--art-selector-max-height);background-color:var(--art-widget-background);transition:all var(--art-transition-duration)ease;flex-direction:column;align-items:center;display:flex;position:absolute;overflow:hidden auto;transform:translateY(10px)}.art-control-selector .art-selector-list .art-selector-item{flex-shrink:0;justify-content:center;align-items:center;width:100%;padding:10px 15px;line-height:1;display:flex}.art-control-selector .art-selector-list .art-selector-item:hover{background-color:#ffffff1a}.art-control-selector .art-selector-list .art-selector-item:hover,.art-control-selector .art-selector-list .art-selector-item.art-current{color:var(--art-theme)}.art-control-selector:hover .art-selector-list{opacity:1;pointer-events:auto;transform:translateY(0)}[class*=hint--]{font-style:normal;display:inline-block;position:relative}[class*=hint--]:before,[class*=hint--]:after{visibility:hidden;opacity:0;z-index:1000000;pointer-events:none;transition:all .3s;position:absolute;transform:translate(0,0)}[class*=hint--]:hover:before,[class*=hint--]:hover:after{visibility:visible;opacity:1;transition-delay:.1s}[class*=hint--]:before{content:"";z-index:1000001;background:0 0;border:6px solid #0000;position:absolute}[class*=hint--]:after{color:#fff;white-space:nowrap;background:#000;padding:8px 10px;font-family:Helvetica Neue,Helvetica,Arial,sans-serif;font-size:12px;line-height:12px}[class*=hint--][aria-label]:after{content:attr(aria-label)}[class*=hint--][data-hint]:after{content:attr(data-hint)}[aria-label=""]:before,[aria-label=""]:after,[data-hint=""]:before,[data-hint=""]:after{display:none!important}.hint--top-left:before,.hint--top-right:before,.hint--top:before{border-top-color:#000}.hint--bottom-left:before,.hint--bottom-right:before,.hint--bottom:before{border-bottom-color:#000}.hint--left:before{border-left-color:#000}.hint--right:before{border-right-color:#000}.hint--top:before{margin-bottom:-11px}.hint--top:before,.hint--top:after{bottom:100%;left:50%}.hint--top:before{left:calc(50% - 6px)}.hint--top:after{transform:translate(-50%)}.hint--top:hover:before{transform:translateY(-8px)}.hint--top:hover:after{transform:translate(-50%)translateY(-8px)}.hint--bottom:before{margin-top:-11px}.hint--bottom:before,.hint--bottom:after{top:100%;left:50%}.hint--bottom:before{left:calc(50% - 6px)}.hint--bottom:after{transform:translate(-50%)}.hint--bottom:hover:before{transform:translateY(8px)}.hint--bottom:hover:after{transform:translate(-50%)translateY(8px)}.hint--right:before{margin-bottom:-6px;margin-left:-11px}.hint--right:after{margin-bottom:-14px}.hint--right:before,.hint--right:after{bottom:50%;left:100%}.hint--right:hover:before,.hint--right:hover:after{transform:translate(8px)}.hint--left:before{margin-bottom:-6px;margin-right:-11px}.hint--left:after{margin-bottom:-14px}.hint--left:before,.hint--left:after{bottom:50%;right:100%}.hint--left:hover:before,.hint--left:hover:after{transform:translate(-8px)}.hint--top-left:before{margin-bottom:-11px}.hint--top-left:before,.hint--top-left:after{bottom:100%;left:50%}.hint--top-left:before{left:calc(50% - 6px)}.hint--top-left:after{margin-left:12px;transform:translate(-100%)}.hint--top-left:hover:before{transform:translateY(-8px)}.hint--top-left:hover:after{transform:translate(-100%)translateY(-8px)}.hint--top-right:before{margin-bottom:-11px}.hint--top-right:before,.hint--top-right:after{bottom:100%;left:50%}.hint--top-right:before{left:calc(50% - 6px)}.hint--top-right:after{margin-left:-12px;transform:translate(0)}.hint--top-right:hover:before,.hint--top-right:hover:after{transform:translateY(-8px)}.hint--bottom-left:before{margin-top:-11px}.hint--bottom-left:before,.hint--bottom-left:after{top:100%;left:50%}.hint--bottom-left:before{left:calc(50% - 6px)}.hint--bottom-left:after{margin-left:12px;transform:translate(-100%)}.hint--bottom-left:hover:before{transform:translateY(8px)}.hint--bottom-left:hover:after{transform:translate(-100%)translateY(8px)}.hint--bottom-right:before{margin-top:-11px}.hint--bottom-right:before,.hint--bottom-right:after{top:100%;left:50%}.hint--bottom-right:before{left:calc(50% - 6px)}.hint--bottom-right:after{margin-left:-12px;transform:translate(0)}.hint--bottom-right:hover:before,.hint--bottom-right:hover:after{transform:translateY(8px)}.hint--small:after,.hint--medium:after,.hint--large:after{white-space:normal;word-wrap:break-word;line-height:1.4em}.hint--small:after{width:80px}.hint--medium:after{width:150px}.hint--large:after{width:300px}[class*=hint--]:after{text-shadow:0 -1px #000;box-shadow:4px 4px 8px #0000004d}.hint--error:after{text-shadow:0 -1px #592726;background-color:#b34e4d}.hint--error.hint--top-left:before,.hint--error.hint--top-right:before,.hint--error.hint--top:before{border-top-color:#b34e4d}.hint--error.hint--bottom-left:before,.hint--error.hint--bottom-right:before,.hint--error.hint--bottom:before{border-bottom-color:#b34e4d}.hint--error.hint--left:before{border-left-color:#b34e4d}.hint--error.hint--right:before{border-right-color:#b34e4d}.hint--warning:after{text-shadow:0 -1px #6c5328;background-color:#c09854}.hint--warning.hint--top-left:before,.hint--warning.hint--top-right:before,.hint--warning.hint--top:before{border-top-color:#c09854}.hint--warning.hint--bottom-left:before,.hint--warning.hint--bottom-right:before,.hint--warning.hint--bottom:before{border-bottom-color:#c09854}.hint--warning.hint--left:before{border-left-color:#c09854}.hint--warning.hint--right:before{border-right-color:#c09854}.hint--info:after{text-shadow:0 -1px #1a3c4d;background-color:#3986ac}.hint--info.hint--top-left:before,.hint--info.hint--top-right:before,.hint--info.hint--top:before{border-top-color:#3986ac}.hint--info.hint--bottom-left:before,.hint--info.hint--bottom-right:before,.hint--info.hint--bottom:before{border-bottom-color:#3986ac}.hint--info.hint--left:before{border-left-color:#3986ac}.hint--info.hint--right:before{border-right-color:#3986ac}.hint--success:after{text-shadow:0 -1px #1a321a;background-color:#458746}.hint--success.hint--top-left:before,.hint--success.hint--top-right:before,.hint--success.hint--top:before{border-top-color:#458746}.hint--success.hint--bottom-left:before,.hint--success.hint--bottom-right:before,.hint--success.hint--bottom:before{border-bottom-color:#458746}.hint--success.hint--left:before{border-left-color:#458746}.hint--success.hint--right:before{border-right-color:#458746}.hint--always:after,.hint--always:before{opacity:1;visibility:visible}.hint--always.hint--top:before{transform:translateY(-8px)}.hint--always.hint--top:after{transform:translate(-50%)translateY(-8px)}.hint--always.hint--top-left:before{transform:translateY(-8px)}.hint--always.hint--top-left:after{transform:translate(-100%)translateY(-8px)}.hint--always.hint--top-right:before,.hint--always.hint--top-right:after{transform:translateY(-8px)}.hint--always.hint--bottom:before{transform:translateY(8px)}.hint--always.hint--bottom:after{transform:translate(-50%)translateY(8px)}.hint--always.hint--bottom-left:before{transform:translateY(8px)}.hint--always.hint--bottom-left:after{transform:translate(-100%)translateY(8px)}.hint--always.hint--bottom-right:before,.hint--always.hint--bottom-right:after{transform:translateY(8px)}.hint--always.hint--left:before,.hint--always.hint--left:after{transform:translate(-8px)}.hint--always.hint--right:before,.hint--always.hint--right:after{transform:translate(8px)}.hint--rounded:after{border-radius:4px}.hint--no-animate:before,.hint--no-animate:after{transition-duration:0s}.hint--bounce:before,.hint--bounce:after{-webkit-transition:opacity .3s,visibility .3s,-webkit-transform .3s cubic-bezier(.71,1.7,.77,1.24);-moz-transition:opacity .3s,visibility .3s,-moz-transform .3s cubic-bezier(.71,1.7,.77,1.24);transition:opacity .3s,visibility .3s,transform .3s cubic-bezier(.71,1.7,.77,1.24)}.hint--no-shadow:before,.hint--no-shadow:after{text-shadow:initial;box-shadow:initial}.hint--no-arrow:before{display:none}.art-video-player.art-mobile{--art-bottom-gap:10px;--art-control-height:38px;--art-control-icon-scale:1;--art-state-size:60px;--art-settings-max-height:180px;--art-selector-max-height:180px;--art-indicator-scale:1;--art-control-opacity:1}.art-video-player.art-mobile .art-controls-left{margin-left:calc(var(--art-padding)/-1)}.art-video-player.art-mobile .art-controls-right{margin-right:calc(var(--art-padding)/-1)}'},{}],"9I0i9":[function(e,t,r){var a;a=function(){function e(t){return(e="function"==typeof Symbol&&"symbol"==typeof Symbol.iterator?function(e){return typeof e}:function(e){return e&&"function"==typeof Symbol&&e.constructor===Symbol&&e!==Symbol.prototype?"symbol":typeof e})(t)}var t=Object.prototype.toString,r=function(r){if(void 0===r)return"undefined";if(null===r)return"null";var i=e(r);if("boolean"===i)return"boolean";if("string"===i)return"string";if("number"===i)return"number";if("symbol"===i)return"symbol";if("function"===i)return"GeneratorFunction"===a(r)?"generatorfunction":"function";if(Array.isArray?Array.isArray(r):r instanceof Array)return"array";if(r.constructor&&"function"==typeof r.constructor.isBuffer&&r.constructor.isBuffer(r))return"buffer";if(function(e){try{if("number"==typeof e.length&&"function"==typeof e.callee)return!0}catch(e){if(-1!==e.message.indexOf("callee"))return!0}return!1}(r))return"arguments";if(r instanceof Date||"function"==typeof r.toDateString&&"function"==typeof r.getDate&&"function"==typeof r.setDate)return"date";if(r instanceof Error||"string"==typeof r.message&&r.constructor&&"number"==typeof r.constructor.stackTraceLimit)return"error";if(r instanceof RegExp||"string"==typeof r.flags&&"boolean"==typeof r.ignoreCase&&"boolean"==typeof r.multiline&&"boolean"==typeof r.global)return"regexp";switch(a(r)){case"Symbol":return"symbol";case"Promise":return"promise";case"WeakMap":return"weakmap";case"WeakSet":return"weakset";case"Map":return"map";case"Set":return"set";case"Int8Array":return"int8array";case"Uint8Array":return"uint8array";case"Uint8ClampedArray":return"uint8clampedarray";case"Int16Array":return"int16array";case"Uint16Array":return"uint16array";case"Int32Array":return"int32array";case"Uint32Array":return"uint32array";case"Float32Array":return"float32array";case"Float64Array":return"float64array"}if("function"==typeof r.throw&&"function"==typeof r.return&&"function"==typeof r.next)return"generator";switch(i=t.call(r)){case"[object Object]":return"object";case"[object Map Iterator]":return"mapiterator";case"[object Set Iterator]":return"setiterator";case"[object String Iterator]":return"stringiterator";case"[object Array Iterator]":return"arrayiterator"}return i.slice(8,-1).toLowerCase().replace(/\s/g,"")};function a(e){return e.constructor?e.constructor.name:null}function i(e,t){var a=2<arguments.length&&void 0!==arguments[2]?arguments[2]:["option"];return o(e,t,a),n(e,t,a),function(e,t,a){var s=r(t),l=r(e);if("object"===s){if("object"!==l)throw Error("[Type Error]: '".concat(a.join("."),"' require 'object' type, but got '").concat(l,"'"));Object.keys(t).forEach(function(r){var s=e[r],l=t[r],c=a.slice();c.push(r),o(s,l,c),n(s,l,c),i(s,l,c)})}if("array"===s){if("array"!==l)throw Error("[Type Error]: '".concat(a.join("."),"' require 'array' type, but got '").concat(l,"'"));e.forEach(function(r,s){var l=e[s],c=t[s]||t[0],u=a.slice();u.push(s),o(l,c,u),n(l,c,u),i(l,c,u)})}}(e,t,a),e}function o(e,t,a){if("string"===r(t)){var i=r(e);if("?"===t[0]&&(t=t.slice(1)+"|undefined"),!(-1<t.indexOf("|")?t.split("|").map(function(e){return e.toLowerCase().trim()}).filter(Boolean).some(function(e){return i===e}):t.toLowerCase().trim()===i))throw Error("[Type Error]: '".concat(a.join("."),"' require '").concat(t,"' type, but got '").concat(i,"'"))}}function n(e,t,a){if("function"===r(t)){var i=t(e,r(e),a);if(!0!==i){var o=r(i);throw"string"===o?Error(i):"error"===o?i:Error("[Validator Error]: The scheme for '".concat(a.join("."),"' validator require return true, but got '").concat(i,"'"))}}}return i.kindOf=r,i},t.exports=a()},{}],"2bGVu":[function(e,t,r){e("@parcel/transformer-js/src/esmodule-helpers.js").defineInteropFlag(r),r.default=class{on(e,t,r){let a=this.e||(this.e={});return(a[e]||(a[e]=[])).push({fn:t,ctx:r}),this}once(e,t,r){let a=this;function i(...o){a.off(e,i),t.apply(r,o)}return i._=t,this.on(e,i,r)}emit(e,...t){let r=((this.e||(this.e={}))[e]||[]).slice();for(let e=0;e<r.length;e+=1)r[e].fn.apply(r[e].ctx,t);return this}off(e,t){let r=this.e||(this.e={}),a=r[e],i=[];if(a&&t)for(let e=0,r=a.length;e<r;e+=1)a[e].fn!==t&&a[e].fn._!==t&&i.push(a[e]);return i.length?r[e]=i:delete r[e],this}}},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],guZOB:[function(e,t,r){r.interopDefault=function(e){return e&&e.__esModule?e:{default:e}},r.defineInteropFlag=function(e){Object.defineProperty(e,"__esModule",{value:!0})},r.exportAll=function(e,t){return Object.keys(e).forEach(function(r){"default"===r||"__esModule"===r||Object.prototype.hasOwnProperty.call(t,r)||Object.defineProperty(t,r,{enumerable:!0,get:function(){return e[r]}})}),t},r.export=function(e,t,r){Object.defineProperty(e,t,{enumerable:!0,get:r})}},{}],h3rH9:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("./dom");a.exportAll(i,r);var o=e("./error");a.exportAll(o,r);var n=e("./subtitle");a.exportAll(n,r);var s=e("./file");a.exportAll(s,r);var l=e("./property");a.exportAll(l,r);var c=e("./time");a.exportAll(c,r);var u=e("./format");a.exportAll(u,r);var p=e("./compatibility");a.exportAll(p,r)},{"./dom":"XgAQE","./error":"2nFlF","./subtitle":"yqFoT","./file":"1VRQn","./property":"3weX2","./time":"7kBIx","./format":"13atT","./compatibility":"luXC1","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],XgAQE:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"query",()=>o),a.export(r,"queryAll",()=>n),a.export(r,"addClass",()=>s),a.export(r,"removeClass",()=>l),a.export(r,"hasClass",()=>c),a.export(r,"append",()=>u),a.export(r,"remove",()=>p),a.export(r,"setStyle",()=>d),a.export(r,"setStyles",()=>f),a.export(r,"getStyle",()=>h),a.export(r,"sublings",()=>m),a.export(r,"inverseClass",()=>g),a.export(r,"tooltip",()=>v),a.export(r,"isInViewport",()=>y),a.export(r,"includeFromEvent",()=>b),a.export(r,"replaceElement",()=>x),a.export(r,"createElement",()=>w),a.export(r,"getIcon",()=>j),a.export(r,"setStyleText",()=>k),a.export(r,"supportsFlex",()=>S),a.export(r,"getRect",()=>I);var i=e("./compatibility");function o(e,t=document){return t.querySelector(e)}function n(e,t=document){return Array.from(t.querySelectorAll(e))}function s(e,t){return e.classList.add(t)}function l(e,t){return e.classList.remove(t)}function c(e,t){return e.classList.contains(t)}function u(e,t){return t instanceof Element?e.appendChild(t):e.insertAdjacentHTML("beforeend",String(t)),e.lastElementChild||e.lastChild}function p(e){return e.parentNode.removeChild(e)}function d(e,t,r){return e.style[t]=r,e}function f(e,t){for(let r in t)d(e,r,t[r]);return e}function h(e,t,r=!0){let a=window.getComputedStyle(e,null).getPropertyValue(t);return r?parseFloat(a):a}function m(e){return Array.from(e.parentElement.children).filter(t=>t!==e)}function g(e,t){m(e).forEach(e=>l(e,t)),s(e,t)}function v(e,t,r="top"){i.isMobile||(e.setAttribute("aria-label",t),s(e,"hint--rounded"),s(e,`hint--${r}`))}function y(e,t=0){let r=e.getBoundingClientRect(),a=window.innerHeight||document.documentElement.clientHeight,i=window.innerWidth||document.documentElement.clientWidth,o=r.top-t<=a&&r.top+r.height+t>=0,n=r.left-t<=i+t&&r.left+r.width+t>=0;return o&&n}function b(e,t){return e.composedPath&&e.composedPath().indexOf(t)>-1}function x(e,t){return t.parentNode.replaceChild(e,t),e}function w(e){return document.createElement(e)}function j(e="",t=""){let r=w("i");return s(r,"art-icon"),s(r,`art-icon-${e}`),u(r,t),r}function k(e,t){let r=document.getElementById(e);if(r)r.textContent=t;else{let r=w("style");r.id=e,r.textContent=t,document.head.appendChild(r)}}function S(){let e=document.createElement("div");return e.style.display="flex","flex"===e.style.display}function I(e){return e.getBoundingClientRect()}},{"./compatibility":"luXC1","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],luXC1:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"userAgent",()=>i),a.export(r,"isSafari",()=>o),a.export(r,"isWechat",()=>n),a.export(r,"isIE",()=>s),a.export(r,"isAndroid",()=>l),a.export(r,"isIOS",()=>c),a.export(r,"isIOS13",()=>u),a.export(r,"isMobile",()=>p),a.export(r,"isBrowser",()=>d);let i="undefined"!=typeof navigator?navigator.userAgent:"",o=/^((?!chrome|android).)*safari/i.test(i),n=/MicroMessenger/i.test(i),s=/MSIE|Trident/i.test(i),l=/android/i.test(i),c=/iPad|iPhone|iPod/i.test(i)&&!window.MSStream,u=c||i.includes("Macintosh")&&navigator.maxTouchPoints>=1,p=/Android|webOS|iPhone|iPad|iPod|BlackBerry|IEMobile|Opera Mini/i.test(i)||u,d="undefined"!=typeof window},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"2nFlF":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"ArtPlayerError",()=>i),a.export(r,"errorHandle",()=>o);class i extends Error{constructor(e,t){super(e),"function"==typeof Error.captureStackTrace&&Error.captureStackTrace(this,t||this.constructor),this.name="ArtPlayerError"}}function o(e,t){if(!e)throw new i(t);return e}},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],yqFoT:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");function i(e){return"WEBVTT \r\n\r\n".concat(e.replace(/(\d\d:\d\d:\d\d)[,.](\d+)/g,(e,t,r)=>{let a=r.slice(0,3);return 1===r.length&&(a=r+"00"),2===r.length&&(a=r+"0"),`${t},${a}`}).replace(/\{\\([ibu])\}/g,"</$1>").replace(/\{\\([ibu])1\}/g,"<$1>").replace(/\{([ibu])\}/g,"<$1>").replace(/\{\/([ibu])\}/g,"</$1>").replace(/(\d\d:\d\d:\d\d),(\d\d\d)/g,"$1.$2").replace(/{[\s\S]*?}/g,"").concat("\r\n\r\n"))}function o(e){return URL.createObjectURL(new Blob([e],{type:"text/vtt"}))}function n(e){let t=RegExp("Dialogue:\\s\\d,(\\d+:\\d\\d:\\d\\d.\\d\\d),(\\d+:\\d\\d:\\d\\d.\\d\\d),([^,]*),([^,]*),(?:[^,]*,){4}([\\s\\S]*)$","i");function r(e=""){return e.split(/[:.]/).map((e,t,r)=>{if(t===r.length-1){if(1===e.length)return`.${e}00`;if(2===e.length)return`.${e}0`}else if(1===e.length)return(0===t?"0":":0")+e;return 0===t?e:t===r.length-1?`.${e}`:`:${e}`}).join("")}return"WEBVTT\n\n"+e.split(/\r?\n/).map(e=>{let a=e.match(t);return a?{start:r(a[1].trim()),end:r(a[2].trim()),text:a[5].replace(/{[\s\S]*?}/g,"").replace(/(\\N)/g,"\n").trim().split(/\r?\n/).map(e=>e.trim()).join("\n")}:null}).filter(e=>e).map((e,t)=>e?t+1+"\n"+`${e.start} --> ${e.end}`+"\n"+`${e.text}`:"").filter(e=>e.trim()).join("\n\n")}a.defineInteropFlag(r),a.export(r,"srtToVtt",()=>i),a.export(r,"vttToBlob",()=>o),a.export(r,"assToVtt",()=>n)},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"1VRQn":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");function i(e,t){let r=document.createElement("a");r.style.display="none",r.href=e,r.download=t,document.body.appendChild(r),r.click(),document.body.removeChild(r)}a.defineInteropFlag(r),a.export(r,"getExt",()=>function e(t){return t.includes("?")?e(t.split("?")[0]):t.includes("#")?e(t.split("#")[0]):t.trim().toLowerCase().split(".").pop()}),a.export(r,"download",()=>i)},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"3weX2":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"def",()=>i),a.export(r,"has",()=>n),a.export(r,"get",()=>s),a.export(r,"mergeDeep",()=>function e(...t){let r=e=>e&&"object"==typeof e&&!Array.isArray(e);return t.reduce((t,a)=>(Object.keys(a).forEach(i=>{let o=t[i],n=a[i];Array.isArray(o)&&Array.isArray(n)?t[i]=o.concat(...n):r(o)&&r(n)?t[i]=e(o,n):t[i]=n}),t),{})});let i=Object.defineProperty,{hasOwnProperty:o}=Object.prototype;function n(e,t){return o.call(e,t)}function s(e,t){return Object.getOwnPropertyDescriptor(e,t)}},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"7kBIx":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");function i(e=0){return new Promise(t=>setTimeout(t,e))}function o(e,t){let r;return function(...a){clearTimeout(r),r=setTimeout(()=>(r=null,e.apply(this,a)),t)}}function n(e,t){let r=!1;return function(...a){r||(e.apply(this,a),r=!0,setTimeout(function(){r=!1},t))}}a.defineInteropFlag(r),a.export(r,"sleep",()=>i),a.export(r,"debounce",()=>o),a.export(r,"throttle",()=>n)},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"13atT":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");function i(e,t,r){return Math.max(Math.min(e,Math.max(t,r)),Math.min(t,r))}function o(e){return e.charAt(0).toUpperCase()+e.slice(1)}function n(e){return["string","number"].includes(typeof e)}function s(e){if(!e)return"00:00";let t=Math.floor(e/3600),r=Math.floor((e-3600*t)/60),a=Math.floor(e-3600*t-60*r);return(t>0?[t,r,a]:[r,a]).map(e=>e<10?`0${e}`:String(e)).join(":")}function l(e){return e.replace(/[&<>'"]/g,e=>({"&":"&amp;","<":"&lt;",">":"&gt;","'":"&#39;",'"':"&quot;"})[e]||e)}function c(e){let t={"&amp;":"&","&lt;":"<","&gt;":">","&#39;":"'","&quot;":'"'},r=RegExp(`(${Object.keys(t).join("|")})`,"g");return e.replace(r,e=>t[e]||e)}a.defineInteropFlag(r),a.export(r,"clamp",()=>i),a.export(r,"capitalize",()=>o),a.export(r,"isStringOrNumber",()=>n),a.export(r,"secondToTime",()=>s),a.export(r,"escape",()=>l),a.export(r,"unescape",()=>c)},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],AdvwB:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"ComponentOption",()=>d);var i=e("../utils");let o="array",n="boolean",s="string",l="number",c="object",u="function";function p(e,t,r){return(0,i.errorHandle)(t===s||t===l||e instanceof Element,`${r.join(".")} require '${s}' or 'Element' type`)}let d={html:p,disable:`?${n}`,name:`?${s}`,index:`?${l}`,style:`?${c}`,click:`?${u}`,mounted:`?${u}`,tooltip:`?${s}|${l}`,width:`?${l}`,selector:`?${o}`,onSelect:`?${u}`,switch:`?${n}`,onSwitch:`?${u}`,range:`?${o}`,onRange:`?${u}`,onChange:`?${u}`};r.default={id:s,container:p,url:s,poster:s,type:s,theme:s,lang:s,volume:l,isLive:n,muted:n,autoplay:n,autoSize:n,autoMini:n,loop:n,flip:n,playbackRate:n,aspectRatio:n,screenshot:n,setting:n,hotkey:n,pip:n,mutex:n,backdrop:n,fullscreen:n,fullscreenWeb:n,subtitleOffset:n,miniProgressBar:n,useSSR:n,playsInline:n,lock:n,fastForward:n,autoPlayback:n,autoOrientation:n,airplay:n,plugins:[u],layers:[d],contextmenu:[d],settings:[d],controls:[{...d,position:(e,t,r)=>{let a=["top","left","right"];return(0,i.errorHandle)(a.includes(e),`${r.join(".")} only accept ${a.toString()} as parameters`)}}],quality:[{default:`?${n}`,html:s,url:s}],highlight:[{time:l,text:s}],thumbnails:{url:s,number:l,column:l,width:l,height:l},subtitle:{url:s,name:s,type:s,style:c,escape:n,encoding:s,onVttLoad:u},moreVideoAttr:c,i18n:c,icons:c,cssVar:c,customType:c}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"9Xmqu":[function(e,t,r){e("@parcel/transformer-js/src/esmodule-helpers.js").defineInteropFlag(r),r.default={propertys:["audioTracks","autoplay","buffered","controller","controls","crossOrigin","currentSrc","currentTime","defaultMuted","defaultPlaybackRate","duration","ended","error","loop","mediaGroup","muted","networkState","paused","playbackRate","played","preload","readyState","seekable","seeking","src","startDate","textTracks","videoTracks","volume"],methods:["addTextTrack","canPlayType","load","play","pause"],events:["abort","canplay","canplaythrough","durationchange","emptied","ended","error","loadeddata","loadedmetadata","loadstart","pause","play","playing","progress","ratechange","seeked","seeking","stalled","suspend","timeupdate","volumechange","waiting"],prototypes:["width","height","videoWidth","videoHeight","poster","webkitDecodedFrameCount","webkitDroppedFrameCount","playsInline","webkitSupportsFullscreen","webkitDisplayingFullscreen","onenterpictureinpicture","onleavepictureinpicture","disablePictureInPicture","cancelVideoFrameCallback","requestVideoFrameCallback","getVideoPlaybackQuality","requestPictureInPicture","webkitEnterFullScreen","webkitEnterFullscreen","webkitExitFullScreen","webkitExitFullscreen"]}},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"2gKYH":[function(e,t,r){e("@parcel/transformer-js/src/esmodule-helpers.js").defineInteropFlag(r);var a=e("./utils");class i{constructor(e){this.art=e;let{option:t,constructor:r}=e;t.container instanceof Element?this.$container=t.container:(this.$container=(0,a.query)(t.container),(0,a.errorHandle)(this.$container,`No container element found by ${t.container}`)),(0,a.errorHandle)((0,a.supportsFlex)(),"The current browser does not support flex layout");let i=this.$container.tagName.toLowerCase();(0,a.errorHandle)("div"===i,`Unsupported container element type, only support 'div' but got '${i}'`),(0,a.errorHandle)(r.instances.every(e=>e.template.$container!==this.$container),"Cannot mount multiple instances on the same dom element"),this.query=this.query.bind(this),this.$container.dataset.artId=e.id,this.init()}static get html(){return`<div class="art-video-player art-subtitle-show art-layer-show art-control-show art-mask-show"><video class="art-video"><track default kind="metadata" src=""></track></video><div class="art-poster"></div><div class="art-subtitle"></div><div class="art-danmuku"></div><div class="art-layers"></div><div class="art-mask"><div class="art-state"></div></div><div class="art-bottom"><div class="art-progress"></div><div class="art-controls"><div class="art-controls-left"></div><div class="art-controls-center"></div><div class="art-controls-right"></div></div></div><div class="art-loading"></div><div class="art-notice"><div class="art-notice-inner"></div></div><div class="art-settings"></div><div class="art-info"><div class="art-info-panel"><div class="art-info-item"><div class="art-info-title">Player version:</div><div class="art-info-content">5.1.6</div></div><div class="art-info-item"><div class="art-info-title">Video url:</div><div class="art-info-content" data-video="src"></div></div><div class="art-info-item"><div class="art-info-title">Video volume:</div><div class="art-info-content" data-video="volume"></div></div><div class="art-info-item"><div class="art-info-title">Video time:</div><div class="art-info-content" data-video="currentTime"></div></div><div class="art-info-item"><div class="art-info-title">Video duration:</div><div class="art-info-content" data-video="duration"></div></div><div class="art-info-item"><div class="art-info-title">Video resolution:</div><div class="art-info-content"><span data-video="videoWidth"></span>x<span data-video="videoHeight"></span></div></div></div><div class="art-info-close">[x]</div></div><div class="art-contextmenus"></div></div>`}query(e){return(0,a.query)(e,this.$container)}init(){let{option:e}=this.art;e.useSSR||(this.$container.innerHTML=i.html),this.$player=this.query(".art-video-player"),this.$video=this.query(".art-video"),this.$track=this.query("track"),this.$poster=this.query(".art-poster"),this.$subtitle=this.query(".art-subtitle"),this.$danmuku=this.query(".art-danmuku"),this.$bottom=this.query(".art-bottom"),this.$progress=this.query(".art-progress"),this.$controls=this.query(".art-controls"),this.$controlsLeft=this.query(".art-controls-left"),this.$controlsCenter=this.query(".art-controls-center"),this.$controlsRight=this.query(".art-controls-right"),this.$layer=this.query(".art-layers"),this.$loading=this.query(".art-loading"),this.$notice=this.query(".art-notice"),this.$noticeInner=this.query(".art-notice-inner"),this.$mask=this.query(".art-mask"),this.$state=this.query(".art-state"),this.$setting=this.query(".art-settings"),this.$info=this.query(".art-info"),this.$infoPanel=this.query(".art-info-panel"),this.$infoClose=this.query(".art-info-close"),this.$contextmenu=this.query(".art-contextmenus"),e.backdrop&&(0,a.addClass)(this.$player,"art-backdrop"),a.isMobile&&(0,a.addClass)(this.$player,"art-mobile")}destroy(e){e?this.$container.innerHTML="":(0,a.addClass)(this.$player,"art-destroy")}}r.default=i},{"./utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"1AdeF":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("../utils"),o=e("./zh-cn"),n=a.interopDefault(o);r.default=class{constructor(e){this.art=e,this.languages={"zh-cn":n.default},this.language={},this.update(e.option.i18n)}init(){let e=this.art.option.lang.toLowerCase();this.language=this.languages[e]||{}}get(e){return this.language[e]||e}update(e){this.languages=(0,i.mergeDeep)(this.languages,e),this.init()}}},{"../utils":"h3rH9","./zh-cn":"3ZSKq","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"3ZSKq":[function(e,t,r){e("@parcel/transformer-js/src/esmodule-helpers.js").defineInteropFlag(r);let a={"Video Info":"统计信息",Close:"关闭","Video Load Failed":"加载失败",Volume:"音量",Play:"播放",Pause:"暂停",Rate:"速度",Mute:"静音","Video Flip":"画面翻转",Horizontal:"水平",Vertical:"垂直",Reconnect:"重新连接","Show Setting":"显示设置","Hide Setting":"隐藏设置",Screenshot:"截图","Play Speed":"播放速度","Aspect Ratio":"画面比例",Default:"默认",Normal:"正常",Open:"打开","Switch Video":"切换","Switch Subtitle":"切换字幕",Fullscreen:"全屏","Exit Fullscreen":"退出全屏","Web Fullscreen":"网页全屏","Exit Web Fullscreen":"退出网页全屏","Mini Player":"迷你播放器","PIP Mode":"开启画中画","Exit PIP Mode":"退出画中画","PIP Not Supported":"不支持画中画","Fullscreen Not Supported":"不支持全屏","Subtitle Offset":"字幕偏移","Last Seen":"上次看到","Jump Play":"跳转播放",AirPlay:"隔空播放","AirPlay Not Available":"隔空播放不可用"};r.default=a,"undefined"!=typeof window&&(window["artplayer-i18n-zh-cn"]=a)},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"556MW":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("./urlMix"),o=a.interopDefault(i),n=e("./attrMix"),s=a.interopDefault(n),l=e("./playMix"),c=a.interopDefault(l),u=e("./pauseMix"),p=a.interopDefault(u),d=e("./toggleMix"),f=a.interopDefault(d),h=e("./seekMix"),m=a.interopDefault(h),g=e("./volumeMix"),v=a.interopDefault(g),y=e("./currentTimeMix"),b=a.interopDefault(y),x=e("./durationMix"),w=a.interopDefault(x),j=e("./switchMix"),k=a.interopDefault(j),S=e("./playbackRateMix"),I=a.interopDefault(S),T=e("./aspectRatioMix"),O=a.interopDefault(T),E=e("./screenshotMix"),M=a.interopDefault(E),$=e("./fullscreenMix"),F=a.interopDefault($),C=e("./fullscreenWebMix"),H=a.interopDefault(C),D=e("./pipMix"),B=a.interopDefault(D),z=e("./loadedMix"),R=a.interopDefault(z),A=e("./playedMix"),L=a.interopDefault(A),P=e("./playingMix"),N=a.interopDefault(P),Z=e("./autoSizeMix"),_=a.interopDefault(Z),q=e("./rectMix"),V=a.interopDefault(q),W=e("./flipMix"),U=a.interopDefault(W),Y=e("./miniMix"),K=a.interopDefault(Y),X=e("./posterMix"),G=a.interopDefault(X),J=e("./autoHeightMix"),Q=a.interopDefault(J),ee=e("./cssVarMix"),et=a.interopDefault(ee),er=e("./themeMix"),ea=a.interopDefault(er),ei=e("./typeMix"),eo=a.interopDefault(ei),en=e("./stateMix"),es=a.interopDefault(en),el=e("./subtitleOffsetMix"),ec=a.interopDefault(el),eu=e("./airplayMix"),ep=a.interopDefault(eu),ed=e("./qualityMix"),ef=a.interopDefault(ed),eh=e("./optionInit"),em=a.interopDefault(eh),eg=e("./eventInit"),ev=a.interopDefault(eg);r.default=class{constructor(e){(0,o.default)(e),(0,s.default)(e),(0,c.default)(e),(0,p.default)(e),(0,f.default)(e),(0,m.default)(e),(0,v.default)(e),(0,b.default)(e),(0,w.default)(e),(0,k.default)(e),(0,I.default)(e),(0,O.default)(e),(0,M.default)(e),(0,F.default)(e),(0,H.default)(e),(0,B.default)(e),(0,R.default)(e),(0,L.default)(e),(0,N.default)(e),(0,_.default)(e),(0,V.default)(e),(0,U.default)(e),(0,K.default)(e),(0,G.default)(e),(0,Q.default)(e),(0,et.default)(e),(0,ea.default)(e),(0,eo.default)(e),(0,es.default)(e),(0,ec.default)(e),(0,ep.default)(e),(0,ef.default)(e),(0,ev.default)(e),(0,em.default)(e)}}},{"./urlMix":"2mRAc","./attrMix":"2EA19","./playMix":"fD2Tc","./pauseMix":"c3LGJ","./toggleMix":"fVsAa","./seekMix":"dmROF","./volumeMix":"9jtfB","./currentTimeMix":"7NCDR","./durationMix":"YS7JL","./switchMix":"dzUqN","./playbackRateMix":"5I2mT","./aspectRatioMix":"7m6R8","./screenshotMix":"2dgtR","./fullscreenMix":"fKDW8","./fullscreenWebMix":"lNvYI","./pipMix":"8j7oC","./loadedMix":"dwVOT","./playedMix":"dDeLx","./playingMix":"ceoBp","./autoSizeMix":"lcWXX","./rectMix":"f7y88","./flipMix":"l4qt5","./miniMix":"9ZPBQ","./posterMix":"5K8hA","./autoHeightMix":"3T5ls","./cssVarMix":"6KfHs","./themeMix":"7lcSc","./typeMix":"8JgTw","./stateMix":"cebt1","./subtitleOffsetMix":"hJvIy","./airplayMix":"4Tp0U","./qualityMix":"3wZgN","./optionInit":"iPdgW","./eventInit":"3mj0J","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"2mRAc":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{option:t,template:{$video:r}}=e;(0,i.def)(e,"url",{get:()=>r.src,async set(a){if(a){let o=e.url,n=t.type||(0,i.getExt)(a),s=t.customType[n];n&&s?(await (0,i.sleep)(),e.loading.show=!0,s.call(e,r,a,e)):(URL.revokeObjectURL(o),r.src=a),o!==e.url&&(e.option.url=a,e.isReady&&o&&e.once("video:canplay",()=>{e.emit("restart",a)}))}else await (0,i.sleep)(),e.loading.show=!0}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"2EA19":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{template:{$video:t}}=e;(0,i.def)(e,"attr",{value(e,r){if(void 0===r)return t[e];t[e]=r}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],fD2Tc:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{i18n:t,notice:r,option:a,constructor:{instances:o},template:{$video:n}}=e;(0,i.def)(e,"play",{value:async function(){let i=await n.play();if(r.show=t.get("Play"),e.emit("play"),a.mutex)for(let t=0;t<o.length;t++){let r=o[t];r!==e&&r.pause()}return i}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],c3LGJ:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{template:{$video:t},i18n:r,notice:a}=e;(0,i.def)(e,"pause",{value(){let i=t.pause();return a.show=r.get("Pause"),e.emit("pause"),i}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],fVsAa:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){(0,i.def)(e,"toggle",{value:()=>e.playing?e.pause():e.play()})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],dmROF:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{notice:t}=e;(0,i.def)(e,"seek",{set(r){e.currentTime=r,e.emit("seek",e.currentTime),e.duration&&(t.show=`${(0,i.secondToTime)(e.currentTime)} / ${(0,i.secondToTime)(e.duration)}`)}}),(0,i.def)(e,"forward",{set(t){e.seek=e.currentTime+t}}),(0,i.def)(e,"backward",{set(t){e.seek=e.currentTime-t}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"9jtfB":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{template:{$video:t},i18n:r,notice:a,storage:o}=e;(0,i.def)(e,"volume",{get:()=>t.volume||0,set:e=>{t.volume=(0,i.clamp)(e,0,1),a.show=`${r.get("Volume")}: ${parseInt(100*t.volume,10)}`,0!==t.volume&&o.set("volume",t.volume)}}),(0,i.def)(e,"muted",{get:()=>t.muted,set:r=>{t.muted=r,e.emit("muted",r)}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"7NCDR":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{$video:t}=e.template;(0,i.def)(e,"currentTime",{get:()=>t.currentTime||0,set:r=>{Number.isNaN(r=parseFloat(r))||(t.currentTime=(0,i.clamp)(r,0,e.duration))}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],YS7JL:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){(0,i.def)(e,"duration",{get:()=>{let{duration:t}=e.template.$video;return t===1/0?0:t||0}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],dzUqN:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){function t(t,r){return new Promise((a,i)=>{if(t===e.url)return;let{playing:o,aspectRatio:n,playbackRate:s}=e;e.pause(),e.url=t,e.notice.show="",e.once("video:error",i),e.once("video:loadedmetadata",()=>{e.currentTime=r}),e.once("video:canplay",async()=>{e.playbackRate=s,e.aspectRatio=n,o&&await e.play(),e.notice.show="",a()})})}(0,i.def)(e,"switchQuality",{value:r=>t(r,e.currentTime)}),(0,i.def)(e,"switchUrl",{value:e=>t(e,0)}),(0,i.def)(e,"switch",{set:e.switchUrl})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"5I2mT":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{template:{$video:t},i18n:r,notice:a}=e;(0,i.def)(e,"playbackRate",{get:()=>t.playbackRate,set(i){i?i!==t.playbackRate&&(t.playbackRate=i,a.show=`${r.get("Rate")}: ${1===i?r.get("Normal"):`${i}x`}`):e.playbackRate=1}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"7m6R8":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{i18n:t,notice:r,template:{$video:a,$player:o}}=e;(0,i.def)(e,"aspectRatio",{get:()=>o.dataset.aspectRatio||"default",set(n){if(n||(n="default"),"default"===n)(0,i.setStyle)(a,"width",null),(0,i.setStyle)(a,"height",null),(0,i.setStyle)(a,"margin",null),delete o.dataset.aspectRatio;else{let e=n.split(":").map(Number),{clientWidth:t,clientHeight:r}=o,s=e[0]/e[1];t/r>s?((0,i.setStyle)(a,"width",`${s*r}px`),(0,i.setStyle)(a,"height","100%"),(0,i.setStyle)(a,"margin","0 auto")):((0,i.setStyle)(a,"width","100%"),(0,i.setStyle)(a,"height",`${t/s}px`),(0,i.setStyle)(a,"margin","auto 0")),o.dataset.aspectRatio=n}r.show=`${t.get("Aspect Ratio")}: ${"default"===n?t.get("Default"):n}`,e.emit("aspectRatio",n)}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"2dgtR":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{notice:t,template:{$video:r}}=e,a=(0,i.createElement)("canvas");(0,i.def)(e,"getDataURL",{value:()=>new Promise((e,i)=>{try{a.width=r.videoWidth,a.height=r.videoHeight,a.getContext("2d").drawImage(r,0,0),e(a.toDataURL("image/png"))}catch(e){t.show=e,i(e)}})}),(0,i.def)(e,"getBlobUrl",{value:()=>new Promise((e,i)=>{try{a.width=r.videoWidth,a.height=r.videoHeight,a.getContext("2d").drawImage(r,0,0),a.toBlob(t=>{e(URL.createObjectURL(t))})}catch(e){t.show=e,i(e)}})}),(0,i.def)(e,"screenshot",{value:async()=>{let t=await e.getDataURL();return(0,i.download)(t,`artplayer_${(0,i.secondToTime)(r.currentTime)}.png`),e.emit("screenshot",t),t}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],fKDW8:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>s);var i=e("../libs/screenfull"),o=a.interopDefault(i),n=e("../utils");function s(e){let{i18n:t,notice:r,template:{$video:a,$player:i}}=e,s=e=>{(0,o.default).on("change",()=>{e.emit("fullscreen",o.default.isFullscreen)}),(0,o.default).on("error",t=>{e.emit("fullscreenError",t)}),(0,n.def)(e,"fullscreen",{get:()=>o.default.isFullscreen,async set(t){t?(e.state="fullscreen",await (0,o.default).request(i),(0,n.addClass)(i,"art-fullscreen")):(await (0,o.default).exit(),(0,n.removeClass)(i,"art-fullscreen")),e.emit("resize")}})},l=e=>{e.proxy(document,"webkitfullscreenchange",()=>{e.emit("fullscreen",e.fullscreen),e.emit("resize")}),(0,n.def)(e,"fullscreen",{get:()=>document.fullscreenElement===a,set(t){t?(e.state="fullscreen",a.webkitEnterFullscreen()):a.webkitExitFullscreen()}})};e.once("video:loadedmetadata",()=>{o.default.isEnabled?s(e):a.webkitSupportsFullscreen?l(e):(0,n.def)(e,"fullscreen",{get:()=>!1,set(){r.show=t.get("Fullscreen Not Supported")}}),(0,n.def)(e,"fullscreen",(0,n.get)(e,"fullscreen"))})}},{"../libs/screenfull":"lUahW","../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],lUahW:[function(e,t,r){e("@parcel/transformer-js/src/esmodule-helpers.js").defineInteropFlag(r);let a=[["requestFullscreen","exitFullscreen","fullscreenElement","fullscreenEnabled","fullscreenchange","fullscreenerror"],["webkitRequestFullscreen","webkitExitFullscreen","webkitFullscreenElement","webkitFullscreenEnabled","webkitfullscreenchange","webkitfullscreenerror"],["webkitRequestFullScreen","webkitCancelFullScreen","webkitCurrentFullScreenElement","webkitCancelFullScreen","webkitfullscreenchange","webkitfullscreenerror"],["mozRequestFullScreen","mozCancelFullScreen","mozFullScreenElement","mozFullScreenEnabled","mozfullscreenchange","mozfullscreenerror"],["msRequestFullscreen","msExitFullscreen","msFullscreenElement","msFullscreenEnabled","MSFullscreenChange","MSFullscreenError"]],i=(()=>{if("undefined"==typeof document)return!1;let e=a[0],t={};for(let r of a)if(r[1]in document){for(let[a,i]of r.entries())t[e[a]]=i;return t}return!1})(),o={change:i.fullscreenchange,error:i.fullscreenerror},n={request:(e=document.documentElement,t)=>new Promise((r,a)=>{let o=()=>{n.off("change",o),r()};n.on("change",o);let s=e[i.requestFullscreen](t);s instanceof Promise&&s.then(o).catch(a)}),exit:()=>new Promise((e,t)=>{if(!n.isFullscreen){e();return}let r=()=>{n.off("change",r),e()};n.on("change",r);let a=document[i.exitFullscreen]();a instanceof Promise&&a.then(r).catch(t)}),toggle:(e,t)=>n.isFullscreen?n.exit():n.request(e,t),onchange(e){n.on("change",e)},onerror(e){n.on("error",e)},on(e,t){let r=o[e];r&&document.addEventListener(r,t,!1)},off(e,t){let r=o[e];r&&document.removeEventListener(r,t,!1)},raw:i};Object.defineProperties(n,{isFullscreen:{get:()=>!!document[i.fullscreenElement]},element:{enumerable:!0,get:()=>document[i.fullscreenElement]},isEnabled:{enumerable:!0,get:()=>!!document[i.fullscreenEnabled]}}),i||(n={isEnabled:!1}),r.default=n},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],lNvYI:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{constructor:t,template:{$container:r,$player:a}}=e,o="";(0,i.def)(e,"fullscreenWeb",{get:()=>(0,i.hasClass)(a,"art-fullscreen-web"),set(n){n?(o=a.style.cssText,t.FULLSCREEN_WEB_IN_BODY&&(0,i.append)(document.body,a),e.state="fullscreenWeb",(0,i.setStyle)(a,"width","100%"),(0,i.setStyle)(a,"height","100%"),(0,i.addClass)(a,"art-fullscreen-web"),e.emit("fullscreenWeb",!0)):(t.FULLSCREEN_WEB_IN_BODY&&(0,i.append)(r,a),o&&(a.style.cssText=o,o=""),(0,i.removeClass)(a,"art-fullscreen-web"),e.emit("fullscreenWeb",!1)),e.emit("resize")}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"8j7oC":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{i18n:t,notice:r,template:{$video:a}}=e;document.pictureInPictureEnabled?function(e){let{template:{$video:t},proxy:r,notice:a}=e;t.disablePictureInPicture=!1,(0,i.def)(e,"pip",{get:()=>document.pictureInPictureElement,set(r){r?(e.state="pip",t.requestPictureInPicture().catch(e=>{throw a.show=e,e})):document.exitPictureInPicture().catch(e=>{throw a.show=e,e})}}),r(t,"enterpictureinpicture",()=>{e.emit("pip",!0)}),r(t,"leavepictureinpicture",()=>{e.emit("pip",!1)})}(e):a.webkitSupportsPresentationMode?function(e){let{$video:t}=e.template;t.webkitSetPresentationMode("inline"),(0,i.def)(e,"pip",{get:()=>"picture-in-picture"===t.webkitPresentationMode,set(r){r?(e.state="pip",t.webkitSetPresentationMode("picture-in-picture"),e.emit("pip",!0)):(t.webkitSetPresentationMode("inline"),e.emit("pip",!1))}})}(e):(0,i.def)(e,"pip",{get:()=>!1,set(){r.show=t.get("PIP Not Supported")}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],dwVOT:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{$video:t}=e.template;(0,i.def)(e,"loaded",{get:()=>e.loadedTime/t.duration}),(0,i.def)(e,"loadedTime",{get:()=>t.buffered.length?t.buffered.end(t.buffered.length-1):0})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],dDeLx:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){(0,i.def)(e,"played",{get:()=>e.currentTime/e.duration})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],ceoBp:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{$video:t}=e.template;(0,i.def)(e,"playing",{get:()=>!!(t.currentTime>0&&!t.paused&&!t.ended&&t.readyState>2)})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],lcWXX:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{$container:t,$player:r,$video:a}=e.template;(0,i.def)(e,"autoSize",{value(){let{videoWidth:o,videoHeight:n}=a,{width:s,height:l}=(0,i.getRect)(t),c=o/n;s/l>c?((0,i.setStyle)(r,"width",`${l*c/s*100}%`),(0,i.setStyle)(r,"height","100%")):((0,i.setStyle)(r,"width","100%"),(0,i.setStyle)(r,"height",`${s/c/l*100}%`)),e.emit("autoSize",{width:e.width,height:e.height})}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],f7y88:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){(0,i.def)(e,"rect",{get:()=>(0,i.getRect)(e.template.$player)});let t=["bottom","height","left","right","top","width"];for(let r=0;r<t.length;r++){let a=t[r];(0,i.def)(e,a,{get:()=>e.rect[a]})}(0,i.def)(e,"x",{get:()=>e.left+window.pageXOffset}),(0,i.def)(e,"y",{get:()=>e.top+window.pageYOffset})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],l4qt5:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{template:{$player:t},i18n:r,notice:a}=e;(0,i.def)(e,"flip",{get:()=>t.dataset.flip||"normal",set(o){o||(o="normal"),"normal"===o?delete t.dataset.flip:t.dataset.flip=o,a.show=`${r.get("Video Flip")}: ${r.get((0,i.capitalize)(o))}`,e.emit("flip",o)}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"9ZPBQ":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{icons:t,proxy:r,storage:a,template:{$player:o,$video:n}}=e,s=!1,l=0,c=0;function u(){let{$mini:t}=e.template;t&&((0,i.removeClass)(o,"art-mini"),(0,i.setStyle)(t,"display","none"),o.prepend(n),e.emit("mini",!1))}function p(t,r){e.playing?((0,i.setStyle)(t,"display","none"),(0,i.setStyle)(r,"display","flex")):((0,i.setStyle)(t,"display","flex"),(0,i.setStyle)(r,"display","none"))}function d(){let{$mini:t}=e.template,r=(0,i.getRect)(t),o=window.innerHeight-r.height-50,n=window.innerWidth-r.width-50;a.set("top",o),a.set("left",n),(0,i.setStyle)(t,"top",`${o}px`),(0,i.setStyle)(t,"left",`${n}px`)}(0,i.def)(e,"mini",{get:()=>(0,i.hasClass)(o,"art-mini"),set(f){if(f){e.state="mini",(0,i.addClass)(o,"art-mini");let f=function(){let{$mini:o}=e.template;if(o)return(0,i.append)(o,n),(0,i.setStyle)(o,"display","flex");{let o=(0,i.createElement)("div");(0,i.addClass)(o,"art-mini-popup"),(0,i.append)(document.body,o),e.template.$mini=o,(0,i.append)(o,n);let d=(0,i.append)(o,'<div class="art-mini-close"></div>');(0,i.append)(d,t.close),r(d,"click",u);let f=(0,i.append)(o,'<div class="art-mini-state"></div>'),h=(0,i.append)(f,t.play),m=(0,i.append)(f,t.pause);return r(h,"click",()=>e.play()),r(m,"click",()=>e.pause()),p(h,m),e.on("video:playing",()=>p(h,m)),e.on("video:pause",()=>p(h,m)),e.on("video:timeupdate",()=>p(h,m)),r(o,"mousedown",e=>{s=0===e.button,l=e.pageX,c=e.pageY}),e.on("document:mousemove",e=>{if(s){(0,i.addClass)(o,"art-mini-droging");let t=e.pageX-l,r=e.pageY-c;(0,i.setStyle)(o,"transform",`translate(${t}px, ${r}px)`)}}),e.on("document:mouseup",()=>{if(s){s=!1,(0,i.removeClass)(o,"art-mini-droging");let e=(0,i.getRect)(o);a.set("left",e.left),a.set("top",e.top),(0,i.setStyle)(o,"left",`${e.left}px`),(0,i.setStyle)(o,"top",`${e.top}px`),(0,i.setStyle)(o,"transform",null)}}),o}}(),h=a.get("top"),m=a.get("left");h&&m?((0,i.setStyle)(f,"top",`${h}px`),(0,i.setStyle)(f,"left",`${m}px`),(0,i.isInViewport)(f)||d()):d(),e.emit("mini",!0)}else u()}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"5K8hA":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{template:{$poster:t}}=e;(0,i.def)(e,"poster",{get:()=>{try{return t.style.backgroundImage.match(/"(.*)"/)[1]}catch(e){return""}},set(e){(0,i.setStyle)(t,"backgroundImage",`url(${e})`)}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"3T5ls":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{template:{$container:t,$video:r}}=e;(0,i.def)(e,"autoHeight",{value(){let{clientWidth:a}=t,{videoHeight:o,videoWidth:n}=r,s=a/n*o;(0,i.setStyle)(t,"height",s+"px"),e.emit("autoHeight",s)}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"6KfHs":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{$player:t}=e.template;(0,i.def)(e,"cssVar",{value:(e,r)=>r?t.style.setProperty(e,r):getComputedStyle(t).getPropertyValue(e)})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"7lcSc":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){(0,i.def)(e,"theme",{get:()=>e.cssVar("--art-theme"),set(t){e.cssVar("--art-theme",t)}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"8JgTw":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){(0,i.def)(e,"type",{get:()=>e.option.type,set(t){e.option.type=t}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],cebt1:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let t=["mini","pip","fullscreen","fullscreenWeb"];(0,i.def)(e,"state",{get:()=>t.find(t=>e[t])||"standard",set(r){for(let a=0;a<t.length;a++){let i=t[a];i!==r&&e[i]&&(e[i]=!1)}}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],hJvIy:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{clamp:t}=e.constructor.utils,{notice:r,template:a,i18n:o}=e,n=0,s=[];e.on("subtitle:switch",()=>{s=[]}),(0,i.def)(e,"subtitleOffset",{get:()=>n,set(i){if(a.$track&&a.$track.track){let l=Array.from(a.$track.track.cues);n=t(i,-5,5);for(let r=0;r<l.length;r++){let a=l[r];s[r]||(s[r]={startTime:a.startTime,endTime:a.endTime}),a.startTime=t(s[r].startTime+n,0,e.duration),a.endTime=t(s[r].endTime+n,0,e.duration)}e.subtitle.update(),r.show=`${o.get("Subtitle Offset")}: ${i}s`,e.emit("subtitleOffset",i)}else e.emit("subtitleOffset",0)}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"4Tp0U":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{i18n:t,notice:r,proxy:a,template:{$video:o}}=e,n=!0;window.WebKitPlaybackTargetAvailabilityEvent&&o.webkitShowPlaybackTargetPicker?a(o,"webkitplaybacktargetavailabilitychanged",e=>{switch(e.availability){case"available":n=!0;break;case"not-available":n=!1}}):n=!1,(0,i.def)(e,"airplay",{value(){n?(o.webkitShowPlaybackTargetPicker(),e.emit("airplay")):r.show=t.get("AirPlay Not Available")}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"3wZgN":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){(0,i.def)(e,"quality",{set(t){let{controls:r,notice:a,i18n:i}=e,o=t.find(e=>e.default)||t[0];r.update({name:"quality",position:"right",index:10,style:{marginRight:"10px"},html:o?o.html:"",selector:t,async onSelect(t){await e.switchQuality(t.url),a.show=`${i.get("Switch Video")}: ${t.html}`}})}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],iPdgW:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{option:t,storage:r,template:{$video:a,$poster:o}}=e;for(let r in t.moreVideoAttr)e.attr(r,t.moreVideoAttr[r]);t.muted&&(e.muted=t.muted),t.volume&&(a.volume=(0,i.clamp)(t.volume,0,1));let n=r.get("volume");for(let r in"number"==typeof n&&(a.volume=(0,i.clamp)(n,0,1)),t.poster&&(0,i.setStyle)(o,"backgroundImage",`url(${t.poster})`),t.autoplay&&(a.autoplay=t.autoplay),t.playsInline&&(a.playsInline=!0,a["webkit-playsinline"]=!0),t.theme&&(t.cssVar["--art-theme"]=t.theme),t.cssVar)e.cssVar(r,t.cssVar[r]);e.url=t.url}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"3mj0J":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>s);var i=e("../config"),o=a.interopDefault(i),n=e("../utils");function s(e){let{i18n:t,notice:r,option:a,constructor:i,proxy:s,template:{$player:l,$video:c,$poster:u}}=e,p=0;for(let t=0;t<o.default.events.length;t++)s(c,o.default.events[t],t=>{e.emit(`video:${t.type}`,t)});e.on("video:canplay",()=>{p=0,e.loading.show=!1}),e.once("video:canplay",()=>{e.loading.show=!1,e.controls.show=!0,e.mask.show=!0,e.isReady=!0,e.emit("ready")}),e.on("video:ended",()=>{a.loop?(e.seek=0,e.play(),e.controls.show=!1,e.mask.show=!1):(e.controls.show=!0,e.mask.show=!0)}),e.on("video:error",async o=>{p<i.RECONNECT_TIME_MAX?(await (0,n.sleep)(i.RECONNECT_SLEEP_TIME),p+=1,e.url=a.url,r.show=`${t.get("Reconnect")}: ${p}`,e.emit("error",o,p)):(e.mask.show=!0,e.loading.show=!1,e.controls.show=!0,(0,n.addClass)(l,"art-error"),await (0,n.sleep)(i.RECONNECT_SLEEP_TIME),r.show=t.get("Video Load Failed"))}),e.on("video:loadedmetadata",()=>{e.emit("resize"),n.isMobile&&(e.loading.show=!1,e.controls.show=!0,e.mask.show=!0)}),e.on("video:loadstart",()=>{e.loading.show=!0,e.mask.show=!1,e.controls.show=!0}),e.on("video:pause",()=>{e.controls.show=!0,e.mask.show=!0}),e.on("video:play",()=>{e.mask.show=!1,(0,n.setStyle)(u,"display","none")}),e.on("video:playing",()=>{e.mask.show=!1}),e.on("video:progress",()=>{e.playing&&(e.loading.show=!1)}),e.on("video:seeked",()=>{e.loading.show=!1,e.mask.show=!0}),e.on("video:seeking",()=>{e.loading.show=!0,e.mask.show=!1}),e.on("video:timeupdate",()=>{e.mask.show=!1}),e.on("video:waiting",()=>{e.loading.show=!0,e.mask.show=!1})}},{"../config":"9Xmqu","../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"14IBq":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("../utils"),o=e("../utils/component"),n=a.interopDefault(o),s=e("./fullscreen"),l=a.interopDefault(s),c=e("./fullscreenWeb"),u=a.interopDefault(c),p=e("./pip"),d=a.interopDefault(p),f=e("./playAndPause"),h=a.interopDefault(f),m=e("./progress"),g=a.interopDefault(m),v=e("./time"),y=a.interopDefault(v),b=e("./volume"),x=a.interopDefault(b),w=e("./setting"),j=a.interopDefault(w),k=e("./thumbnails"),S=a.interopDefault(k),I=e("./screenshot"),T=a.interopDefault(I),O=e("./airplay"),E=a.interopDefault(O);class M extends n.default{constructor(e){super(e),this.isHover=!1,this.name="control",this.timer=Date.now();let{constructor:t}=e,{$player:r,$bottom:a}=this.art.template;e.on("mousemove",()=>{i.isMobile||(this.show=!0)}),e.on("click",()=>{i.isMobile?this.toggle():this.show=!0}),e.on("document:mousemove",e=>{this.isHover=(0,i.includeFromEvent)(e,a)}),e.on("video:timeupdate",()=>{!e.setting.show&&!this.isHover&&!e.isInput&&e.playing&&this.show&&Date.now()-this.timer>=t.CONTROL_HIDE_TIME&&(this.show=!1)}),e.on("control",e=>{e?((0,i.removeClass)(r,"art-hide-cursor"),(0,i.addClass)(r,"art-hover"),this.timer=Date.now()):((0,i.addClass)(r,"art-hide-cursor"),(0,i.removeClass)(r,"art-hover"))}),this.init()}init(){let{option:e}=this.art;e.isLive||this.add((0,g.default)({name:"progress",position:"top",index:10})),e.thumbnails.url&&!e.isLive&&this.add((0,S.default)({name:"thumbnails",position:"top",index:20})),this.add((0,h.default)({name:"playAndPause",position:"left",index:10})),this.add((0,x.default)({name:"volume",position:"left",index:20})),e.isLive||this.add((0,y.default)({name:"time",position:"left",index:30})),e.quality.length&&(0,i.sleep)().then(()=>{this.art.quality=e.quality}),e.screenshot&&!i.isMobile&&this.add((0,T.default)({name:"screenshot",position:"right",index:20})),e.setting&&this.add((0,j.default)({name:"setting",position:"right",index:30})),e.pip&&this.add((0,d.default)({name:"pip",position:"right",index:40})),e.airplay&&window.WebKitPlaybackTargetAvailabilityEvent&&this.add((0,E.default)({name:"airplay",position:"right",index:50})),e.fullscreenWeb&&this.add((0,u.default)({name:"fullscreenWeb",position:"right",index:60})),e.fullscreen&&this.add((0,l.default)({name:"fullscreen",position:"right",index:70}));for(let t=0;t<e.controls.length;t++)this.add(e.controls[t])}add(e){let t="function"==typeof e?e(this.art):e,{$progress:r,$controlsLeft:a,$controlsRight:o}=this.art.template;switch(t.position){case"top":this.$parent=r;break;case"left":this.$parent=a;break;case"right":this.$parent=o;break;default:(0,i.errorHandle)(!1,"Control option.position must one of 'top', 'left', 'right'")}super.add(t)}}r.default=M},{"../utils":"h3rH9","../utils/component":"guki8","./fullscreen":"cxHNK","./fullscreenWeb":"66eEC","./pip":"kCFkA","./playAndPause":"iRhgD","./progress":"aBBSH","./time":"7H0CE","./volume":"lMwFm","./setting":"8BrCu","./thumbnails":"2HiWx","./screenshot":"c1GeG","./airplay":"6GRju","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],guki8:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("./dom"),o=e("./format"),n=e("./error"),s=e("option-validator"),l=a.interopDefault(s),c=e("../scheme");r.default=class{constructor(e){this.id=0,this.art=e,this.cache=new Map,this.add=this.add.bind(this),this.remove=this.remove.bind(this),this.update=this.update.bind(this)}get show(){return(0,i.hasClass)(this.art.template.$player,`art-${this.name}-show`)}set show(e){let{$player:t}=this.art.template,r=`art-${this.name}-show`;e?(0,i.addClass)(t,r):(0,i.removeClass)(t,r),this.art.emit(this.name,e)}toggle(){this.show=!this.show}add(e){let t="function"==typeof e?e(this.art):e;if(t.html=t.html||"",(0,l.default)(t,c.ComponentOption),!this.$parent||!this.name||t.disable)return;let r=t.name||`${this.name}${this.id}`,a=this.cache.get(r);(0,n.errorHandle)(!a,`Can't add an existing [${r}] to the [${this.name}]`),this.id+=1;let o=(0,i.createElement)("div");(0,i.addClass)(o,`art-${this.name}`),(0,i.addClass)(o,`art-${this.name}-${r}`);let s=Array.from(this.$parent.children);o.dataset.index=t.index||this.id;let u=s.find(e=>Number(e.dataset.index)>=Number(o.dataset.index));u?u.insertAdjacentElement("beforebegin",o):(0,i.append)(this.$parent,o),t.html&&(0,i.append)(o,t.html),t.style&&(0,i.setStyles)(o,t.style),t.tooltip&&(0,i.tooltip)(o,t.tooltip);let p=[];if(t.click){let e=this.art.events.proxy(o,"click",e=>{e.preventDefault(),t.click.call(this.art,this,e)});p.push(e)}return t.selector&&["left","right"].includes(t.position)&&this.addSelector(t,o,p),this[r]=o,this.cache.set(r,{$ref:o,events:p,option:t}),t.mounted&&t.mounted.call(this.art,o),o}addSelector(e,t,r){let{hover:a,proxy:n}=this.art.events;(0,i.addClass)(t,"art-control-selector");let s=(0,i.createElement)("div");(0,i.addClass)(s,"art-selector-value"),(0,i.append)(s,e.html),t.innerText="",(0,i.append)(t,s);let l=e.selector.map((e,t)=>`<div class="art-selector-item ${e.default?"art-current":""}" data-index="${t}">${e.html}</div>`).join(""),c=(0,i.createElement)("div");(0,i.addClass)(c,"art-selector-list"),(0,i.append)(c,l),(0,i.append)(t,c);let u=()=>{let e=(0,i.getStyle)(t,"width"),r=(0,i.getStyle)(c,"width");c.style.left=`${e/2-r/2}px`};a(t,u);let p=n(c,"click",async t=>{let r=(t.composedPath()||[]).find(e=>(0,i.hasClass)(e,"art-selector-item"));if(!r)return;(0,i.inverseClass)(r,"art-current");let a=Number(r.dataset.index),n=e.selector[a]||{};if(s.innerText=r.innerText,e.onSelect){let a=await e.onSelect.call(this.art,n,r,t);(0,o.isStringOrNumber)(a)&&(s.innerHTML=a)}u()});r.push(p)}remove(e){let t=this.cache.get(e);(0,n.errorHandle)(t,`Can't find [${e}] from the [${this.name}]`),t.option.beforeUnmount&&t.option.beforeUnmount.call(this.art,t.$ref);for(let e=0;e<t.events.length;e++)this.art.events.remove(t.events[e]);this.cache.delete(e),delete this[e],(0,i.remove)(t.$ref)}update(e){let t=this.cache.get(e.name);return t&&(e=Object.assign(t.option,e),this.remove(e.name)),this.add(e)}}},{"./dom":"XgAQE","./format":"13atT","./error":"2nFlF","option-validator":"9I0i9","../scheme":"AdvwB","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],cxHNK:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){return t=>({...e,tooltip:t.i18n.get("Fullscreen"),mounted:e=>{let{proxy:r,icons:a,i18n:o}=t,n=(0,i.append)(e,a.fullscreenOn),s=(0,i.append)(e,a.fullscreenOff);(0,i.setStyle)(s,"display","none"),r(e,"click",()=>{t.fullscreen=!t.fullscreen}),t.on("fullscreen",t=>{t?((0,i.tooltip)(e,o.get("Exit Fullscreen")),(0,i.setStyle)(n,"display","none"),(0,i.setStyle)(s,"display","inline-flex")):((0,i.tooltip)(e,o.get("Fullscreen")),(0,i.setStyle)(n,"display","inline-flex"),(0,i.setStyle)(s,"display","none"))})}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"66eEC":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){return t=>({...e,tooltip:t.i18n.get("Web Fullscreen"),mounted:e=>{let{proxy:r,icons:a,i18n:o}=t,n=(0,i.append)(e,a.fullscreenWebOn),s=(0,i.append)(e,a.fullscreenWebOff);(0,i.setStyle)(s,"display","none"),r(e,"click",()=>{t.fullscreenWeb=!t.fullscreenWeb}),t.on("fullscreenWeb",t=>{t?((0,i.tooltip)(e,o.get("Exit Web Fullscreen")),(0,i.setStyle)(n,"display","none"),(0,i.setStyle)(s,"display","inline-flex")):((0,i.tooltip)(e,o.get("Web Fullscreen")),(0,i.setStyle)(n,"display","inline-flex"),(0,i.setStyle)(s,"display","none"))})}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],kCFkA:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){return t=>({...e,tooltip:t.i18n.get("PIP Mode"),mounted:e=>{let{proxy:r,icons:a,i18n:o}=t;(0,i.append)(e,a.pip),r(e,"click",()=>{t.pip=!t.pip}),t.on("pip",t=>{(0,i.tooltip)(e,o.get(t?"Exit PIP Mode":"PIP Mode"))})}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],iRhgD:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){return t=>({...e,mounted:e=>{let{proxy:r,icons:a,i18n:o}=t,n=(0,i.append)(e,a.play),s=(0,i.append)(e,a.pause);function l(){(0,i.setStyle)(n,"display","flex"),(0,i.setStyle)(s,"display","none")}function c(){(0,i.setStyle)(n,"display","none"),(0,i.setStyle)(s,"display","flex")}(0,i.tooltip)(n,o.get("Play")),(0,i.tooltip)(s,o.get("Pause")),r(n,"click",()=>{t.play()}),r(s,"click",()=>{t.pause()}),t.playing?c():l(),t.on("video:playing",()=>{c()}),t.on("video:pause",()=>{l()})}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],aBBSH:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"getPosFromEvent",()=>o),a.export(r,"setCurrentTime",()=>n),a.export(r,"default",()=>s);var i=e("../utils");function o(e,t){let{$progress:r}=e.template,{left:a}=(0,i.getRect)(r),o=i.isMobile?t.touches[0].clientX:t.clientX,n=(0,i.clamp)(o-a,0,r.clientWidth),s=n/r.clientWidth*e.duration,l=(0,i.secondToTime)(s),c=(0,i.clamp)(n/r.clientWidth,0,1);return{second:s,time:l,width:n,percentage:c}}function n(e,t){if(e.isRotate){let r=t.touches[0].clientY/e.height,a=r*e.duration;e.emit("setBar","played",r,t),e.seek=a}else{let{second:r,percentage:a}=o(e,t);e.emit("setBar","played",a,t),e.seek=r}}function s(e){return t=>{let{icons:r,option:a,proxy:s}=t;return{...e,html:`<div class="art-control-progress-inner"><div class="art-progress-hover"></div><div class="art-progress-loaded"></div><div class="art-progress-played"></div><div class="art-progress-highlight"></div><div class="art-progress-indicator"></div><div class="art-progress-tip"></div></div>`,mounted:e=>{let l=null,c=!1,u=(0,i.query)(".art-progress-hover",e),p=(0,i.query)(".art-progress-loaded",e),d=(0,i.query)(".art-progress-played",e),f=(0,i.query)(".art-progress-highlight",e),h=(0,i.query)(".art-progress-indicator",e),m=(0,i.query)(".art-progress-tip",e);function g(r,a){let{width:n,time:s}=a||o(t,r);m.innerText=s;let l=m.clientWidth;n<=l/2?(0,i.setStyle)(m,"left",0):n>e.clientWidth-l/2?(0,i.setStyle)(m,"left",`${e.clientWidth-l}px`):(0,i.setStyle)(m,"left",`${n-l/2}px`)}r.indicator?(0,i.append)(h,r.indicator):(0,i.setStyle)(h,"backgroundColor","var(--art-theme)"),t.on("setBar",function(r,a,o){let n="played"===r&&o&&i.isMobile;"loaded"===r&&(0,i.setStyle)(p,"width",`${100*a}%`),"hover"===r&&(0,i.setStyle)(u,"width",`${100*a}%`),"played"===r&&((0,i.setStyle)(d,"width",`${100*a}%`),(0,i.setStyle)(h,"left",`${100*a}%`)),n&&((0,i.setStyle)(m,"display","flex"),g(o,{width:e.clientWidth*a,time:(0,i.secondToTime)(a*t.duration)}),clearTimeout(l),l=setTimeout(()=>{(0,i.setStyle)(m,"display","none")},500))}),t.on("video:loadedmetadata",function(){f.innerText="";for(let e=0;e<a.highlight.length;e++){let r=a.highlight[e],o=(0,i.clamp)(r.time,0,t.duration)/t.duration*100,n=`<span data-text="${r.text}" data-time="${r.time}" style="left: ${o}%"></span>`;(0,i.append)(f,n)}}),t.on("video:progress",()=>{t.emit("setBar","loaded",t.loaded)}),t.constructor.USE_RAF?t.on("raf",()=>{t.emit("setBar","played",t.played)}):t.on("video:timeupdate",()=>{t.emit("setBar","played",t.played)}),t.on("video:ended",()=>{t.emit("setBar","played",1)}),t.emit("setBar","loaded",t.loaded||0),i.isMobile||(s(e,"click",e=>{e.target!==h&&n(t,e)}),s(e,"mousemove",r=>{let{percentage:a}=o(t,r);t.emit("setBar","hover",a,r),(0,i.setStyle)(m,"display","flex"),(0,i.includeFromEvent)(r,f)?function(r){let{width:a}=o(t,r),{text:n}=r.target.dataset;m.innerText=n;let s=m.clientWidth;a<=s/2?(0,i.setStyle)(m,"left",0):a>e.clientWidth-s/2?(0,i.setStyle)(m,"left",`${e.clientWidth-s}px`):(0,i.setStyle)(m,"left",`${a-s/2}px`)}(r):g(r)}),s(e,"mouseleave",e=>{(0,i.setStyle)(m,"display","none"),t.emit("setBar","hover",0,e)}),s(e,"mousedown",e=>{c=0===e.button}),t.on("document:mousemove",e=>{if(c){let{second:r,percentage:a}=o(t,e);t.emit("setBar","played",a,e),t.seek=r}}),t.on("document:mouseup",()=>{c&&(c=!1)}))}}}}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"7H0CE":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){return t=>({...e,style:i.isMobile?{fontSize:"12px",padding:"0 5px"}:{cursor:"auto",padding:"0 10px"},mounted:e=>{function r(){let r=`${(0,i.secondToTime)(t.currentTime)} / ${(0,i.secondToTime)(t.duration)}`;r!==e.innerText&&(e.innerText=r)}r();let a=["video:loadedmetadata","video:timeupdate","video:progress"];for(let e=0;e<a.length;e++)t.on(a[e],r)}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],lMwFm:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){return t=>({...e,mounted:e=>{let{proxy:r,icons:a}=t,o=(0,i.append)(e,a.volume),n=(0,i.append)(e,a.volumeClose),s=(0,i.append)(e,'<div class="art-volume-panel"></div>'),l=(0,i.append)(s,'<div class="art-volume-inner"></div>'),c=(0,i.append)(l,'<div class="art-volume-val"></div>'),u=(0,i.append)(l,'<div class="art-volume-slider"></div>'),p=(0,i.append)(u,'<div class="art-volume-handle"></div>'),d=(0,i.append)(p,'<div class="art-volume-loaded"></div>'),f=(0,i.append)(u,'<div class="art-volume-indicator"></div>');function h(e){let{top:t,height:r}=(0,i.getRect)(u);return 1-(e.clientY-t)/r}function m(){if(t.muted||0===t.volume)(0,i.setStyle)(o,"display","none"),(0,i.setStyle)(n,"display","flex"),(0,i.setStyle)(f,"top","100%"),(0,i.setStyle)(d,"top","100%"),c.innerText=0;else{let e=100*t.volume;(0,i.setStyle)(o,"display","flex"),(0,i.setStyle)(n,"display","none"),(0,i.setStyle)(f,"top",`${100-e}%`),(0,i.setStyle)(d,"top",`${100-e}%`),c.innerText=Math.floor(e)}}if(m(),t.on("video:volumechange",m),r(o,"click",()=>{t.muted=!0}),r(n,"click",()=>{t.muted=!1}),i.isMobile)(0,i.setStyle)(s,"display","none");else{let e=!1;r(u,"mousedown",r=>{e=0===r.button,t.volume=h(r)}),t.on("document:mousemove",r=>{e&&(t.muted=!1,t.volume=h(r))}),t.on("document:mouseup",()=>{e&&(e=!1)})}}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"8BrCu":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){return t=>({...e,tooltip:t.i18n.get("Show Setting"),mounted:e=>{let{proxy:r,icons:a,i18n:o}=t;(0,i.append)(e,a.setting),r(e,"click",()=>{t.setting.toggle(),t.setting.updateStyle()}),t.on("setting",t=>{(0,i.tooltip)(e,o.get(t?"Hide Setting":"Show Setting"))})}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"2HiWx":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){return t=>({...e,mounted:e=>{let{option:r,events:{loadImg:a},template:{$progress:o,$video:n}}=t,s=null,l=null,c=!1,u=!1;t.on("setBar",async(t,p,d)=>{let f="played"===t&&d&&i.isMobile;if("hover"===t||f){if(c||(c=!0,l=await a(r.thumbnails.url),u=!0),!u)return;let t=o.clientWidth*p;(0,i.setStyle)(e,"display","flex"),t>0&&t<o.clientWidth?function(t){let{url:a,number:s,column:c,width:u,height:p}=r.thumbnails,d=u||l.naturalWidth/c,f=p||d/(n.videoWidth/n.videoHeight),h=Math.floor(t/(o.clientWidth/s)),m=Math.ceil(h/c)-1;(0,i.setStyle)(e,"backgroundImage",`url(${a})`),(0,i.setStyle)(e,"height",`${f}px`),(0,i.setStyle)(e,"width",`${d}px`),(0,i.setStyle)(e,"backgroundPosition",`-${(h%c||c-1)*d}px -${m*f}px`),t<=d/2?(0,i.setStyle)(e,"left",0):t>o.clientWidth-d/2?(0,i.setStyle)(e,"left",`${o.clientWidth-d}px`):(0,i.setStyle)(e,"left",`${t-d/2}px`)}(t):i.isMobile||(0,i.setStyle)(e,"display","none"),f&&(clearTimeout(s),s=setTimeout(()=>{(0,i.setStyle)(e,"display","none")},500))}})}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],c1GeG:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){return t=>({...e,tooltip:t.i18n.get("Screenshot"),mounted:e=>{let{proxy:r,icons:a}=t;(0,i.append)(e,a.screenshot),r(e,"click",()=>{t.screenshot()})}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"6GRju":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){return t=>({...e,tooltip:t.i18n.get("AirPlay"),mounted:e=>{let{proxy:r,icons:a}=t;(0,i.append)(e,a.airplay),r(e,"click",()=>t.airplay())}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"7iUum":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("../utils"),o=e("../utils/component"),n=a.interopDefault(o),s=e("./playbackRate"),l=a.interopDefault(s),c=e("./aspectRatio"),u=a.interopDefault(c),p=e("./flip"),d=a.interopDefault(p),f=e("./info"),h=a.interopDefault(f),m=e("./version"),g=a.interopDefault(m),v=e("./close"),y=a.interopDefault(v);class b extends n.default{constructor(e){super(e),this.name="contextmenu",this.$parent=e.template.$contextmenu,i.isMobile||this.init()}init(){let{option:e,proxy:t,template:{$player:r,$contextmenu:a}}=this.art;e.playbackRate&&this.add((0,l.default)({name:"playbackRate",index:10})),e.aspectRatio&&this.add((0,u.default)({name:"aspectRatio",index:20})),e.flip&&this.add((0,d.default)({name:"flip",index:30})),this.add((0,h.default)({name:"info",index:40})),this.add((0,g.default)({name:"version",index:50})),this.add((0,y.default)({name:"close",index:60}));for(let t=0;t<e.contextmenu.length;t++)this.add(e.contextmenu[t]);t(r,"contextmenu",e=>{if(e.preventDefault(),!this.art.constructor.CONTEXTMENU)return;this.show=!0;let t=e.clientX,o=e.clientY,{height:n,width:s,left:l,top:c}=(0,i.getRect)(r),{height:u,width:p}=(0,i.getRect)(a),d=t-l,f=o-c;t+p>l+s&&(d=s-p),o+u>c+n&&(f=n-u),(0,i.setStyles)(a,{top:`${f}px`,left:`${d}px`})}),t(r,"click",e=>{(0,i.includeFromEvent)(e,a)||(this.show=!1)}),this.art.on("blur",()=>{this.show=!1})}}r.default=b},{"../utils":"h3rH9","../utils/component":"guki8","./playbackRate":"f1W36","./aspectRatio":"afxZC","./flip":"9jCuX","./info":"k8wIZ","./version":"bb0TU","./close":"9zTkI","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],f1W36:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){return t=>{let{i18n:r,constructor:{PLAYBACK_RATE:a}}=t,o=a.map(e=>`<span data-value="${e}">${1===e?r.get("Normal"):e.toFixed(1)}</span>`).join("");return{...e,html:`${r.get("Play Speed")}: ${o}`,click:(e,r)=>{let{value:a}=r.target.dataset;a&&(t.playbackRate=Number(a),e.show=!1)},mounted:e=>{let r=(0,i.query)('[data-value="1"]',e);r&&(0,i.inverseClass)(r,"art-current"),t.on("video:ratechange",()=>{let r=(0,i.queryAll)("span",e).find(e=>Number(e.dataset.value)===t.playbackRate);r&&(0,i.inverseClass)(r,"art-current")})}}}}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],afxZC:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){return t=>{let{i18n:r,constructor:{ASPECT_RATIO:a}}=t,o=a.map(e=>`<span data-value="${e}">${"default"===e?r.get("Default"):e}</span>`).join("");return{...e,html:`${r.get("Aspect Ratio")}: ${o}`,click:(e,r)=>{let{value:a}=r.target.dataset;a&&(t.aspectRatio=a,e.show=!1)},mounted:e=>{let r=(0,i.query)('[data-value="default"]',e);r&&(0,i.inverseClass)(r,"art-current"),t.on("aspectRatio",t=>{let r=(0,i.queryAll)("span",e).find(e=>e.dataset.value===t);r&&(0,i.inverseClass)(r,"art-current")})}}}}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"9jCuX":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){return t=>{let{i18n:r,constructor:{FLIP:a}}=t,o=a.map(e=>`<span data-value="${e}">${r.get((0,i.capitalize)(e))}</span>`).join("");return{...e,html:`${r.get("Video Flip")}: ${o}`,click:(e,r)=>{let{value:a}=r.target.dataset;a&&(t.flip=a.toLowerCase(),e.show=!1)},mounted:e=>{let r=(0,i.query)('[data-value="normal"]',e);r&&(0,i.inverseClass)(r,"art-current"),t.on("flip",t=>{let r=(0,i.queryAll)("span",e).find(e=>e.dataset.value===t);r&&(0,i.inverseClass)(r,"art-current")})}}}}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],k8wIZ:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");function i(e){return t=>({...e,html:t.i18n.get("Video Info"),click:e=>{t.info.show=!0,e.show=!1}})}a.defineInteropFlag(r),a.export(r,"default",()=>i)},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],bb0TU:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");function i(e){return{...e,html:'<a href="https://artplayer.org" target="_blank">ArtPlayer 5.1.6</a>'}}a.defineInteropFlag(r),a.export(r,"default",()=>i)},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"9zTkI":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");function i(e){return t=>({...e,html:t.i18n.get("Close"),click:e=>{e.show=!1}})}a.defineInteropFlag(r),a.export(r,"default",()=>i)},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],hD2Lg:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("./utils"),o=e("./utils/component"),n=a.interopDefault(o);class s extends n.default{constructor(e){super(e),this.name="info",i.isMobile||this.init()}init(){let{proxy:e,constructor:t,template:{$infoPanel:r,$infoClose:a,$video:o}}=this.art;e(a,"click",()=>{this.show=!1});let n=null,s=(0,i.queryAll)("[data-video]",r)||[];this.art.on("destroy",()=>clearTimeout(n)),function e(){for(let e=0;e<s.length;e++){let t=s[e],r=o[t.dataset.video],a="number"==typeof r?r.toFixed(2):r;t.innerText!==a&&(t.innerText=a)}n=setTimeout(e,t.INFO_LOOP_TIME)}()}}r.default=s},{"./utils":"h3rH9","./utils/component":"guki8","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],lum0D:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("./utils"),o=e("./utils/component"),n=a.interopDefault(o),s=e("option-validator"),l=a.interopDefault(s),c=e("./scheme"),u=a.interopDefault(c);class p extends n.default{constructor(e){super(e),this.name="subtitle",this.eventDestroy=()=>null,this.init(e.option.subtitle);let t=!1;e.on("video:timeupdate",()=>{if(!this.url)return;let e=this.art.template.$video.webkitDisplayingFullscreen;"boolean"==typeof e&&e!==t&&(t=e,this.createTrack(e?"subtitles":"metadata",this.url))})}get url(){return this.art.template.$track.src}set url(e){this.switch(e)}get textTrack(){return this.art.template.$video.textTracks[0]}get activeCue(){return this.textTrack.activeCues[0]}style(e,t){let{$subtitle:r}=this.art.template;return"object"==typeof e?(0,i.setStyles)(r,e):(0,i.setStyle)(r,e,t)}update(){let{$subtitle:e}=this.art.template;e.innerHTML="",this.activeCue&&(this.art.option.subtitle.escape?e.innerHTML=this.activeCue.text.split(/\r?\n/).map(e=>`<div class="art-subtitle-line">${(0,i.escape)(e)}</div>`).join(""):e.innerHTML=this.activeCue.text,this.art.emit("subtitleUpdate",this.activeCue.text))}async switch(e,t={}){let{i18n:r,notice:a,option:i}=this.art,o={...i.subtitle,...t,url:e},n=await this.init(o);return t.name&&(a.show=`${r.get("Switch Subtitle")}: ${t.name}`),n}createTrack(e,t){let{template:r,proxy:a,option:o}=this.art,{$video:n,$track:s}=r,l=(0,i.createElement)("track");l.default=!0,l.kind=e,l.src=t,l.label=o.subtitle.name||"Artplayer",l.track.mode="hidden",this.eventDestroy(),(0,i.remove)(s),(0,i.append)(n,l),r.$track=l,this.eventDestroy=a(this.textTrack,"cuechange",()=>this.update())}async init(e){let{notice:t,template:{$subtitle:r}}=this.art;if((0,l.default)(e,u.default.subtitle),e.url)return this.style(e.style),fetch(e.url).then(e=>e.arrayBuffer()).then(t=>{let r=new TextDecoder(e.encoding).decode(t);switch(this.art.emit("subtitleLoad",e.url),e.type||(0,i.getExt)(e.url)){case"srt":{let t=(0,i.srtToVtt)(r),a=e.onVttLoad(t);return(0,i.vttToBlob)(a)}case"ass":{let t=(0,i.assToVtt)(r),a=e.onVttLoad(t);return(0,i.vttToBlob)(a)}case"vtt":{let t=e.onVttLoad(r);return(0,i.vttToBlob)(t)}default:return e.url}}).then(e=>(r.innerHTML="",this.url===e||(URL.revokeObjectURL(this.url),this.createTrack("metadata",e),this.art.emit("subtitleSwitch",e)),e)).catch(e=>{throw r.innerHTML="",t.show=e,e})}}r.default=p},{"./utils":"h3rH9","./utils/component":"guki8","option-validator":"9I0i9","./scheme":"AdvwB","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"1Epl5":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("../utils/error"),o=e("./clickInit"),n=a.interopDefault(o),s=e("./hoverInit"),l=a.interopDefault(s),c=e("./moveInit"),u=a.interopDefault(c),p=e("./resizeInit"),d=a.interopDefault(p),f=e("./gestureInit"),h=a.interopDefault(f),m=e("./viewInit"),g=a.interopDefault(m),v=e("./documentInit"),y=a.interopDefault(v),b=e("./updateInit"),x=a.interopDefault(b);r.default=class{constructor(e){this.destroyEvents=[],this.proxy=this.proxy.bind(this),this.hover=this.hover.bind(this),this.loadImg=this.loadImg.bind(this),(0,n.default)(e,this),(0,l.default)(e,this),(0,u.default)(e,this),(0,d.default)(e,this),(0,h.default)(e,this),(0,g.default)(e,this),(0,y.default)(e,this),(0,x.default)(e,this)}proxy(e,t,r,a={}){if(Array.isArray(t))return t.map(t=>this.proxy(e,t,r,a));e.addEventListener(t,r,a);let i=()=>e.removeEventListener(t,r,a);return this.destroyEvents.push(i),i}hover(e,t,r){t&&this.proxy(e,"mouseenter",t),r&&this.proxy(e,"mouseleave",r)}loadImg(e){return new Promise((t,r)=>{let a;if(e instanceof HTMLImageElement)a=e;else{if("string"!=typeof e)return r(new i.ArtPlayerError("Unable to get Image"));(a=new Image).src=e}if(a.complete)return t(a);this.proxy(a,"load",()=>t(a)),this.proxy(a,"error",()=>r(new i.ArtPlayerError(`Failed to load Image: ${a.src}`)))})}remove(e){let t=this.destroyEvents.indexOf(e);t>-1&&(e(),this.destroyEvents.splice(t,1))}destroy(){for(let e=0;e<this.destroyEvents.length;e++)this.destroyEvents[e]()}}},{"../utils/error":"2nFlF","./clickInit":"gzL6e","./hoverInit":"kpTJf","./moveInit":"ef6qz","./resizeInit":"9TXOX","./gestureInit":"dePMU","./viewInit":"hDyWF","./documentInit":"7RjDP","./updateInit":"8SmBT","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],gzL6e:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e,t){let{constructor:r,template:{$player:a,$video:o}}=e;t.proxy(document,["click","contextmenu"],t=>{(0,i.includeFromEvent)(t,a)?(e.isInput="INPUT"===t.target.tagName,e.isFocus=!0,e.emit("focus",t)):(e.isInput=!1,e.isFocus=!1,e.emit("blur",t))});let n=[];t.proxy(o,"click",t=>{let a=Date.now();n.push(a);let{MOBILE_CLICK_PLAY:o,DBCLICK_TIME:s,MOBILE_DBCLICK_PLAY:l,DBCLICK_FULLSCREEN:c}=r,u=n.filter(e=>a-e<=s);switch(u.length){case 1:e.emit("click",t),i.isMobile?!e.isLock&&o&&e.toggle():e.toggle(),n=u;break;case 2:e.emit("dblclick",t),i.isMobile?!e.isLock&&l&&e.toggle():c&&(e.fullscreen=!e.fullscreen),n=[];break;default:n=[]}})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],kpTJf:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e,t){let{$player:r}=e.template;t.hover(r,t=>{(0,i.addClass)(r,"art-hover"),e.emit("hover",!0,t)},t=>{(0,i.removeClass)(r,"art-hover"),e.emit("hover",!1,t)})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],ef6qz:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");function i(e,t){let{$player:r}=e.template;t.proxy(r,"mousemove",t=>{e.emit("mousemove",t)})}a.defineInteropFlag(r),a.export(r,"default",()=>i)},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"9TXOX":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e,t){let{option:r,constructor:a}=e;e.on("resize",()=>{let{aspectRatio:t,notice:a}=e;"standard"===e.state&&r.autoSize&&e.autoSize(),e.aspectRatio=t,a.show=""});let o=(0,i.debounce)(()=>e.emit("resize"),a.RESIZE_TIME);t.proxy(window,["orientationchange","resize"],()=>o()),screen&&screen.orientation&&screen.orientation.onchange&&t.proxy(screen.orientation,"change",()=>o())}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],dePMU:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>n);var i=e("../utils"),o=e("../control/progress");function n(e,t){if(i.isMobile&&!e.option.isLive){let{$video:r,$progress:a}=e.template,n=null,s=!1,l=0,c=0,u=0,p=t=>{if(1===t.touches.length&&!e.isLock){n===a&&(0,o.setCurrentTime)(e,t),s=!0;let{pageX:r,pageY:i}=t.touches[0];l=r,c=i,u=e.currentTime}},d=t=>{if(1===t.touches.length&&s&&e.duration){let{pageX:a,pageY:o}=t.touches[0],s=function(e,t,r,a){var i=t-a,o=r-e,n=0;if(2>Math.abs(o)&&2>Math.abs(i))return n;var s=180*Math.atan2(i,o)/Math.PI;return s>=-45&&s<45?n=4:s>=45&&s<135?n=1:s>=-135&&s<-45?n=2:(s>=135&&s<=180||s>=-180&&s<-135)&&(n=3),n}(l,c,a,o),p=[3,4].includes(s),d=[1,2].includes(s);if(p&&!e.isRotate||d&&e.isRotate){let s=(0,i.clamp)((a-l)/e.width,-1,1),p=(0,i.clamp)((o-c)/e.height,-1,1),d=e.isRotate?p:s,f=n===r?e.constructor.TOUCH_MOVE_RATIO:1,h=(0,i.clamp)(u+e.duration*d*f,0,e.duration);e.seek=h,e.emit("setBar","played",(0,i.clamp)(h/e.duration,0,1),t),e.notice.show=`${(0,i.secondToTime)(h)} / ${(0,i.secondToTime)(e.duration)}`}}};t.proxy(a,"touchstart",e=>{n=a,p(e)}),t.proxy(r,"touchstart",e=>{n=r,p(e)}),t.proxy(r,"touchmove",d),t.proxy(a,"touchmove",d),t.proxy(document,"touchend",()=>{s&&(l=0,c=0,u=0,s=!1,n=null)})}}},{"../utils":"h3rH9","../control/progress":"aBBSH","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],hDyWF:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e,t){let{option:r,constructor:a,template:{$container:o}}=e,n=(0,i.throttle)(()=>{e.emit("view",(0,i.isInViewport)(o,a.SCROLL_GAP))},a.SCROLL_TIME);t.proxy(window,"scroll",()=>n()),e.on("view",t=>{r.autoMini&&(e.mini=!t)})}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"7RjDP":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");function i(e,t){t.proxy(document,"mousemove",t=>{e.emit("document:mousemove",t)}),t.proxy(document,"mouseup",t=>{e.emit("document:mouseup",t)})}a.defineInteropFlag(r),a.export(r,"default",()=>i)},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"8SmBT":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");function i(e){if(e.constructor.USE_RAF){let t=null;!function r(){e.playing&&e.emit("raf"),e.isDestroy||(t=requestAnimationFrame(r))}(),e.on("destroy",()=>{cancelAnimationFrame(t)})}}a.defineInteropFlag(r),a.export(r,"default",()=>i)},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],eTow4:[function(e,t,r){e("@parcel/transformer-js/src/esmodule-helpers.js").defineInteropFlag(r);var a=e("./utils");r.default=class{constructor(e){this.art=e,this.keys={},e.option.hotkey&&!a.isMobile&&this.init()}init(){let{proxy:e,constructor:t}=this.art;this.add(27,()=>{this.art.fullscreenWeb&&(this.art.fullscreenWeb=!1)}),this.add(32,()=>{this.art.toggle()}),this.add(37,()=>{this.art.backward=t.SEEK_STEP}),this.add(38,()=>{this.art.volume+=t.VOLUME_STEP}),this.add(39,()=>{this.art.forward=t.SEEK_STEP}),this.add(40,()=>{this.art.volume-=t.VOLUME_STEP}),e(window,"keydown",e=>{if(this.art.isFocus){let t=document.activeElement.tagName.toUpperCase(),r=document.activeElement.getAttribute("contenteditable");if("INPUT"!==t&&"TEXTAREA"!==t&&""!==r&&"true"!==r&&!e.altKey&&!e.ctrlKey&&!e.metaKey&&!e.shiftKey){let t=this.keys[e.keyCode];if(t){e.preventDefault();for(let r=0;r<t.length;r++)t[r].call(this.art,e);this.art.emit("hotkey",e)}}}})}add(e,t){return this.keys[e]?this.keys[e].push(t):this.keys[e]=[t],this}remove(e,t){if(this.keys[e]){let r=this.keys[e].indexOf(t);-1!==r&&this.keys[e].splice(r,1)}return this}}},{"./utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"4fDoD":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("./utils/component"),o=a.interopDefault(i);class n extends o.default{constructor(e){super(e);let{option:t,template:{$layer:r}}=e;this.name="layer",this.$parent=r;for(let e=0;e<t.layers.length;e++)this.add(t.layers[e])}}r.default=n},{"./utils/component":"guki8","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],fE0Sp:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("./utils"),o=e("./utils/component"),n=a.interopDefault(o);class s extends n.default{constructor(e){super(e),this.name="loading",(0,i.append)(e.template.$loading,e.icons.loading)}}r.default=s},{"./utils":"h3rH9","./utils/component":"guki8","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"9PuGy":[function(e,t,r){e("@parcel/transformer-js/src/esmodule-helpers.js").defineInteropFlag(r);var a=e("./utils");r.default=class{constructor(e){this.art=e,this.timer=null}set show(e){let{constructor:t,template:{$player:r,$noticeInner:i}}=this.art;e?(i.innerText=e instanceof Error?e.message.trim():e,(0,a.addClass)(r,"art-notice-show"),clearTimeout(this.timer),this.timer=setTimeout(()=>{i.innerText="",(0,a.removeClass)(r,"art-notice-show")},t.NOTICE_TIME)):(0,a.removeClass)(r,"art-notice-show")}}},{"./utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"2etr0":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("./utils"),o=e("./utils/component"),n=a.interopDefault(o);class s extends n.default{constructor(e){super(e),this.name="mask";let{template:t,icons:r,events:a}=e,o=(0,i.append)(t.$state,r.state),n=(0,i.append)(t.$state,r.error);(0,i.setStyle)(n,"display","none"),e.on("destroy",()=>{(0,i.setStyle)(o,"display","none"),(0,i.setStyle)(n,"display",null)}),a.proxy(t.$state,"click",()=>e.play())}}r.default=s},{"./utils":"h3rH9","./utils/component":"guki8","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"6dYSr":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("../utils"),o=e("bundle-text:./loading.svg"),n=a.interopDefault(o),s=e("bundle-text:./state.svg"),l=a.interopDefault(s),c=e("bundle-text:./check.svg"),u=a.interopDefault(c),p=e("bundle-text:./play.svg"),d=a.interopDefault(p),f=e("bundle-text:./pause.svg"),h=a.interopDefault(f),m=e("bundle-text:./volume.svg"),g=a.interopDefault(m),v=e("bundle-text:./volume-close.svg"),y=a.interopDefault(v),b=e("bundle-text:./screenshot.svg"),x=a.interopDefault(b),w=e("bundle-text:./setting.svg"),j=a.interopDefault(w),k=e("bundle-text:./arrow-left.svg"),S=a.interopDefault(k),I=e("bundle-text:./arrow-right.svg"),T=a.interopDefault(I),O=e("bundle-text:./playback-rate.svg"),E=a.interopDefault(O),M=e("bundle-text:./aspect-ratio.svg"),$=a.interopDefault(M),F=e("bundle-text:./config.svg"),C=a.interopDefault(F),H=e("bundle-text:./pip.svg"),D=a.interopDefault(H),B=e("bundle-text:./lock.svg"),z=a.interopDefault(B),R=e("bundle-text:./unlock.svg"),A=a.interopDefault(R),L=e("bundle-text:./fullscreen-off.svg"),P=a.interopDefault(L),N=e("bundle-text:./fullscreen-on.svg"),Z=a.interopDefault(N),_=e("bundle-text:./fullscreen-web-off.svg"),q=a.interopDefault(_),V=e("bundle-text:./fullscreen-web-on.svg"),W=a.interopDefault(V),U=e("bundle-text:./switch-on.svg"),Y=a.interopDefault(U),K=e("bundle-text:./switch-off.svg"),X=a.interopDefault(K),G=e("bundle-text:./flip.svg"),J=a.interopDefault(G),Q=e("bundle-text:./error.svg"),ee=a.interopDefault(Q),et=e("bundle-text:./close.svg"),er=a.interopDefault(et),ea=e("bundle-text:./airplay.svg"),ei=a.interopDefault(ea);r.default=class{constructor(e){let t={loading:n.default,state:l.default,play:d.default,pause:h.default,check:u.default,volume:g.default,volumeClose:y.default,screenshot:x.default,setting:j.default,pip:D.default,arrowLeft:S.default,arrowRight:T.default,playbackRate:E.default,aspectRatio:$.default,config:C.default,lock:z.default,flip:J.default,unlock:A.default,fullscreenOff:P.default,fullscreenOn:Z.default,fullscreenWebOff:q.default,fullscreenWebOn:W.default,switchOn:Y.default,switchOff:X.default,error:ee.default,close:er.default,airplay:ei.default,...e.option.icons};for(let e in t)(0,i.def)(this,e,{get:()=>(0,i.getIcon)(e,t[e])})}}},{"../utils":"h3rH9","bundle-text:./loading.svg":"fY5Gt","bundle-text:./state.svg":"iNfLt","bundle-text:./check.svg":"jtE9u","bundle-text:./play.svg":"elgfY","bundle-text:./pause.svg":"eKokJ","bundle-text:./volume.svg":"hNB4y","bundle-text:./volume-close.svg":"i9vta","bundle-text:./screenshot.svg":"kB3Mf","bundle-text:./setting.svg":"3MONs","bundle-text:./arrow-left.svg":"iMCpk","bundle-text:./arrow-right.svg":"3oe4L","bundle-text:./playback-rate.svg":"liE22","bundle-text:./aspect-ratio.svg":"8HqYc","bundle-text:./config.svg":"hYAAH","bundle-text:./pip.svg":"jmNrH","bundle-text:./lock.svg":"cIqko","bundle-text:./unlock.svg":"65zy4","bundle-text:./fullscreen-off.svg":"jaJRT","bundle-text:./fullscreen-on.svg":"cRY1X","bundle-text:./fullscreen-web-off.svg":"3aVGL","bundle-text:./fullscreen-web-on.svg":"4DiVn","bundle-text:./switch-on.svg":"kwdKE","bundle-text:./switch-off.svg":"bWfXZ","bundle-text:./flip.svg":"h3zZ9","bundle-text:./error.svg":"7Oyth","bundle-text:./close.svg":"U5Jcy","bundle-text:./airplay.svg":"jK5Fx","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],fY5Gt:[function(e,t,r){t.exports='<svg xmlns="http://www.w3.org/2000/svg" width="50" height="50" viewBox="0 0 100 100" preserveAspectRatio="xMidYMid" class="uil-default"><path fill="none" class="bk" d="M0 0h100v100H0z"/><rect x="47" y="40" width="6" height="20" rx="5" ry="5" fill="#fff" transform="translate(0 -30)"><animate attributeName="opacity" from="1" to="0" dur="1s" begin="-1s" repeatCount="indefinite"/></rect><rect x="47" y="40" width="6" height="20" rx="5" ry="5" fill="#fff" transform="rotate(30 105.98 65)"><animate attributeName="opacity" from="1" to="0" dur="1s" begin="-0.9166666666666666s" repeatCount="indefinite"/></rect><rect x="47" y="40" width="6" height="20" rx="5" ry="5" fill="#fff" transform="rotate(60 75.98 65)"><animate attributeName="opacity" from="1" to="0" dur="1s" begin="-0.8333333333333334s" repeatCount="indefinite"/></rect><rect x="47" y="40" width="6" height="20" rx="5" ry="5" fill="#fff" transform="rotate(90 65 65)"><animate attributeName="opacity" from="1" to="0" dur="1s" begin="-0.75s" repeatCount="indefinite"/></rect><rect x="47" y="40" width="6" height="20" rx="5" ry="5" fill="#fff" transform="rotate(120 58.66 65)"><animate attributeName="opacity" from="1" to="0" dur="1s" begin="-0.6666666666666666s" repeatCount="indefinite"/></rect><rect x="47" y="40" width="6" height="20" rx="5" ry="5" fill="#fff" transform="rotate(150 54.02 65)"><animate attributeName="opacity" from="1" to="0" dur="1s" begin="-0.5833333333333334s" repeatCount="indefinite"/></rect><rect x="47" y="40" width="6" height="20" rx="5" ry="5" fill="#fff" transform="rotate(180 50 65)"><animate attributeName="opacity" from="1" to="0" dur="1s" begin="-0.5s" repeatCount="indefinite"/></rect><rect x="47" y="40" width="6" height="20" rx="5" ry="5" fill="#fff" transform="rotate(-150 45.98 65)"><animate attributeName="opacity" from="1" to="0" dur="1s" begin="-0.4166666666666667s" repeatCount="indefinite"/></rect><rect x="47" y="40" width="6" height="20" rx="5" ry="5" fill="#fff" transform="rotate(-120 41.34 65)"><animate attributeName="opacity" from="1" to="0" dur="1s" begin="-0.3333333333333333s" repeatCount="indefinite"/></rect><rect x="47" y="40" width="6" height="20" rx="5" ry="5" fill="#fff" transform="rotate(-90 35 65)"><animate attributeName="opacity" from="1" to="0" dur="1s" begin="-0.25s" repeatCount="indefinite"/></rect><rect x="47" y="40" width="6" height="20" rx="5" ry="5" fill="#fff" transform="rotate(-60 24.02 65)"><animate attributeName="opacity" from="1" to="0" dur="1s" begin="-0.16666666666666666s" repeatCount="indefinite"/></rect><rect x="47" y="40" width="6" height="20" rx="5" ry="5" fill="#fff" transform="rotate(-30 -5.98 65)"><animate attributeName="opacity" from="1" to="0" dur="1s" begin="-0.08333333333333333s" repeatCount="indefinite"/></rect></svg>'},{}],iNfLt:[function(e,t,r){t.exports='<svg xmlns="http://www.w3.org/2000/svg" width="80" height="80" viewBox="0 0 24 24"><path fill="#fff" d="M9.5 9.325v5.35q0 .575.525.875t1.025-.05l4.15-2.65q.475-.3.475-.85t-.475-.85L11.05 8.5q-.5-.35-1.025-.05t-.525.875ZM12 22q-2.075 0-3.9-.788t-3.175-2.137q-1.35-1.35-2.137-3.175T2 12q0-2.075.788-3.9t2.137-3.175q1.35-1.35 3.175-2.137T12 2q2.075 0 3.9.788t3.175 2.137q1.35 1.35 2.138 3.175T22 12q0 2.075-.788 3.9t-2.137 3.175q-1.35 1.35-3.175 2.138T12 22Z"/></svg>'},{}],jtE9u:[function(e,t,r){t.exports='<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" style="width:100%;height:100%"><path d="M9 16.2 4.8 12l-1.4 1.4L9 19 21 7l-1.4-1.4L9 16.2z" fill="#fff"/></svg>'},{}],elgfY:[function(e,t,r){t.exports='<svg xmlns="http://www.w3.org/2000/svg" height="22" width="22"><path d="M17.982 9.275 8.06 3.27A2.013 2.013 0 0 0 5 4.994v12.011a2.017 2.017 0 0 0 3.06 1.725l9.922-6.005a2.017 2.017 0 0 0 0-3.45z"/></svg>'},{}],eKokJ:[function(e,t,r){t.exports='<svg xmlns="http://www.w3.org/2000/svg" height="22" width="22"><path d="M7 3a2 2 0 0 0-2 2v12a2 2 0 1 0 4 0V5a2 2 0 0 0-2-2zm8 0a2 2 0 0 0-2 2v12a2 2 0 1 0 4 0V5a2 2 0 0 0-2-2z"/></svg>'},{}],hNB4y:[function(e,t,r){t.exports='<svg xmlns="http://www.w3.org/2000/svg" height="22" width="22"><path d="M10.188 4.65 6 8H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h1l4.188 3.35a.5.5 0 0 0 .812-.39V5.04a.498.498 0 0 0-.812-.39zm4.258-.872a1 1 0 0 0-.862 1.804 6.002 6.002 0 0 1-.007 10.838 1 1 0 0 0 .86 1.806A8.001 8.001 0 0 0 19 11a8.001 8.001 0 0 0-4.554-7.222z"/><path d="M15 11a3.998 3.998 0 0 0-2-3.465v6.93A3.998 3.998 0 0 0 15 11z"/></svg>'},{}],i9vta:[function(e,t,r){t.exports='<svg xmlns="http://www.w3.org/2000/svg" height="22" width="22"><path d="M15 11a3.998 3.998 0 0 0-2-3.465v2.636l1.865 1.865A4.02 4.02 0 0 0 15 11z"/><path d="M13.583 5.583A5.998 5.998 0 0 1 17 11a6 6 0 0 1-.585 2.587l1.477 1.477a8.001 8.001 0 0 0-3.446-11.286 1 1 0 0 0-.863 1.805zm5.195 13.195-2.121-2.121-1.414-1.414-1.415-1.415L13 13l-2-2-3.889-3.889-3.889-3.889a.999.999 0 1 0-1.414 1.414L5.172 8H5a2 2 0 0 0-2 2v2a2 2 0 0 0 2 2h1l4.188 3.35a.5.5 0 0 0 .812-.39v-3.131l2.587 2.587-.01.005a1 1 0 0 0 .86 1.806c.215-.102.424-.214.627-.333l2.3 2.3a1.001 1.001 0 0 0 1.414-1.416zM11 5.04a.5.5 0 0 0-.813-.39L8.682 5.854 11 8.172V5.04z"/></svg>'},{}],kB3Mf:[function(e,t,r){t.exports='<svg xmlns="http://www.w3.org/2000/svg" height="22" width="22" viewBox="0 0 50 50"><path d="M19.402 6a5 5 0 0 0-4.902 4.012L14.098 12H9a5 5 0 0 0-5 5v21a5 5 0 0 0 5 5h32a5 5 0 0 0 5-5V17a5 5 0 0 0-5-5h-5.098l-.402-1.988A5 5 0 0 0 30.598 6ZM25 17c5.52 0 10 4.48 10 10s-4.48 10-10 10-10-4.48-10-10 4.48-10 10-10Zm0 2c-4.41 0-8 3.59-8 8s3.59 8 8 8 8-3.59 8-8-3.59-8-8-8Z"/></svg>'},{}],"3MONs":[function(e,t,r){t.exports='<svg xmlns="http://www.w3.org/2000/svg" height="22" width="22"><circle cx="11" cy="11" r="2"/><path d="M19.164 8.861 17.6 8.6a6.978 6.978 0 0 0-1.186-2.099l.574-1.533a1 1 0 0 0-.436-1.217l-1.997-1.153a1.001 1.001 0 0 0-1.272.23l-1.008 1.225a7.04 7.04 0 0 0-2.55.001L8.716 2.829a1 1 0 0 0-1.272-.23L5.447 3.751a1 1 0 0 0-.436 1.217l.574 1.533A6.997 6.997 0 0 0 4.4 8.6l-1.564.261A.999.999 0 0 0 2 9.847v2.306c0 .489.353.906.836.986l1.613.269a7 7 0 0 0 1.228 2.075l-.558 1.487a1 1 0 0 0 .436 1.217l1.997 1.153c.423.244.961.147 1.272-.23l1.04-1.263a7.089 7.089 0 0 0 2.272 0l1.04 1.263a1 1 0 0 0 1.272.23l1.997-1.153a1 1 0 0 0 .436-1.217l-.557-1.487c.521-.61.94-1.31 1.228-2.075l1.613-.269a.999.999 0 0 0 .835-.986V9.847a.999.999 0 0 0-.836-.986zM11 15a4 4 0 1 1 0-8 4 4 0 0 1 0 8z"/></svg>'},{}],iMCpk:[function(e,t,r){t.exports='<svg xmlns="http://www.w3.org/2000/svg" height="32" width="32"><path d="m19.41 20.09-4.58-4.59 4.58-4.59L18 9.5l-6 6 6 6z" fill="#fff"/></svg>'},{}],"3oe4L":[function(e,t,r){t.exports='<svg xmlns="http://www.w3.org/2000/svg" height="32" width="32"><path d="m12.59 20.34 4.58-4.59-4.58-4.59L14 9.75l6 6-6 6z" fill="#fff"/></svg>'},{}],liE22:[function(e,t,r){t.exports='<svg height="24" width="24"><path d="M10 8v8l6-4-6-4zM6.3 5l-.6-.8C7.2 3 9 2.2 11 2l.1 1c-1.8.2-3.4.9-4.8 2zM5 6.3l-.8-.6C3 7.2 2.2 9 2 11l1 .1c.2-1.8.9-3.4 2-4.8zm0 11.4c-1.1-1.4-1.8-3.1-2-4.8L2 13c.2 2 1 3.8 2.2 5.4l.8-.7zm6.1 3.3c-1.8-.2-3.4-.9-4.8-2l-.6.8C7.2 21 9 21.8 11 22l.1-1zM22 12c0-5.2-3.9-9.4-9-10l-.1 1c4.6.5 8.1 4.3 8.1 9s-3.5 8.5-8.1 9l.1 1c5.2-.5 9-4.8 9-10z" fill="#fff" style="--darkreader-inline-fill:#a8a6a4"/></svg>'},{}],"8HqYc":[function(e,t,r){t.exports='<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 88 88" style="width:100%;height:100%;transform:translate(0,0)"><defs><clipPath id="__lottie_element_216"><path d="M0 0h88v88H0z"/></clipPath></defs><g style="display:block" clip-path="url(\'#__lottie_element_216\')"><path fill="#FFF" d="m12.438-12.702-2.82 2.82c-.79.79-.79 2.05 0 2.83l7.07 7.07-7.07 7.07c-.79.79-.79 2.05 0 2.83l2.82 2.83c.79.78 2.05.78 2.83 0l11.32-11.31c.78-.78.78-2.05 0-2.83l-11.32-11.31c-.78-.79-2.04-.79-2.83 0zm-24.88 0c-.74-.74-1.92-.78-2.7-.12l-.13.12-11.31 11.31a2 2 0 0 0-.12 2.7l.12.13 11.31 11.31a2 2 0 0 0 2.7.12l.13-.12 2.83-2.83c.74-.74.78-1.91.11-2.7l-.11-.13-7.07-7.07 7.07-7.07c.74-.74.78-1.91.11-2.7l-.11-.13-2.83-2.82zM28-28c4.42 0 8 3.58 8 8v40c0 4.42-3.58 8-8 8h-56c-4.42 0-8-3.58-8-8v-40c0-4.42 3.58-8 8-8h56z" style="--darkreader-inline-fill:#a8a6a4" transform="translate(44 44)"/></g></svg>'},{}],hYAAH:[function(e,t,r){t.exports='<svg height="24" width="24"><path d="M15 17h6v1h-6v-1zm-4 0H3v1h8v2h1v-5h-1v2zm3-9h1V3h-1v2H3v1h11v2zm4-3v1h3V5h-3zM6 14h1V9H6v2H3v1h3v2zm4-2h11v-1H10v1z" fill="#fff" style="--darkreader-inline-fill:#a8a6a4"/></svg>'},{}],jmNrH:[function(e,t,r){t.exports='<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 36 36" height="32" width="32"><path d="M25 17h-8v6h8v-6Zm4 8V10.98C29 9.88 28.1 9 27 9H9c-1.1 0-2 .88-2 1.98V25c0 1.1.9 2 2 2h18c1.1 0 2-.9 2-2Zm-2 .02H9V10.97h18v14.05Z"/></svg>'},{}],cIqko:[function(e,t,r){t.exports='<svg class="icon" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" width="20" height="20"><path d="M298.667 426.667v-85.334a213.333 213.333 0 1 1 426.666 0v85.334H768A85.333 85.333 0 0 1 853.333 512v256A85.333 85.333 0 0 1 768 853.333H256A85.333 85.333 0 0 1 170.667 768V512A85.333 85.333 0 0 1 256 426.667h42.667zM512 213.333a128 128 0 0 0-128 128v85.334h256v-85.334a128 128 0 0 0-128-128z" fill="#fff"/></svg>'},{}],"65zy4":[function(e,t,r){t.exports='<svg class="icon" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" width="20" height="20"><path d="m666.752 194.517-49.365 74.112A128 128 0 0 0 384 341.333l.043 85.334h384A85.333 85.333 0 0 1 853.376 512v256a85.333 85.333 0 0 1-85.333 85.333H256A85.333 85.333 0 0 1 170.667 768V512A85.333 85.333 0 0 1 256 426.667h42.667v-85.334a213.333 213.333 0 0 1 368.085-146.816z" fill="#fff"/></svg>'},{}],jaJRT:[function(e,t,r){t.exports='<svg class="icon" width="22" height="22" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg"><path fill="#fff" d="M768 298.667h170.667V384h-256V128H768v170.667zM341.333 384h-256v-85.333H256V128h85.333v256zM768 725.333V896h-85.333V640h256v85.333H768zM341.333 640v256H256V725.333H85.333V640h256z"/></svg>'},{}],cRY1X:[function(e,t,r){t.exports='<svg class="icon" width="22" height="22" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg"><path fill="#fff" d="M625.778 256H768v142.222h113.778v-256h-256V256zM256 398.222V256h142.222V142.222h-256v256H256zm512 227.556V768H625.778v113.778h256v-256H768zM398.222 768H256V625.778H142.222v256h256V768z"/></svg>'},{}],"3aVGL":[function(e,t,r){t.exports='<svg class="icon" width="18" height="18" viewBox="0 0 1152 1024" xmlns="http://www.w3.org/2000/svg"><path fill="#fff" d="M1075.2 0H76.8A76.8 76.8 0 0 0 0 76.8v870.4a76.8 76.8 0 0 0 76.8 76.8h998.4a76.8 76.8 0 0 0 76.8-76.8V76.8A76.8 76.8 0 0 0 1075.2 0zM1024 128v768H128V128h896zM896 512a64 64 0 0 1 7.488 127.552L896 640H768v128a64 64 0 0 1-56.512 63.552L704 832a64 64 0 0 1-63.552-56.512L640 768V582.592c0-34.496 25.024-66.112 61.632-70.208l8-.384H896zm-640 0a64 64 0 0 1-7.488-127.552L256 384h128V256a64 64 0 0 1 56.512-63.552L448 192a64 64 0 0 1 63.552 56.512L512 256v185.408c0 34.432-25.024 66.112-61.632 70.144l-8 .448H256z"/></svg>'},{}],"4DiVn":[function(e,t,r){t.exports='<svg class="icon" width="18" height="18" viewBox="0 0 1152 1024" xmlns="http://www.w3.org/2000/svg"><path fill="#fff" d="M1075.2 0H76.8A76.8 76.8 0 0 0 0 76.8v870.4a76.8 76.8 0 0 0 76.8 76.8h998.4a76.8 76.8 0 0 0 76.8-76.8V76.8A76.8 76.8 0 0 0 1075.2 0zM1024 128v768H128V128h896zm-576 64a64 64 0 0 1 7.488 127.552L448 320H320v128a64 64 0 0 1-56.512 63.552L256 512a64 64 0 0 1-63.552-56.512L192 448V262.592c0-34.432 25.024-66.112 61.632-70.144l8-.448H448zm256 640a64 64 0 0 1-7.488-127.552L704 704h128V576a64 64 0 0 1 56.512-63.552L896 512a64 64 0 0 1 63.552 56.512L960 576v185.408c0 34.496-25.024 66.112-61.632 70.208l-8 .384H704z"/></svg>'},{}],kwdKE:[function(e,t,r){t.exports='<svg class="icon" width="26" height="26" viewBox="0 0 1664 1024" xmlns="http://www.w3.org/2000/svg"><path fill="#648FFC" d="M1152 0H512a512 512 0 0 0 0 1024h640a512 512 0 0 0 0-1024zm0 960a448 448 0 1 1 448-448 448 448 0 0 1-448 448z"/></svg>'},{}],bWfXZ:[function(e,t,r){t.exports='<svg class="icon" width="26" height="26" viewBox="0 0 1740 1024" xmlns="http://www.w3.org/2000/svg"><path fill="#fff" d="M511.898 1024h670.515c282.419-.41 511.18-229.478 511.18-511.898 0-282.419-228.761-511.488-511.18-511.897H511.898C229.478.615.717 229.683.717 512.102c0 282.42 228.761 511.488 511.18 511.898zm-.564-975.36A464.589 464.589 0 1 1 48.026 513.024 463.872 463.872 0 0 1 511.334 48.435v.205z"/></svg>'},{}],h3zZ9:[function(e,t,r){t.exports='<svg class="icon" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" width="24" height="24"><path d="M554.667 810.667V896h-85.334v-85.333h85.334zm-384-632.662a42.667 42.667 0 0 1 34.986 18.219l203.904 291.328a42.667 42.667 0 0 1 0 48.896L205.611 827.776A42.667 42.667 0 0 1 128 803.328V220.672a42.667 42.667 0 0 1 42.667-42.667zm682.666 0a42.667 42.667 0 0 1 42.368 37.718l.299 4.949v582.656a42.667 42.667 0 0 1-74.24 28.63l-3.413-4.182-203.904-291.328a42.667 42.667 0 0 1-3.03-43.861l3.03-5.035 203.946-291.328a42.667 42.667 0 0 1 34.944-18.219zM554.667 640v85.333h-85.334V640h85.334zm-358.4-320.896V716.8L335.957 512 196.31 319.104zm358.4 150.23v85.333h-85.334v-85.334h85.334zm0-170.667V384h-85.334v-85.333h85.334zm0-170.667v85.333h-85.334V128h85.334z" fill="#fff"/></svg>'},{}],"7Oyth":[function(e,t,r){t.exports='<svg viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" width="50" height="50"><path d="M593.818 168.55 949.82 763.76c26.153 43.746 10.732 99.738-34.447 125.052-14.397 8.069-30.72 12.308-47.37 12.308H155.976c-52.224 0-94.536-40.96-94.536-91.505 0-16.097 4.383-31.928 12.718-45.875l356.004-595.19c26.173-43.724 84.009-58.654 129.208-33.341a93.082 93.082 0 0 1 34.448 33.341zM512 819.2a61.44 61.44 0 1 0 0-122.88 61.44 61.44 0 0 0 0 122.88zm0-512a72.315 72.315 0 0 0-71.762 81.306l25.723 205.721a46.408 46.408 0 0 0 92.078 0l25.723-205.742A72.315 72.315 0 0 0 512 307.2z"/></svg>'},{}],U5Jcy:[function(e,t,r){t.exports='<svg class="icon" viewBox="0 0 1024 1024" xmlns="http://www.w3.org/2000/svg" width="22" height="22"><path d="m571.733 512 268.8-268.8c17.067-17.067 17.067-42.667 0-59.733-17.066-17.067-42.666-17.067-59.733 0L512 452.267l-268.8-268.8c-17.067-17.067-42.667-17.067-59.733 0-17.067 17.066-17.067 42.666 0 59.733l268.8 268.8-268.8 268.8c-17.067 17.067-17.067 42.667 0 59.733 8.533 8.534 19.2 12.8 29.866 12.8s21.334-4.266 29.867-12.8l268.8-268.8 268.8 268.8c8.533 8.534 19.2 12.8 29.867 12.8s21.333-4.266 29.866-12.8c17.067-17.066 17.067-42.666 0-59.733L571.733 512z"/></svg>'},{}],jK5Fx:[function(e,t,r){t.exports='<svg width="18" height="18" xmlns="http://www.w3.org/2000/svg"><g fill="#fff"><path d="M16 1H2a1 1 0 0 0-1 1v10a1 1 0 0 0 1 1h3v-2H3V3h12v8h-2v2h3a1 1 0 0 0 1-1V2a1 1 0 0 0-1-1Z"/><path d="M4 17h10l-5-6z"/></g></svg>'},{}],bRHiA:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("./flip"),o=a.interopDefault(i),n=e("./aspectRatio"),s=a.interopDefault(n),l=e("./playbackRate"),c=a.interopDefault(l),u=e("./subtitleOffset"),p=a.interopDefault(u),d=e("../utils/component"),f=a.interopDefault(d),h=e("../utils/error"),m=e("../utils");class g extends f.default{constructor(e){super(e);let{option:t,controls:r,template:{$setting:a}}=e;this.name="setting",this.$parent=a,this.option=[],this.events=[],this.cache=new Map,t.setting&&(this.init(),e.on("blur",()=>{this.show&&(this.show=!1,this.render(this.option))}),e.on("focus",e=>{let t=(0,m.includeFromEvent)(e,r.setting),a=(0,m.includeFromEvent)(e,this.$parent);!this.show||t||a||(this.show=!1,this.render(this.option))}))}static makeRecursion(e,t,r){for(let a=0;a<e.length;a++){let i=e[a];i.$parentItem=t,i.$parentList=r,g.makeRecursion(i.selector||[],i,e)}return e}get defaultSettings(){let e=[],{option:t}=this.art;return t.playbackRate&&e.push((0,c.default)(this.art)),t.aspectRatio&&e.push((0,s.default)(this.art)),t.flip&&e.push((0,o.default)(this.art)),t.subtitleOffset&&e.push((0,p.default)(this.art)),e}init(){let{option:e}=this.art,t=[...this.defaultSettings,...e.settings];this.option=g.makeRecursion(t),this.destroy(),this.render(this.option)}destroy(){for(let e=0;e<this.events.length;e++)this.art.events.remove(this.events[e]);this.$parent.innerHTML="",this.events=[],this.cache=new Map}find(e="",t=this.option){for(let r=0;r<t.length;r++){let a=t[r];if(a.name===e)return a;{let t=this.find(e,a.selector||[]);if(t)return t}}}remove(e){let t=this.find(e);(0,h.errorHandle)(t,`Can't find [${e}] from the [setting]`);let r=t.$parentItem?t.$parentItem.selector:this.option;return r.splice(r.indexOf(t),1),this.option=g.makeRecursion(this.option),this.destroy(),this.render(this.option),this.option}update(e){let t=this.find(e.name);return t?(Object.assign(t,e),this.option=g.makeRecursion(this.option),this.destroy(),this.render(this.option)):this.add(e),this.option}add(e){return this.option.push(e),this.option=g.makeRecursion(this.option),this.destroy(),this.render(this.option),this.option}creatHeader(e){let{icons:t,proxy:r,constructor:a}=this.art,i=(0,m.createElement)("div");(0,m.setStyle)(i,"height",`${a.SETTING_ITEM_HEIGHT}px`),(0,m.addClass)(i,"art-setting-item"),(0,m.addClass)(i,"art-setting-item-back");let o=(0,m.append)(i,'<div class="art-setting-item-left"></div>'),n=(0,m.createElement)("div");(0,m.addClass)(n,"art-setting-item-left-icon"),(0,m.append)(n,t.arrowLeft),(0,m.append)(o,n),(0,m.append)(o,e.$parentItem.html);let s=r(i,"click",()=>this.render(e.$parentList));return this.events.push(s),i}creatItem(e,t){let{icons:r,proxy:a,constructor:i}=this.art,o=(0,m.createElement)("div");(0,m.addClass)(o,"art-setting-item"),(0,m.setStyle)(o,"height",`${i.SETTING_ITEM_HEIGHT}px`),(0,m.isStringOrNumber)(t.name)&&(o.dataset.name=t.name),(0,m.isStringOrNumber)(t.value)&&(o.dataset.value=t.value);let n=(0,m.append)(o,'<div class="art-setting-item-left"></div>'),s=(0,m.append)(o,'<div class="art-setting-item-right"></div>'),l=(0,m.createElement)("div");switch((0,m.addClass)(l,"art-setting-item-left-icon"),e){case"switch":case"range":(0,m.append)(l,(0,m.isStringOrNumber)(t.icon)||t.icon instanceof Element?t.icon:r.config);break;case"selector":t.selector&&t.selector.length?(0,m.append)(l,(0,m.isStringOrNumber)(t.icon)||t.icon instanceof Element?t.icon:r.config):(0,m.append)(l,r.check)}(0,m.append)(n,l),t.$icon=l,(0,m.def)(t,"icon",{configurable:!0,get:()=>l.innerHTML,set(e){(0,m.isStringOrNumber)(e)&&(l.innerHTML=e)}});let c=(0,m.createElement)("div");(0,m.addClass)(c,"art-setting-item-left-text"),(0,m.append)(c,t.html||""),(0,m.append)(n,c),t.$html=c,(0,m.def)(t,"html",{configurable:!0,get:()=>c.innerHTML,set(e){(0,m.isStringOrNumber)(e)&&(c.innerHTML=e)}});let u=(0,m.createElement)("div");switch((0,m.addClass)(u,"art-setting-item-right-tooltip"),(0,m.append)(u,t.tooltip||""),(0,m.append)(s,u),t.$tooltip=u,(0,m.def)(t,"tooltip",{configurable:!0,get:()=>u.innerHTML,set(e){(0,m.isStringOrNumber)(e)&&(u.innerHTML=e)}}),e){case"switch":{let e=(0,m.createElement)("div");(0,m.addClass)(e,"art-setting-item-right-icon");let a=(0,m.append)(e,r.switchOn),i=(0,m.append)(e,r.switchOff);(0,m.setStyle)(t.switch?i:a,"display","none"),(0,m.append)(s,e),t.$switch=t.switch,(0,m.def)(t,"switch",{configurable:!0,get:()=>t.$switch,set(e){t.$switch=e,e?((0,m.setStyle)(i,"display","none"),(0,m.setStyle)(a,"display",null)):((0,m.setStyle)(i,"display",null),(0,m.setStyle)(a,"display","none"))}});break}case"range":{let e=(0,m.createElement)("div");(0,m.addClass)(e,"art-setting-item-right-icon");let r=(0,m.append)(e,'<input type="range">');r.value=t.range[0]||0,r.min=t.range[1]||0,r.max=t.range[2]||10,r.step=t.range[3]||1,(0,m.addClass)(r,"art-setting-range"),(0,m.append)(s,e),t.$range=r,(0,m.def)(t,"range",{configurable:!0,get:()=>r.valueAsNumber,set(e){r.value=Number(e)}})}break;case"selector":if(t.selector&&t.selector.length){let e=(0,m.createElement)("div");(0,m.addClass)(e,"art-setting-item-right-icon"),(0,m.append)(e,r.arrowRight),(0,m.append)(s,e)}}switch(e){case"switch":if(t.onSwitch){let e=a(o,"click",async e=>{t.switch=await t.onSwitch.call(this.art,t,o,e)});this.events.push(e)}break;case"range":if(t.$range){if(t.onRange){let e=a(t.$range,"change",async e=>{t.tooltip=await t.onRange.call(this.art,t,o,e)});this.events.push(e)}if(t.onChange){let e=a(t.$range,"input",async e=>{t.tooltip=await t.onChange.call(this.art,t,o,e)});this.events.push(e)}}break;case"selector":{let e=a(o,"click",async e=>{if(t.selector&&t.selector.length)this.render(t.selector,t.width);else{(0,m.inverseClass)(o,"art-current");for(let e=0;e<t.$parentItem.selector.length;e++){let r=t.$parentItem.selector[e];r.default=r===t}if(t.$parentList&&this.render(t.$parentList),t.$parentItem&&t.$parentItem.onSelect){let r=await t.$parentItem.onSelect.call(this.art,t,o,e);t.$parentItem.$tooltip&&(0,m.isStringOrNumber)(r)&&(t.$parentItem.$tooltip.innerHTML=r)}}});this.events.push(e),t.default&&(0,m.addClass)(o,"art-current")}}return o}updateStyle(e){let{controls:t,constructor:r,template:{$player:a,$setting:i}}=this.art;if(t.setting&&!m.isMobile){let o=e||r.SETTING_WIDTH,{left:n,width:s}=(0,m.getRect)(t.setting),{left:l,width:c}=(0,m.getRect)(a),u=n-l+s/2-o/2;u+o>c?((0,m.setStyle)(i,"left",null),(0,m.setStyle)(i,"right",null)):((0,m.setStyle)(i,"left",`${u}px`),(0,m.setStyle)(i,"right","auto"))}}render(e,t){let{constructor:r}=this.art;if(this.cache.has(e)){let t=this.cache.get(e);(0,m.inverseClass)(t,"art-current"),(0,m.setStyle)(this.$parent,"width",`${t.dataset.width}px`),(0,m.setStyle)(this.$parent,"height",`${t.dataset.height}px`),this.updateStyle(Number(t.dataset.width))}else{let a=(0,m.createElement)("div");(0,m.addClass)(a,"art-setting-panel"),a.dataset.width=t||r.SETTING_WIDTH,a.dataset.height=e.length*r.SETTING_ITEM_HEIGHT,e[0]&&e[0].$parentItem&&((0,m.append)(a,this.creatHeader(e[0])),a.dataset.height=Number(a.dataset.height)+r.SETTING_ITEM_HEIGHT);for(let t=0;t<e.length;t++){let r=e[t];(0,m.has)(r,"switch")?(0,m.append)(a,this.creatItem("switch",r)):(0,m.has)(r,"range")?(0,m.append)(a,this.creatItem("range",r)):(0,m.append)(a,this.creatItem("selector",r))}(0,m.append)(this.$parent,a),this.cache.set(e,a),(0,m.inverseClass)(a,"art-current"),(0,m.setStyle)(this.$parent,"width",`${a.dataset.width}px`),(0,m.setStyle)(this.$parent,"height",`${a.dataset.height}px`),this.updateStyle(Number(a.dataset.width)),e[0]&&e[0].$parentItem&&e[0].$parentItem.mounted&&e[0].$parentItem.mounted.call(this.art,a,e[0].$parentItem)}}}r.default=g},{"./flip":"bNOaj","./aspectRatio":"5lAsp","./playbackRate":"e6hsR","./subtitleOffset":"fFNEr","../utils/component":"guki8","../utils/error":"2nFlF","../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],bNOaj:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{i18n:t,icons:r,constructor:{SETTING_ITEM_WIDTH:a,FLIP:o}}=e;function n(e,r,a){r&&(r.innerText=t.get((0,i.capitalize)(a)));let o=(0,i.queryAll)(".art-setting-item",e).find(e=>e.dataset.value===a);o&&(0,i.inverseClass)(o,"art-current")}return{width:a,name:"flip",html:t.get("Video Flip"),tooltip:t.get((0,i.capitalize)(e.flip)),icon:r.flip,selector:o.map(r=>({value:r,name:`aspect-ratio-${r}`,default:r===e.flip,html:t.get((0,i.capitalize)(r))})),onSelect:t=>(e.flip=t.value,t.html),mounted:(t,r)=>{n(t,r.$tooltip,e.flip),e.on("flip",()=>{n(t,r.$tooltip,e.flip)})}}}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"5lAsp":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{i18n:t,icons:r,constructor:{SETTING_ITEM_WIDTH:a,ASPECT_RATIO:o}}=e;function n(e){return"default"===e?t.get("Default"):e}function s(e,t,r){t&&(t.innerText=n(r));let a=(0,i.queryAll)(".art-setting-item",e).find(e=>e.dataset.value===r);a&&(0,i.inverseClass)(a,"art-current")}return{width:a,name:"aspect-ratio",html:t.get("Aspect Ratio"),icon:r.aspectRatio,tooltip:n(e.aspectRatio),selector:o.map(t=>({value:t,name:`aspect-ratio-${t}`,default:t===e.aspectRatio,html:n(t)})),onSelect:t=>(e.aspectRatio=t.value,t.html),mounted:(t,r)=>{s(t,r.$tooltip,e.aspectRatio),e.on("aspectRatio",()=>{s(t,r.$tooltip,e.aspectRatio)})}}}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],e6hsR:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{i18n:t,icons:r,constructor:{SETTING_ITEM_WIDTH:a,PLAYBACK_RATE:o}}=e;function n(e){return 1===e?t.get("Normal"):e.toFixed(1)}function s(e,t,r){t&&(t.innerText=n(r));let a=(0,i.queryAll)(".art-setting-item",e).find(e=>Number(e.dataset.value)===r);a&&(0,i.inverseClass)(a,"art-current")}return{width:a,name:"playback-rate",html:t.get("Play Speed"),tooltip:n(e.playbackRate),icon:r.playbackRate,selector:o.map(t=>({value:t,name:`aspect-ratio-${t}`,default:t===e.playbackRate,html:n(t)})),onSelect:t=>(e.playbackRate=t.value,t.html),mounted:(t,r)=>{s(t,r.$tooltip,e.playbackRate),e.on("video:ratechange",()=>{s(t,r.$tooltip,e.playbackRate)})}}}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],fFNEr:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");function i(e){let{i18n:t,icons:r,constructor:a}=e;return{width:a.SETTING_ITEM_WIDTH,name:"subtitle-offset",html:t.get("Subtitle Offset"),icon:r.subtitle,tooltip:"0s",range:[0,-5,5,.1],onChange:t=>(e.subtitleOffset=t.range,t.range+"s")}}a.defineInteropFlag(r),a.export(r,"default",()=>i)},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],f2Thp:[function(e,t,r){e("@parcel/transformer-js/src/esmodule-helpers.js").defineInteropFlag(r),r.default=class{constructor(){this.name="artplayer_settings",this.settings={}}get(e){try{let t=JSON.parse(window.localStorage.getItem(this.name))||{};return e?t[e]:t}catch(t){return e?this.settings[e]:this.settings}}set(e,t){try{let r=Object.assign({},this.get(),{[e]:t});window.localStorage.setItem(this.name,JSON.stringify(r))}catch(r){this.settings[e]=t}}del(e){try{let t=this.get();delete t[e],window.localStorage.setItem(this.name,JSON.stringify(t))}catch(t){delete this.settings[e]}}clear(){try{window.localStorage.removeItem(this.name)}catch(e){this.settings={}}}}},{"@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"96ThS":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r);var i=e("../utils"),o=e("./miniProgressBar"),n=a.interopDefault(o),s=e("./autoOrientation"),l=a.interopDefault(s),c=e("./autoPlayback"),u=a.interopDefault(c),p=e("./fastForward"),d=a.interopDefault(p),f=e("./lock"),h=a.interopDefault(f);r.default=class{constructor(e){this.art=e,this.id=0;let{option:t}=e;t.miniProgressBar&&!t.isLive&&this.add(n.default),t.lock&&i.isMobile&&this.add(h.default),t.autoPlayback&&!t.isLive&&this.add(u.default),t.autoOrientation&&i.isMobile&&this.add(l.default),t.fastForward&&i.isMobile&&!t.isLive&&this.add(d.default);for(let e=0;e<t.plugins.length;e++)this.add(t.plugins[e])}add(e){this.id+=1;let t=e.call(this.art,this.art);return t instanceof Promise?t.then(t=>this.next(e,t)):this.next(e,t)}next(e,t){let r=t&&t.name||e.name||`plugin${this.id}`;return(0,i.errorHandle)(!(0,i.has)(this,r),`Cannot add a plugin that already has the same name: ${r}`),(0,i.def)(this,r,{value:t}),this}}},{"../utils":"h3rH9","./miniProgressBar":"iBx4M","./autoOrientation":"2O9qO","./autoPlayback":"iiOc1","./fastForward":"d9NUE","./lock":"5dnKh","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],iBx4M:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){return e.on("control",t=>{t?(0,i.removeClass)(e.template.$player,"art-mini-progress-bar"):(0,i.addClass)(e.template.$player,"art-mini-progress-bar")}),{name:"mini-progress-bar"}}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"2O9qO":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{constructor:t,template:{$player:r,$video:a}}=e;return e.on("fullscreenWeb",o=>{if(o){let{videoWidth:o,videoHeight:n}=a,{clientWidth:s,clientHeight:l}=document.documentElement;(o>n&&s<l||o<n&&s>l)&&setTimeout(()=>{(0,i.setStyle)(r,"width",`${l}px`),(0,i.setStyle)(r,"height",`${s}px`),(0,i.setStyle)(r,"transform-origin","0 0"),(0,i.setStyle)(r,"transform",`rotate(90deg) translate(0, -${s}px)`),(0,i.addClass)(r,"art-auto-orientation"),e.isRotate=!0,e.emit("resize")},t.AUTO_ORIENTATION_TIME)}else(0,i.hasClass)(r,"art-auto-orientation")&&((0,i.removeClass)(r,"art-auto-orientation"),e.isRotate=!1,e.emit("resize"))}),e.on("fullscreen",async e=>{if(!screen?.orientation?.lock)return;let t=screen.orientation.type;if(e){let{videoWidth:e,videoHeight:o}=a,{clientWidth:n,clientHeight:s}=document.documentElement;if(e>o&&n<s||e<o&&n>s){let e=t.startsWith("portrait")?"landscape":"portrait";await screen.orientation.lock(e),(0,i.addClass)(r,"art-auto-orientation-fullscreen")}}else(0,i.hasClass)(r,"art-auto-orientation-fullscreen")&&(await screen.orientation.lock(t),(0,i.removeClass)(r,"art-auto-orientation-fullscreen"))}),{name:"autoOrientation",get state(){return(0,i.hasClass)(r,"art-auto-orientation")}}}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],iiOc1:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{i18n:t,icons:r,storage:a,constructor:o,proxy:n,template:{$poster:s}}=e,l=e.layers.add({name:"auto-playback",html:`<div class="art-auto-playback-close"></div><div class="art-auto-playback-last"></div><div class="art-auto-playback-jump"></div>`}),c=(0,i.query)(".art-auto-playback-last",l),u=(0,i.query)(".art-auto-playback-jump",l),p=(0,i.query)(".art-auto-playback-close",l);return e.on("video:timeupdate",()=>{if(e.playing){let t=a.get("times")||{},r=Object.keys(t);r.length>o.AUTO_PLAYBACK_MAX&&delete t[r[0]],t[e.option.id||e.option.url]=e.currentTime,a.set("times",t)}}),e.on("ready",()=>{let d=(a.get("times")||{})[e.option.id||e.option.url];d&&d>=o.AUTO_PLAYBACK_MIN&&((0,i.append)(p,r.close),(0,i.setStyle)(l,"display","flex"),c.innerText=`${t.get("Last Seen")} ${(0,i.secondToTime)(d)}`,u.innerText=t.get("Jump Play"),n(p,"click",()=>{(0,i.setStyle)(l,"display","none")}),n(u,"click",()=>{e.seek=d,e.play(),(0,i.setStyle)(s,"display","none"),(0,i.setStyle)(l,"display","none")}),e.once("video:timeupdate",()=>{setTimeout(()=>{(0,i.setStyle)(l,"display","none")},o.AUTO_PLAYBACK_TIMEOUT)}))}),{name:"auto-playback",get times(){return a.get("times")||{}},clear:()=>a.del("times"),delete(e){let t=a.get("times")||{};return delete t[e],a.set("times",t),t}}}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],d9NUE:[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{constructor:t,proxy:r,template:{$player:a,$video:o}}=e,n=null,s=!1,l=1,c=()=>{clearTimeout(n),s&&(s=!1,e.playbackRate=l,(0,i.removeClass)(a,"art-fast-forward"))};return r(o,"touchstart",r=>{1===r.touches.length&&e.playing&&!e.isLock&&(n=setTimeout(()=>{s=!0,l=e.playbackRate,e.playbackRate=t.FAST_FORWARD_VALUE,(0,i.addClass)(a,"art-fast-forward")},t.FAST_FORWARD_TIME))}),r(document,"touchmove",c),r(document,"touchend",c),{name:"fastForward",get state(){return(0,i.hasClass)(a,"art-fast-forward")}}}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}],"5dnKh":[function(e,t,r){var a=e("@parcel/transformer-js/src/esmodule-helpers.js");a.defineInteropFlag(r),a.export(r,"default",()=>o);var i=e("../utils");function o(e){let{layers:t,icons:r,template:{$player:a}}=e;function o(){return(0,i.hasClass)(a,"art-lock")}function n(){(0,i.addClass)(a,"art-lock"),e.isLock=!0,e.emit("lock",!0)}function s(){(0,i.removeClass)(a,"art-lock"),e.isLock=!1,e.emit("lock",!1)}return t.add({name:"lock",mounted(t){let a=(0,i.append)(t,r.lock),o=(0,i.append)(t,r.unlock);(0,i.setStyle)(a,"display","none"),e.on("lock",e=>{e?((0,i.setStyle)(a,"display","inline-flex"),(0,i.setStyle)(o,"display","none")):((0,i.setStyle)(a,"display","none"),(0,i.setStyle)(o,"display","inline-flex"))})},click(){o()?s():n()}}),{name:"lock",get state(){return o()},set state(value){value?n():s()}}}},{"../utils":"h3rH9","@parcel/transformer-js/src/esmodule-helpers.js":"guZOB"}]},["abjMI"],"abjMI","parcelRequireb749");

/***/ }),

/***/ "./node_modules/base64-js/index.js":
/*!*****************************************!*\
  !*** ./node_modules/base64-js/index.js ***!
  \*****************************************/
/***/ ((__unused_webpack_module, exports) => {

"use strict";


exports.byteLength = byteLength
exports.toByteArray = toByteArray
exports.fromByteArray = fromByteArray

var lookup = []
var revLookup = []
var Arr = typeof Uint8Array !== 'undefined' ? Uint8Array : Array

var code = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/'
for (var i = 0, len = code.length; i < len; ++i) {
  lookup[i] = code[i]
  revLookup[code.charCodeAt(i)] = i
}

// Support decoding URL-safe base64 strings, as Node.js does.
// See: https://en.wikipedia.org/wiki/Base64#URL_applications
revLookup['-'.charCodeAt(0)] = 62
revLookup['_'.charCodeAt(0)] = 63

function getLens (b64) {
  var len = b64.length

  if (len % 4 > 0) {
    throw new Error('Invalid string. Length must be a multiple of 4')
  }

  // Trim off extra bytes after placeholder bytes are found
  // See: https://github.com/beatgammit/base64-js/issues/42
  var validLen = b64.indexOf('=')
  if (validLen === -1) validLen = len

  var placeHoldersLen = validLen === len
    ? 0
    : 4 - (validLen % 4)

  return [validLen, placeHoldersLen]
}

// base64 is 4/3 + up to two characters of the original data
function byteLength (b64) {
  var lens = getLens(b64)
  var validLen = lens[0]
  var placeHoldersLen = lens[1]
  return ((validLen + placeHoldersLen) * 3 / 4) - placeHoldersLen
}

function _byteLength (b64, validLen, placeHoldersLen) {
  return ((validLen + placeHoldersLen) * 3 / 4) - placeHoldersLen
}

function toByteArray (b64) {
  var tmp
  var lens = getLens(b64)
  var validLen = lens[0]
  var placeHoldersLen = lens[1]

  var arr = new Arr(_byteLength(b64, validLen, placeHoldersLen))

  var curByte = 0

  // if there are placeholders, only get up to the last complete 4 chars
  var len = placeHoldersLen > 0
    ? validLen - 4
    : validLen

  var i
  for (i = 0; i < len; i += 4) {
    tmp =
      (revLookup[b64.charCodeAt(i)] << 18) |
      (revLookup[b64.charCodeAt(i + 1)] << 12) |
      (revLookup[b64.charCodeAt(i + 2)] << 6) |
      revLookup[b64.charCodeAt(i + 3)]
    arr[curByte++] = (tmp >> 16) & 0xFF
    arr[curByte++] = (tmp >> 8) & 0xFF
    arr[curByte++] = tmp & 0xFF
  }

  if (placeHoldersLen === 2) {
    tmp =
      (revLookup[b64.charCodeAt(i)] << 2) |
      (revLookup[b64.charCodeAt(i + 1)] >> 4)
    arr[curByte++] = tmp & 0xFF
  }

  if (placeHoldersLen === 1) {
    tmp =
      (revLookup[b64.charCodeAt(i)] << 10) |
      (revLookup[b64.charCodeAt(i + 1)] << 4) |
      (revLookup[b64.charCodeAt(i + 2)] >> 2)
    arr[curByte++] = (tmp >> 8) & 0xFF
    arr[curByte++] = tmp & 0xFF
  }

  return arr
}

function tripletToBase64 (num) {
  return lookup[num >> 18 & 0x3F] +
    lookup[num >> 12 & 0x3F] +
    lookup[num >> 6 & 0x3F] +
    lookup[num & 0x3F]
}

function encodeChunk (uint8, start, end) {
  var tmp
  var output = []
  for (var i = start; i < end; i += 3) {
    tmp =
      ((uint8[i] << 16) & 0xFF0000) +
      ((uint8[i + 1] << 8) & 0xFF00) +
      (uint8[i + 2] & 0xFF)
    output.push(tripletToBase64(tmp))
  }
  return output.join('')
}

function fromByteArray (uint8) {
  var tmp
  var len = uint8.length
  var extraBytes = len % 3 // if we have 1 byte left, pad 2 bytes
  var parts = []
  var maxChunkLength = 16383 // must be multiple of 3

  // go through the array every three bytes, we'll deal with trailing stuff later
  for (var i = 0, len2 = len - extraBytes; i < len2; i += maxChunkLength) {
    parts.push(encodeChunk(uint8, i, (i + maxChunkLength) > len2 ? len2 : (i + maxChunkLength)))
  }

  // pad the end with zeros, but make sure to not forget the extra bytes
  if (extraBytes === 1) {
    tmp = uint8[len - 1]
    parts.push(
      lookup[tmp >> 2] +
      lookup[(tmp << 4) & 0x3F] +
      '=='
    )
  } else if (extraBytes === 2) {
    tmp = (uint8[len - 2] << 8) + uint8[len - 1]
    parts.push(
      lookup[tmp >> 10] +
      lookup[(tmp >> 4) & 0x3F] +
      lookup[(tmp << 2) & 0x3F] +
      '='
    )
  }

  return parts.join('')
}


/***/ }),

/***/ "./node_modules/buffer/index.js":
/*!**************************************!*\
  !*** ./node_modules/buffer/index.js ***!
  \**************************************/
/***/ ((__unused_webpack_module, exports, __webpack_require__) => {

"use strict";
/*!
 * The buffer module from node.js, for the browser.
 *
 * @author   Feross Aboukhadijeh <https://feross.org>
 * @license  MIT
 */
/* eslint-disable no-proto */



const base64 = __webpack_require__(/*! base64-js */ "./node_modules/base64-js/index.js")
const ieee754 = __webpack_require__(/*! ieee754 */ "./node_modules/ieee754/index.js")
const customInspectSymbol =
  (typeof Symbol === 'function' && typeof Symbol['for'] === 'function') // eslint-disable-line dot-notation
    ? Symbol['for']('nodejs.util.inspect.custom') // eslint-disable-line dot-notation
    : null

exports.Buffer = Buffer
exports.SlowBuffer = SlowBuffer
exports.INSPECT_MAX_BYTES = 50

const K_MAX_LENGTH = 0x7fffffff
exports.kMaxLength = K_MAX_LENGTH

/**
 * If `Buffer.TYPED_ARRAY_SUPPORT`:
 *   === true    Use Uint8Array implementation (fastest)
 *   === false   Print warning and recommend using `buffer` v4.x which has an Object
 *               implementation (most compatible, even IE6)
 *
 * Browsers that support typed arrays are IE 10+, Firefox 4+, Chrome 7+, Safari 5.1+,
 * Opera 11.6+, iOS 4.2+.
 *
 * We report that the browser does not support typed arrays if the are not subclassable
 * using __proto__. Firefox 4-29 lacks support for adding new properties to `Uint8Array`
 * (See: https://bugzilla.mozilla.org/show_bug.cgi?id=695438). IE 10 lacks support
 * for __proto__ and has a buggy typed array implementation.
 */
Buffer.TYPED_ARRAY_SUPPORT = typedArraySupport()

if (!Buffer.TYPED_ARRAY_SUPPORT && typeof console !== 'undefined' &&
    typeof console.error === 'function') {
  console.error(
    'This browser lacks typed array (Uint8Array) support which is required by ' +
    '`buffer` v5.x. Use `buffer` v4.x if you require old browser support.'
  )
}

function typedArraySupport () {
  // Can typed array instances can be augmented?
  try {
    const arr = new Uint8Array(1)
    const proto = { foo: function () { return 42 } }
    Object.setPrototypeOf(proto, Uint8Array.prototype)
    Object.setPrototypeOf(arr, proto)
    return arr.foo() === 42
  } catch (e) {
    return false
  }
}

Object.defineProperty(Buffer.prototype, 'parent', {
  enumerable: true,
  get: function () {
    if (!Buffer.isBuffer(this)) return undefined
    return this.buffer
  }
})

Object.defineProperty(Buffer.prototype, 'offset', {
  enumerable: true,
  get: function () {
    if (!Buffer.isBuffer(this)) return undefined
    return this.byteOffset
  }
})

function createBuffer (length) {
  if (length > K_MAX_LENGTH) {
    throw new RangeError('The value "' + length + '" is invalid for option "size"')
  }
  // Return an augmented `Uint8Array` instance
  const buf = new Uint8Array(length)
  Object.setPrototypeOf(buf, Buffer.prototype)
  return buf
}

/**
 * The Buffer constructor returns instances of `Uint8Array` that have their
 * prototype changed to `Buffer.prototype`. Furthermore, `Buffer` is a subclass of
 * `Uint8Array`, so the returned instances will have all the node `Buffer` methods
 * and the `Uint8Array` methods. Square bracket notation works as expected -- it
 * returns a single octet.
 *
 * The `Uint8Array` prototype remains unmodified.
 */

function Buffer (arg, encodingOrOffset, length) {
  // Common case.
  if (typeof arg === 'number') {
    if (typeof encodingOrOffset === 'string') {
      throw new TypeError(
        'The "string" argument must be of type string. Received type number'
      )
    }
    return allocUnsafe(arg)
  }
  return from(arg, encodingOrOffset, length)
}

Buffer.poolSize = 8192 // not used by this implementation

function from (value, encodingOrOffset, length) {
  if (typeof value === 'string') {
    return fromString(value, encodingOrOffset)
  }

  if (ArrayBuffer.isView(value)) {
    return fromArrayView(value)
  }

  if (value == null) {
    throw new TypeError(
      'The first argument must be one of type string, Buffer, ArrayBuffer, Array, ' +
      'or Array-like Object. Received type ' + (typeof value)
    )
  }

  if (isInstance(value, ArrayBuffer) ||
      (value && isInstance(value.buffer, ArrayBuffer))) {
    return fromArrayBuffer(value, encodingOrOffset, length)
  }

  if (typeof SharedArrayBuffer !== 'undefined' &&
      (isInstance(value, SharedArrayBuffer) ||
      (value && isInstance(value.buffer, SharedArrayBuffer)))) {
    return fromArrayBuffer(value, encodingOrOffset, length)
  }

  if (typeof value === 'number') {
    throw new TypeError(
      'The "value" argument must not be of type number. Received type number'
    )
  }

  const valueOf = value.valueOf && value.valueOf()
  if (valueOf != null && valueOf !== value) {
    return Buffer.from(valueOf, encodingOrOffset, length)
  }

  const b = fromObject(value)
  if (b) return b

  if (typeof Symbol !== 'undefined' && Symbol.toPrimitive != null &&
      typeof value[Symbol.toPrimitive] === 'function') {
    return Buffer.from(value[Symbol.toPrimitive]('string'), encodingOrOffset, length)
  }

  throw new TypeError(
    'The first argument must be one of type string, Buffer, ArrayBuffer, Array, ' +
    'or Array-like Object. Received type ' + (typeof value)
  )
}

/**
 * Functionally equivalent to Buffer(arg, encoding) but throws a TypeError
 * if value is a number.
 * Buffer.from(str[, encoding])
 * Buffer.from(array)
 * Buffer.from(buffer)
 * Buffer.from(arrayBuffer[, byteOffset[, length]])
 **/
Buffer.from = function (value, encodingOrOffset, length) {
  return from(value, encodingOrOffset, length)
}

// Note: Change prototype *after* Buffer.from is defined to workaround Chrome bug:
// https://github.com/feross/buffer/pull/148
Object.setPrototypeOf(Buffer.prototype, Uint8Array.prototype)
Object.setPrototypeOf(Buffer, Uint8Array)

function assertSize (size) {
  if (typeof size !== 'number') {
    throw new TypeError('"size" argument must be of type number')
  } else if (size < 0) {
    throw new RangeError('The value "' + size + '" is invalid for option "size"')
  }
}

function alloc (size, fill, encoding) {
  assertSize(size)
  if (size <= 0) {
    return createBuffer(size)
  }
  if (fill !== undefined) {
    // Only pay attention to encoding if it's a string. This
    // prevents accidentally sending in a number that would
    // be interpreted as a start offset.
    return typeof encoding === 'string'
      ? createBuffer(size).fill(fill, encoding)
      : createBuffer(size).fill(fill)
  }
  return createBuffer(size)
}

/**
 * Creates a new filled Buffer instance.
 * alloc(size[, fill[, encoding]])
 **/
Buffer.alloc = function (size, fill, encoding) {
  return alloc(size, fill, encoding)
}

function allocUnsafe (size) {
  assertSize(size)
  return createBuffer(size < 0 ? 0 : checked(size) | 0)
}

/**
 * Equivalent to Buffer(num), by default creates a non-zero-filled Buffer instance.
 * */
Buffer.allocUnsafe = function (size) {
  return allocUnsafe(size)
}
/**
 * Equivalent to SlowBuffer(num), by default creates a non-zero-filled Buffer instance.
 */
Buffer.allocUnsafeSlow = function (size) {
  return allocUnsafe(size)
}

function fromString (string, encoding) {
  if (typeof encoding !== 'string' || encoding === '') {
    encoding = 'utf8'
  }

  if (!Buffer.isEncoding(encoding)) {
    throw new TypeError('Unknown encoding: ' + encoding)
  }

  const length = byteLength(string, encoding) | 0
  let buf = createBuffer(length)

  const actual = buf.write(string, encoding)

  if (actual !== length) {
    // Writing a hex string, for example, that contains invalid characters will
    // cause everything after the first invalid character to be ignored. (e.g.
    // 'abxxcd' will be treated as 'ab')
    buf = buf.slice(0, actual)
  }

  return buf
}

function fromArrayLike (array) {
  const length = array.length < 0 ? 0 : checked(array.length) | 0
  const buf = createBuffer(length)
  for (let i = 0; i < length; i += 1) {
    buf[i] = array[i] & 255
  }
  return buf
}

function fromArrayView (arrayView) {
  if (isInstance(arrayView, Uint8Array)) {
    const copy = new Uint8Array(arrayView)
    return fromArrayBuffer(copy.buffer, copy.byteOffset, copy.byteLength)
  }
  return fromArrayLike(arrayView)
}

function fromArrayBuffer (array, byteOffset, length) {
  if (byteOffset < 0 || array.byteLength < byteOffset) {
    throw new RangeError('"offset" is outside of buffer bounds')
  }

  if (array.byteLength < byteOffset + (length || 0)) {
    throw new RangeError('"length" is outside of buffer bounds')
  }

  let buf
  if (byteOffset === undefined && length === undefined) {
    buf = new Uint8Array(array)
  } else if (length === undefined) {
    buf = new Uint8Array(array, byteOffset)
  } else {
    buf = new Uint8Array(array, byteOffset, length)
  }

  // Return an augmented `Uint8Array` instance
  Object.setPrototypeOf(buf, Buffer.prototype)

  return buf
}

function fromObject (obj) {
  if (Buffer.isBuffer(obj)) {
    const len = checked(obj.length) | 0
    const buf = createBuffer(len)

    if (buf.length === 0) {
      return buf
    }

    obj.copy(buf, 0, 0, len)
    return buf
  }

  if (obj.length !== undefined) {
    if (typeof obj.length !== 'number' || numberIsNaN(obj.length)) {
      return createBuffer(0)
    }
    return fromArrayLike(obj)
  }

  if (obj.type === 'Buffer' && Array.isArray(obj.data)) {
    return fromArrayLike(obj.data)
  }
}

function checked (length) {
  // Note: cannot use `length < K_MAX_LENGTH` here because that fails when
  // length is NaN (which is otherwise coerced to zero.)
  if (length >= K_MAX_LENGTH) {
    throw new RangeError('Attempt to allocate Buffer larger than maximum ' +
                         'size: 0x' + K_MAX_LENGTH.toString(16) + ' bytes')
  }
  return length | 0
}

function SlowBuffer (length) {
  if (+length != length) { // eslint-disable-line eqeqeq
    length = 0
  }
  return Buffer.alloc(+length)
}

Buffer.isBuffer = function isBuffer (b) {
  return b != null && b._isBuffer === true &&
    b !== Buffer.prototype // so Buffer.isBuffer(Buffer.prototype) will be false
}

Buffer.compare = function compare (a, b) {
  if (isInstance(a, Uint8Array)) a = Buffer.from(a, a.offset, a.byteLength)
  if (isInstance(b, Uint8Array)) b = Buffer.from(b, b.offset, b.byteLength)
  if (!Buffer.isBuffer(a) || !Buffer.isBuffer(b)) {
    throw new TypeError(
      'The "buf1", "buf2" arguments must be one of type Buffer or Uint8Array'
    )
  }

  if (a === b) return 0

  let x = a.length
  let y = b.length

  for (let i = 0, len = Math.min(x, y); i < len; ++i) {
    if (a[i] !== b[i]) {
      x = a[i]
      y = b[i]
      break
    }
  }

  if (x < y) return -1
  if (y < x) return 1
  return 0
}

Buffer.isEncoding = function isEncoding (encoding) {
  switch (String(encoding).toLowerCase()) {
    case 'hex':
    case 'utf8':
    case 'utf-8':
    case 'ascii':
    case 'latin1':
    case 'binary':
    case 'base64':
    case 'ucs2':
    case 'ucs-2':
    case 'utf16le':
    case 'utf-16le':
      return true
    default:
      return false
  }
}

Buffer.concat = function concat (list, length) {
  if (!Array.isArray(list)) {
    throw new TypeError('"list" argument must be an Array of Buffers')
  }

  if (list.length === 0) {
    return Buffer.alloc(0)
  }

  let i
  if (length === undefined) {
    length = 0
    for (i = 0; i < list.length; ++i) {
      length += list[i].length
    }
  }

  const buffer = Buffer.allocUnsafe(length)
  let pos = 0
  for (i = 0; i < list.length; ++i) {
    let buf = list[i]
    if (isInstance(buf, Uint8Array)) {
      if (pos + buf.length > buffer.length) {
        if (!Buffer.isBuffer(buf)) buf = Buffer.from(buf)
        buf.copy(buffer, pos)
      } else {
        Uint8Array.prototype.set.call(
          buffer,
          buf,
          pos
        )
      }
    } else if (!Buffer.isBuffer(buf)) {
      throw new TypeError('"list" argument must be an Array of Buffers')
    } else {
      buf.copy(buffer, pos)
    }
    pos += buf.length
  }
  return buffer
}

function byteLength (string, encoding) {
  if (Buffer.isBuffer(string)) {
    return string.length
  }
  if (ArrayBuffer.isView(string) || isInstance(string, ArrayBuffer)) {
    return string.byteLength
  }
  if (typeof string !== 'string') {
    throw new TypeError(
      'The "string" argument must be one of type string, Buffer, or ArrayBuffer. ' +
      'Received type ' + typeof string
    )
  }

  const len = string.length
  const mustMatch = (arguments.length > 2 && arguments[2] === true)
  if (!mustMatch && len === 0) return 0

  // Use a for loop to avoid recursion
  let loweredCase = false
  for (;;) {
    switch (encoding) {
      case 'ascii':
      case 'latin1':
      case 'binary':
        return len
      case 'utf8':
      case 'utf-8':
        return utf8ToBytes(string).length
      case 'ucs2':
      case 'ucs-2':
      case 'utf16le':
      case 'utf-16le':
        return len * 2
      case 'hex':
        return len >>> 1
      case 'base64':
        return base64ToBytes(string).length
      default:
        if (loweredCase) {
          return mustMatch ? -1 : utf8ToBytes(string).length // assume utf8
        }
        encoding = ('' + encoding).toLowerCase()
        loweredCase = true
    }
  }
}
Buffer.byteLength = byteLength

function slowToString (encoding, start, end) {
  let loweredCase = false

  // No need to verify that "this.length <= MAX_UINT32" since it's a read-only
  // property of a typed array.

  // This behaves neither like String nor Uint8Array in that we set start/end
  // to their upper/lower bounds if the value passed is out of range.
  // undefined is handled specially as per ECMA-262 6th Edition,
  // Section 13.3.3.7 Runtime Semantics: KeyedBindingInitialization.
  if (start === undefined || start < 0) {
    start = 0
  }
  // Return early if start > this.length. Done here to prevent potential uint32
  // coercion fail below.
  if (start > this.length) {
    return ''
  }

  if (end === undefined || end > this.length) {
    end = this.length
  }

  if (end <= 0) {
    return ''
  }

  // Force coercion to uint32. This will also coerce falsey/NaN values to 0.
  end >>>= 0
  start >>>= 0

  if (end <= start) {
    return ''
  }

  if (!encoding) encoding = 'utf8'

  while (true) {
    switch (encoding) {
      case 'hex':
        return hexSlice(this, start, end)

      case 'utf8':
      case 'utf-8':
        return utf8Slice(this, start, end)

      case 'ascii':
        return asciiSlice(this, start, end)

      case 'latin1':
      case 'binary':
        return latin1Slice(this, start, end)

      case 'base64':
        return base64Slice(this, start, end)

      case 'ucs2':
      case 'ucs-2':
      case 'utf16le':
      case 'utf-16le':
        return utf16leSlice(this, start, end)

      default:
        if (loweredCase) throw new TypeError('Unknown encoding: ' + encoding)
        encoding = (encoding + '').toLowerCase()
        loweredCase = true
    }
  }
}

// This property is used by `Buffer.isBuffer` (and the `is-buffer` npm package)
// to detect a Buffer instance. It's not possible to use `instanceof Buffer`
// reliably in a browserify context because there could be multiple different
// copies of the 'buffer' package in use. This method works even for Buffer
// instances that were created from another copy of the `buffer` package.
// See: https://github.com/feross/buffer/issues/154
Buffer.prototype._isBuffer = true

function swap (b, n, m) {
  const i = b[n]
  b[n] = b[m]
  b[m] = i
}

Buffer.prototype.swap16 = function swap16 () {
  const len = this.length
  if (len % 2 !== 0) {
    throw new RangeError('Buffer size must be a multiple of 16-bits')
  }
  for (let i = 0; i < len; i += 2) {
    swap(this, i, i + 1)
  }
  return this
}

Buffer.prototype.swap32 = function swap32 () {
  const len = this.length
  if (len % 4 !== 0) {
    throw new RangeError('Buffer size must be a multiple of 32-bits')
  }
  for (let i = 0; i < len; i += 4) {
    swap(this, i, i + 3)
    swap(this, i + 1, i + 2)
  }
  return this
}

Buffer.prototype.swap64 = function swap64 () {
  const len = this.length
  if (len % 8 !== 0) {
    throw new RangeError('Buffer size must be a multiple of 64-bits')
  }
  for (let i = 0; i < len; i += 8) {
    swap(this, i, i + 7)
    swap(this, i + 1, i + 6)
    swap(this, i + 2, i + 5)
    swap(this, i + 3, i + 4)
  }
  return this
}

Buffer.prototype.toString = function toString () {
  const length = this.length
  if (length === 0) return ''
  if (arguments.length === 0) return utf8Slice(this, 0, length)
  return slowToString.apply(this, arguments)
}

Buffer.prototype.toLocaleString = Buffer.prototype.toString

Buffer.prototype.equals = function equals (b) {
  if (!Buffer.isBuffer(b)) throw new TypeError('Argument must be a Buffer')
  if (this === b) return true
  return Buffer.compare(this, b) === 0
}

Buffer.prototype.inspect = function inspect () {
  let str = ''
  const max = exports.INSPECT_MAX_BYTES
  str = this.toString('hex', 0, max).replace(/(.{2})/g, '$1 ').trim()
  if (this.length > max) str += ' ... '
  return '<Buffer ' + str + '>'
}
if (customInspectSymbol) {
  Buffer.prototype[customInspectSymbol] = Buffer.prototype.inspect
}

Buffer.prototype.compare = function compare (target, start, end, thisStart, thisEnd) {
  if (isInstance(target, Uint8Array)) {
    target = Buffer.from(target, target.offset, target.byteLength)
  }
  if (!Buffer.isBuffer(target)) {
    throw new TypeError(
      'The "target" argument must be one of type Buffer or Uint8Array. ' +
      'Received type ' + (typeof target)
    )
  }

  if (start === undefined) {
    start = 0
  }
  if (end === undefined) {
    end = target ? target.length : 0
  }
  if (thisStart === undefined) {
    thisStart = 0
  }
  if (thisEnd === undefined) {
    thisEnd = this.length
  }

  if (start < 0 || end > target.length || thisStart < 0 || thisEnd > this.length) {
    throw new RangeError('out of range index')
  }

  if (thisStart >= thisEnd && start >= end) {
    return 0
  }
  if (thisStart >= thisEnd) {
    return -1
  }
  if (start >= end) {
    return 1
  }

  start >>>= 0
  end >>>= 0
  thisStart >>>= 0
  thisEnd >>>= 0

  if (this === target) return 0

  let x = thisEnd - thisStart
  let y = end - start
  const len = Math.min(x, y)

  const thisCopy = this.slice(thisStart, thisEnd)
  const targetCopy = target.slice(start, end)

  for (let i = 0; i < len; ++i) {
    if (thisCopy[i] !== targetCopy[i]) {
      x = thisCopy[i]
      y = targetCopy[i]
      break
    }
  }

  if (x < y) return -1
  if (y < x) return 1
  return 0
}

// Finds either the first index of `val` in `buffer` at offset >= `byteOffset`,
// OR the last index of `val` in `buffer` at offset <= `byteOffset`.
//
// Arguments:
// - buffer - a Buffer to search
// - val - a string, Buffer, or number
// - byteOffset - an index into `buffer`; will be clamped to an int32
// - encoding - an optional encoding, relevant is val is a string
// - dir - true for indexOf, false for lastIndexOf
function bidirectionalIndexOf (buffer, val, byteOffset, encoding, dir) {
  // Empty buffer means no match
  if (buffer.length === 0) return -1

  // Normalize byteOffset
  if (typeof byteOffset === 'string') {
    encoding = byteOffset
    byteOffset = 0
  } else if (byteOffset > 0x7fffffff) {
    byteOffset = 0x7fffffff
  } else if (byteOffset < -0x80000000) {
    byteOffset = -0x80000000
  }
  byteOffset = +byteOffset // Coerce to Number.
  if (numberIsNaN(byteOffset)) {
    // byteOffset: it it's undefined, null, NaN, "foo", etc, search whole buffer
    byteOffset = dir ? 0 : (buffer.length - 1)
  }

  // Normalize byteOffset: negative offsets start from the end of the buffer
  if (byteOffset < 0) byteOffset = buffer.length + byteOffset
  if (byteOffset >= buffer.length) {
    if (dir) return -1
    else byteOffset = buffer.length - 1
  } else if (byteOffset < 0) {
    if (dir) byteOffset = 0
    else return -1
  }

  // Normalize val
  if (typeof val === 'string') {
    val = Buffer.from(val, encoding)
  }

  // Finally, search either indexOf (if dir is true) or lastIndexOf
  if (Buffer.isBuffer(val)) {
    // Special case: looking for empty string/buffer always fails
    if (val.length === 0) {
      return -1
    }
    return arrayIndexOf(buffer, val, byteOffset, encoding, dir)
  } else if (typeof val === 'number') {
    val = val & 0xFF // Search for a byte value [0-255]
    if (typeof Uint8Array.prototype.indexOf === 'function') {
      if (dir) {
        return Uint8Array.prototype.indexOf.call(buffer, val, byteOffset)
      } else {
        return Uint8Array.prototype.lastIndexOf.call(buffer, val, byteOffset)
      }
    }
    return arrayIndexOf(buffer, [val], byteOffset, encoding, dir)
  }

  throw new TypeError('val must be string, number or Buffer')
}

function arrayIndexOf (arr, val, byteOffset, encoding, dir) {
  let indexSize = 1
  let arrLength = arr.length
  let valLength = val.length

  if (encoding !== undefined) {
    encoding = String(encoding).toLowerCase()
    if (encoding === 'ucs2' || encoding === 'ucs-2' ||
        encoding === 'utf16le' || encoding === 'utf-16le') {
      if (arr.length < 2 || val.length < 2) {
        return -1
      }
      indexSize = 2
      arrLength /= 2
      valLength /= 2
      byteOffset /= 2
    }
  }

  function read (buf, i) {
    if (indexSize === 1) {
      return buf[i]
    } else {
      return buf.readUInt16BE(i * indexSize)
    }
  }

  let i
  if (dir) {
    let foundIndex = -1
    for (i = byteOffset; i < arrLength; i++) {
      if (read(arr, i) === read(val, foundIndex === -1 ? 0 : i - foundIndex)) {
        if (foundIndex === -1) foundIndex = i
        if (i - foundIndex + 1 === valLength) return foundIndex * indexSize
      } else {
        if (foundIndex !== -1) i -= i - foundIndex
        foundIndex = -1
      }
    }
  } else {
    if (byteOffset + valLength > arrLength) byteOffset = arrLength - valLength
    for (i = byteOffset; i >= 0; i--) {
      let found = true
      for (let j = 0; j < valLength; j++) {
        if (read(arr, i + j) !== read(val, j)) {
          found = false
          break
        }
      }
      if (found) return i
    }
  }

  return -1
}

Buffer.prototype.includes = function includes (val, byteOffset, encoding) {
  return this.indexOf(val, byteOffset, encoding) !== -1
}

Buffer.prototype.indexOf = function indexOf (val, byteOffset, encoding) {
  return bidirectionalIndexOf(this, val, byteOffset, encoding, true)
}

Buffer.prototype.lastIndexOf = function lastIndexOf (val, byteOffset, encoding) {
  return bidirectionalIndexOf(this, val, byteOffset, encoding, false)
}

function hexWrite (buf, string, offset, length) {
  offset = Number(offset) || 0
  const remaining = buf.length - offset
  if (!length) {
    length = remaining
  } else {
    length = Number(length)
    if (length > remaining) {
      length = remaining
    }
  }

  const strLen = string.length

  if (length > strLen / 2) {
    length = strLen / 2
  }
  let i
  for (i = 0; i < length; ++i) {
    const parsed = parseInt(string.substr(i * 2, 2), 16)
    if (numberIsNaN(parsed)) return i
    buf[offset + i] = parsed
  }
  return i
}

function utf8Write (buf, string, offset, length) {
  return blitBuffer(utf8ToBytes(string, buf.length - offset), buf, offset, length)
}

function asciiWrite (buf, string, offset, length) {
  return blitBuffer(asciiToBytes(string), buf, offset, length)
}

function base64Write (buf, string, offset, length) {
  return blitBuffer(base64ToBytes(string), buf, offset, length)
}

function ucs2Write (buf, string, offset, length) {
  return blitBuffer(utf16leToBytes(string, buf.length - offset), buf, offset, length)
}

Buffer.prototype.write = function write (string, offset, length, encoding) {
  // Buffer#write(string)
  if (offset === undefined) {
    encoding = 'utf8'
    length = this.length
    offset = 0
  // Buffer#write(string, encoding)
  } else if (length === undefined && typeof offset === 'string') {
    encoding = offset
    length = this.length
    offset = 0
  // Buffer#write(string, offset[, length][, encoding])
  } else if (isFinite(offset)) {
    offset = offset >>> 0
    if (isFinite(length)) {
      length = length >>> 0
      if (encoding === undefined) encoding = 'utf8'
    } else {
      encoding = length
      length = undefined
    }
  } else {
    throw new Error(
      'Buffer.write(string, encoding, offset[, length]) is no longer supported'
    )
  }

  const remaining = this.length - offset
  if (length === undefined || length > remaining) length = remaining

  if ((string.length > 0 && (length < 0 || offset < 0)) || offset > this.length) {
    throw new RangeError('Attempt to write outside buffer bounds')
  }

  if (!encoding) encoding = 'utf8'

  let loweredCase = false
  for (;;) {
    switch (encoding) {
      case 'hex':
        return hexWrite(this, string, offset, length)

      case 'utf8':
      case 'utf-8':
        return utf8Write(this, string, offset, length)

      case 'ascii':
      case 'latin1':
      case 'binary':
        return asciiWrite(this, string, offset, length)

      case 'base64':
        // Warning: maxLength not taken into account in base64Write
        return base64Write(this, string, offset, length)

      case 'ucs2':
      case 'ucs-2':
      case 'utf16le':
      case 'utf-16le':
        return ucs2Write(this, string, offset, length)

      default:
        if (loweredCase) throw new TypeError('Unknown encoding: ' + encoding)
        encoding = ('' + encoding).toLowerCase()
        loweredCase = true
    }
  }
}

Buffer.prototype.toJSON = function toJSON () {
  return {
    type: 'Buffer',
    data: Array.prototype.slice.call(this._arr || this, 0)
  }
}

function base64Slice (buf, start, end) {
  if (start === 0 && end === buf.length) {
    return base64.fromByteArray(buf)
  } else {
    return base64.fromByteArray(buf.slice(start, end))
  }
}

function utf8Slice (buf, start, end) {
  end = Math.min(buf.length, end)
  const res = []

  let i = start
  while (i < end) {
    const firstByte = buf[i]
    let codePoint = null
    let bytesPerSequence = (firstByte > 0xEF)
      ? 4
      : (firstByte > 0xDF)
          ? 3
          : (firstByte > 0xBF)
              ? 2
              : 1

    if (i + bytesPerSequence <= end) {
      let secondByte, thirdByte, fourthByte, tempCodePoint

      switch (bytesPerSequence) {
        case 1:
          if (firstByte < 0x80) {
            codePoint = firstByte
          }
          break
        case 2:
          secondByte = buf[i + 1]
          if ((secondByte & 0xC0) === 0x80) {
            tempCodePoint = (firstByte & 0x1F) << 0x6 | (secondByte & 0x3F)
            if (tempCodePoint > 0x7F) {
              codePoint = tempCodePoint
            }
          }
          break
        case 3:
          secondByte = buf[i + 1]
          thirdByte = buf[i + 2]
          if ((secondByte & 0xC0) === 0x80 && (thirdByte & 0xC0) === 0x80) {
            tempCodePoint = (firstByte & 0xF) << 0xC | (secondByte & 0x3F) << 0x6 | (thirdByte & 0x3F)
            if (tempCodePoint > 0x7FF && (tempCodePoint < 0xD800 || tempCodePoint > 0xDFFF)) {
              codePoint = tempCodePoint
            }
          }
          break
        case 4:
          secondByte = buf[i + 1]
          thirdByte = buf[i + 2]
          fourthByte = buf[i + 3]
          if ((secondByte & 0xC0) === 0x80 && (thirdByte & 0xC0) === 0x80 && (fourthByte & 0xC0) === 0x80) {
            tempCodePoint = (firstByte & 0xF) << 0x12 | (secondByte & 0x3F) << 0xC | (thirdByte & 0x3F) << 0x6 | (fourthByte & 0x3F)
            if (tempCodePoint > 0xFFFF && tempCodePoint < 0x110000) {
              codePoint = tempCodePoint
            }
          }
      }
    }

    if (codePoint === null) {
      // we did not generate a valid codePoint so insert a
      // replacement char (U+FFFD) and advance only 1 byte
      codePoint = 0xFFFD
      bytesPerSequence = 1
    } else if (codePoint > 0xFFFF) {
      // encode to utf16 (surrogate pair dance)
      codePoint -= 0x10000
      res.push(codePoint >>> 10 & 0x3FF | 0xD800)
      codePoint = 0xDC00 | codePoint & 0x3FF
    }

    res.push(codePoint)
    i += bytesPerSequence
  }

  return decodeCodePointsArray(res)
}

// Based on http://stackoverflow.com/a/22747272/680742, the browser with
// the lowest limit is Chrome, with 0x10000 args.
// We go 1 magnitude less, for safety
const MAX_ARGUMENTS_LENGTH = 0x1000

function decodeCodePointsArray (codePoints) {
  const len = codePoints.length
  if (len <= MAX_ARGUMENTS_LENGTH) {
    return String.fromCharCode.apply(String, codePoints) // avoid extra slice()
  }

  // Decode in chunks to avoid "call stack size exceeded".
  let res = ''
  let i = 0
  while (i < len) {
    res += String.fromCharCode.apply(
      String,
      codePoints.slice(i, i += MAX_ARGUMENTS_LENGTH)
    )
  }
  return res
}

function asciiSlice (buf, start, end) {
  let ret = ''
  end = Math.min(buf.length, end)

  for (let i = start; i < end; ++i) {
    ret += String.fromCharCode(buf[i] & 0x7F)
  }
  return ret
}

function latin1Slice (buf, start, end) {
  let ret = ''
  end = Math.min(buf.length, end)

  for (let i = start; i < end; ++i) {
    ret += String.fromCharCode(buf[i])
  }
  return ret
}

function hexSlice (buf, start, end) {
  const len = buf.length

  if (!start || start < 0) start = 0
  if (!end || end < 0 || end > len) end = len

  let out = ''
  for (let i = start; i < end; ++i) {
    out += hexSliceLookupTable[buf[i]]
  }
  return out
}

function utf16leSlice (buf, start, end) {
  const bytes = buf.slice(start, end)
  let res = ''
  // If bytes.length is odd, the last 8 bits must be ignored (same as node.js)
  for (let i = 0; i < bytes.length - 1; i += 2) {
    res += String.fromCharCode(bytes[i] + (bytes[i + 1] * 256))
  }
  return res
}

Buffer.prototype.slice = function slice (start, end) {
  const len = this.length
  start = ~~start
  end = end === undefined ? len : ~~end

  if (start < 0) {
    start += len
    if (start < 0) start = 0
  } else if (start > len) {
    start = len
  }

  if (end < 0) {
    end += len
    if (end < 0) end = 0
  } else if (end > len) {
    end = len
  }

  if (end < start) end = start

  const newBuf = this.subarray(start, end)
  // Return an augmented `Uint8Array` instance
  Object.setPrototypeOf(newBuf, Buffer.prototype)

  return newBuf
}

/*
 * Need to make sure that buffer isn't trying to write out of bounds.
 */
function checkOffset (offset, ext, length) {
  if ((offset % 1) !== 0 || offset < 0) throw new RangeError('offset is not uint')
  if (offset + ext > length) throw new RangeError('Trying to access beyond buffer length')
}

Buffer.prototype.readUintLE =
Buffer.prototype.readUIntLE = function readUIntLE (offset, byteLength, noAssert) {
  offset = offset >>> 0
  byteLength = byteLength >>> 0
  if (!noAssert) checkOffset(offset, byteLength, this.length)

  let val = this[offset]
  let mul = 1
  let i = 0
  while (++i < byteLength && (mul *= 0x100)) {
    val += this[offset + i] * mul
  }

  return val
}

Buffer.prototype.readUintBE =
Buffer.prototype.readUIntBE = function readUIntBE (offset, byteLength, noAssert) {
  offset = offset >>> 0
  byteLength = byteLength >>> 0
  if (!noAssert) {
    checkOffset(offset, byteLength, this.length)
  }

  let val = this[offset + --byteLength]
  let mul = 1
  while (byteLength > 0 && (mul *= 0x100)) {
    val += this[offset + --byteLength] * mul
  }

  return val
}

Buffer.prototype.readUint8 =
Buffer.prototype.readUInt8 = function readUInt8 (offset, noAssert) {
  offset = offset >>> 0
  if (!noAssert) checkOffset(offset, 1, this.length)
  return this[offset]
}

Buffer.prototype.readUint16LE =
Buffer.prototype.readUInt16LE = function readUInt16LE (offset, noAssert) {
  offset = offset >>> 0
  if (!noAssert) checkOffset(offset, 2, this.length)
  return this[offset] | (this[offset + 1] << 8)
}

Buffer.prototype.readUint16BE =
Buffer.prototype.readUInt16BE = function readUInt16BE (offset, noAssert) {
  offset = offset >>> 0
  if (!noAssert) checkOffset(offset, 2, this.length)
  return (this[offset] << 8) | this[offset + 1]
}

Buffer.prototype.readUint32LE =
Buffer.prototype.readUInt32LE = function readUInt32LE (offset, noAssert) {
  offset = offset >>> 0
  if (!noAssert) checkOffset(offset, 4, this.length)

  return ((this[offset]) |
      (this[offset + 1] << 8) |
      (this[offset + 2] << 16)) +
      (this[offset + 3] * 0x1000000)
}

Buffer.prototype.readUint32BE =
Buffer.prototype.readUInt32BE = function readUInt32BE (offset, noAssert) {
  offset = offset >>> 0
  if (!noAssert) checkOffset(offset, 4, this.length)

  return (this[offset] * 0x1000000) +
    ((this[offset + 1] << 16) |
    (this[offset + 2] << 8) |
    this[offset + 3])
}

Buffer.prototype.readBigUInt64LE = defineBigIntMethod(function readBigUInt64LE (offset) {
  offset = offset >>> 0
  validateNumber(offset, 'offset')
  const first = this[offset]
  const last = this[offset + 7]
  if (first === undefined || last === undefined) {
    boundsError(offset, this.length - 8)
  }

  const lo = first +
    this[++offset] * 2 ** 8 +
    this[++offset] * 2 ** 16 +
    this[++offset] * 2 ** 24

  const hi = this[++offset] +
    this[++offset] * 2 ** 8 +
    this[++offset] * 2 ** 16 +
    last * 2 ** 24

  return BigInt(lo) + (BigInt(hi) << BigInt(32))
})

Buffer.prototype.readBigUInt64BE = defineBigIntMethod(function readBigUInt64BE (offset) {
  offset = offset >>> 0
  validateNumber(offset, 'offset')
  const first = this[offset]
  const last = this[offset + 7]
  if (first === undefined || last === undefined) {
    boundsError(offset, this.length - 8)
  }

  const hi = first * 2 ** 24 +
    this[++offset] * 2 ** 16 +
    this[++offset] * 2 ** 8 +
    this[++offset]

  const lo = this[++offset] * 2 ** 24 +
    this[++offset] * 2 ** 16 +
    this[++offset] * 2 ** 8 +
    last

  return (BigInt(hi) << BigInt(32)) + BigInt(lo)
})

Buffer.prototype.readIntLE = function readIntLE (offset, byteLength, noAssert) {
  offset = offset >>> 0
  byteLength = byteLength >>> 0
  if (!noAssert) checkOffset(offset, byteLength, this.length)

  let val = this[offset]
  let mul = 1
  let i = 0
  while (++i < byteLength && (mul *= 0x100)) {
    val += this[offset + i] * mul
  }
  mul *= 0x80

  if (val >= mul) val -= Math.pow(2, 8 * byteLength)

  return val
}

Buffer.prototype.readIntBE = function readIntBE (offset, byteLength, noAssert) {
  offset = offset >>> 0
  byteLength = byteLength >>> 0
  if (!noAssert) checkOffset(offset, byteLength, this.length)

  let i = byteLength
  let mul = 1
  let val = this[offset + --i]
  while (i > 0 && (mul *= 0x100)) {
    val += this[offset + --i] * mul
  }
  mul *= 0x80

  if (val >= mul) val -= Math.pow(2, 8 * byteLength)

  return val
}

Buffer.prototype.readInt8 = function readInt8 (offset, noAssert) {
  offset = offset >>> 0
  if (!noAssert) checkOffset(offset, 1, this.length)
  if (!(this[offset] & 0x80)) return (this[offset])
  return ((0xff - this[offset] + 1) * -1)
}

Buffer.prototype.readInt16LE = function readInt16LE (offset, noAssert) {
  offset = offset >>> 0
  if (!noAssert) checkOffset(offset, 2, this.length)
  const val = this[offset] | (this[offset + 1] << 8)
  return (val & 0x8000) ? val | 0xFFFF0000 : val
}

Buffer.prototype.readInt16BE = function readInt16BE (offset, noAssert) {
  offset = offset >>> 0
  if (!noAssert) checkOffset(offset, 2, this.length)
  const val = this[offset + 1] | (this[offset] << 8)
  return (val & 0x8000) ? val | 0xFFFF0000 : val
}

Buffer.prototype.readInt32LE = function readInt32LE (offset, noAssert) {
  offset = offset >>> 0
  if (!noAssert) checkOffset(offset, 4, this.length)

  return (this[offset]) |
    (this[offset + 1] << 8) |
    (this[offset + 2] << 16) |
    (this[offset + 3] << 24)
}

Buffer.prototype.readInt32BE = function readInt32BE (offset, noAssert) {
  offset = offset >>> 0
  if (!noAssert) checkOffset(offset, 4, this.length)

  return (this[offset] << 24) |
    (this[offset + 1] << 16) |
    (this[offset + 2] << 8) |
    (this[offset + 3])
}

Buffer.prototype.readBigInt64LE = defineBigIntMethod(function readBigInt64LE (offset) {
  offset = offset >>> 0
  validateNumber(offset, 'offset')
  const first = this[offset]
  const last = this[offset + 7]
  if (first === undefined || last === undefined) {
    boundsError(offset, this.length - 8)
  }

  const val = this[offset + 4] +
    this[offset + 5] * 2 ** 8 +
    this[offset + 6] * 2 ** 16 +
    (last << 24) // Overflow

  return (BigInt(val) << BigInt(32)) +
    BigInt(first +
    this[++offset] * 2 ** 8 +
    this[++offset] * 2 ** 16 +
    this[++offset] * 2 ** 24)
})

Buffer.prototype.readBigInt64BE = defineBigIntMethod(function readBigInt64BE (offset) {
  offset = offset >>> 0
  validateNumber(offset, 'offset')
  const first = this[offset]
  const last = this[offset + 7]
  if (first === undefined || last === undefined) {
    boundsError(offset, this.length - 8)
  }

  const val = (first << 24) + // Overflow
    this[++offset] * 2 ** 16 +
    this[++offset] * 2 ** 8 +
    this[++offset]

  return (BigInt(val) << BigInt(32)) +
    BigInt(this[++offset] * 2 ** 24 +
    this[++offset] * 2 ** 16 +
    this[++offset] * 2 ** 8 +
    last)
})

Buffer.prototype.readFloatLE = function readFloatLE (offset, noAssert) {
  offset = offset >>> 0
  if (!noAssert) checkOffset(offset, 4, this.length)
  return ieee754.read(this, offset, true, 23, 4)
}

Buffer.prototype.readFloatBE = function readFloatBE (offset, noAssert) {
  offset = offset >>> 0
  if (!noAssert) checkOffset(offset, 4, this.length)
  return ieee754.read(this, offset, false, 23, 4)
}

Buffer.prototype.readDoubleLE = function readDoubleLE (offset, noAssert) {
  offset = offset >>> 0
  if (!noAssert) checkOffset(offset, 8, this.length)
  return ieee754.read(this, offset, true, 52, 8)
}

Buffer.prototype.readDoubleBE = function readDoubleBE (offset, noAssert) {
  offset = offset >>> 0
  if (!noAssert) checkOffset(offset, 8, this.length)
  return ieee754.read(this, offset, false, 52, 8)
}

function checkInt (buf, value, offset, ext, max, min) {
  if (!Buffer.isBuffer(buf)) throw new TypeError('"buffer" argument must be a Buffer instance')
  if (value > max || value < min) throw new RangeError('"value" argument is out of bounds')
  if (offset + ext > buf.length) throw new RangeError('Index out of range')
}

Buffer.prototype.writeUintLE =
Buffer.prototype.writeUIntLE = function writeUIntLE (value, offset, byteLength, noAssert) {
  value = +value
  offset = offset >>> 0
  byteLength = byteLength >>> 0
  if (!noAssert) {
    const maxBytes = Math.pow(2, 8 * byteLength) - 1
    checkInt(this, value, offset, byteLength, maxBytes, 0)
  }

  let mul = 1
  let i = 0
  this[offset] = value & 0xFF
  while (++i < byteLength && (mul *= 0x100)) {
    this[offset + i] = (value / mul) & 0xFF
  }

  return offset + byteLength
}

Buffer.prototype.writeUintBE =
Buffer.prototype.writeUIntBE = function writeUIntBE (value, offset, byteLength, noAssert) {
  value = +value
  offset = offset >>> 0
  byteLength = byteLength >>> 0
  if (!noAssert) {
    const maxBytes = Math.pow(2, 8 * byteLength) - 1
    checkInt(this, value, offset, byteLength, maxBytes, 0)
  }

  let i = byteLength - 1
  let mul = 1
  this[offset + i] = value & 0xFF
  while (--i >= 0 && (mul *= 0x100)) {
    this[offset + i] = (value / mul) & 0xFF
  }

  return offset + byteLength
}

Buffer.prototype.writeUint8 =
Buffer.prototype.writeUInt8 = function writeUInt8 (value, offset, noAssert) {
  value = +value
  offset = offset >>> 0
  if (!noAssert) checkInt(this, value, offset, 1, 0xff, 0)
  this[offset] = (value & 0xff)
  return offset + 1
}

Buffer.prototype.writeUint16LE =
Buffer.prototype.writeUInt16LE = function writeUInt16LE (value, offset, noAssert) {
  value = +value
  offset = offset >>> 0
  if (!noAssert) checkInt(this, value, offset, 2, 0xffff, 0)
  this[offset] = (value & 0xff)
  this[offset + 1] = (value >>> 8)
  return offset + 2
}

Buffer.prototype.writeUint16BE =
Buffer.prototype.writeUInt16BE = function writeUInt16BE (value, offset, noAssert) {
  value = +value
  offset = offset >>> 0
  if (!noAssert) checkInt(this, value, offset, 2, 0xffff, 0)
  this[offset] = (value >>> 8)
  this[offset + 1] = (value & 0xff)
  return offset + 2
}

Buffer.prototype.writeUint32LE =
Buffer.prototype.writeUInt32LE = function writeUInt32LE (value, offset, noAssert) {
  value = +value
  offset = offset >>> 0
  if (!noAssert) checkInt(this, value, offset, 4, 0xffffffff, 0)
  this[offset + 3] = (value >>> 24)
  this[offset + 2] = (value >>> 16)
  this[offset + 1] = (value >>> 8)
  this[offset] = (value & 0xff)
  return offset + 4
}

Buffer.prototype.writeUint32BE =
Buffer.prototype.writeUInt32BE = function writeUInt32BE (value, offset, noAssert) {
  value = +value
  offset = offset >>> 0
  if (!noAssert) checkInt(this, value, offset, 4, 0xffffffff, 0)
  this[offset] = (value >>> 24)
  this[offset + 1] = (value >>> 16)
  this[offset + 2] = (value >>> 8)
  this[offset + 3] = (value & 0xff)
  return offset + 4
}

function wrtBigUInt64LE (buf, value, offset, min, max) {
  checkIntBI(value, min, max, buf, offset, 7)

  let lo = Number(value & BigInt(0xffffffff))
  buf[offset++] = lo
  lo = lo >> 8
  buf[offset++] = lo
  lo = lo >> 8
  buf[offset++] = lo
  lo = lo >> 8
  buf[offset++] = lo
  let hi = Number(value >> BigInt(32) & BigInt(0xffffffff))
  buf[offset++] = hi
  hi = hi >> 8
  buf[offset++] = hi
  hi = hi >> 8
  buf[offset++] = hi
  hi = hi >> 8
  buf[offset++] = hi
  return offset
}

function wrtBigUInt64BE (buf, value, offset, min, max) {
  checkIntBI(value, min, max, buf, offset, 7)

  let lo = Number(value & BigInt(0xffffffff))
  buf[offset + 7] = lo
  lo = lo >> 8
  buf[offset + 6] = lo
  lo = lo >> 8
  buf[offset + 5] = lo
  lo = lo >> 8
  buf[offset + 4] = lo
  let hi = Number(value >> BigInt(32) & BigInt(0xffffffff))
  buf[offset + 3] = hi
  hi = hi >> 8
  buf[offset + 2] = hi
  hi = hi >> 8
  buf[offset + 1] = hi
  hi = hi >> 8
  buf[offset] = hi
  return offset + 8
}

Buffer.prototype.writeBigUInt64LE = defineBigIntMethod(function writeBigUInt64LE (value, offset = 0) {
  return wrtBigUInt64LE(this, value, offset, BigInt(0), BigInt('0xffffffffffffffff'))
})

Buffer.prototype.writeBigUInt64BE = defineBigIntMethod(function writeBigUInt64BE (value, offset = 0) {
  return wrtBigUInt64BE(this, value, offset, BigInt(0), BigInt('0xffffffffffffffff'))
})

Buffer.prototype.writeIntLE = function writeIntLE (value, offset, byteLength, noAssert) {
  value = +value
  offset = offset >>> 0
  if (!noAssert) {
    const limit = Math.pow(2, (8 * byteLength) - 1)

    checkInt(this, value, offset, byteLength, limit - 1, -limit)
  }

  let i = 0
  let mul = 1
  let sub = 0
  this[offset] = value & 0xFF
  while (++i < byteLength && (mul *= 0x100)) {
    if (value < 0 && sub === 0 && this[offset + i - 1] !== 0) {
      sub = 1
    }
    this[offset + i] = ((value / mul) >> 0) - sub & 0xFF
  }

  return offset + byteLength
}

Buffer.prototype.writeIntBE = function writeIntBE (value, offset, byteLength, noAssert) {
  value = +value
  offset = offset >>> 0
  if (!noAssert) {
    const limit = Math.pow(2, (8 * byteLength) - 1)

    checkInt(this, value, offset, byteLength, limit - 1, -limit)
  }

  let i = byteLength - 1
  let mul = 1
  let sub = 0
  this[offset + i] = value & 0xFF
  while (--i >= 0 && (mul *= 0x100)) {
    if (value < 0 && sub === 0 && this[offset + i + 1] !== 0) {
      sub = 1
    }
    this[offset + i] = ((value / mul) >> 0) - sub & 0xFF
  }

  return offset + byteLength
}

Buffer.prototype.writeInt8 = function writeInt8 (value, offset, noAssert) {
  value = +value
  offset = offset >>> 0
  if (!noAssert) checkInt(this, value, offset, 1, 0x7f, -0x80)
  if (value < 0) value = 0xff + value + 1
  this[offset] = (value & 0xff)
  return offset + 1
}

Buffer.prototype.writeInt16LE = function writeInt16LE (value, offset, noAssert) {
  value = +value
  offset = offset >>> 0
  if (!noAssert) checkInt(this, value, offset, 2, 0x7fff, -0x8000)
  this[offset] = (value & 0xff)
  this[offset + 1] = (value >>> 8)
  return offset + 2
}

Buffer.prototype.writeInt16BE = function writeInt16BE (value, offset, noAssert) {
  value = +value
  offset = offset >>> 0
  if (!noAssert) checkInt(this, value, offset, 2, 0x7fff, -0x8000)
  this[offset] = (value >>> 8)
  this[offset + 1] = (value & 0xff)
  return offset + 2
}

Buffer.prototype.writeInt32LE = function writeInt32LE (value, offset, noAssert) {
  value = +value
  offset = offset >>> 0
  if (!noAssert) checkInt(this, value, offset, 4, 0x7fffffff, -0x80000000)
  this[offset] = (value & 0xff)
  this[offset + 1] = (value >>> 8)
  this[offset + 2] = (value >>> 16)
  this[offset + 3] = (value >>> 24)
  return offset + 4
}

Buffer.prototype.writeInt32BE = function writeInt32BE (value, offset, noAssert) {
  value = +value
  offset = offset >>> 0
  if (!noAssert) checkInt(this, value, offset, 4, 0x7fffffff, -0x80000000)
  if (value < 0) value = 0xffffffff + value + 1
  this[offset] = (value >>> 24)
  this[offset + 1] = (value >>> 16)
  this[offset + 2] = (value >>> 8)
  this[offset + 3] = (value & 0xff)
  return offset + 4
}

Buffer.prototype.writeBigInt64LE = defineBigIntMethod(function writeBigInt64LE (value, offset = 0) {
  return wrtBigUInt64LE(this, value, offset, -BigInt('0x8000000000000000'), BigInt('0x7fffffffffffffff'))
})

Buffer.prototype.writeBigInt64BE = defineBigIntMethod(function writeBigInt64BE (value, offset = 0) {
  return wrtBigUInt64BE(this, value, offset, -BigInt('0x8000000000000000'), BigInt('0x7fffffffffffffff'))
})

function checkIEEE754 (buf, value, offset, ext, max, min) {
  if (offset + ext > buf.length) throw new RangeError('Index out of range')
  if (offset < 0) throw new RangeError('Index out of range')
}

function writeFloat (buf, value, offset, littleEndian, noAssert) {
  value = +value
  offset = offset >>> 0
  if (!noAssert) {
    checkIEEE754(buf, value, offset, 4, 3.4028234663852886e+38, -3.4028234663852886e+38)
  }
  ieee754.write(buf, value, offset, littleEndian, 23, 4)
  return offset + 4
}

Buffer.prototype.writeFloatLE = function writeFloatLE (value, offset, noAssert) {
  return writeFloat(this, value, offset, true, noAssert)
}

Buffer.prototype.writeFloatBE = function writeFloatBE (value, offset, noAssert) {
  return writeFloat(this, value, offset, false, noAssert)
}

function writeDouble (buf, value, offset, littleEndian, noAssert) {
  value = +value
  offset = offset >>> 0
  if (!noAssert) {
    checkIEEE754(buf, value, offset, 8, 1.7976931348623157E+308, -1.7976931348623157E+308)
  }
  ieee754.write(buf, value, offset, littleEndian, 52, 8)
  return offset + 8
}

Buffer.prototype.writeDoubleLE = function writeDoubleLE (value, offset, noAssert) {
  return writeDouble(this, value, offset, true, noAssert)
}

Buffer.prototype.writeDoubleBE = function writeDoubleBE (value, offset, noAssert) {
  return writeDouble(this, value, offset, false, noAssert)
}

// copy(targetBuffer, targetStart=0, sourceStart=0, sourceEnd=buffer.length)
Buffer.prototype.copy = function copy (target, targetStart, start, end) {
  if (!Buffer.isBuffer(target)) throw new TypeError('argument should be a Buffer')
  if (!start) start = 0
  if (!end && end !== 0) end = this.length
  if (targetStart >= target.length) targetStart = target.length
  if (!targetStart) targetStart = 0
  if (end > 0 && end < start) end = start

  // Copy 0 bytes; we're done
  if (end === start) return 0
  if (target.length === 0 || this.length === 0) return 0

  // Fatal error conditions
  if (targetStart < 0) {
    throw new RangeError('targetStart out of bounds')
  }
  if (start < 0 || start >= this.length) throw new RangeError('Index out of range')
  if (end < 0) throw new RangeError('sourceEnd out of bounds')

  // Are we oob?
  if (end > this.length) end = this.length
  if (target.length - targetStart < end - start) {
    end = target.length - targetStart + start
  }

  const len = end - start

  if (this === target && typeof Uint8Array.prototype.copyWithin === 'function') {
    // Use built-in when available, missing from IE11
    this.copyWithin(targetStart, start, end)
  } else {
    Uint8Array.prototype.set.call(
      target,
      this.subarray(start, end),
      targetStart
    )
  }

  return len
}

// Usage:
//    buffer.fill(number[, offset[, end]])
//    buffer.fill(buffer[, offset[, end]])
//    buffer.fill(string[, offset[, end]][, encoding])
Buffer.prototype.fill = function fill (val, start, end, encoding) {
  // Handle string cases:
  if (typeof val === 'string') {
    if (typeof start === 'string') {
      encoding = start
      start = 0
      end = this.length
    } else if (typeof end === 'string') {
      encoding = end
      end = this.length
    }
    if (encoding !== undefined && typeof encoding !== 'string') {
      throw new TypeError('encoding must be a string')
    }
    if (typeof encoding === 'string' && !Buffer.isEncoding(encoding)) {
      throw new TypeError('Unknown encoding: ' + encoding)
    }
    if (val.length === 1) {
      const code = val.charCodeAt(0)
      if ((encoding === 'utf8' && code < 128) ||
          encoding === 'latin1') {
        // Fast path: If `val` fits into a single byte, use that numeric value.
        val = code
      }
    }
  } else if (typeof val === 'number') {
    val = val & 255
  } else if (typeof val === 'boolean') {
    val = Number(val)
  }

  // Invalid ranges are not set to a default, so can range check early.
  if (start < 0 || this.length < start || this.length < end) {
    throw new RangeError('Out of range index')
  }

  if (end <= start) {
    return this
  }

  start = start >>> 0
  end = end === undefined ? this.length : end >>> 0

  if (!val) val = 0

  let i
  if (typeof val === 'number') {
    for (i = start; i < end; ++i) {
      this[i] = val
    }
  } else {
    const bytes = Buffer.isBuffer(val)
      ? val
      : Buffer.from(val, encoding)
    const len = bytes.length
    if (len === 0) {
      throw new TypeError('The value "' + val +
        '" is invalid for argument "value"')
    }
    for (i = 0; i < end - start; ++i) {
      this[i + start] = bytes[i % len]
    }
  }

  return this
}

// CUSTOM ERRORS
// =============

// Simplified versions from Node, changed for Buffer-only usage
const errors = {}
function E (sym, getMessage, Base) {
  errors[sym] = class NodeError extends Base {
    constructor () {
      super()

      Object.defineProperty(this, 'message', {
        value: getMessage.apply(this, arguments),
        writable: true,
        configurable: true
      })

      // Add the error code to the name to include it in the stack trace.
      this.name = `${this.name} [${sym}]`
      // Access the stack to generate the error message including the error code
      // from the name.
      this.stack // eslint-disable-line no-unused-expressions
      // Reset the name to the actual name.
      delete this.name
    }

    get code () {
      return sym
    }

    set code (value) {
      Object.defineProperty(this, 'code', {
        configurable: true,
        enumerable: true,
        value,
        writable: true
      })
    }

    toString () {
      return `${this.name} [${sym}]: ${this.message}`
    }
  }
}

E('ERR_BUFFER_OUT_OF_BOUNDS',
  function (name) {
    if (name) {
      return `${name} is outside of buffer bounds`
    }

    return 'Attempt to access memory outside buffer bounds'
  }, RangeError)
E('ERR_INVALID_ARG_TYPE',
  function (name, actual) {
    return `The "${name}" argument must be of type number. Received type ${typeof actual}`
  }, TypeError)
E('ERR_OUT_OF_RANGE',
  function (str, range, input) {
    let msg = `The value of "${str}" is out of range.`
    let received = input
    if (Number.isInteger(input) && Math.abs(input) > 2 ** 32) {
      received = addNumericalSeparator(String(input))
    } else if (typeof input === 'bigint') {
      received = String(input)
      if (input > BigInt(2) ** BigInt(32) || input < -(BigInt(2) ** BigInt(32))) {
        received = addNumericalSeparator(received)
      }
      received += 'n'
    }
    msg += ` It must be ${range}. Received ${received}`
    return msg
  }, RangeError)

function addNumericalSeparator (val) {
  let res = ''
  let i = val.length
  const start = val[0] === '-' ? 1 : 0
  for (; i >= start + 4; i -= 3) {
    res = `_${val.slice(i - 3, i)}${res}`
  }
  return `${val.slice(0, i)}${res}`
}

// CHECK FUNCTIONS
// ===============

function checkBounds (buf, offset, byteLength) {
  validateNumber(offset, 'offset')
  if (buf[offset] === undefined || buf[offset + byteLength] === undefined) {
    boundsError(offset, buf.length - (byteLength + 1))
  }
}

function checkIntBI (value, min, max, buf, offset, byteLength) {
  if (value > max || value < min) {
    const n = typeof min === 'bigint' ? 'n' : ''
    let range
    if (byteLength > 3) {
      if (min === 0 || min === BigInt(0)) {
        range = `>= 0${n} and < 2${n} ** ${(byteLength + 1) * 8}${n}`
      } else {
        range = `>= -(2${n} ** ${(byteLength + 1) * 8 - 1}${n}) and < 2 ** ` +
                `${(byteLength + 1) * 8 - 1}${n}`
      }
    } else {
      range = `>= ${min}${n} and <= ${max}${n}`
    }
    throw new errors.ERR_OUT_OF_RANGE('value', range, value)
  }
  checkBounds(buf, offset, byteLength)
}

function validateNumber (value, name) {
  if (typeof value !== 'number') {
    throw new errors.ERR_INVALID_ARG_TYPE(name, 'number', value)
  }
}

function boundsError (value, length, type) {
  if (Math.floor(value) !== value) {
    validateNumber(value, type)
    throw new errors.ERR_OUT_OF_RANGE(type || 'offset', 'an integer', value)
  }

  if (length < 0) {
    throw new errors.ERR_BUFFER_OUT_OF_BOUNDS()
  }

  throw new errors.ERR_OUT_OF_RANGE(type || 'offset',
                                    `>= ${type ? 1 : 0} and <= ${length}`,
                                    value)
}

// HELPER FUNCTIONS
// ================

const INVALID_BASE64_RE = /[^+/0-9A-Za-z-_]/g

function base64clean (str) {
  // Node takes equal signs as end of the Base64 encoding
  str = str.split('=')[0]
  // Node strips out invalid characters like \n and \t from the string, base64-js does not
  str = str.trim().replace(INVALID_BASE64_RE, '')
  // Node converts strings with length < 2 to ''
  if (str.length < 2) return ''
  // Node allows for non-padded base64 strings (missing trailing ===), base64-js does not
  while (str.length % 4 !== 0) {
    str = str + '='
  }
  return str
}

function utf8ToBytes (string, units) {
  units = units || Infinity
  let codePoint
  const length = string.length
  let leadSurrogate = null
  const bytes = []

  for (let i = 0; i < length; ++i) {
    codePoint = string.charCodeAt(i)

    // is surrogate component
    if (codePoint > 0xD7FF && codePoint < 0xE000) {
      // last char was a lead
      if (!leadSurrogate) {
        // no lead yet
        if (codePoint > 0xDBFF) {
          // unexpected trail
          if ((units -= 3) > -1) bytes.push(0xEF, 0xBF, 0xBD)
          continue
        } else if (i + 1 === length) {
          // unpaired lead
          if ((units -= 3) > -1) bytes.push(0xEF, 0xBF, 0xBD)
          continue
        }

        // valid lead
        leadSurrogate = codePoint

        continue
      }

      // 2 leads in a row
      if (codePoint < 0xDC00) {
        if ((units -= 3) > -1) bytes.push(0xEF, 0xBF, 0xBD)
        leadSurrogate = codePoint
        continue
      }

      // valid surrogate pair
      codePoint = (leadSurrogate - 0xD800 << 10 | codePoint - 0xDC00) + 0x10000
    } else if (leadSurrogate) {
      // valid bmp char, but last char was a lead
      if ((units -= 3) > -1) bytes.push(0xEF, 0xBF, 0xBD)
    }

    leadSurrogate = null

    // encode utf8
    if (codePoint < 0x80) {
      if ((units -= 1) < 0) break
      bytes.push(codePoint)
    } else if (codePoint < 0x800) {
      if ((units -= 2) < 0) break
      bytes.push(
        codePoint >> 0x6 | 0xC0,
        codePoint & 0x3F | 0x80
      )
    } else if (codePoint < 0x10000) {
      if ((units -= 3) < 0) break
      bytes.push(
        codePoint >> 0xC | 0xE0,
        codePoint >> 0x6 & 0x3F | 0x80,
        codePoint & 0x3F | 0x80
      )
    } else if (codePoint < 0x110000) {
      if ((units -= 4) < 0) break
      bytes.push(
        codePoint >> 0x12 | 0xF0,
        codePoint >> 0xC & 0x3F | 0x80,
        codePoint >> 0x6 & 0x3F | 0x80,
        codePoint & 0x3F | 0x80
      )
    } else {
      throw new Error('Invalid code point')
    }
  }

  return bytes
}

function asciiToBytes (str) {
  const byteArray = []
  for (let i = 0; i < str.length; ++i) {
    // Node's code seems to be doing this and not & 0x7F..
    byteArray.push(str.charCodeAt(i) & 0xFF)
  }
  return byteArray
}

function utf16leToBytes (str, units) {
  let c, hi, lo
  const byteArray = []
  for (let i = 0; i < str.length; ++i) {
    if ((units -= 2) < 0) break

    c = str.charCodeAt(i)
    hi = c >> 8
    lo = c % 256
    byteArray.push(lo)
    byteArray.push(hi)
  }

  return byteArray
}

function base64ToBytes (str) {
  return base64.toByteArray(base64clean(str))
}

function blitBuffer (src, dst, offset, length) {
  let i
  for (i = 0; i < length; ++i) {
    if ((i + offset >= dst.length) || (i >= src.length)) break
    dst[i + offset] = src[i]
  }
  return i
}

// ArrayBuffer or Uint8Array objects from other contexts (i.e. iframes) do not pass
// the `instanceof` check but they should be treated as of that type.
// See: https://github.com/feross/buffer/issues/166
function isInstance (obj, type) {
  return obj instanceof type ||
    (obj != null && obj.constructor != null && obj.constructor.name != null &&
      obj.constructor.name === type.name)
}
function numberIsNaN (obj) {
  // For IE11 support
  return obj !== obj // eslint-disable-line no-self-compare
}

// Create lookup table for `toString('hex')`
// See: https://github.com/feross/buffer/issues/219
const hexSliceLookupTable = (function () {
  const alphabet = '0123456789abcdef'
  const table = new Array(256)
  for (let i = 0; i < 16; ++i) {
    const i16 = i * 16
    for (let j = 0; j < 16; ++j) {
      table[i16 + j] = alphabet[i] + alphabet[j]
    }
  }
  return table
})()

// Return not function with Error if BigInt not supported
function defineBigIntMethod (fn) {
  return typeof BigInt === 'undefined' ? BufferBigIntNotDefined : fn
}

function BufferBigIntNotDefined () {
  throw new Error('BigInt not supported')
}


/***/ }),

/***/ "./node_modules/builtin-status-codes/browser.js":
/*!******************************************************!*\
  !*** ./node_modules/builtin-status-codes/browser.js ***!
  \******************************************************/
/***/ ((module) => {

module.exports = {
  "100": "Continue",
  "101": "Switching Protocols",
  "102": "Processing",
  "200": "OK",
  "201": "Created",
  "202": "Accepted",
  "203": "Non-Authoritative Information",
  "204": "No Content",
  "205": "Reset Content",
  "206": "Partial Content",
  "207": "Multi-Status",
  "208": "Already Reported",
  "226": "IM Used",
  "300": "Multiple Choices",
  "301": "Moved Permanently",
  "302": "Found",
  "303": "See Other",
  "304": "Not Modified",
  "305": "Use Proxy",
  "307": "Temporary Redirect",
  "308": "Permanent Redirect",
  "400": "Bad Request",
  "401": "Unauthorized",
  "402": "Payment Required",
  "403": "Forbidden",
  "404": "Not Found",
  "405": "Method Not Allowed",
  "406": "Not Acceptable",
  "407": "Proxy Authentication Required",
  "408": "Request Timeout",
  "409": "Conflict",
  "410": "Gone",
  "411": "Length Required",
  "412": "Precondition Failed",
  "413": "Payload Too Large",
  "414": "URI Too Long",
  "415": "Unsupported Media Type",
  "416": "Range Not Satisfiable",
  "417": "Expectation Failed",
  "418": "I'm a teapot",
  "421": "Misdirected Request",
  "422": "Unprocessable Entity",
  "423": "Locked",
  "424": "Failed Dependency",
  "425": "Unordered Collection",
  "426": "Upgrade Required",
  "428": "Precondition Required",
  "429": "Too Many Requests",
  "431": "Request Header Fields Too Large",
  "451": "Unavailable For Legal Reasons",
  "500": "Internal Server Error",
  "501": "Not Implemented",
  "502": "Bad Gateway",
  "503": "Service Unavailable",
  "504": "Gateway Timeout",
  "505": "HTTP Version Not Supported",
  "506": "Variant Also Negotiates",
  "507": "Insufficient Storage",
  "508": "Loop Detected",
  "509": "Bandwidth Limit Exceeded",
  "510": "Not Extended",
  "511": "Network Authentication Required"
}


/***/ }),

/***/ "./node_modules/call-bind/callBound.js":
/*!*********************************************!*\
  !*** ./node_modules/call-bind/callBound.js ***!
  \*********************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var GetIntrinsic = __webpack_require__(/*! get-intrinsic */ "./node_modules/get-intrinsic/index.js");

var callBind = __webpack_require__(/*! ./ */ "./node_modules/call-bind/index.js");

var $indexOf = callBind(GetIntrinsic('String.prototype.indexOf'));

module.exports = function callBoundIntrinsic(name, allowMissing) {
	var intrinsic = GetIntrinsic(name, !!allowMissing);
	if (typeof intrinsic === 'function' && $indexOf(name, '.prototype.') > -1) {
		return callBind(intrinsic);
	}
	return intrinsic;
};


/***/ }),

/***/ "./node_modules/call-bind/index.js":
/*!*****************************************!*\
  !*** ./node_modules/call-bind/index.js ***!
  \*****************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var bind = __webpack_require__(/*! function-bind */ "./node_modules/function-bind/index.js");
var GetIntrinsic = __webpack_require__(/*! get-intrinsic */ "./node_modules/get-intrinsic/index.js");
var setFunctionLength = __webpack_require__(/*! set-function-length */ "./node_modules/set-function-length/index.js");

var $TypeError = GetIntrinsic('%TypeError%');
var $apply = GetIntrinsic('%Function.prototype.apply%');
var $call = GetIntrinsic('%Function.prototype.call%');
var $reflectApply = GetIntrinsic('%Reflect.apply%', true) || bind.call($call, $apply);

var $defineProperty = GetIntrinsic('%Object.defineProperty%', true);
var $max = GetIntrinsic('%Math.max%');

if ($defineProperty) {
	try {
		$defineProperty({}, 'a', { value: 1 });
	} catch (e) {
		// IE 8 has a broken defineProperty
		$defineProperty = null;
	}
}

module.exports = function callBind(originalFunction) {
	if (typeof originalFunction !== 'function') {
		throw new $TypeError('a function is required');
	}
	var func = $reflectApply(bind, $call, arguments);
	return setFunctionLength(
		func,
		1 + $max(0, originalFunction.length - (arguments.length - 1)),
		true
	);
};

var applyBind = function applyBind() {
	return $reflectApply(bind, $apply, arguments);
};

if ($defineProperty) {
	$defineProperty(module.exports, 'apply', { value: applyBind });
} else {
	module.exports.apply = applyBind;
}


/***/ }),

/***/ "./node_modules/define-data-property/index.js":
/*!****************************************************!*\
  !*** ./node_modules/define-data-property/index.js ***!
  \****************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var hasPropertyDescriptors = __webpack_require__(/*! has-property-descriptors */ "./node_modules/has-property-descriptors/index.js")();

var GetIntrinsic = __webpack_require__(/*! get-intrinsic */ "./node_modules/get-intrinsic/index.js");

var $defineProperty = hasPropertyDescriptors && GetIntrinsic('%Object.defineProperty%', true);
if ($defineProperty) {
	try {
		$defineProperty({}, 'a', { value: 1 });
	} catch (e) {
		// IE 8 has a broken defineProperty
		$defineProperty = false;
	}
}

var $SyntaxError = GetIntrinsic('%SyntaxError%');
var $TypeError = GetIntrinsic('%TypeError%');

var gopd = __webpack_require__(/*! gopd */ "./node_modules/gopd/index.js");

/** @type {(obj: Record<PropertyKey, unknown>, property: PropertyKey, value: unknown, nonEnumerable?: boolean | null, nonWritable?: boolean | null, nonConfigurable?: boolean | null, loose?: boolean) => void} */
module.exports = function defineDataProperty(
	obj,
	property,
	value
) {
	if (!obj || (typeof obj !== 'object' && typeof obj !== 'function')) {
		throw new $TypeError('`obj` must be an object or a function`');
	}
	if (typeof property !== 'string' && typeof property !== 'symbol') {
		throw new $TypeError('`property` must be a string or a symbol`');
	}
	if (arguments.length > 3 && typeof arguments[3] !== 'boolean' && arguments[3] !== null) {
		throw new $TypeError('`nonEnumerable`, if provided, must be a boolean or null');
	}
	if (arguments.length > 4 && typeof arguments[4] !== 'boolean' && arguments[4] !== null) {
		throw new $TypeError('`nonWritable`, if provided, must be a boolean or null');
	}
	if (arguments.length > 5 && typeof arguments[5] !== 'boolean' && arguments[5] !== null) {
		throw new $TypeError('`nonConfigurable`, if provided, must be a boolean or null');
	}
	if (arguments.length > 6 && typeof arguments[6] !== 'boolean') {
		throw new $TypeError('`loose`, if provided, must be a boolean');
	}

	var nonEnumerable = arguments.length > 3 ? arguments[3] : null;
	var nonWritable = arguments.length > 4 ? arguments[4] : null;
	var nonConfigurable = arguments.length > 5 ? arguments[5] : null;
	var loose = arguments.length > 6 ? arguments[6] : false;

	/* @type {false | TypedPropertyDescriptor<unknown>} */
	var desc = !!gopd && gopd(obj, property);

	if ($defineProperty) {
		$defineProperty(obj, property, {
			configurable: nonConfigurable === null && desc ? desc.configurable : !nonConfigurable,
			enumerable: nonEnumerable === null && desc ? desc.enumerable : !nonEnumerable,
			value: value,
			writable: nonWritable === null && desc ? desc.writable : !nonWritable
		});
	} else if (loose || (!nonEnumerable && !nonWritable && !nonConfigurable)) {
		// must fall back to [[Set]], and was not explicitly asked to make non-enumerable, non-writable, or non-configurable
		obj[property] = value; // eslint-disable-line no-param-reassign
	} else {
		throw new $SyntaxError('This environment does not support defining a property as non-configurable, non-writable, or non-enumerable.');
	}
};


/***/ }),

/***/ "./node_modules/events/events.js":
/*!***************************************!*\
  !*** ./node_modules/events/events.js ***!
  \***************************************/
/***/ ((module) => {

"use strict";
// Copyright Joyent, Inc. and other Node contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to permit
// persons to whom the Software is furnished to do so, subject to the
// following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN
// NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
// USE OR OTHER DEALINGS IN THE SOFTWARE.



var R = typeof Reflect === 'object' ? Reflect : null
var ReflectApply = R && typeof R.apply === 'function'
  ? R.apply
  : function ReflectApply(target, receiver, args) {
    return Function.prototype.apply.call(target, receiver, args);
  }

var ReflectOwnKeys
if (R && typeof R.ownKeys === 'function') {
  ReflectOwnKeys = R.ownKeys
} else if (Object.getOwnPropertySymbols) {
  ReflectOwnKeys = function ReflectOwnKeys(target) {
    return Object.getOwnPropertyNames(target)
      .concat(Object.getOwnPropertySymbols(target));
  };
} else {
  ReflectOwnKeys = function ReflectOwnKeys(target) {
    return Object.getOwnPropertyNames(target);
  };
}

function ProcessEmitWarning(warning) {
  if (console && console.warn) console.warn(warning);
}

var NumberIsNaN = Number.isNaN || function NumberIsNaN(value) {
  return value !== value;
}

function EventEmitter() {
  EventEmitter.init.call(this);
}
module.exports = EventEmitter;
module.exports.once = once;

// Backwards-compat with node 0.10.x
EventEmitter.EventEmitter = EventEmitter;

EventEmitter.prototype._events = undefined;
EventEmitter.prototype._eventsCount = 0;
EventEmitter.prototype._maxListeners = undefined;

// By default EventEmitters will print a warning if more than 10 listeners are
// added to it. This is a useful default which helps finding memory leaks.
var defaultMaxListeners = 10;

function checkListener(listener) {
  if (typeof listener !== 'function') {
    throw new TypeError('The "listener" argument must be of type Function. Received type ' + typeof listener);
  }
}

Object.defineProperty(EventEmitter, 'defaultMaxListeners', {
  enumerable: true,
  get: function() {
    return defaultMaxListeners;
  },
  set: function(arg) {
    if (typeof arg !== 'number' || arg < 0 || NumberIsNaN(arg)) {
      throw new RangeError('The value of "defaultMaxListeners" is out of range. It must be a non-negative number. Received ' + arg + '.');
    }
    defaultMaxListeners = arg;
  }
});

EventEmitter.init = function() {

  if (this._events === undefined ||
      this._events === Object.getPrototypeOf(this)._events) {
    this._events = Object.create(null);
    this._eventsCount = 0;
  }

  this._maxListeners = this._maxListeners || undefined;
};

// Obviously not all Emitters should be limited to 10. This function allows
// that to be increased. Set to zero for unlimited.
EventEmitter.prototype.setMaxListeners = function setMaxListeners(n) {
  if (typeof n !== 'number' || n < 0 || NumberIsNaN(n)) {
    throw new RangeError('The value of "n" is out of range. It must be a non-negative number. Received ' + n + '.');
  }
  this._maxListeners = n;
  return this;
};

function _getMaxListeners(that) {
  if (that._maxListeners === undefined)
    return EventEmitter.defaultMaxListeners;
  return that._maxListeners;
}

EventEmitter.prototype.getMaxListeners = function getMaxListeners() {
  return _getMaxListeners(this);
};

EventEmitter.prototype.emit = function emit(type) {
  var args = [];
  for (var i = 1; i < arguments.length; i++) args.push(arguments[i]);
  var doError = (type === 'error');

  var events = this._events;
  if (events !== undefined)
    doError = (doError && events.error === undefined);
  else if (!doError)
    return false;

  // If there is no 'error' event listener then throw.
  if (doError) {
    var er;
    if (args.length > 0)
      er = args[0];
    if (er instanceof Error) {
      // Note: The comments on the `throw` lines are intentional, they show
      // up in Node's output if this results in an unhandled exception.
      throw er; // Unhandled 'error' event
    }
    // At least give some kind of context to the user
    var err = new Error('Unhandled error.' + (er ? ' (' + er.message + ')' : ''));
    err.context = er;
    throw err; // Unhandled 'error' event
  }

  var handler = events[type];

  if (handler === undefined)
    return false;

  if (typeof handler === 'function') {
    ReflectApply(handler, this, args);
  } else {
    var len = handler.length;
    var listeners = arrayClone(handler, len);
    for (var i = 0; i < len; ++i)
      ReflectApply(listeners[i], this, args);
  }

  return true;
};

function _addListener(target, type, listener, prepend) {
  var m;
  var events;
  var existing;

  checkListener(listener);

  events = target._events;
  if (events === undefined) {
    events = target._events = Object.create(null);
    target._eventsCount = 0;
  } else {
    // To avoid recursion in the case that type === "newListener"! Before
    // adding it to the listeners, first emit "newListener".
    if (events.newListener !== undefined) {
      target.emit('newListener', type,
                  listener.listener ? listener.listener : listener);

      // Re-assign `events` because a newListener handler could have caused the
      // this._events to be assigned to a new object
      events = target._events;
    }
    existing = events[type];
  }

  if (existing === undefined) {
    // Optimize the case of one listener. Don't need the extra array object.
    existing = events[type] = listener;
    ++target._eventsCount;
  } else {
    if (typeof existing === 'function') {
      // Adding the second element, need to change to array.
      existing = events[type] =
        prepend ? [listener, existing] : [existing, listener];
      // If we've already got an array, just append.
    } else if (prepend) {
      existing.unshift(listener);
    } else {
      existing.push(listener);
    }

    // Check for listener leak
    m = _getMaxListeners(target);
    if (m > 0 && existing.length > m && !existing.warned) {
      existing.warned = true;
      // No error code for this since it is a Warning
      // eslint-disable-next-line no-restricted-syntax
      var w = new Error('Possible EventEmitter memory leak detected. ' +
                          existing.length + ' ' + String(type) + ' listeners ' +
                          'added. Use emitter.setMaxListeners() to ' +
                          'increase limit');
      w.name = 'MaxListenersExceededWarning';
      w.emitter = target;
      w.type = type;
      w.count = existing.length;
      ProcessEmitWarning(w);
    }
  }

  return target;
}

EventEmitter.prototype.addListener = function addListener(type, listener) {
  return _addListener(this, type, listener, false);
};

EventEmitter.prototype.on = EventEmitter.prototype.addListener;

EventEmitter.prototype.prependListener =
    function prependListener(type, listener) {
      return _addListener(this, type, listener, true);
    };

function onceWrapper() {
  if (!this.fired) {
    this.target.removeListener(this.type, this.wrapFn);
    this.fired = true;
    if (arguments.length === 0)
      return this.listener.call(this.target);
    return this.listener.apply(this.target, arguments);
  }
}

function _onceWrap(target, type, listener) {
  var state = { fired: false, wrapFn: undefined, target: target, type: type, listener: listener };
  var wrapped = onceWrapper.bind(state);
  wrapped.listener = listener;
  state.wrapFn = wrapped;
  return wrapped;
}

EventEmitter.prototype.once = function once(type, listener) {
  checkListener(listener);
  this.on(type, _onceWrap(this, type, listener));
  return this;
};

EventEmitter.prototype.prependOnceListener =
    function prependOnceListener(type, listener) {
      checkListener(listener);
      this.prependListener(type, _onceWrap(this, type, listener));
      return this;
    };

// Emits a 'removeListener' event if and only if the listener was removed.
EventEmitter.prototype.removeListener =
    function removeListener(type, listener) {
      var list, events, position, i, originalListener;

      checkListener(listener);

      events = this._events;
      if (events === undefined)
        return this;

      list = events[type];
      if (list === undefined)
        return this;

      if (list === listener || list.listener === listener) {
        if (--this._eventsCount === 0)
          this._events = Object.create(null);
        else {
          delete events[type];
          if (events.removeListener)
            this.emit('removeListener', type, list.listener || listener);
        }
      } else if (typeof list !== 'function') {
        position = -1;

        for (i = list.length - 1; i >= 0; i--) {
          if (list[i] === listener || list[i].listener === listener) {
            originalListener = list[i].listener;
            position = i;
            break;
          }
        }

        if (position < 0)
          return this;

        if (position === 0)
          list.shift();
        else {
          spliceOne(list, position);
        }

        if (list.length === 1)
          events[type] = list[0];

        if (events.removeListener !== undefined)
          this.emit('removeListener', type, originalListener || listener);
      }

      return this;
    };

EventEmitter.prototype.off = EventEmitter.prototype.removeListener;

EventEmitter.prototype.removeAllListeners =
    function removeAllListeners(type) {
      var listeners, events, i;

      events = this._events;
      if (events === undefined)
        return this;

      // not listening for removeListener, no need to emit
      if (events.removeListener === undefined) {
        if (arguments.length === 0) {
          this._events = Object.create(null);
          this._eventsCount = 0;
        } else if (events[type] !== undefined) {
          if (--this._eventsCount === 0)
            this._events = Object.create(null);
          else
            delete events[type];
        }
        return this;
      }

      // emit removeListener for all listeners on all events
      if (arguments.length === 0) {
        var keys = Object.keys(events);
        var key;
        for (i = 0; i < keys.length; ++i) {
          key = keys[i];
          if (key === 'removeListener') continue;
          this.removeAllListeners(key);
        }
        this.removeAllListeners('removeListener');
        this._events = Object.create(null);
        this._eventsCount = 0;
        return this;
      }

      listeners = events[type];

      if (typeof listeners === 'function') {
        this.removeListener(type, listeners);
      } else if (listeners !== undefined) {
        // LIFO order
        for (i = listeners.length - 1; i >= 0; i--) {
          this.removeListener(type, listeners[i]);
        }
      }

      return this;
    };

function _listeners(target, type, unwrap) {
  var events = target._events;

  if (events === undefined)
    return [];

  var evlistener = events[type];
  if (evlistener === undefined)
    return [];

  if (typeof evlistener === 'function')
    return unwrap ? [evlistener.listener || evlistener] : [evlistener];

  return unwrap ?
    unwrapListeners(evlistener) : arrayClone(evlistener, evlistener.length);
}

EventEmitter.prototype.listeners = function listeners(type) {
  return _listeners(this, type, true);
};

EventEmitter.prototype.rawListeners = function rawListeners(type) {
  return _listeners(this, type, false);
};

EventEmitter.listenerCount = function(emitter, type) {
  if (typeof emitter.listenerCount === 'function') {
    return emitter.listenerCount(type);
  } else {
    return listenerCount.call(emitter, type);
  }
};

EventEmitter.prototype.listenerCount = listenerCount;
function listenerCount(type) {
  var events = this._events;

  if (events !== undefined) {
    var evlistener = events[type];

    if (typeof evlistener === 'function') {
      return 1;
    } else if (evlistener !== undefined) {
      return evlistener.length;
    }
  }

  return 0;
}

EventEmitter.prototype.eventNames = function eventNames() {
  return this._eventsCount > 0 ? ReflectOwnKeys(this._events) : [];
};

function arrayClone(arr, n) {
  var copy = new Array(n);
  for (var i = 0; i < n; ++i)
    copy[i] = arr[i];
  return copy;
}

function spliceOne(list, index) {
  for (; index + 1 < list.length; index++)
    list[index] = list[index + 1];
  list.pop();
}

function unwrapListeners(arr) {
  var ret = new Array(arr.length);
  for (var i = 0; i < ret.length; ++i) {
    ret[i] = arr[i].listener || arr[i];
  }
  return ret;
}

function once(emitter, name) {
  return new Promise(function (resolve, reject) {
    function errorListener(err) {
      emitter.removeListener(name, resolver);
      reject(err);
    }

    function resolver() {
      if (typeof emitter.removeListener === 'function') {
        emitter.removeListener('error', errorListener);
      }
      resolve([].slice.call(arguments));
    };

    eventTargetAgnosticAddListener(emitter, name, resolver, { once: true });
    if (name !== 'error') {
      addErrorHandlerIfEventEmitter(emitter, errorListener, { once: true });
    }
  });
}

function addErrorHandlerIfEventEmitter(emitter, handler, flags) {
  if (typeof emitter.on === 'function') {
    eventTargetAgnosticAddListener(emitter, 'error', handler, flags);
  }
}

function eventTargetAgnosticAddListener(emitter, name, listener, flags) {
  if (typeof emitter.on === 'function') {
    if (flags.once) {
      emitter.once(name, listener);
    } else {
      emitter.on(name, listener);
    }
  } else if (typeof emitter.addEventListener === 'function') {
    // EventTarget does not have `error` event semantics like Node
    // EventEmitters, we do not listen for `error` events here.
    emitter.addEventListener(name, function wrapListener(arg) {
      // IE does not have builtin `{ once: true }` support so we
      // have to do it manually.
      if (flags.once) {
        emitter.removeEventListener(name, wrapListener);
      }
      listener(arg);
    });
  } else {
    throw new TypeError('The "emitter" argument must be of type EventEmitter. Received type ' + typeof emitter);
  }
}


/***/ }),

/***/ "./node_modules/function-bind/implementation.js":
/*!******************************************************!*\
  !*** ./node_modules/function-bind/implementation.js ***!
  \******************************************************/
/***/ ((module) => {

"use strict";


/* eslint no-invalid-this: 1 */

var ERROR_MESSAGE = 'Function.prototype.bind called on incompatible ';
var toStr = Object.prototype.toString;
var max = Math.max;
var funcType = '[object Function]';

var concatty = function concatty(a, b) {
    var arr = [];

    for (var i = 0; i < a.length; i += 1) {
        arr[i] = a[i];
    }
    for (var j = 0; j < b.length; j += 1) {
        arr[j + a.length] = b[j];
    }

    return arr;
};

var slicy = function slicy(arrLike, offset) {
    var arr = [];
    for (var i = offset || 0, j = 0; i < arrLike.length; i += 1, j += 1) {
        arr[j] = arrLike[i];
    }
    return arr;
};

var joiny = function (arr, joiner) {
    var str = '';
    for (var i = 0; i < arr.length; i += 1) {
        str += arr[i];
        if (i + 1 < arr.length) {
            str += joiner;
        }
    }
    return str;
};

module.exports = function bind(that) {
    var target = this;
    if (typeof target !== 'function' || toStr.apply(target) !== funcType) {
        throw new TypeError(ERROR_MESSAGE + target);
    }
    var args = slicy(arguments, 1);

    var bound;
    var binder = function () {
        if (this instanceof bound) {
            var result = target.apply(
                this,
                concatty(args, arguments)
            );
            if (Object(result) === result) {
                return result;
            }
            return this;
        }
        return target.apply(
            that,
            concatty(args, arguments)
        );

    };

    var boundLength = max(0, target.length - args.length);
    var boundArgs = [];
    for (var i = 0; i < boundLength; i++) {
        boundArgs[i] = '$' + i;
    }

    bound = Function('binder', 'return function (' + joiny(boundArgs, ',') + '){ return binder.apply(this,arguments); }')(binder);

    if (target.prototype) {
        var Empty = function Empty() {};
        Empty.prototype = target.prototype;
        bound.prototype = new Empty();
        Empty.prototype = null;
    }

    return bound;
};


/***/ }),

/***/ "./node_modules/function-bind/index.js":
/*!*********************************************!*\
  !*** ./node_modules/function-bind/index.js ***!
  \*********************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var implementation = __webpack_require__(/*! ./implementation */ "./node_modules/function-bind/implementation.js");

module.exports = Function.prototype.bind || implementation;


/***/ }),

/***/ "./node_modules/get-intrinsic/index.js":
/*!*********************************************!*\
  !*** ./node_modules/get-intrinsic/index.js ***!
  \*********************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var undefined;

var $SyntaxError = SyntaxError;
var $Function = Function;
var $TypeError = TypeError;

// eslint-disable-next-line consistent-return
var getEvalledConstructor = function (expressionSyntax) {
	try {
		return $Function('"use strict"; return (' + expressionSyntax + ').constructor;')();
	} catch (e) {}
};

var $gOPD = Object.getOwnPropertyDescriptor;
if ($gOPD) {
	try {
		$gOPD({}, '');
	} catch (e) {
		$gOPD = null; // this is IE 8, which has a broken gOPD
	}
}

var throwTypeError = function () {
	throw new $TypeError();
};
var ThrowTypeError = $gOPD
	? (function () {
		try {
			// eslint-disable-next-line no-unused-expressions, no-caller, no-restricted-properties
			arguments.callee; // IE 8 does not throw here
			return throwTypeError;
		} catch (calleeThrows) {
			try {
				// IE 8 throws on Object.getOwnPropertyDescriptor(arguments, '')
				return $gOPD(arguments, 'callee').get;
			} catch (gOPDthrows) {
				return throwTypeError;
			}
		}
	}())
	: throwTypeError;

var hasSymbols = __webpack_require__(/*! has-symbols */ "./node_modules/has-symbols/index.js")();
var hasProto = __webpack_require__(/*! has-proto */ "./node_modules/has-proto/index.js")();

var getProto = Object.getPrototypeOf || (
	hasProto
		? function (x) { return x.__proto__; } // eslint-disable-line no-proto
		: null
);

var needsEval = {};

var TypedArray = typeof Uint8Array === 'undefined' || !getProto ? undefined : getProto(Uint8Array);

var INTRINSICS = {
	'%AggregateError%': typeof AggregateError === 'undefined' ? undefined : AggregateError,
	'%Array%': Array,
	'%ArrayBuffer%': typeof ArrayBuffer === 'undefined' ? undefined : ArrayBuffer,
	'%ArrayIteratorPrototype%': hasSymbols && getProto ? getProto([][Symbol.iterator]()) : undefined,
	'%AsyncFromSyncIteratorPrototype%': undefined,
	'%AsyncFunction%': needsEval,
	'%AsyncGenerator%': needsEval,
	'%AsyncGeneratorFunction%': needsEval,
	'%AsyncIteratorPrototype%': needsEval,
	'%Atomics%': typeof Atomics === 'undefined' ? undefined : Atomics,
	'%BigInt%': typeof BigInt === 'undefined' ? undefined : BigInt,
	'%BigInt64Array%': typeof BigInt64Array === 'undefined' ? undefined : BigInt64Array,
	'%BigUint64Array%': typeof BigUint64Array === 'undefined' ? undefined : BigUint64Array,
	'%Boolean%': Boolean,
	'%DataView%': typeof DataView === 'undefined' ? undefined : DataView,
	'%Date%': Date,
	'%decodeURI%': decodeURI,
	'%decodeURIComponent%': decodeURIComponent,
	'%encodeURI%': encodeURI,
	'%encodeURIComponent%': encodeURIComponent,
	'%Error%': Error,
	'%eval%': eval, // eslint-disable-line no-eval
	'%EvalError%': EvalError,
	'%Float32Array%': typeof Float32Array === 'undefined' ? undefined : Float32Array,
	'%Float64Array%': typeof Float64Array === 'undefined' ? undefined : Float64Array,
	'%FinalizationRegistry%': typeof FinalizationRegistry === 'undefined' ? undefined : FinalizationRegistry,
	'%Function%': $Function,
	'%GeneratorFunction%': needsEval,
	'%Int8Array%': typeof Int8Array === 'undefined' ? undefined : Int8Array,
	'%Int16Array%': typeof Int16Array === 'undefined' ? undefined : Int16Array,
	'%Int32Array%': typeof Int32Array === 'undefined' ? undefined : Int32Array,
	'%isFinite%': isFinite,
	'%isNaN%': isNaN,
	'%IteratorPrototype%': hasSymbols && getProto ? getProto(getProto([][Symbol.iterator]())) : undefined,
	'%JSON%': typeof JSON === 'object' ? JSON : undefined,
	'%Map%': typeof Map === 'undefined' ? undefined : Map,
	'%MapIteratorPrototype%': typeof Map === 'undefined' || !hasSymbols || !getProto ? undefined : getProto(new Map()[Symbol.iterator]()),
	'%Math%': Math,
	'%Number%': Number,
	'%Object%': Object,
	'%parseFloat%': parseFloat,
	'%parseInt%': parseInt,
	'%Promise%': typeof Promise === 'undefined' ? undefined : Promise,
	'%Proxy%': typeof Proxy === 'undefined' ? undefined : Proxy,
	'%RangeError%': RangeError,
	'%ReferenceError%': ReferenceError,
	'%Reflect%': typeof Reflect === 'undefined' ? undefined : Reflect,
	'%RegExp%': RegExp,
	'%Set%': typeof Set === 'undefined' ? undefined : Set,
	'%SetIteratorPrototype%': typeof Set === 'undefined' || !hasSymbols || !getProto ? undefined : getProto(new Set()[Symbol.iterator]()),
	'%SharedArrayBuffer%': typeof SharedArrayBuffer === 'undefined' ? undefined : SharedArrayBuffer,
	'%String%': String,
	'%StringIteratorPrototype%': hasSymbols && getProto ? getProto(''[Symbol.iterator]()) : undefined,
	'%Symbol%': hasSymbols ? Symbol : undefined,
	'%SyntaxError%': $SyntaxError,
	'%ThrowTypeError%': ThrowTypeError,
	'%TypedArray%': TypedArray,
	'%TypeError%': $TypeError,
	'%Uint8Array%': typeof Uint8Array === 'undefined' ? undefined : Uint8Array,
	'%Uint8ClampedArray%': typeof Uint8ClampedArray === 'undefined' ? undefined : Uint8ClampedArray,
	'%Uint16Array%': typeof Uint16Array === 'undefined' ? undefined : Uint16Array,
	'%Uint32Array%': typeof Uint32Array === 'undefined' ? undefined : Uint32Array,
	'%URIError%': URIError,
	'%WeakMap%': typeof WeakMap === 'undefined' ? undefined : WeakMap,
	'%WeakRef%': typeof WeakRef === 'undefined' ? undefined : WeakRef,
	'%WeakSet%': typeof WeakSet === 'undefined' ? undefined : WeakSet
};

if (getProto) {
	try {
		null.error; // eslint-disable-line no-unused-expressions
	} catch (e) {
		// https://github.com/tc39/proposal-shadowrealm/pull/384#issuecomment-1364264229
		var errorProto = getProto(getProto(e));
		INTRINSICS['%Error.prototype%'] = errorProto;
	}
}

var doEval = function doEval(name) {
	var value;
	if (name === '%AsyncFunction%') {
		value = getEvalledConstructor('async function () {}');
	} else if (name === '%GeneratorFunction%') {
		value = getEvalledConstructor('function* () {}');
	} else if (name === '%AsyncGeneratorFunction%') {
		value = getEvalledConstructor('async function* () {}');
	} else if (name === '%AsyncGenerator%') {
		var fn = doEval('%AsyncGeneratorFunction%');
		if (fn) {
			value = fn.prototype;
		}
	} else if (name === '%AsyncIteratorPrototype%') {
		var gen = doEval('%AsyncGenerator%');
		if (gen && getProto) {
			value = getProto(gen.prototype);
		}
	}

	INTRINSICS[name] = value;

	return value;
};

var LEGACY_ALIASES = {
	'%ArrayBufferPrototype%': ['ArrayBuffer', 'prototype'],
	'%ArrayPrototype%': ['Array', 'prototype'],
	'%ArrayProto_entries%': ['Array', 'prototype', 'entries'],
	'%ArrayProto_forEach%': ['Array', 'prototype', 'forEach'],
	'%ArrayProto_keys%': ['Array', 'prototype', 'keys'],
	'%ArrayProto_values%': ['Array', 'prototype', 'values'],
	'%AsyncFunctionPrototype%': ['AsyncFunction', 'prototype'],
	'%AsyncGenerator%': ['AsyncGeneratorFunction', 'prototype'],
	'%AsyncGeneratorPrototype%': ['AsyncGeneratorFunction', 'prototype', 'prototype'],
	'%BooleanPrototype%': ['Boolean', 'prototype'],
	'%DataViewPrototype%': ['DataView', 'prototype'],
	'%DatePrototype%': ['Date', 'prototype'],
	'%ErrorPrototype%': ['Error', 'prototype'],
	'%EvalErrorPrototype%': ['EvalError', 'prototype'],
	'%Float32ArrayPrototype%': ['Float32Array', 'prototype'],
	'%Float64ArrayPrototype%': ['Float64Array', 'prototype'],
	'%FunctionPrototype%': ['Function', 'prototype'],
	'%Generator%': ['GeneratorFunction', 'prototype'],
	'%GeneratorPrototype%': ['GeneratorFunction', 'prototype', 'prototype'],
	'%Int8ArrayPrototype%': ['Int8Array', 'prototype'],
	'%Int16ArrayPrototype%': ['Int16Array', 'prototype'],
	'%Int32ArrayPrototype%': ['Int32Array', 'prototype'],
	'%JSONParse%': ['JSON', 'parse'],
	'%JSONStringify%': ['JSON', 'stringify'],
	'%MapPrototype%': ['Map', 'prototype'],
	'%NumberPrototype%': ['Number', 'prototype'],
	'%ObjectPrototype%': ['Object', 'prototype'],
	'%ObjProto_toString%': ['Object', 'prototype', 'toString'],
	'%ObjProto_valueOf%': ['Object', 'prototype', 'valueOf'],
	'%PromisePrototype%': ['Promise', 'prototype'],
	'%PromiseProto_then%': ['Promise', 'prototype', 'then'],
	'%Promise_all%': ['Promise', 'all'],
	'%Promise_reject%': ['Promise', 'reject'],
	'%Promise_resolve%': ['Promise', 'resolve'],
	'%RangeErrorPrototype%': ['RangeError', 'prototype'],
	'%ReferenceErrorPrototype%': ['ReferenceError', 'prototype'],
	'%RegExpPrototype%': ['RegExp', 'prototype'],
	'%SetPrototype%': ['Set', 'prototype'],
	'%SharedArrayBufferPrototype%': ['SharedArrayBuffer', 'prototype'],
	'%StringPrototype%': ['String', 'prototype'],
	'%SymbolPrototype%': ['Symbol', 'prototype'],
	'%SyntaxErrorPrototype%': ['SyntaxError', 'prototype'],
	'%TypedArrayPrototype%': ['TypedArray', 'prototype'],
	'%TypeErrorPrototype%': ['TypeError', 'prototype'],
	'%Uint8ArrayPrototype%': ['Uint8Array', 'prototype'],
	'%Uint8ClampedArrayPrototype%': ['Uint8ClampedArray', 'prototype'],
	'%Uint16ArrayPrototype%': ['Uint16Array', 'prototype'],
	'%Uint32ArrayPrototype%': ['Uint32Array', 'prototype'],
	'%URIErrorPrototype%': ['URIError', 'prototype'],
	'%WeakMapPrototype%': ['WeakMap', 'prototype'],
	'%WeakSetPrototype%': ['WeakSet', 'prototype']
};

var bind = __webpack_require__(/*! function-bind */ "./node_modules/function-bind/index.js");
var hasOwn = __webpack_require__(/*! hasown */ "./node_modules/hasown/index.js");
var $concat = bind.call(Function.call, Array.prototype.concat);
var $spliceApply = bind.call(Function.apply, Array.prototype.splice);
var $replace = bind.call(Function.call, String.prototype.replace);
var $strSlice = bind.call(Function.call, String.prototype.slice);
var $exec = bind.call(Function.call, RegExp.prototype.exec);

/* adapted from https://github.com/lodash/lodash/blob/4.17.15/dist/lodash.js#L6735-L6744 */
var rePropName = /[^%.[\]]+|\[(?:(-?\d+(?:\.\d+)?)|(["'])((?:(?!\2)[^\\]|\\.)*?)\2)\]|(?=(?:\.|\[\])(?:\.|\[\]|%$))/g;
var reEscapeChar = /\\(\\)?/g; /** Used to match backslashes in property paths. */
var stringToPath = function stringToPath(string) {
	var first = $strSlice(string, 0, 1);
	var last = $strSlice(string, -1);
	if (first === '%' && last !== '%') {
		throw new $SyntaxError('invalid intrinsic syntax, expected closing `%`');
	} else if (last === '%' && first !== '%') {
		throw new $SyntaxError('invalid intrinsic syntax, expected opening `%`');
	}
	var result = [];
	$replace(string, rePropName, function (match, number, quote, subString) {
		result[result.length] = quote ? $replace(subString, reEscapeChar, '$1') : number || match;
	});
	return result;
};
/* end adaptation */

var getBaseIntrinsic = function getBaseIntrinsic(name, allowMissing) {
	var intrinsicName = name;
	var alias;
	if (hasOwn(LEGACY_ALIASES, intrinsicName)) {
		alias = LEGACY_ALIASES[intrinsicName];
		intrinsicName = '%' + alias[0] + '%';
	}

	if (hasOwn(INTRINSICS, intrinsicName)) {
		var value = INTRINSICS[intrinsicName];
		if (value === needsEval) {
			value = doEval(intrinsicName);
		}
		if (typeof value === 'undefined' && !allowMissing) {
			throw new $TypeError('intrinsic ' + name + ' exists, but is not available. Please file an issue!');
		}

		return {
			alias: alias,
			name: intrinsicName,
			value: value
		};
	}

	throw new $SyntaxError('intrinsic ' + name + ' does not exist!');
};

module.exports = function GetIntrinsic(name, allowMissing) {
	if (typeof name !== 'string' || name.length === 0) {
		throw new $TypeError('intrinsic name must be a non-empty string');
	}
	if (arguments.length > 1 && typeof allowMissing !== 'boolean') {
		throw new $TypeError('"allowMissing" argument must be a boolean');
	}

	if ($exec(/^%?[^%]*%?$/, name) === null) {
		throw new $SyntaxError('`%` may not be present anywhere but at the beginning and end of the intrinsic name');
	}
	var parts = stringToPath(name);
	var intrinsicBaseName = parts.length > 0 ? parts[0] : '';

	var intrinsic = getBaseIntrinsic('%' + intrinsicBaseName + '%', allowMissing);
	var intrinsicRealName = intrinsic.name;
	var value = intrinsic.value;
	var skipFurtherCaching = false;

	var alias = intrinsic.alias;
	if (alias) {
		intrinsicBaseName = alias[0];
		$spliceApply(parts, $concat([0, 1], alias));
	}

	for (var i = 1, isOwn = true; i < parts.length; i += 1) {
		var part = parts[i];
		var first = $strSlice(part, 0, 1);
		var last = $strSlice(part, -1);
		if (
			(
				(first === '"' || first === "'" || first === '`')
				|| (last === '"' || last === "'" || last === '`')
			)
			&& first !== last
		) {
			throw new $SyntaxError('property names with quotes must have matching quotes');
		}
		if (part === 'constructor' || !isOwn) {
			skipFurtherCaching = true;
		}

		intrinsicBaseName += '.' + part;
		intrinsicRealName = '%' + intrinsicBaseName + '%';

		if (hasOwn(INTRINSICS, intrinsicRealName)) {
			value = INTRINSICS[intrinsicRealName];
		} else if (value != null) {
			if (!(part in value)) {
				if (!allowMissing) {
					throw new $TypeError('base intrinsic for ' + name + ' exists, but the property is not available.');
				}
				return void undefined;
			}
			if ($gOPD && (i + 1) >= parts.length) {
				var desc = $gOPD(value, part);
				isOwn = !!desc;

				// By convention, when a data property is converted to an accessor
				// property to emulate a data property that does not suffer from
				// the override mistake, that accessor's getter is marked with
				// an `originalValue` property. Here, when we detect this, we
				// uphold the illusion by pretending to see that original data
				// property, i.e., returning the value rather than the getter
				// itself.
				if (isOwn && 'get' in desc && !('originalValue' in desc.get)) {
					value = desc.get;
				} else {
					value = value[part];
				}
			} else {
				isOwn = hasOwn(value, part);
				value = value[part];
			}

			if (isOwn && !skipFurtherCaching) {
				INTRINSICS[intrinsicRealName] = value;
			}
		}
	}
	return value;
};


/***/ }),

/***/ "./node_modules/gopd/index.js":
/*!************************************!*\
  !*** ./node_modules/gopd/index.js ***!
  \************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var GetIntrinsic = __webpack_require__(/*! get-intrinsic */ "./node_modules/get-intrinsic/index.js");

var $gOPD = GetIntrinsic('%Object.getOwnPropertyDescriptor%', true);

if ($gOPD) {
	try {
		$gOPD([], 'length');
	} catch (e) {
		// IE 8 has a broken gOPD
		$gOPD = null;
	}
}

module.exports = $gOPD;


/***/ }),

/***/ "./node_modules/has-property-descriptors/index.js":
/*!********************************************************!*\
  !*** ./node_modules/has-property-descriptors/index.js ***!
  \********************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var GetIntrinsic = __webpack_require__(/*! get-intrinsic */ "./node_modules/get-intrinsic/index.js");

var $defineProperty = GetIntrinsic('%Object.defineProperty%', true);

var hasPropertyDescriptors = function hasPropertyDescriptors() {
	if ($defineProperty) {
		try {
			$defineProperty({}, 'a', { value: 1 });
			return true;
		} catch (e) {
			// IE 8 has a broken defineProperty
			return false;
		}
	}
	return false;
};

hasPropertyDescriptors.hasArrayLengthDefineBug = function hasArrayLengthDefineBug() {
	// node v0.6 has a bug where array lengths can be Set but not Defined
	if (!hasPropertyDescriptors()) {
		return null;
	}
	try {
		return $defineProperty([], 'length', { value: 1 }).length !== 1;
	} catch (e) {
		// In Firefox 4-22, defining length on an array throws an exception.
		return true;
	}
};

module.exports = hasPropertyDescriptors;


/***/ }),

/***/ "./node_modules/has-proto/index.js":
/*!*****************************************!*\
  !*** ./node_modules/has-proto/index.js ***!
  \*****************************************/
/***/ ((module) => {

"use strict";


var test = {
	foo: {}
};

var $Object = Object;

module.exports = function hasProto() {
	return { __proto__: test }.foo === test.foo && !({ __proto__: null } instanceof $Object);
};


/***/ }),

/***/ "./node_modules/has-symbols/index.js":
/*!*******************************************!*\
  !*** ./node_modules/has-symbols/index.js ***!
  \*******************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var origSymbol = typeof Symbol !== 'undefined' && Symbol;
var hasSymbolSham = __webpack_require__(/*! ./shams */ "./node_modules/has-symbols/shams.js");

module.exports = function hasNativeSymbols() {
	if (typeof origSymbol !== 'function') { return false; }
	if (typeof Symbol !== 'function') { return false; }
	if (typeof origSymbol('foo') !== 'symbol') { return false; }
	if (typeof Symbol('bar') !== 'symbol') { return false; }

	return hasSymbolSham();
};


/***/ }),

/***/ "./node_modules/has-symbols/shams.js":
/*!*******************************************!*\
  !*** ./node_modules/has-symbols/shams.js ***!
  \*******************************************/
/***/ ((module) => {

"use strict";


/* eslint complexity: [2, 18], max-statements: [2, 33] */
module.exports = function hasSymbols() {
	if (typeof Symbol !== 'function' || typeof Object.getOwnPropertySymbols !== 'function') { return false; }
	if (typeof Symbol.iterator === 'symbol') { return true; }

	var obj = {};
	var sym = Symbol('test');
	var symObj = Object(sym);
	if (typeof sym === 'string') { return false; }

	if (Object.prototype.toString.call(sym) !== '[object Symbol]') { return false; }
	if (Object.prototype.toString.call(symObj) !== '[object Symbol]') { return false; }

	// temp disabled per https://github.com/ljharb/object.assign/issues/17
	// if (sym instanceof Symbol) { return false; }
	// temp disabled per https://github.com/WebReflection/get-own-property-symbols/issues/4
	// if (!(symObj instanceof Symbol)) { return false; }

	// if (typeof Symbol.prototype.toString !== 'function') { return false; }
	// if (String(sym) !== Symbol.prototype.toString.call(sym)) { return false; }

	var symVal = 42;
	obj[sym] = symVal;
	for (sym in obj) { return false; } // eslint-disable-line no-restricted-syntax, no-unreachable-loop
	if (typeof Object.keys === 'function' && Object.keys(obj).length !== 0) { return false; }

	if (typeof Object.getOwnPropertyNames === 'function' && Object.getOwnPropertyNames(obj).length !== 0) { return false; }

	var syms = Object.getOwnPropertySymbols(obj);
	if (syms.length !== 1 || syms[0] !== sym) { return false; }

	if (!Object.prototype.propertyIsEnumerable.call(obj, sym)) { return false; }

	if (typeof Object.getOwnPropertyDescriptor === 'function') {
		var descriptor = Object.getOwnPropertyDescriptor(obj, sym);
		if (descriptor.value !== symVal || descriptor.enumerable !== true) { return false; }
	}

	return true;
};


/***/ }),

/***/ "./node_modules/hasown/index.js":
/*!**************************************!*\
  !*** ./node_modules/hasown/index.js ***!
  \**************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var call = Function.prototype.call;
var $hasOwn = Object.prototype.hasOwnProperty;
var bind = __webpack_require__(/*! function-bind */ "./node_modules/function-bind/index.js");

/** @type {(o: {}, p: PropertyKey) => p is keyof o} */
module.exports = bind.call(call, $hasOwn);


/***/ }),

/***/ "./node_modules/ieee754/index.js":
/*!***************************************!*\
  !*** ./node_modules/ieee754/index.js ***!
  \***************************************/
/***/ ((__unused_webpack_module, exports) => {

/*! ieee754. BSD-3-Clause License. Feross Aboukhadijeh <https://feross.org/opensource> */
exports.read = function (buffer, offset, isLE, mLen, nBytes) {
  var e, m
  var eLen = (nBytes * 8) - mLen - 1
  var eMax = (1 << eLen) - 1
  var eBias = eMax >> 1
  var nBits = -7
  var i = isLE ? (nBytes - 1) : 0
  var d = isLE ? -1 : 1
  var s = buffer[offset + i]

  i += d

  e = s & ((1 << (-nBits)) - 1)
  s >>= (-nBits)
  nBits += eLen
  for (; nBits > 0; e = (e * 256) + buffer[offset + i], i += d, nBits -= 8) {}

  m = e & ((1 << (-nBits)) - 1)
  e >>= (-nBits)
  nBits += mLen
  for (; nBits > 0; m = (m * 256) + buffer[offset + i], i += d, nBits -= 8) {}

  if (e === 0) {
    e = 1 - eBias
  } else if (e === eMax) {
    return m ? NaN : ((s ? -1 : 1) * Infinity)
  } else {
    m = m + Math.pow(2, mLen)
    e = e - eBias
  }
  return (s ? -1 : 1) * m * Math.pow(2, e - mLen)
}

exports.write = function (buffer, value, offset, isLE, mLen, nBytes) {
  var e, m, c
  var eLen = (nBytes * 8) - mLen - 1
  var eMax = (1 << eLen) - 1
  var eBias = eMax >> 1
  var rt = (mLen === 23 ? Math.pow(2, -24) - Math.pow(2, -77) : 0)
  var i = isLE ? 0 : (nBytes - 1)
  var d = isLE ? 1 : -1
  var s = value < 0 || (value === 0 && 1 / value < 0) ? 1 : 0

  value = Math.abs(value)

  if (isNaN(value) || value === Infinity) {
    m = isNaN(value) ? 1 : 0
    e = eMax
  } else {
    e = Math.floor(Math.log(value) / Math.LN2)
    if (value * (c = Math.pow(2, -e)) < 1) {
      e--
      c *= 2
    }
    if (e + eBias >= 1) {
      value += rt / c
    } else {
      value += rt * Math.pow(2, 1 - eBias)
    }
    if (value * c >= 2) {
      e++
      c /= 2
    }

    if (e + eBias >= eMax) {
      m = 0
      e = eMax
    } else if (e + eBias >= 1) {
      m = ((value * c) - 1) * Math.pow(2, mLen)
      e = e + eBias
    } else {
      m = value * Math.pow(2, eBias - 1) * Math.pow(2, mLen)
      e = 0
    }
  }

  for (; mLen >= 8; buffer[offset + i] = m & 0xff, i += d, m /= 256, mLen -= 8) {}

  e = (e << mLen) | m
  eLen += mLen
  for (; eLen > 0; buffer[offset + i] = e & 0xff, i += d, e /= 256, eLen -= 8) {}

  buffer[offset + i - d] |= s * 128
}


/***/ }),

/***/ "./node_modules/inherits/inherits_browser.js":
/*!***************************************************!*\
  !*** ./node_modules/inherits/inherits_browser.js ***!
  \***************************************************/
/***/ ((module) => {

if (typeof Object.create === 'function') {
  // implementation from standard node.js 'util' module
  module.exports = function inherits(ctor, superCtor) {
    if (superCtor) {
      ctor.super_ = superCtor
      ctor.prototype = Object.create(superCtor.prototype, {
        constructor: {
          value: ctor,
          enumerable: false,
          writable: true,
          configurable: true
        }
      })
    }
  };
} else {
  // old school shim for old browsers
  module.exports = function inherits(ctor, superCtor) {
    if (superCtor) {
      ctor.super_ = superCtor
      var TempCtor = function () {}
      TempCtor.prototype = superCtor.prototype
      ctor.prototype = new TempCtor()
      ctor.prototype.constructor = ctor
    }
  }
}


/***/ }),

/***/ "./node_modules/mpegts.js/dist/mpegts.js":
/*!***********************************************!*\
  !*** ./node_modules/mpegts.js/dist/mpegts.js ***!
  \***********************************************/
/***/ ((module) => {

!function(e,t){ true?module.exports=t():0}(window,(function(){return function(e){var t={};function i(n){if(t[n])return t[n].exports;var a=t[n]={i:n,l:!1,exports:{}};return e[n].call(a.exports,a,a.exports,i),a.l=!0,a.exports}return i.m=e,i.c=t,i.d=function(e,t,n){i.o(e,t)||Object.defineProperty(e,t,{enumerable:!0,get:n})},i.r=function(e){"undefined"!=typeof Symbol&&Symbol.toStringTag&&Object.defineProperty(e,Symbol.toStringTag,{value:"Module"}),Object.defineProperty(e,"__esModule",{value:!0})},i.t=function(e,t){if(1&t&&(e=i(e)),8&t)return e;if(4&t&&"object"==typeof e&&e&&e.__esModule)return e;var n=Object.create(null);if(i.r(n),Object.defineProperty(n,"default",{enumerable:!0,value:e}),2&t&&"string"!=typeof e)for(var a in e)i.d(n,a,function(t){return e[t]}.bind(null,a));return n},i.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return i.d(t,"a",t),t},i.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},i.p="",i(i.s=14)}([function(e,t,i){"use strict";var n=i(6),a=i.n(n),r=function(){function e(){}return e.e=function(t,i){t&&!e.FORCE_GLOBAL_TAG||(t=e.GLOBAL_TAG);var n="["+t+"] > "+i;e.ENABLE_CALLBACK&&e.emitter.emit("log","error",n),e.ENABLE_ERROR&&(console.error?console.error(n):console.warn?console.warn(n):console.log(n))},e.i=function(t,i){t&&!e.FORCE_GLOBAL_TAG||(t=e.GLOBAL_TAG);var n="["+t+"] > "+i;e.ENABLE_CALLBACK&&e.emitter.emit("log","info",n),e.ENABLE_INFO&&(console.info?console.info(n):console.log(n))},e.w=function(t,i){t&&!e.FORCE_GLOBAL_TAG||(t=e.GLOBAL_TAG);var n="["+t+"] > "+i;e.ENABLE_CALLBACK&&e.emitter.emit("log","warn",n),e.ENABLE_WARN&&(console.warn?console.warn(n):console.log(n))},e.d=function(t,i){t&&!e.FORCE_GLOBAL_TAG||(t=e.GLOBAL_TAG);var n="["+t+"] > "+i;e.ENABLE_CALLBACK&&e.emitter.emit("log","debug",n),e.ENABLE_DEBUG&&(console.debug?console.debug(n):console.log(n))},e.v=function(t,i){t&&!e.FORCE_GLOBAL_TAG||(t=e.GLOBAL_TAG);var n="["+t+"] > "+i;e.ENABLE_CALLBACK&&e.emitter.emit("log","verbose",n),e.ENABLE_VERBOSE&&console.log(n)},e}();r.GLOBAL_TAG="mpegts.js",r.FORCE_GLOBAL_TAG=!1,r.ENABLE_ERROR=!0,r.ENABLE_INFO=!0,r.ENABLE_WARN=!0,r.ENABLE_DEBUG=!0,r.ENABLE_VERBOSE=!0,r.ENABLE_CALLBACK=!1,r.emitter=new a.a,t.a=r},function(e,t,i){"use strict";t.a={IO_ERROR:"io_error",DEMUX_ERROR:"demux_error",INIT_SEGMENT:"init_segment",MEDIA_SEGMENT:"media_segment",LOADING_COMPLETE:"loading_complete",RECOVERED_EARLY_EOF:"recovered_early_eof",MEDIA_INFO:"media_info",METADATA_ARRIVED:"metadata_arrived",SCRIPTDATA_ARRIVED:"scriptdata_arrived",TIMED_ID3_METADATA_ARRIVED:"timed_id3_metadata_arrived",SMPTE2038_METADATA_ARRIVED:"smpte2038_metadata_arrived",SCTE35_METADATA_ARRIVED:"scte35_metadata_arrived",PES_PRIVATE_DATA_DESCRIPTOR:"pes_private_data_descriptor",PES_PRIVATE_DATA_ARRIVED:"pes_private_data_arrived",STATISTICS_INFO:"statistics_info",RECOMMEND_SEEKPOINT:"recommend_seekpoint"}},function(e,t,i){"use strict";i.d(t,"c",(function(){return a})),i.d(t,"b",(function(){return r})),i.d(t,"a",(function(){return s}));var n=i(3),a={kIdle:0,kConnecting:1,kBuffering:2,kError:3,kComplete:4},r={OK:"OK",EXCEPTION:"Exception",HTTP_STATUS_CODE_INVALID:"HttpStatusCodeInvalid",CONNECTING_TIMEOUT:"ConnectingTimeout",EARLY_EOF:"EarlyEof",UNRECOVERABLE_EARLY_EOF:"UnrecoverableEarlyEof"},s=function(){function e(e){this._type=e||"undefined",this._status=a.kIdle,this._needStash=!1,this._onContentLengthKnown=null,this._onURLRedirect=null,this._onDataArrival=null,this._onError=null,this._onComplete=null}return e.prototype.destroy=function(){this._status=a.kIdle,this._onContentLengthKnown=null,this._onURLRedirect=null,this._onDataArrival=null,this._onError=null,this._onComplete=null},e.prototype.isWorking=function(){return this._status===a.kConnecting||this._status===a.kBuffering},Object.defineProperty(e.prototype,"type",{get:function(){return this._type},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"status",{get:function(){return this._status},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"needStashBuffer",{get:function(){return this._needStash},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onContentLengthKnown",{get:function(){return this._onContentLengthKnown},set:function(e){this._onContentLengthKnown=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onURLRedirect",{get:function(){return this._onURLRedirect},set:function(e){this._onURLRedirect=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onDataArrival",{get:function(){return this._onDataArrival},set:function(e){this._onDataArrival=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onError",{get:function(){return this._onError},set:function(e){this._onError=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onComplete",{get:function(){return this._onComplete},set:function(e){this._onComplete=e},enumerable:!1,configurable:!0}),e.prototype.open=function(e,t){throw new n.c("Unimplemented abstract function!")},e.prototype.abort=function(){throw new n.c("Unimplemented abstract function!")},e}()},function(e,t,i){"use strict";i.d(t,"d",(function(){return r})),i.d(t,"a",(function(){return s})),i.d(t,"b",(function(){return o})),i.d(t,"c",(function(){return d}));var n,a=(n=function(e,t){return(n=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(e,t){e.__proto__=t}||function(e,t){for(var i in t)t.hasOwnProperty(i)&&(e[i]=t[i])})(e,t)},function(e,t){function i(){this.constructor=e}n(e,t),e.prototype=null===t?Object.create(t):(i.prototype=t.prototype,new i)}),r=function(){function e(e){this._message=e}return Object.defineProperty(e.prototype,"name",{get:function(){return"RuntimeException"},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"message",{get:function(){return this._message},enumerable:!1,configurable:!0}),e.prototype.toString=function(){return this.name+": "+this.message},e}(),s=function(e){function t(t){return e.call(this,t)||this}return a(t,e),Object.defineProperty(t.prototype,"name",{get:function(){return"IllegalStateException"},enumerable:!1,configurable:!0}),t}(r),o=function(e){function t(t){return e.call(this,t)||this}return a(t,e),Object.defineProperty(t.prototype,"name",{get:function(){return"InvalidArgumentException"},enumerable:!1,configurable:!0}),t}(r),d=function(e){function t(t){return e.call(this,t)||this}return a(t,e),Object.defineProperty(t.prototype,"name",{get:function(){return"NotImplementedException"},enumerable:!1,configurable:!0}),t}(r)},function(e,t,i){"use strict";var n={};!function(){var e=self.navigator.userAgent.toLowerCase(),t=/(edge)\/([\w.]+)/.exec(e)||/(opr)[\/]([\w.]+)/.exec(e)||/(chrome)[ \/]([\w.]+)/.exec(e)||/(iemobile)[\/]([\w.]+)/.exec(e)||/(version)(applewebkit)[ \/]([\w.]+).*(safari)[ \/]([\w.]+)/.exec(e)||/(webkit)[ \/]([\w.]+).*(version)[ \/]([\w.]+).*(safari)[ \/]([\w.]+)/.exec(e)||/(webkit)[ \/]([\w.]+)/.exec(e)||/(opera)(?:.*version|)[ \/]([\w.]+)/.exec(e)||/(msie) ([\w.]+)/.exec(e)||e.indexOf("trident")>=0&&/(rv)(?::| )([\w.]+)/.exec(e)||e.indexOf("compatible")<0&&/(firefox)[ \/]([\w.]+)/.exec(e)||[],i=/(ipad)/.exec(e)||/(ipod)/.exec(e)||/(windows phone)/.exec(e)||/(iphone)/.exec(e)||/(kindle)/.exec(e)||/(android)/.exec(e)||/(windows)/.exec(e)||/(mac)/.exec(e)||/(linux)/.exec(e)||/(cros)/.exec(e)||[],a={browser:t[5]||t[3]||t[1]||"",version:t[2]||t[4]||"0",majorVersion:t[4]||t[2]||"0",platform:i[0]||""},r={};if(a.browser){r[a.browser]=!0;var s=a.majorVersion.split(".");r.version={major:parseInt(a.majorVersion,10),string:a.version},s.length>1&&(r.version.minor=parseInt(s[1],10)),s.length>2&&(r.version.build=parseInt(s[2],10))}if(a.platform&&(r[a.platform]=!0),(r.chrome||r.opr||r.safari)&&(r.webkit=!0),r.rv||r.iemobile){r.rv&&delete r.rv;a.browser="msie",r.msie=!0}if(r.edge){delete r.edge;a.browser="msedge",r.msedge=!0}if(r.opr){a.browser="opera",r.opera=!0}if(r.safari&&r.android){a.browser="android",r.android=!0}for(var o in r.name=a.browser,r.platform=a.platform,n)n.hasOwnProperty(o)&&delete n[o];Object.assign(n,r)}(),t.a=n},function(e,t,i){"use strict";t.a={OK:"OK",FORMAT_ERROR:"FormatError",FORMAT_UNSUPPORTED:"FormatUnsupported",CODEC_UNSUPPORTED:"CodecUnsupported"}},function(e,t,i){"use strict";var n,a="object"==typeof Reflect?Reflect:null,r=a&&"function"==typeof a.apply?a.apply:function(e,t,i){return Function.prototype.apply.call(e,t,i)};n=a&&"function"==typeof a.ownKeys?a.ownKeys:Object.getOwnPropertySymbols?function(e){return Object.getOwnPropertyNames(e).concat(Object.getOwnPropertySymbols(e))}:function(e){return Object.getOwnPropertyNames(e)};var s=Number.isNaN||function(e){return e!=e};function o(){o.init.call(this)}e.exports=o,e.exports.once=function(e,t){return new Promise((function(i,n){function a(i){e.removeListener(t,r),n(i)}function r(){"function"==typeof e.removeListener&&e.removeListener("error",a),i([].slice.call(arguments))}g(e,t,r,{once:!0}),"error"!==t&&function(e,t,i){"function"==typeof e.on&&g(e,"error",t,i)}(e,a,{once:!0})}))},o.EventEmitter=o,o.prototype._events=void 0,o.prototype._eventsCount=0,o.prototype._maxListeners=void 0;var d=10;function _(e){if("function"!=typeof e)throw new TypeError('The "listener" argument must be of type Function. Received type '+typeof e)}function h(e){return void 0===e._maxListeners?o.defaultMaxListeners:e._maxListeners}function c(e,t,i,n){var a,r,s,o;if(_(i),void 0===(r=e._events)?(r=e._events=Object.create(null),e._eventsCount=0):(void 0!==r.newListener&&(e.emit("newListener",t,i.listener?i.listener:i),r=e._events),s=r[t]),void 0===s)s=r[t]=i,++e._eventsCount;else if("function"==typeof s?s=r[t]=n?[i,s]:[s,i]:n?s.unshift(i):s.push(i),(a=h(e))>0&&s.length>a&&!s.warned){s.warned=!0;var d=new Error("Possible EventEmitter memory leak detected. "+s.length+" "+String(t)+" listeners added. Use emitter.setMaxListeners() to increase limit");d.name="MaxListenersExceededWarning",d.emitter=e,d.type=t,d.count=s.length,o=d,console&&console.warn&&console.warn(o)}return e}function u(){if(!this.fired)return this.target.removeListener(this.type,this.wrapFn),this.fired=!0,0===arguments.length?this.listener.call(this.target):this.listener.apply(this.target,arguments)}function l(e,t,i){var n={fired:!1,wrapFn:void 0,target:e,type:t,listener:i},a=u.bind(n);return a.listener=i,n.wrapFn=a,a}function f(e,t,i){var n=e._events;if(void 0===n)return[];var a=n[t];return void 0===a?[]:"function"==typeof a?i?[a.listener||a]:[a]:i?function(e){for(var t=new Array(e.length),i=0;i<t.length;++i)t[i]=e[i].listener||e[i];return t}(a):m(a,a.length)}function p(e){var t=this._events;if(void 0!==t){var i=t[e];if("function"==typeof i)return 1;if(void 0!==i)return i.length}return 0}function m(e,t){for(var i=new Array(t),n=0;n<t;++n)i[n]=e[n];return i}function g(e,t,i,n){if("function"==typeof e.on)n.once?e.once(t,i):e.on(t,i);else{if("function"!=typeof e.addEventListener)throw new TypeError('The "emitter" argument must be of type EventEmitter. Received type '+typeof e);e.addEventListener(t,(function a(r){n.once&&e.removeEventListener(t,a),i(r)}))}}Object.defineProperty(o,"defaultMaxListeners",{enumerable:!0,get:function(){return d},set:function(e){if("number"!=typeof e||e<0||s(e))throw new RangeError('The value of "defaultMaxListeners" is out of range. It must be a non-negative number. Received '+e+".");d=e}}),o.init=function(){void 0!==this._events&&this._events!==Object.getPrototypeOf(this)._events||(this._events=Object.create(null),this._eventsCount=0),this._maxListeners=this._maxListeners||void 0},o.prototype.setMaxListeners=function(e){if("number"!=typeof e||e<0||s(e))throw new RangeError('The value of "n" is out of range. It must be a non-negative number. Received '+e+".");return this._maxListeners=e,this},o.prototype.getMaxListeners=function(){return h(this)},o.prototype.emit=function(e){for(var t=[],i=1;i<arguments.length;i++)t.push(arguments[i]);var n="error"===e,a=this._events;if(void 0!==a)n=n&&void 0===a.error;else if(!n)return!1;if(n){var s;if(t.length>0&&(s=t[0]),s instanceof Error)throw s;var o=new Error("Unhandled error."+(s?" ("+s.message+")":""));throw o.context=s,o}var d=a[e];if(void 0===d)return!1;if("function"==typeof d)r(d,this,t);else{var _=d.length,h=m(d,_);for(i=0;i<_;++i)r(h[i],this,t)}return!0},o.prototype.addListener=function(e,t){return c(this,e,t,!1)},o.prototype.on=o.prototype.addListener,o.prototype.prependListener=function(e,t){return c(this,e,t,!0)},o.prototype.once=function(e,t){return _(t),this.on(e,l(this,e,t)),this},o.prototype.prependOnceListener=function(e,t){return _(t),this.prependListener(e,l(this,e,t)),this},o.prototype.removeListener=function(e,t){var i,n,a,r,s;if(_(t),void 0===(n=this._events))return this;if(void 0===(i=n[e]))return this;if(i===t||i.listener===t)0==--this._eventsCount?this._events=Object.create(null):(delete n[e],n.removeListener&&this.emit("removeListener",e,i.listener||t));else if("function"!=typeof i){for(a=-1,r=i.length-1;r>=0;r--)if(i[r]===t||i[r].listener===t){s=i[r].listener,a=r;break}if(a<0)return this;0===a?i.shift():function(e,t){for(;t+1<e.length;t++)e[t]=e[t+1];e.pop()}(i,a),1===i.length&&(n[e]=i[0]),void 0!==n.removeListener&&this.emit("removeListener",e,s||t)}return this},o.prototype.off=o.prototype.removeListener,o.prototype.removeAllListeners=function(e){var t,i,n;if(void 0===(i=this._events))return this;if(void 0===i.removeListener)return 0===arguments.length?(this._events=Object.create(null),this._eventsCount=0):void 0!==i[e]&&(0==--this._eventsCount?this._events=Object.create(null):delete i[e]),this;if(0===arguments.length){var a,r=Object.keys(i);for(n=0;n<r.length;++n)"removeListener"!==(a=r[n])&&this.removeAllListeners(a);return this.removeAllListeners("removeListener"),this._events=Object.create(null),this._eventsCount=0,this}if("function"==typeof(t=i[e]))this.removeListener(e,t);else if(void 0!==t)for(n=t.length-1;n>=0;n--)this.removeListener(e,t[n]);return this},o.prototype.listeners=function(e){return f(this,e,!0)},o.prototype.rawListeners=function(e){return f(this,e,!1)},o.listenerCount=function(e,t){return"function"==typeof e.listenerCount?e.listenerCount(t):p.call(e,t)},o.prototype.listenerCount=p,o.prototype.eventNames=function(){return this._eventsCount>0?n(this._events):[]}},function(e,t,i){"use strict";i.d(t,"d",(function(){return n})),i.d(t,"b",(function(){return a})),i.d(t,"a",(function(){return r})),i.d(t,"c",(function(){return s}));var n=function(e,t,i,n,a){this.dts=e,this.pts=t,this.duration=i,this.originalDts=n,this.isSyncPoint=a,this.fileposition=null},a=function(){function e(){this.beginDts=0,this.endDts=0,this.beginPts=0,this.endPts=0,this.originalBeginDts=0,this.originalEndDts=0,this.syncPoints=[],this.firstSample=null,this.lastSample=null}return e.prototype.appendSyncPoint=function(e){e.isSyncPoint=!0,this.syncPoints.push(e)},e}(),r=function(){function e(){this._list=[]}return e.prototype.clear=function(){this._list=[]},e.prototype.appendArray=function(e){var t=this._list;0!==e.length&&(t.length>0&&e[0].originalDts<t[t.length-1].originalDts&&this.clear(),Array.prototype.push.apply(t,e))},e.prototype.getLastSyncPointBeforeDts=function(e){if(0==this._list.length)return null;var t=this._list,i=0,n=t.length-1,a=0,r=0,s=n;for(e<t[0].dts&&(i=0,r=s+1);r<=s;){if((a=r+Math.floor((s-r)/2))===n||e>=t[a].dts&&e<t[a+1].dts){i=a;break}t[a].dts<e?r=a+1:s=a-1}return this._list[i]},e}(),s=function(){function e(e){this._type=e,this._list=[],this._lastAppendLocation=-1}return Object.defineProperty(e.prototype,"type",{get:function(){return this._type},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"length",{get:function(){return this._list.length},enumerable:!1,configurable:!0}),e.prototype.isEmpty=function(){return 0===this._list.length},e.prototype.clear=function(){this._list=[],this._lastAppendLocation=-1},e.prototype._searchNearestSegmentBefore=function(e){var t=this._list;if(0===t.length)return-2;var i=t.length-1,n=0,a=0,r=i,s=0;if(e<t[0].originalBeginDts)return s=-1;for(;a<=r;){if((n=a+Math.floor((r-a)/2))===i||e>t[n].lastSample.originalDts&&e<t[n+1].originalBeginDts){s=n;break}t[n].originalBeginDts<e?a=n+1:r=n-1}return s},e.prototype._searchNearestSegmentAfter=function(e){return this._searchNearestSegmentBefore(e)+1},e.prototype.append=function(e){var t=this._list,i=e,n=this._lastAppendLocation,a=0;-1!==n&&n<t.length&&i.originalBeginDts>=t[n].lastSample.originalDts&&(n===t.length-1||n<t.length-1&&i.originalBeginDts<t[n+1].originalBeginDts)?a=n+1:t.length>0&&(a=this._searchNearestSegmentBefore(i.originalBeginDts)+1),this._lastAppendLocation=a,this._list.splice(a,0,i)},e.prototype.getLastSegmentBefore=function(e){var t=this._searchNearestSegmentBefore(e);return t>=0?this._list[t]:null},e.prototype.getLastSampleBefore=function(e){var t=this.getLastSegmentBefore(e);return null!=t?t.lastSample:null},e.prototype.getLastSyncPointBefore=function(e){for(var t=this._searchNearestSegmentBefore(e),i=this._list[t].syncPoints;0===i.length&&t>0;)t--,i=this._list[t].syncPoints;return i.length>0?i[i.length-1]:null},e}()},function(e,t,i){"use strict";var n=function(){function e(){this.mimeType=null,this.duration=null,this.hasAudio=null,this.hasVideo=null,this.audioCodec=null,this.videoCodec=null,this.audioDataRate=null,this.videoDataRate=null,this.audioSampleRate=null,this.audioChannelCount=null,this.width=null,this.height=null,this.fps=null,this.profile=null,this.level=null,this.refFrames=null,this.chromaFormat=null,this.sarNum=null,this.sarDen=null,this.metadata=null,this.segments=null,this.segmentCount=null,this.hasKeyframesIndex=null,this.keyframesIndex=null}return e.prototype.isComplete=function(){var e=!1===this.hasAudio||!0===this.hasAudio&&null!=this.audioCodec&&null!=this.audioSampleRate&&null!=this.audioChannelCount,t=!1===this.hasVideo||!0===this.hasVideo&&null!=this.videoCodec&&null!=this.width&&null!=this.height&&null!=this.fps&&null!=this.profile&&null!=this.level&&null!=this.refFrames&&null!=this.chromaFormat&&null!=this.sarNum&&null!=this.sarDen;return null!=this.mimeType&&e&&t},e.prototype.isSeekable=function(){return!0===this.hasKeyframesIndex},e.prototype.getNearestKeyframe=function(e){if(null==this.keyframesIndex)return null;var t=this.keyframesIndex,i=this._search(t.times,e);return{index:i,milliseconds:t.times[i],fileposition:t.filepositions[i]}},e.prototype._search=function(e,t){var i=0,n=e.length-1,a=0,r=0,s=n;for(t<e[0]&&(i=0,r=s+1);r<=s;){if((a=r+Math.floor((s-r)/2))===n||t>=e[a]&&t<e[a+1]){i=a;break}e[a]<t?r=a+1:s=a-1}return i},e}();t.a=n},function(e,t,i){"use strict";var n=i(6),a=i.n(n),r=i(0),s=function(){function e(){}return Object.defineProperty(e,"forceGlobalTag",{get:function(){return r.a.FORCE_GLOBAL_TAG},set:function(t){r.a.FORCE_GLOBAL_TAG=t,e._notifyChange()},enumerable:!1,configurable:!0}),Object.defineProperty(e,"globalTag",{get:function(){return r.a.GLOBAL_TAG},set:function(t){r.a.GLOBAL_TAG=t,e._notifyChange()},enumerable:!1,configurable:!0}),Object.defineProperty(e,"enableAll",{get:function(){return r.a.ENABLE_VERBOSE&&r.a.ENABLE_DEBUG&&r.a.ENABLE_INFO&&r.a.ENABLE_WARN&&r.a.ENABLE_ERROR},set:function(t){r.a.ENABLE_VERBOSE=t,r.a.ENABLE_DEBUG=t,r.a.ENABLE_INFO=t,r.a.ENABLE_WARN=t,r.a.ENABLE_ERROR=t,e._notifyChange()},enumerable:!1,configurable:!0}),Object.defineProperty(e,"enableDebug",{get:function(){return r.a.ENABLE_DEBUG},set:function(t){r.a.ENABLE_DEBUG=t,e._notifyChange()},enumerable:!1,configurable:!0}),Object.defineProperty(e,"enableVerbose",{get:function(){return r.a.ENABLE_VERBOSE},set:function(t){r.a.ENABLE_VERBOSE=t,e._notifyChange()},enumerable:!1,configurable:!0}),Object.defineProperty(e,"enableInfo",{get:function(){return r.a.ENABLE_INFO},set:function(t){r.a.ENABLE_INFO=t,e._notifyChange()},enumerable:!1,configurable:!0}),Object.defineProperty(e,"enableWarn",{get:function(){return r.a.ENABLE_WARN},set:function(t){r.a.ENABLE_WARN=t,e._notifyChange()},enumerable:!1,configurable:!0}),Object.defineProperty(e,"enableError",{get:function(){return r.a.ENABLE_ERROR},set:function(t){r.a.ENABLE_ERROR=t,e._notifyChange()},enumerable:!1,configurable:!0}),e.getConfig=function(){return{globalTag:r.a.GLOBAL_TAG,forceGlobalTag:r.a.FORCE_GLOBAL_TAG,enableVerbose:r.a.ENABLE_VERBOSE,enableDebug:r.a.ENABLE_DEBUG,enableInfo:r.a.ENABLE_INFO,enableWarn:r.a.ENABLE_WARN,enableError:r.a.ENABLE_ERROR,enableCallback:r.a.ENABLE_CALLBACK}},e.applyConfig=function(e){r.a.GLOBAL_TAG=e.globalTag,r.a.FORCE_GLOBAL_TAG=e.forceGlobalTag,r.a.ENABLE_VERBOSE=e.enableVerbose,r.a.ENABLE_DEBUG=e.enableDebug,r.a.ENABLE_INFO=e.enableInfo,r.a.ENABLE_WARN=e.enableWarn,r.a.ENABLE_ERROR=e.enableError,r.a.ENABLE_CALLBACK=e.enableCallback},e._notifyChange=function(){var t=e.emitter;if(t.listenerCount("change")>0){var i=e.getConfig();t.emit("change",i)}},e.registerListener=function(t){e.emitter.addListener("change",t)},e.removeListener=function(t){e.emitter.removeListener("change",t)},e.addLogListener=function(t){r.a.emitter.addListener("log",t),r.a.emitter.listenerCount("log")>0&&(r.a.ENABLE_CALLBACK=!0,e._notifyChange())},e.removeLogListener=function(t){r.a.emitter.removeListener("log",t),0===r.a.emitter.listenerCount("log")&&(r.a.ENABLE_CALLBACK=!1,e._notifyChange())},e}();s.emitter=new a.a,t.a=s},function(e,t,i){"use strict";var n=i(6),a=i.n(n),r=i(0),s=i(4),o=i(8);function d(e,t,i){var n=e;if(t+i<n.length){for(;i--;)if(128!=(192&n[++t]))return!1;return!0}return!1}var _,h=function(e){for(var t=[],i=e,n=0,a=e.length;n<a;)if(i[n]<128)t.push(String.fromCharCode(i[n])),++n;else{if(i[n]<192);else if(i[n]<224){if(d(i,n,1))if((r=(31&i[n])<<6|63&i[n+1])>=128){t.push(String.fromCharCode(65535&r)),n+=2;continue}}else if(i[n]<240){if(d(i,n,2))if((r=(15&i[n])<<12|(63&i[n+1])<<6|63&i[n+2])>=2048&&55296!=(63488&r)){t.push(String.fromCharCode(65535&r)),n+=3;continue}}else if(i[n]<248){var r;if(d(i,n,3))if((r=(7&i[n])<<18|(63&i[n+1])<<12|(63&i[n+2])<<6|63&i[n+3])>65536&&r<1114112){r-=65536,t.push(String.fromCharCode(r>>>10|55296)),t.push(String.fromCharCode(1023&r|56320)),n+=4;continue}}t.push(String.fromCharCode(65533)),++n}return t.join("")},c=i(3),u=(_=new ArrayBuffer(2),new DataView(_).setInt16(0,256,!0),256===new Int16Array(_)[0]),l=function(){function e(){}return e.parseScriptData=function(t,i,n){var a={};try{var s=e.parseValue(t,i,n),o=e.parseValue(t,i+s.size,n-s.size);a[s.data]=o.data}catch(e){r.a.e("AMF",e.toString())}return a},e.parseObject=function(t,i,n){if(n<3)throw new c.a("Data not enough when parse ScriptDataObject");var a=e.parseString(t,i,n),r=e.parseValue(t,i+a.size,n-a.size),s=r.objectEnd;return{data:{name:a.data,value:r.data},size:a.size+r.size,objectEnd:s}},e.parseVariable=function(t,i,n){return e.parseObject(t,i,n)},e.parseString=function(e,t,i){if(i<2)throw new c.a("Data not enough when parse String");var n=new DataView(e,t,i).getUint16(0,!u);return{data:n>0?h(new Uint8Array(e,t+2,n)):"",size:2+n}},e.parseLongString=function(e,t,i){if(i<4)throw new c.a("Data not enough when parse LongString");var n=new DataView(e,t,i).getUint32(0,!u);return{data:n>0?h(new Uint8Array(e,t+4,n)):"",size:4+n}},e.parseDate=function(e,t,i){if(i<10)throw new c.a("Data size invalid when parse Date");var n=new DataView(e,t,i),a=n.getFloat64(0,!u),r=n.getInt16(8,!u);return{data:new Date(a+=60*r*1e3),size:10}},e.parseValue=function(t,i,n){if(n<1)throw new c.a("Data not enough when parse Value");var a,s=new DataView(t,i,n),o=1,d=s.getUint8(0),_=!1;try{switch(d){case 0:a=s.getFloat64(1,!u),o+=8;break;case 1:a=!!s.getUint8(1),o+=1;break;case 2:var h=e.parseString(t,i+1,n-1);a=h.data,o+=h.size;break;case 3:a={};var l=0;for(9==(16777215&s.getUint32(n-4,!u))&&(l=3);o<n-4;){var f=e.parseObject(t,i+o,n-o-l);if(f.objectEnd)break;a[f.data.name]=f.data.value,o+=f.size}if(o<=n-3)9===(16777215&s.getUint32(o-1,!u))&&(o+=3);break;case 8:a={},o+=4;l=0;for(9==(16777215&s.getUint32(n-4,!u))&&(l=3);o<n-8;){var p=e.parseVariable(t,i+o,n-o-l);if(p.objectEnd)break;a[p.data.name]=p.data.value,o+=p.size}if(o<=n-3)9===(16777215&s.getUint32(o-1,!u))&&(o+=3);break;case 9:a=void 0,o=1,_=!0;break;case 10:a=[];var m=s.getUint32(1,!u);o+=4;for(var g=0;g<m;g++){var v=e.parseValue(t,i+o,n-o);a.push(v.data),o+=v.size}break;case 11:var y=e.parseDate(t,i+1,n-1);a=y.data,o+=y.size;break;case 12:var b=e.parseString(t,i+1,n-1);a=b.data,o+=b.size;break;default:o=n,r.a.w("AMF","Unsupported AMF value type "+d)}}catch(e){r.a.e("AMF",e.toString())}return{data:a,size:o,objectEnd:_}},e}(),f=function(){function e(e){this.TAG="ExpGolomb",this._buffer=e,this._buffer_index=0,this._total_bytes=e.byteLength,this._total_bits=8*e.byteLength,this._current_word=0,this._current_word_bits_left=0}return e.prototype.destroy=function(){this._buffer=null},e.prototype._fillCurrentWord=function(){var e=this._total_bytes-this._buffer_index;if(e<=0)throw new c.a("ExpGolomb: _fillCurrentWord() but no bytes available");var t=Math.min(4,e),i=new Uint8Array(4);i.set(this._buffer.subarray(this._buffer_index,this._buffer_index+t)),this._current_word=new DataView(i.buffer).getUint32(0,!1),this._buffer_index+=t,this._current_word_bits_left=8*t},e.prototype.readBits=function(e){if(e>32)throw new c.b("ExpGolomb: readBits() bits exceeded max 32bits!");if(e<=this._current_word_bits_left){var t=this._current_word>>>32-e;return this._current_word<<=e,this._current_word_bits_left-=e,t}var i=this._current_word_bits_left?this._current_word:0;i>>>=32-this._current_word_bits_left;var n=e-this._current_word_bits_left;this._fillCurrentWord();var a=Math.min(n,this._current_word_bits_left),r=this._current_word>>>32-a;return this._current_word<<=a,this._current_word_bits_left-=a,i=i<<a|r},e.prototype.readBool=function(){return 1===this.readBits(1)},e.prototype.readByte=function(){return this.readBits(8)},e.prototype._skipLeadingZero=function(){var e;for(e=0;e<this._current_word_bits_left;e++)if(0!=(this._current_word&2147483648>>>e))return this._current_word<<=e,this._current_word_bits_left-=e,e;return this._fillCurrentWord(),e+this._skipLeadingZero()},e.prototype.readUEG=function(){var e=this._skipLeadingZero();return this.readBits(e+1)-1},e.prototype.readSEG=function(){var e=this.readUEG();return 1&e?e+1>>>1:-1*(e>>>1)},e}(),p=function(){function e(){}return e._ebsp2rbsp=function(e){for(var t=e,i=t.byteLength,n=new Uint8Array(i),a=0,r=0;r<i;r++)r>=2&&3===t[r]&&0===t[r-1]&&0===t[r-2]||(n[a]=t[r],a++);return new Uint8Array(n.buffer,0,a)},e.parseSPS=function(t){for(var i=t.subarray(1,4),n="avc1.",a=0;a<3;a++){var r=i[a].toString(16);r.length<2&&(r="0"+r),n+=r}var s=e._ebsp2rbsp(t),o=new f(s);o.readByte();var d=o.readByte();o.readByte();var _=o.readByte();o.readUEG();var h=e.getProfileString(d),c=e.getLevelString(_),u=1,l=420,p=8,m=8;if((100===d||110===d||122===d||244===d||44===d||83===d||86===d||118===d||128===d||138===d||144===d)&&(3===(u=o.readUEG())&&o.readBits(1),u<=3&&(l=[0,420,422,444][u]),p=o.readUEG()+8,m=o.readUEG()+8,o.readBits(1),o.readBool()))for(var g=3!==u?8:12,v=0;v<g;v++)o.readBool()&&(v<6?e._skipScalingList(o,16):e._skipScalingList(o,64));o.readUEG();var y=o.readUEG();if(0===y)o.readUEG();else if(1===y){o.readBits(1),o.readSEG(),o.readSEG();var b=o.readUEG();for(v=0;v<b;v++)o.readSEG()}var S=o.readUEG();o.readBits(1);var E=o.readUEG(),A=o.readUEG(),R=o.readBits(1);0===R&&o.readBits(1),o.readBits(1);var T=0,L=0,w=0,k=0;o.readBool()&&(T=o.readUEG(),L=o.readUEG(),w=o.readUEG(),k=o.readUEG());var D=1,C=1,B=0,I=!0,O=0,P=0;if(o.readBool()){if(o.readBool()){var M=o.readByte();M>0&&M<16?(D=[1,12,10,16,40,24,20,32,80,18,15,64,160,4,3,2][M-1],C=[1,11,11,11,33,11,11,11,33,11,11,33,99,3,2,1][M-1]):255===M&&(D=o.readByte()<<8|o.readByte(),C=o.readByte()<<8|o.readByte())}if(o.readBool()&&o.readBool(),o.readBool()&&(o.readBits(4),o.readBool()&&o.readBits(24)),o.readBool()&&(o.readUEG(),o.readUEG()),o.readBool()){var x=o.readBits(32),U=o.readBits(32);I=o.readBool(),B=(O=U)/(P=2*x)}}var N=1;1===D&&1===C||(N=D/C);var G=0,V=0;0===u?(G=1,V=2-R):(G=3===u?1:2,V=(1===u?2:1)*(2-R));var F=16*(E+1),j=16*(A+1)*(2-R);F-=(T+L)*G,j-=(w+k)*V;var z=Math.ceil(F*N);return o.destroy(),o=null,{codec_mimetype:n,profile_idc:d,level_idc:_,profile_string:h,level_string:c,chroma_format_idc:u,bit_depth:p,bit_depth_luma:p,bit_depth_chroma:m,ref_frames:S,chroma_format:l,chroma_format_string:e.getChromaFormatString(l),frame_rate:{fixed:I,fps:B,fps_den:P,fps_num:O},sar_ratio:{width:D,height:C},codec_size:{width:F,height:j},present_size:{width:z,height:j}}},e._skipScalingList=function(e,t){for(var i=8,n=8,a=0;a<t;a++)0!==n&&(n=(i+e.readSEG()+256)%256),i=0===n?i:n},e.getProfileString=function(e){switch(e){case 66:return"Baseline";case 77:return"Main";case 88:return"Extended";case 100:return"High";case 110:return"High10";case 122:return"High422";case 244:return"High444";default:return"Unknown"}},e.getLevelString=function(e){return(e/10).toFixed(1)},e.getChromaFormatString=function(e){switch(e){case 420:return"4:2:0";case 422:return"4:2:2";case 444:return"4:4:4";default:return"Unknown"}},e}(),m=i(5),g=function(){function e(){}return e._ebsp2rbsp=function(e){for(var t=e,i=t.byteLength,n=new Uint8Array(i),a=0,r=0;r<i;r++)r>=2&&3===t[r]&&0===t[r-1]&&0===t[r-2]||(n[a]=t[r],a++);return new Uint8Array(n.buffer,0,a)},e.parseVPS=function(t){var i=e._ebsp2rbsp(t),n=new f(i);n.readByte(),n.readByte();n.readBits(4);n.readBits(2);n.readBits(6);return{num_temporal_layers:n.readBits(3)+1,temporal_id_nested:n.readBool()}},e.parseSPS=function(t){var i=e._ebsp2rbsp(t),n=new f(i);n.readByte(),n.readByte();for(var a=0,r=0,s=0,o=0,d=(n.readBits(4),n.readBits(3)),_=(n.readBool(),n.readBits(2)),h=n.readBool(),c=n.readBits(5),u=n.readByte(),l=n.readByte(),p=n.readByte(),m=n.readByte(),g=n.readByte(),v=n.readByte(),y=n.readByte(),b=n.readByte(),S=n.readByte(),E=n.readByte(),A=n.readByte(),R=[],T=[],L=0;L<d;L++)R.push(n.readBool()),T.push(n.readBool());if(d>0)for(L=d;L<8;L++)n.readBits(2);for(L=0;L<d;L++)R[L]&&(n.readByte(),n.readByte(),n.readByte(),n.readByte(),n.readByte(),n.readByte(),n.readByte(),n.readByte(),n.readByte(),n.readByte(),n.readByte()),T[L]&&n.readByte();n.readUEG();var w=n.readUEG();3==w&&n.readBits(1);var k=n.readUEG(),D=n.readUEG();n.readBool()&&(a+=n.readUEG(),r+=n.readUEG(),s+=n.readUEG(),o+=n.readUEG());var C=n.readUEG(),B=n.readUEG(),I=n.readUEG();for(L=n.readBool()?0:d;L<=d;L++)n.readUEG(),n.readUEG(),n.readUEG();n.readUEG(),n.readUEG(),n.readUEG(),n.readUEG(),n.readUEG(),n.readUEG();if(n.readBool()&&n.readBool())for(var O=0;O<4;O++)for(var P=0;P<(3===O?2:6);P++){if(n.readBool()){var M=Math.min(64,1<<4+(O<<1));O>1&&n.readSEG();for(L=0;L<M;L++)n.readSEG()}else n.readUEG()}n.readBool(),n.readBool();n.readBool()&&(n.readByte(),n.readUEG(),n.readUEG(),n.readBool());var x=n.readUEG(),U=0;for(L=0;L<x;L++){var N=!1;if(0!==L&&(N=n.readBool()),N){L===x&&n.readUEG(),n.readBool(),n.readUEG();for(var G=0,V=0;V<=U;V++){var F=n.readBool(),j=!1;F||(j=n.readBool()),(F||j)&&G++}U=G}else{var z=n.readUEG(),H=n.readUEG();U=z+H;for(V=0;V<z;V++)n.readUEG(),n.readBool();for(V=0;V<H;V++)n.readUEG(),n.readBool()}}if(n.readBool()){var q=n.readUEG();for(L=0;L<q;L++){for(V=0;V<I+4;V++)n.readBits(1);n.readBits(1)}}var K=0,W=1,X=1,Y=!1,J=1,Z=1;n.readBool(),n.readBool();if(n.readBool()){if(n.readBool()){var Q=n.readByte();Q>0&&Q<=16?(W=[1,12,10,16,40,24,20,32,80,18,15,64,160,4,3,2][Q-1],X=[1,11,11,11,33,11,11,11,33,11,11,33,99,3,2,1][Q-1]):255===Q&&(W=n.readBits(16),X=n.readBits(16))}if(n.readBool()&&n.readBool(),n.readBool())n.readBits(3),n.readBool(),n.readBool()&&(n.readByte(),n.readByte(),n.readByte());n.readBool()&&(n.readUEG(),n.readUEG());n.readBool(),n.readBool(),n.readBool();if(n.readBool()&&(n.readUEG(),n.readUEG(),n.readUEG(),n.readUEG()),n.readBool())if(J=n.readBits(32),Z=n.readBits(32),n.readBool())if(n.readUEG(),n.readBool()){var $=!1,ee=!1,te=!1;if($=n.readBool(),ee=n.readBool(),$||ee){(te=n.readBool())&&(n.readByte(),n.readBits(5),n.readBool(),n.readBits(5));n.readBits(4),n.readBits(4);te&&n.readBits(4),n.readBits(5),n.readBits(5),n.readBits(5)}for(L=0;L<=d;L++){var ie=n.readBool();Y=ie;var ne=!1,ae=1;ie||(ne=n.readBool());var re=!1;if(ne?n.readSEG():re=n.readBool(),re||(ae=n.readUEG()+1),$)for(V=0;V<ae;V++)n.readUEG(),n.readUEG(),te&&(n.readUEG(),n.readUEG());if(ee)for(V=0;V<ae;V++)n.readUEG(),n.readUEG(),te&&(n.readUEG(),n.readUEG())}}if(n.readBool()){n.readBool(),n.readBool(),n.readBool();K=n.readUEG();n.readUEG(),n.readUEG(),n.readUEG(),n.readUEG()}}n.readBool();var se="hvc1."+c+".1.L"+A+".B0",oe=k-(a+r)*(1===w||2===w?2:1),de=D-(s+o)*(1===w?2:1),_e=1;return 1!==W&&1!==X&&(_e=W/X),n.destroy(),n=null,{codec_mimetype:se,level_string:e.getLevelString(A),profile_idc:c,bit_depth:C+8,ref_frames:1,chroma_format:w,chroma_format_string:e.getChromaFormatString(w),general_level_idc:A,general_profile_space:_,general_tier_flag:h,general_profile_idc:c,general_profile_compatibility_flags_1:u,general_profile_compatibility_flags_2:l,general_profile_compatibility_flags_3:p,general_profile_compatibility_flags_4:m,general_constraint_indicator_flags_1:g,general_constraint_indicator_flags_2:v,general_constraint_indicator_flags_3:y,general_constraint_indicator_flags_4:b,general_constraint_indicator_flags_5:S,general_constraint_indicator_flags_6:E,min_spatial_segmentation_idc:K,constant_frame_rate:0,chroma_format_idc:w,bit_depth_luma_minus8:C,bit_depth_chroma_minus8:B,frame_rate:{fixed:Y,fps:Z/J,fps_den:J,fps_num:Z},sar_ratio:{width:W,height:X},codec_size:{width:oe,height:de},present_size:{width:oe*_e,height:de}}},e.parsePPS=function(t){var i=e._ebsp2rbsp(t),n=new f(i);n.readByte(),n.readByte();n.readUEG(),n.readUEG(),n.readBool(),n.readBool(),n.readBits(3),n.readBool(),n.readBool(),n.readUEG(),n.readUEG(),n.readSEG(),n.readBool(),n.readBool();if(n.readBool())n.readUEG();n.readSEG(),n.readSEG(),n.readBool(),n.readBool(),n.readBool(),n.readBool();var a=n.readBool(),r=n.readBool(),s=1;return r&&a?s=0:r?s=3:a&&(s=2),{parallelismType:s}},e.getChromaFormatString=function(e){switch(e){case 0:return"4:0:0";case 1:return"4:2:0";case 2:return"4:2:2";case 3:return"4:4:4";default:return"Unknown"}},e.getProfileString=function(e){switch(e){case 1:return"Main";case 2:return"Main10";case 3:return"MainSP";case 4:return"Rext";case 9:return"SCC";default:return"Unknown"}},e.getLevelString=function(e){return(e/30).toFixed(1)},e}();function v(e){return e.byteOffset%2==0&&e.byteLength%2==0}function y(e){return e.byteOffset%4==0&&e.byteLength%4==0}function b(e,t){for(var i=0;i<e.length;i++)if(e[i]!==t[i])return!1;return!0}var S=function(e,t){return e.byteLength===t.byteLength&&(y(e)&&y(t)?function(e,t){return b(new Uint32Array(e.buffer,e.byteOffset,e.byteLength/4),new Uint32Array(t.buffer,t.byteOffset,t.byteLength/4))}(e,t):v(e)&&v(t)?function(e,t){return b(new Uint16Array(e.buffer,e.byteOffset,e.byteLength/2),new Uint16Array(t.buffer,t.byteOffset,t.byteLength/2))}(e,t):function(e,t){return b(e,t)}(e,t))};var E,A=function(){function e(e,t){this.TAG="FLVDemuxer",this._config=t,this._onError=null,this._onMediaInfo=null,this._onMetaDataArrived=null,this._onScriptDataArrived=null,this._onTrackMetadata=null,this._onDataAvailable=null,this._dataOffset=e.dataOffset,this._firstParse=!0,this._dispatch=!1,this._hasAudio=e.hasAudioTrack,this._hasVideo=e.hasVideoTrack,this._hasAudioFlagOverrided=!1,this._hasVideoFlagOverrided=!1,this._audioInitialMetadataDispatched=!1,this._videoInitialMetadataDispatched=!1,this._mediaInfo=new o.a,this._mediaInfo.hasAudio=this._hasAudio,this._mediaInfo.hasVideo=this._hasVideo,this._metadata=null,this._audioMetadata=null,this._videoMetadata=null,this._naluLengthSize=4,this._timestampBase=0,this._timescale=1e3,this._duration=0,this._durationOverrided=!1,this._referenceFrameRate={fixed:!0,fps:23.976,fps_num:23976,fps_den:1e3},this._flvSoundRateTable=[5500,11025,22050,44100,48e3],this._mpegSamplingRates=[96e3,88200,64e3,48e3,44100,32e3,24e3,22050,16e3,12e3,11025,8e3,7350],this._mpegAudioV10SampleRateTable=[44100,48e3,32e3,0],this._mpegAudioV20SampleRateTable=[22050,24e3,16e3,0],this._mpegAudioV25SampleRateTable=[11025,12e3,8e3,0],this._mpegAudioL1BitRateTable=[0,32,64,96,128,160,192,224,256,288,320,352,384,416,448,-1],this._mpegAudioL2BitRateTable=[0,32,48,56,64,80,96,112,128,160,192,224,256,320,384,-1],this._mpegAudioL3BitRateTable=[0,32,40,48,56,64,80,96,112,128,160,192,224,256,320,-1],this._videoTrack={type:"video",id:1,sequenceNumber:0,samples:[],length:0},this._audioTrack={type:"audio",id:2,sequenceNumber:0,samples:[],length:0},this._littleEndian=function(){var e=new ArrayBuffer(2);return new DataView(e).setInt16(0,256,!0),256===new Int16Array(e)[0]}()}return e.prototype.destroy=function(){this._mediaInfo=null,this._metadata=null,this._audioMetadata=null,this._videoMetadata=null,this._videoTrack=null,this._audioTrack=null,this._onError=null,this._onMediaInfo=null,this._onMetaDataArrived=null,this._onScriptDataArrived=null,this._onTrackMetadata=null,this._onDataAvailable=null},e.probe=function(e){var t=new Uint8Array(e);if(t.byteLength<9)return{needMoreData:!0};var i={match:!1};if(70!==t[0]||76!==t[1]||86!==t[2]||1!==t[3])return i;var n,a,r=(4&t[4])>>>2!=0,s=0!=(1&t[4]),o=(n=t)[a=5]<<24|n[a+1]<<16|n[a+2]<<8|n[a+3];return o<9?i:{match:!0,consumed:o,dataOffset:o,hasAudioTrack:r,hasVideoTrack:s}},e.prototype.bindDataSource=function(e){return e.onDataArrival=this.parseChunks.bind(this),this},Object.defineProperty(e.prototype,"onTrackMetadata",{get:function(){return this._onTrackMetadata},set:function(e){this._onTrackMetadata=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onMediaInfo",{get:function(){return this._onMediaInfo},set:function(e){this._onMediaInfo=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onMetaDataArrived",{get:function(){return this._onMetaDataArrived},set:function(e){this._onMetaDataArrived=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onScriptDataArrived",{get:function(){return this._onScriptDataArrived},set:function(e){this._onScriptDataArrived=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onError",{get:function(){return this._onError},set:function(e){this._onError=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onDataAvailable",{get:function(){return this._onDataAvailable},set:function(e){this._onDataAvailable=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"timestampBase",{get:function(){return this._timestampBase},set:function(e){this._timestampBase=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"overridedDuration",{get:function(){return this._duration},set:function(e){this._durationOverrided=!0,this._duration=e,this._mediaInfo.duration=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"overridedHasAudio",{set:function(e){this._hasAudioFlagOverrided=!0,this._hasAudio=e,this._mediaInfo.hasAudio=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"overridedHasVideo",{set:function(e){this._hasVideoFlagOverrided=!0,this._hasVideo=e,this._mediaInfo.hasVideo=e},enumerable:!1,configurable:!0}),e.prototype.resetMediaInfo=function(){this._mediaInfo=new o.a},e.prototype._isInitialMetadataDispatched=function(){return this._hasAudio&&this._hasVideo?this._audioInitialMetadataDispatched&&this._videoInitialMetadataDispatched:this._hasAudio&&!this._hasVideo?this._audioInitialMetadataDispatched:!(this._hasAudio||!this._hasVideo)&&this._videoInitialMetadataDispatched},e.prototype.parseChunks=function(t,i){if(!(this._onError&&this._onMediaInfo&&this._onTrackMetadata&&this._onDataAvailable))throw new c.a("Flv: onError & onMediaInfo & onTrackMetadata & onDataAvailable callback must be specified");var n=0,a=this._littleEndian;if(0===i){if(!(t.byteLength>13))return 0;n=e.probe(t).dataOffset}this._firstParse&&(this._firstParse=!1,i+n!==this._dataOffset&&r.a.w(this.TAG,"First time parsing but chunk byteStart invalid!"),0!==(s=new DataView(t,n)).getUint32(0,!a)&&r.a.w(this.TAG,"PrevTagSize0 !== 0 !!!"),n+=4);for(;n<t.byteLength;){this._dispatch=!0;var s=new DataView(t,n);if(n+11+4>t.byteLength)break;var o=s.getUint8(0),d=16777215&s.getUint32(0,!a);if(n+11+d+4>t.byteLength)break;if(8===o||9===o||18===o){var _=s.getUint8(4),h=s.getUint8(5),u=s.getUint8(6)|h<<8|_<<16|s.getUint8(7)<<24;0!==(16777215&s.getUint32(7,!a))&&r.a.w(this.TAG,"Meet tag which has StreamID != 0!");var l=n+11;switch(o){case 8:this._parseAudioData(t,l,d,u);break;case 9:this._parseVideoData(t,l,d,u,i+n);break;case 18:this._parseScriptData(t,l,d)}var f=s.getUint32(11+d,!a);f!==11+d&&r.a.w(this.TAG,"Invalid PrevTagSize "+f),n+=11+d+4}else r.a.w(this.TAG,"Unsupported tag type "+o+", skipped"),n+=11+d+4}return this._isInitialMetadataDispatched()&&this._dispatch&&(this._audioTrack.length||this._videoTrack.length)&&this._onDataAvailable(this._audioTrack,this._videoTrack),n},e.prototype._parseScriptData=function(e,t,i){var n=l.parseScriptData(e,t,i);if(n.hasOwnProperty("onMetaData")){if(null==n.onMetaData||"object"!=typeof n.onMetaData)return void r.a.w(this.TAG,"Invalid onMetaData structure!");this._metadata&&r.a.w(this.TAG,"Found another onMetaData tag!"),this._metadata=n;var a=this._metadata.onMetaData;if(this._onMetaDataArrived&&this._onMetaDataArrived(Object.assign({},a)),"boolean"==typeof a.hasAudio&&!1===this._hasAudioFlagOverrided&&(this._hasAudio=a.hasAudio,this._mediaInfo.hasAudio=this._hasAudio),"boolean"==typeof a.hasVideo&&!1===this._hasVideoFlagOverrided&&(this._hasVideo=a.hasVideo,this._mediaInfo.hasVideo=this._hasVideo),"number"==typeof a.audiodatarate&&(this._mediaInfo.audioDataRate=a.audiodatarate),"number"==typeof a.videodatarate&&(this._mediaInfo.videoDataRate=a.videodatarate),"number"==typeof a.width&&(this._mediaInfo.width=a.width),"number"==typeof a.height&&(this._mediaInfo.height=a.height),"number"==typeof a.duration){if(!this._durationOverrided){var s=Math.floor(a.duration*this._timescale);this._duration=s,this._mediaInfo.duration=s}}else this._mediaInfo.duration=0;if("number"==typeof a.framerate){var o=Math.floor(1e3*a.framerate);if(o>0){var d=o/1e3;this._referenceFrameRate.fixed=!0,this._referenceFrameRate.fps=d,this._referenceFrameRate.fps_num=o,this._referenceFrameRate.fps_den=1e3,this._mediaInfo.fps=d}}if("object"==typeof a.keyframes){this._mediaInfo.hasKeyframesIndex=!0;var _=a.keyframes;this._mediaInfo.keyframesIndex=this._parseKeyframesIndex(_),a.keyframes=null}else this._mediaInfo.hasKeyframesIndex=!1;this._dispatch=!1,this._mediaInfo.metadata=a,r.a.v(this.TAG,"Parsed onMetaData"),this._mediaInfo.isComplete()&&this._onMediaInfo(this._mediaInfo)}Object.keys(n).length>0&&this._onScriptDataArrived&&this._onScriptDataArrived(Object.assign({},n))},e.prototype._parseKeyframesIndex=function(e){for(var t=[],i=[],n=1;n<e.times.length;n++){var a=this._timestampBase+Math.floor(1e3*e.times[n]);t.push(a),i.push(e.filepositions[n])}return{times:t,filepositions:i}},e.prototype._parseAudioData=function(e,t,i,n){if(i<=1)r.a.w(this.TAG,"Flv: Invalid audio packet, missing SoundData payload!");else if(!0!==this._hasAudioFlagOverrided||!1!==this._hasAudio){this._littleEndian;var a=new DataView(e,t,i).getUint8(0),s=a>>>4;if(2===s||10===s){var o=0,d=(12&a)>>>2;if(d>=0&&d<=4){o=this._flvSoundRateTable[d];var _=1&a,h=this._audioMetadata,c=this._audioTrack;if(h||(!1===this._hasAudio&&!1===this._hasAudioFlagOverrided&&(this._hasAudio=!0,this._mediaInfo.hasAudio=!0),(h=this._audioMetadata={}).type="audio",h.id=c.id,h.timescale=this._timescale,h.duration=this._duration,h.audioSampleRate=o,h.channelCount=0===_?1:2),10===s){var u=this._parseAACAudioData(e,t+1,i-1);if(null==u)return;if(0===u.packetType){if(h.config){if(S(u.data.config,h.config))return;r.a.w(this.TAG,"AudioSpecificConfig has been changed, re-generate initialization segment")}var l=u.data;h.audioSampleRate=l.samplingRate,h.channelCount=l.channelCount,h.codec=l.codec,h.originalCodec=l.originalCodec,h.config=l.config,h.refSampleDuration=1024/h.audioSampleRate*h.timescale,r.a.v(this.TAG,"Parsed AudioSpecificConfig"),this._isInitialMetadataDispatched()?this._dispatch&&(this._audioTrack.length||this._videoTrack.length)&&this._onDataAvailable(this._audioTrack,this._videoTrack):this._audioInitialMetadataDispatched=!0,this._dispatch=!1,this._onTrackMetadata("audio",h),(g=this._mediaInfo).audioCodec=h.originalCodec,g.audioSampleRate=h.audioSampleRate,g.audioChannelCount=h.channelCount,g.hasVideo?null!=g.videoCodec&&(g.mimeType='video/x-flv; codecs="'+g.videoCodec+","+g.audioCodec+'"'):g.mimeType='video/x-flv; codecs="'+g.audioCodec+'"',g.isComplete()&&this._onMediaInfo(g)}else if(1===u.packetType){var f=this._timestampBase+n,p={unit:u.data,length:u.data.byteLength,dts:f,pts:f};c.samples.push(p),c.length+=u.data.length}else r.a.e(this.TAG,"Flv: Unsupported AAC data type "+u.packetType)}else if(2===s){if(!h.codec){var g;if(null==(l=this._parseMP3AudioData(e,t+1,i-1,!0)))return;h.audioSampleRate=l.samplingRate,h.channelCount=l.channelCount,h.codec=l.codec,h.originalCodec=l.originalCodec,h.refSampleDuration=1152/h.audioSampleRate*h.timescale,r.a.v(this.TAG,"Parsed MPEG Audio Frame Header"),this._audioInitialMetadataDispatched=!0,this._onTrackMetadata("audio",h),(g=this._mediaInfo).audioCodec=h.codec,g.audioSampleRate=h.audioSampleRate,g.audioChannelCount=h.channelCount,g.audioDataRate=l.bitRate,g.hasVideo?null!=g.videoCodec&&(g.mimeType='video/x-flv; codecs="'+g.videoCodec+","+g.audioCodec+'"'):g.mimeType='video/x-flv; codecs="'+g.audioCodec+'"',g.isComplete()&&this._onMediaInfo(g)}var v=this._parseMP3AudioData(e,t+1,i-1,!1);if(null==v)return;f=this._timestampBase+n;var y={unit:v,length:v.byteLength,dts:f,pts:f};c.samples.push(y),c.length+=v.length}}else this._onError(m.a.FORMAT_ERROR,"Flv: Invalid audio sample rate idx: "+d)}else this._onError(m.a.CODEC_UNSUPPORTED,"Flv: Unsupported audio codec idx: "+s)}},e.prototype._parseAACAudioData=function(e,t,i){if(!(i<=1)){var n={},a=new Uint8Array(e,t,i);return n.packetType=a[0],0===a[0]?n.data=this._parseAACAudioSpecificConfig(e,t+1,i-1):n.data=a.subarray(1),n}r.a.w(this.TAG,"Flv: Invalid AAC packet, missing AACPacketType or/and Data!")},e.prototype._parseAACAudioSpecificConfig=function(e,t,i){var n,a,r=new Uint8Array(e,t,i),s=null,o=0,d=null;if(o=n=r[0]>>>3,(a=(7&r[0])<<1|r[1]>>>7)<0||a>=this._mpegSamplingRates.length)this._onError(m.a.FORMAT_ERROR,"Flv: AAC invalid sampling frequency index!");else{var _=this._mpegSamplingRates[a],h=(120&r[1])>>>3;if(!(h<0||h>=8)){5===o&&(d=(7&r[1])<<1|r[2]>>>7,(124&r[2])>>>2);var c=self.navigator.userAgent.toLowerCase();return-1!==c.indexOf("firefox")?a>=6?(o=5,s=new Array(4),d=a-3):(o=2,s=new Array(2),d=a):-1!==c.indexOf("android")?(o=2,s=new Array(2),d=a):(o=5,d=a,s=new Array(4),a>=6?d=a-3:1===h&&(o=2,s=new Array(2),d=a)),s[0]=o<<3,s[0]|=(15&a)>>>1,s[1]=(15&a)<<7,s[1]|=(15&h)<<3,5===o&&(s[1]|=(15&d)>>>1,s[2]=(1&d)<<7,s[2]|=8,s[3]=0),{config:s,samplingRate:_,channelCount:h,codec:"mp4a.40."+o,originalCodec:"mp4a.40."+n}}this._onError(m.a.FORMAT_ERROR,"Flv: AAC invalid channel configuration")}},e.prototype._parseMP3AudioData=function(e,t,i,n){if(!(i<4)){this._littleEndian;var a=new Uint8Array(e,t,i),s=null;if(n){if(255!==a[0])return;var o=a[1]>>>3&3,d=(6&a[1])>>1,_=(240&a[2])>>>4,h=(12&a[2])>>>2,c=3!==(a[3]>>>6&3)?2:1,u=0,l=0;switch(o){case 0:u=this._mpegAudioV25SampleRateTable[h];break;case 2:u=this._mpegAudioV20SampleRateTable[h];break;case 3:u=this._mpegAudioV10SampleRateTable[h]}switch(d){case 1:34,_<this._mpegAudioL3BitRateTable.length&&(l=this._mpegAudioL3BitRateTable[_]);break;case 2:33,_<this._mpegAudioL2BitRateTable.length&&(l=this._mpegAudioL2BitRateTable[_]);break;case 3:32,_<this._mpegAudioL1BitRateTable.length&&(l=this._mpegAudioL1BitRateTable[_])}s={bitRate:l,samplingRate:u,channelCount:c,codec:"mp3",originalCodec:"mp3"}}else s=a;return s}r.a.w(this.TAG,"Flv: Invalid MP3 packet, header missing!")},e.prototype._parseVideoData=function(e,t,i,n,a){if(i<=1)r.a.w(this.TAG,"Flv: Invalid video packet, missing VideoData payload!");else if(!0!==this._hasVideoFlagOverrided||!1!==this._hasVideo){var s=new Uint8Array(e,t,i)[0],o=(112&s)>>>4;if(0!=(128&s)){var d=15&s,_=String.fromCharCode.apply(String,new Uint8Array(e,t,i).slice(1,5));if("hvc1"!==_)return void this._onError(m.a.CODEC_UNSUPPORTED,"Flv: Unsupported codec in video frame: "+_);this._parseEnhancedHEVCVideoPacket(e,t+5,i-5,n,a,o,d)}else{var h=15&s;if(7===h)this._parseAVCVideoPacket(e,t+1,i-1,n,a,o);else{if(12!==h)return void this._onError(m.a.CODEC_UNSUPPORTED,"Flv: Unsupported codec in video frame: "+h);this._parseHEVCVideoPacket(e,t+1,i-1,n,a,o)}}}},e.prototype._parseAVCVideoPacket=function(e,t,i,n,a,s){if(i<4)r.a.w(this.TAG,"Flv: Invalid AVC packet, missing AVCPacketType or/and CompositionTime");else{var o=this._littleEndian,d=new DataView(e,t,i),_=d.getUint8(0),h=(16777215&d.getUint32(0,!o))<<8>>8;if(0===_)this._parseAVCDecoderConfigurationRecord(e,t+4,i-4);else if(1===_)this._parseAVCVideoData(e,t+4,i-4,n,a,s,h);else if(2!==_)return void this._onError(m.a.FORMAT_ERROR,"Flv: Invalid video packet type "+_)}},e.prototype._parseHEVCVideoPacket=function(e,t,i,n,a,s){if(i<4)r.a.w(this.TAG,"Flv: Invalid HEVC packet, missing HEVCPacketType or/and CompositionTime");else{var o=this._littleEndian,d=new DataView(e,t,i),_=d.getUint8(0),h=(16777215&d.getUint32(0,!o))<<8>>8;if(0===_)this._parseHEVCDecoderConfigurationRecord(e,t+4,i-4);else if(1===_)this._parseHEVCVideoData(e,t+4,i-4,n,a,s,h);else if(2!==_)return void this._onError(m.a.FORMAT_ERROR,"Flv: Invalid video packet type "+_)}},e.prototype._parseEnhancedHEVCVideoPacket=function(e,t,i,n,a,s,o){if(i<4)r.a.w(this.TAG,"Flv: Invalid HEVC packet, missing HEVCPacketType or/and CompositionTime");else{var d=this._littleEndian,_=new DataView(e,t,i);if(0===o)this._parseHEVCDecoderConfigurationRecord(e,t,i);else if(1===o){var h=(4294967040&_.getUint32(0,!d))>>8;this._parseHEVCVideoData(e,t+3,i-3,n,a,s,h)}else if(3===o)this._parseHEVCVideoData(e,t,i,n,a,s,0);else if(2!==o)return void this._onError(m.a.FORMAT_ERROR,"Flv: Invalid video packet type "+o)}},e.prototype._parseAVCDecoderConfigurationRecord=function(e,t,i){if(i<7)r.a.w(this.TAG,"Flv: Invalid AVCDecoderConfigurationRecord, lack of data!");else{var n=this._videoMetadata,a=this._videoTrack,s=this._littleEndian,o=new DataView(e,t,i);if(n){if(void 0!==n.avcc){var d=new Uint8Array(e,t,i);if(S(d,n.avcc))return;r.a.w(this.TAG,"AVCDecoderConfigurationRecord has been changed, re-generate initialization segment")}}else!1===this._hasVideo&&!1===this._hasVideoFlagOverrided&&(this._hasVideo=!0,this._mediaInfo.hasVideo=!0),(n=this._videoMetadata={}).type="video",n.id=a.id,n.timescale=this._timescale,n.duration=this._duration;var _=o.getUint8(0),h=o.getUint8(1);o.getUint8(2),o.getUint8(3);if(1===_&&0!==h)if(this._naluLengthSize=1+(3&o.getUint8(4)),3===this._naluLengthSize||4===this._naluLengthSize){var c=31&o.getUint8(5);if(0!==c){c>1&&r.a.w(this.TAG,"Flv: Strange AVCDecoderConfigurationRecord: SPS Count = "+c);for(var u=6,l=0;l<c;l++){var f=o.getUint16(u,!s);if(u+=2,0!==f){var g=new Uint8Array(e,t+u,f);u+=f;var v=p.parseSPS(g);if(0===l){n.codecWidth=v.codec_size.width,n.codecHeight=v.codec_size.height,n.presentWidth=v.present_size.width,n.presentHeight=v.present_size.height,n.profile=v.profile_string,n.level=v.level_string,n.bitDepth=v.bit_depth,n.chromaFormat=v.chroma_format,n.sarRatio=v.sar_ratio,n.frameRate=v.frame_rate,!1!==v.frame_rate.fixed&&0!==v.frame_rate.fps_num&&0!==v.frame_rate.fps_den||(n.frameRate=this._referenceFrameRate);var y=n.frameRate.fps_den,b=n.frameRate.fps_num;n.refSampleDuration=n.timescale*(y/b);for(var E=g.subarray(1,4),A="avc1.",R=0;R<3;R++){var T=E[R].toString(16);T.length<2&&(T="0"+T),A+=T}n.codec=A;var L=this._mediaInfo;L.width=n.codecWidth,L.height=n.codecHeight,L.fps=n.frameRate.fps,L.profile=n.profile,L.level=n.level,L.refFrames=v.ref_frames,L.chromaFormat=v.chroma_format_string,L.sarNum=n.sarRatio.width,L.sarDen=n.sarRatio.height,L.videoCodec=A,L.hasAudio?null!=L.audioCodec&&(L.mimeType='video/x-flv; codecs="'+L.videoCodec+","+L.audioCodec+'"'):L.mimeType='video/x-flv; codecs="'+L.videoCodec+'"',L.isComplete()&&this._onMediaInfo(L)}}}var w=o.getUint8(u);if(0!==w){w>1&&r.a.w(this.TAG,"Flv: Strange AVCDecoderConfigurationRecord: PPS Count = "+w),u++;for(l=0;l<w;l++){f=o.getUint16(u,!s);u+=2,0!==f&&(u+=f)}n.avcc=new Uint8Array(i),n.avcc.set(new Uint8Array(e,t,i),0),r.a.v(this.TAG,"Parsed AVCDecoderConfigurationRecord"),this._isInitialMetadataDispatched()?this._dispatch&&(this._audioTrack.length||this._videoTrack.length)&&this._onDataAvailable(this._audioTrack,this._videoTrack):this._videoInitialMetadataDispatched=!0,this._dispatch=!1,this._onTrackMetadata("video",n)}else this._onError(m.a.FORMAT_ERROR,"Flv: Invalid AVCDecoderConfigurationRecord: No PPS")}else this._onError(m.a.FORMAT_ERROR,"Flv: Invalid AVCDecoderConfigurationRecord: No SPS")}else this._onError(m.a.FORMAT_ERROR,"Flv: Strange NaluLengthSizeMinusOne: "+(this._naluLengthSize-1));else this._onError(m.a.FORMAT_ERROR,"Flv: Invalid AVCDecoderConfigurationRecord")}},e.prototype._parseHEVCDecoderConfigurationRecord=function(e,t,i){if(i<22)r.a.w(this.TAG,"Flv: Invalid HEVCDecoderConfigurationRecord, lack of data!");else{var n=this._videoMetadata,a=this._videoTrack,s=this._littleEndian,o=new DataView(e,t,i);if(n){if(void 0!==n.hvcc){var d=new Uint8Array(e,t,i);if(S(d,n.hvcc))return;r.a.w(this.TAG,"HEVCDecoderConfigurationRecord has been changed, re-generate initialization segment")}}else!1===this._hasVideo&&!1===this._hasVideoFlagOverrided&&(this._hasVideo=!0,this._mediaInfo.hasVideo=!0),(n=this._videoMetadata={}).type="video",n.id=a.id,n.timescale=this._timescale,n.duration=this._duration;var _=o.getUint8(0),h=31&o.getUint8(1);if(1===_&&0!==h)if(this._naluLengthSize=1+(3&o.getUint8(21)),3===this._naluLengthSize||4===this._naluLengthSize){for(var c=o.getUint8(22),u=0,l=23;u<c;u++){var f=63&o.getUint8(l+0),p=o.getUint16(l+1,!s);l+=3;for(var v=0;v<p;v++){var y=o.getUint16(l+0,!s);if(0===v)if(33===f){l+=2;var b=new Uint8Array(e,t+l,y),E=g.parseSPS(b);n.codecWidth=E.codec_size.width,n.codecHeight=E.codec_size.height,n.presentWidth=E.present_size.width,n.presentHeight=E.present_size.height,n.profile=E.profile_string,n.level=E.level_string,n.bitDepth=E.bit_depth,n.chromaFormat=E.chroma_format,n.sarRatio=E.sar_ratio,n.frameRate=E.frame_rate,!1!==E.frame_rate.fixed&&0!==E.frame_rate.fps_num&&0!==E.frame_rate.fps_den||(n.frameRate=this._referenceFrameRate);var A=n.frameRate.fps_den,R=n.frameRate.fps_num;n.refSampleDuration=n.timescale*(A/R),n.codec=E.codec_mimetype;var T=this._mediaInfo;T.width=n.codecWidth,T.height=n.codecHeight,T.fps=n.frameRate.fps,T.profile=n.profile,T.level=n.level,T.refFrames=E.ref_frames,T.chromaFormat=E.chroma_format_string,T.sarNum=n.sarRatio.width,T.sarDen=n.sarRatio.height,T.videoCodec=E.codec_mimetype,T.hasAudio?null!=T.audioCodec&&(T.mimeType='video/x-flv; codecs="'+T.videoCodec+","+T.audioCodec+'"'):T.mimeType='video/x-flv; codecs="'+T.videoCodec+'"',T.isComplete()&&this._onMediaInfo(T),l+=y}else l+=2+y;else l+=2+y}}n.hvcc=new Uint8Array(i),n.hvcc.set(new Uint8Array(e,t,i),0),r.a.v(this.TAG,"Parsed HEVCDecoderConfigurationRecord"),this._isInitialMetadataDispatched()?this._dispatch&&(this._audioTrack.length||this._videoTrack.length)&&this._onDataAvailable(this._audioTrack,this._videoTrack):this._videoInitialMetadataDispatched=!0,this._dispatch=!1,this._onTrackMetadata("video",n)}else this._onError(m.a.FORMAT_ERROR,"Flv: Strange NaluLengthSizeMinusOne: "+(this._naluLengthSize-1));else this._onError(m.a.FORMAT_ERROR,"Flv: Invalid HEVCDecoderConfigurationRecord")}},e.prototype._parseAVCVideoData=function(e,t,i,n,a,s,o){for(var d=this._littleEndian,_=new DataView(e,t,i),h=[],c=0,u=0,l=this._naluLengthSize,f=this._timestampBase+n,p=1===s;u<i;){if(u+4>=i){r.a.w(this.TAG,"Malformed Nalu near timestamp "+f+", offset = "+u+", dataSize = "+i);break}var m=_.getUint32(u,!d);if(3===l&&(m>>>=8),m>i-l)return void r.a.w(this.TAG,"Malformed Nalus near timestamp "+f+", NaluSize > DataSize!");var g=31&_.getUint8(u+l);5===g&&(p=!0);var v=new Uint8Array(e,t+u,l+m),y={type:g,data:v};h.push(y),c+=v.byteLength,u+=l+m}if(h.length){var b=this._videoTrack,S={units:h,length:c,isKeyframe:p,dts:f,cts:o,pts:f+o};p&&(S.fileposition=a),b.samples.push(S),b.length+=c}},e.prototype._parseHEVCVideoData=function(e,t,i,n,a,s,o){for(var d=this._littleEndian,_=new DataView(e,t,i),h=[],c=0,u=0,l=this._naluLengthSize,f=this._timestampBase+n,p=1===s;u<i;){if(u+4>=i){r.a.w(this.TAG,"Malformed Nalu near timestamp "+f+", offset = "+u+", dataSize = "+i);break}var m=_.getUint32(u,!d);if(3===l&&(m>>>=8),m>i-l)return void r.a.w(this.TAG,"Malformed Nalus near timestamp "+f+", NaluSize > DataSize!");var g=31&_.getUint8(u+l);19!==g&&20!==g||(p=!0);var v=new Uint8Array(e,t+u,l+m),y={type:g,data:v};h.push(y),c+=v.byteLength,u+=l+m}if(h.length){var b=this._videoTrack,S={units:h,length:c,isKeyframe:p,dts:f,cts:o,pts:f+o};p&&(S.fileposition=a),b.samples.push(S),b.length+=c}},e}(),R=function(){function e(){}return e.prototype.destroy=function(){this.onError=null,this.onMediaInfo=null,this.onMetaDataArrived=null,this.onTrackMetadata=null,this.onDataAvailable=null,this.onTimedID3Metadata=null,this.onSMPTE2038Metadata=null,this.onSCTE35Metadata=null,this.onPESPrivateData=null,this.onPESPrivateDataDescriptor=null},e}(),T=function(){this.program_pmt_pid={}};!function(e){e[e.kMPEG1Audio=3]="kMPEG1Audio",e[e.kMPEG2Audio=4]="kMPEG2Audio",e[e.kPESPrivateData=6]="kPESPrivateData",e[e.kADTSAAC=15]="kADTSAAC",e[e.kLOASAAC=17]="kLOASAAC",e[e.kAC3=129]="kAC3",e[e.kID3=21]="kID3",e[e.kSCTE35=134]="kSCTE35",e[e.kH264=27]="kH264",e[e.kH265=36]="kH265"}(E||(E={}));var L,w=function(){this.pid_stream_type={},this.common_pids={h264:void 0,h265:void 0,adts_aac:void 0,loas_aac:void 0,opus:void 0,ac3:void 0,mp3:void 0},this.pes_private_data_pids={},this.timed_id3_pids={},this.scte_35_pids={},this.smpte2038_pids={}},k=function(){},D=function(){},C=function(){this.slices=[],this.total_length=0,this.expected_length=0,this.file_position=0};!function(e){e[e.kUnspecified=0]="kUnspecified",e[e.kSliceNonIDR=1]="kSliceNonIDR",e[e.kSliceDPA=2]="kSliceDPA",e[e.kSliceDPB=3]="kSliceDPB",e[e.kSliceDPC=4]="kSliceDPC",e[e.kSliceIDR=5]="kSliceIDR",e[e.kSliceSEI=6]="kSliceSEI",e[e.kSliceSPS=7]="kSliceSPS",e[e.kSlicePPS=8]="kSlicePPS",e[e.kSliceAUD=9]="kSliceAUD",e[e.kEndOfSequence=10]="kEndOfSequence",e[e.kEndOfStream=11]="kEndOfStream",e[e.kFiller=12]="kFiller",e[e.kSPSExt=13]="kSPSExt",e[e.kReserved0=14]="kReserved0"}(L||(L={}));var B,I,O=function(){},P=function(e){var t=e.data.byteLength;this.type=e.type,this.data=new Uint8Array(4+t),new DataView(this.data.buffer).setUint32(0,t),this.data.set(e.data,4)},M=function(){function e(e){this.TAG="H264AnnexBParser",this.current_startcode_offset_=0,this.eof_flag_=!1,this.data_=e,this.current_startcode_offset_=this.findNextStartCodeOffset(0),this.eof_flag_&&r.a.e(this.TAG,"Could not find H264 startcode until payload end!")}return e.prototype.findNextStartCodeOffset=function(e){for(var t=e,i=this.data_;;){if(t+3>=i.byteLength)return this.eof_flag_=!0,i.byteLength;var n=i[t+0]<<24|i[t+1]<<16|i[t+2]<<8|i[t+3],a=i[t+0]<<16|i[t+1]<<8|i[t+2];if(1===n||1===a)return t;t++}},e.prototype.readNextNaluPayload=function(){for(var e=this.data_,t=null;null==t&&!this.eof_flag_;){var i=this.current_startcode_offset_,n=31&e[i+=1===(e[i]<<24|e[i+1]<<16|e[i+2]<<8|e[i+3])?4:3],a=(128&e[i])>>>7,r=this.findNextStartCodeOffset(i);if(this.current_startcode_offset_=r,!(n>=L.kReserved0)&&0===a){var s=e.subarray(i,r);(t=new O).type=n,t.data=s}}return t},e}(),x=function(){function e(e,t,i){var n=8+e.byteLength+1+2+t.byteLength,a=!1;66!==e[3]&&77!==e[3]&&88!==e[3]&&(a=!0,n+=4);var r=this.data=new Uint8Array(n);r[0]=1,r[1]=e[1],r[2]=e[2],r[3]=e[3],r[4]=255,r[5]=225;var s=e.byteLength;r[6]=s>>>8,r[7]=255&s;var o=8;r.set(e,8),r[o+=s]=1;var d=t.byteLength;r[o+1]=d>>>8,r[o+2]=255&d,r.set(t,o+3),o+=3+d,a&&(r[o]=252|i.chroma_format_idc,r[o+1]=248|i.bit_depth_luma-8,r[o+2]=248|i.bit_depth_chroma-8,r[o+3]=0,o+=4)}return e.prototype.getData=function(){return this.data},e}();!function(e){e[e.kNull=0]="kNull",e[e.kAACMain=1]="kAACMain",e[e.kAAC_LC=2]="kAAC_LC",e[e.kAAC_SSR=3]="kAAC_SSR",e[e.kAAC_LTP=4]="kAAC_LTP",e[e.kAAC_SBR=5]="kAAC_SBR",e[e.kAAC_Scalable=6]="kAAC_Scalable",e[e.kLayer1=32]="kLayer1",e[e.kLayer2=33]="kLayer2",e[e.kLayer3=34]="kLayer3"}(B||(B={})),function(e){e[e.k96000Hz=0]="k96000Hz",e[e.k88200Hz=1]="k88200Hz",e[e.k64000Hz=2]="k64000Hz",e[e.k48000Hz=3]="k48000Hz",e[e.k44100Hz=4]="k44100Hz",e[e.k32000Hz=5]="k32000Hz",e[e.k24000Hz=6]="k24000Hz",e[e.k22050Hz=7]="k22050Hz",e[e.k16000Hz=8]="k16000Hz",e[e.k12000Hz=9]="k12000Hz",e[e.k11025Hz=10]="k11025Hz",e[e.k8000Hz=11]="k8000Hz",e[e.k7350Hz=12]="k7350Hz"}(I||(I={}));var U,N,G=[96e3,88200,64e3,48e3,44100,32e3,24e3,22050,16e3,12e3,11025,8e3,7350],V=(U=function(e,t){return(U=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(e,t){e.__proto__=t}||function(e,t){for(var i in t)t.hasOwnProperty(i)&&(e[i]=t[i])})(e,t)},function(e,t){function i(){this.constructor=e}U(e,t),e.prototype=null===t?Object.create(t):(i.prototype=t.prototype,new i)}),F=function(){},j=function(e){function t(){return null!==e&&e.apply(this,arguments)||this}return V(t,e),t}(F),z=function(){function e(e){this.TAG="AACADTSParser",this.data_=e,this.current_syncword_offset_=this.findNextSyncwordOffset(0),this.eof_flag_&&r.a.e(this.TAG,"Could not found ADTS syncword until payload end")}return e.prototype.findNextSyncwordOffset=function(e){for(var t=e,i=this.data_;;){if(t+7>=i.byteLength)return this.eof_flag_=!0,i.byteLength;if(4095===(i[t+0]<<8|i[t+1])>>>4)return t;t++}},e.prototype.readNextAACFrame=function(){for(var e=this.data_,t=null;null==t&&!this.eof_flag_;){var i=this.current_syncword_offset_,n=(8&e[i+1])>>>3,a=(6&e[i+1])>>>1,r=1&e[i+1],s=(192&e[i+2])>>>6,o=(60&e[i+2])>>>2,d=(1&e[i+2])<<2|(192&e[i+3])>>>6,_=(3&e[i+3])<<11|e[i+4]<<3|(224&e[i+5])>>>5;e[i+6];if(i+_>this.data_.byteLength){this.eof_flag_=!0,this.has_last_incomplete_data=!0;break}var h=1===r?7:9,c=_-h;i+=h;var u=this.findNextSyncwordOffset(i+c);if(this.current_syncword_offset_=u,(0===n||1===n)&&0===a){var l=e.subarray(i,i+c);(t=new F).audio_object_type=s+1,t.sampling_freq_index=o,t.sampling_frequency=G[o],t.channel_config=d,t.data=l}}return t},e.prototype.hasIncompleteData=function(){return this.has_last_incomplete_data},e.prototype.getIncompleteData=function(){return this.has_last_incomplete_data?this.data_.subarray(this.current_syncword_offset_):null},e}(),H=function(){function e(e){this.TAG="AACLOASParser",this.data_=e,this.current_syncword_offset_=this.findNextSyncwordOffset(0),this.eof_flag_&&r.a.e(this.TAG,"Could not found LOAS syncword until payload end")}return e.prototype.findNextSyncwordOffset=function(e){for(var t=e,i=this.data_;;){if(t+1>=i.byteLength)return this.eof_flag_=!0,i.byteLength;if(695===(i[t+0]<<3|i[t+1]>>>5))return t;t++}},e.prototype.getLATMValue=function(e){for(var t=e.readBits(2),i=0,n=0;n<=t;n++)i<<=8,i|=e.readByte();return i},e.prototype.readNextAACFrame=function(e){for(var t=this.data_,i=null;null==i&&!this.eof_flag_;){var n=this.current_syncword_offset_,a=(31&t[n+1])<<8|t[n+2];if(n+3+a>=this.data_.byteLength){this.eof_flag_=!0,this.has_last_incomplete_data=!0;break}var s=new f(t.subarray(n+3,n+3+a)),o=null;if(s.readBool()){if(null==e){r.a.w(this.TAG,"StreamMuxConfig Missing"),this.current_syncword_offset_=this.findNextSyncwordOffset(n+3+a),s.destroy();continue}o=e}else{var d=s.readBool();if(d&&s.readBool()){r.a.e(this.TAG,"audioMuxVersionA is Not Supported"),s.destroy();break}if(d&&this.getLATMValue(s),!s.readBool()){r.a.e(this.TAG,"allStreamsSameTimeFraming zero is Not Supported"),s.destroy();break}if(0!==s.readBits(6)){r.a.e(this.TAG,"more than 2 numSubFrames Not Supported"),s.destroy();break}if(0!==s.readBits(4)){r.a.e(this.TAG,"more than 2 numProgram Not Supported"),s.destroy();break}if(0!==s.readBits(3)){r.a.e(this.TAG,"more than 2 numLayer Not Supported"),s.destroy();break}var _=d?this.getLATMValue(s):0,h=s.readBits(5);_-=5;var c=s.readBits(4);_-=4;var u=s.readBits(4);_-=4,s.readBits(3),(_-=3)>0&&s.readBits(_);var l=s.readBits(3);if(0!==l){r.a.e(this.TAG,"frameLengthType = "+l+". Only frameLengthType = 0 Supported"),s.destroy();break}s.readByte();var p=s.readBool();if(p)if(d)this.getLATMValue(s);else{for(var m=0;;){m<<=8;var g=s.readBool();if(m+=s.readByte(),!g)break}console.log(m)}s.readBool()&&s.readByte(),(o=new j).audio_object_type=h,o.sampling_freq_index=c,o.sampling_frequency=G[o.sampling_freq_index],o.channel_config=u,o.other_data_present=p}for(var v=0;;){var y=s.readByte();if(v+=y,255!==y)break}for(var b=new Uint8Array(v),S=0;S<v;S++)b[S]=s.readByte();(i=new j).audio_object_type=o.audio_object_type,i.sampling_freq_index=o.sampling_freq_index,i.sampling_frequency=G[o.sampling_freq_index],i.channel_config=o.channel_config,i.other_data_present=o.other_data_present,i.data=b,this.current_syncword_offset_=this.findNextSyncwordOffset(n+3+a)}return i},e.prototype.hasIncompleteData=function(){return this.has_last_incomplete_data},e.prototype.getIncompleteData=function(){return this.has_last_incomplete_data?this.data_.subarray(this.current_syncword_offset_):null},e}(),q=function(e){var t=null,i=e.audio_object_type,n=e.audio_object_type,a=e.sampling_freq_index,r=e.channel_config,s=0,o=navigator.userAgent.toLowerCase();-1!==o.indexOf("firefox")?a>=6?(n=5,t=new Array(4),s=a-3):(n=2,t=new Array(2),s=a):-1!==o.indexOf("android")?(n=2,t=new Array(2),s=a):(n=5,s=a,t=new Array(4),a>=6?s=a-3:1===r&&(n=2,t=new Array(2),s=a)),t[0]=n<<3,t[0]|=(15&a)>>>1,t[1]=(15&a)<<7,t[1]|=(15&r)<<3,5===n&&(t[1]|=(15&s)>>>1,t[2]=(1&s)<<7,t[2]|=8,t[3]=0),this.config=t,this.sampling_rate=G[a],this.channel_count=r,this.codec_mimetype="mp4a.40."+n,this.original_codec_mimetype="mp4a.40."+i},K=function(){},W=function(){};!function(e){e[e.kSpliceNull=0]="kSpliceNull",e[e.kSpliceSchedule=4]="kSpliceSchedule",e[e.kSpliceInsert=5]="kSpliceInsert",e[e.kTimeSignal=6]="kTimeSignal",e[e.kBandwidthReservation=7]="kBandwidthReservation",e[e.kPrivateCommand=255]="kPrivateCommand"}(N||(N={}));var X,Y=function(e){var t=e.readBool();return t?(e.readBits(6),{time_specified_flag:t,pts_time:4*e.readBits(31)+e.readBits(2)}):(e.readBits(7),{time_specified_flag:t})},J=function(e){var t=e.readBool();return e.readBits(6),{auto_return:t,duration:4*e.readBits(31)+e.readBits(2)}},Z=function(e,t){var i=t.readBits(8);return e?{component_tag:i}:{component_tag:i,splice_time:Y(t)}},Q=function(e){return{component_tag:e.readBits(8),utc_splice_time:e.readBits(32)}},$=function(e){var t=e.readBits(32),i=e.readBool();e.readBits(7);var n={splice_event_id:t,splice_event_cancel_indicator:i};if(i)return n;if(n.out_of_network_indicator=e.readBool(),n.program_splice_flag=e.readBool(),n.duration_flag=e.readBool(),e.readBits(5),n.program_splice_flag)n.utc_splice_time=e.readBits(32);else{n.component_count=e.readBits(8),n.components=[];for(var a=0;a<n.component_count;a++)n.components.push(Q(e))}return n.duration_flag&&(n.break_duration=J(e)),n.unique_program_id=e.readBits(16),n.avail_num=e.readBits(8),n.avails_expected=e.readBits(8),n},ee=function(e,t,i,n){return{descriptor_tag:e,descriptor_length:t,identifier:i,provider_avail_id:n.readBits(32)}},te=function(e,t,i,n){var a=n.readBits(8),r=n.readBits(3);n.readBits(5);for(var s="",o=0;o<r;o++)s+=String.fromCharCode(n.readBits(8));return{descriptor_tag:e,descriptor_length:t,identifier:i,preroll:a,dtmf_count:r,DTMF_char:s}},ie=function(e){var t=e.readBits(8);return e.readBits(7),{component_tag:t,pts_offset:4*e.readBits(31)+e.readBits(2)}},ne=function(e,t,i,n){var a=n.readBits(32),r=n.readBool();n.readBits(7);var s={descriptor_tag:e,descriptor_length:t,identifier:i,segmentation_event_id:a,segmentation_event_cancel_indicator:r};if(r)return s;if(s.program_segmentation_flag=n.readBool(),s.segmentation_duration_flag=n.readBool(),s.delivery_not_restricted_flag=n.readBool(),s.delivery_not_restricted_flag?n.readBits(5):(s.web_delivery_allowed_flag=n.readBool(),s.no_regional_blackout_flag=n.readBool(),s.archive_allowed_flag=n.readBool(),s.device_restrictions=n.readBits(2)),!s.program_segmentation_flag){s.component_count=n.readBits(8),s.components=[];for(var o=0;o<s.component_count;o++)s.components.push(ie(n))}s.segmentation_duration_flag&&(s.segmentation_duration=n.readBits(40)),s.segmentation_upid_type=n.readBits(8),s.segmentation_upid_length=n.readBits(8);var d=new Uint8Array(s.segmentation_upid_length);for(o=0;o<s.segmentation_upid_length;o++)d[o]=n.readBits(8);return s.segmentation_upid=d.buffer,s.segmentation_type_id=n.readBits(8),s.segment_num=n.readBits(8),s.segments_expected=n.readBits(8),52!==s.segmentation_type_id&&54!==s.segmentation_type_id&&56!==s.segmentation_type_id&&58!==s.segmentation_type_id||(s.sub_segment_num=n.readBits(8),s.sub_segments_expected=n.readBits(8)),s},ae=function(e,t,i,n){return{descriptor_tag:e,descriptor_length:t,identifier:i,TAI_seconds:n.readBits(48),TAI_ns:n.readBits(32),UTC_offset:n.readBits(16)}},re=function(e){return{component_tag:e.readBits(8),ISO_code:String.fromCharCode(e.readBits(8),e.readBits(8),e.readBits(8)),Bit_Stream_Mode:e.readBits(3),Num_Channels:e.readBits(4),Full_Srvc_Audio:e.readBool()}},se=function(e,t,i,n){for(var a=n.readBits(4),r=[],s=0;s<a;s++)r.push(re(n));return{descriptor_tag:e,descriptor_length:t,identifier:i,audio_count:a,components:r}},oe=function(e){var t=new f(e),i=t.readBits(8),n=t.readBool(),a=t.readBool();t.readBits(2);var r=t.readBits(12),s=t.readBits(8),o=t.readBool(),d=t.readBits(6),_=4*t.readBits(31)+t.readBits(2),h=t.readBits(8),c=t.readBits(12),u=t.readBits(12),l=t.readBits(8),p=null;l===N.kSpliceNull?p={}:l===N.kSpliceSchedule?p=function(e){for(var t=e.readBits(8),i=[],n=0;n<t;n++)i.push($(e));return{splice_count:t,events:i}}(t):l===N.kSpliceInsert?p=function(e){var t=e.readBits(32),i=e.readBool();e.readBits(7);var n={splice_event_id:t,splice_event_cancel_indicator:i};if(i)return n;if(n.out_of_network_indicator=e.readBool(),n.program_splice_flag=e.readBool(),n.duration_flag=e.readBool(),n.splice_immediate_flag=e.readBool(),e.readBits(4),n.program_splice_flag&&!n.splice_immediate_flag&&(n.splice_time=Y(e)),!n.program_splice_flag){n.component_count=e.readBits(8),n.components=[];for(var a=0;a<n.component_count;a++)n.components.push(Z(n.splice_immediate_flag,e))}return n.duration_flag&&(n.break_duration=J(e)),n.unique_program_id=e.readBits(16),n.avail_num=e.readBits(8),n.avails_expected=e.readBits(8),n}(t):l===N.kTimeSignal?p=function(e){return{splice_time:Y(e)}}(t):l===N.kBandwidthReservation?p={}:l===N.kPrivateCommand?p=function(e,t){for(var i=String.fromCharCode(t.readBits(8),t.readBits(8),t.readBits(8),t.readBits(8)),n=new Uint8Array(e-4),a=0;a<e-4;a++)n[a]=t.readBits(8);return{identifier:i,private_data:n.buffer}}(u,t):t.readBits(8*u);for(var m=[],g=t.readBits(16),v=0;v<g;){var y=t.readBits(8),b=t.readBits(8),S=String.fromCharCode(t.readBits(8),t.readBits(8),t.readBits(8),t.readBits(8));0===y?m.push(ee(y,b,S,t)):1===y?m.push(te(y,b,S,t)):2===y?m.push(ne(y,b,S,t)):3===y?m.push(ae(y,b,S,t)):4===y?m.push(se(y,b,S,t)):t.readBits(8*(b-4)),v+=2+b}var E={table_id:i,section_syntax_indicator:n,private_indicator:a,section_length:r,protocol_version:s,encrypted_packet:o,encryption_algorithm:d,pts_adjustment:_,cw_index:h,tier:c,splice_command_length:u,splice_command_type:l,splice_command:p,descriptor_loop_length:g,splice_descriptors:m,E_CRC32:o?t.readBits(32):void 0,CRC32:t.readBits(32)};if(l===N.kSpliceInsert){var A=p;if(A.splice_event_cancel_indicator)return{splice_command_type:l,detail:E,data:e};if(A.program_splice_flag&&!A.splice_immediate_flag){var R=A.duration_flag?A.break_duration.auto_return:void 0,T=A.duration_flag?A.break_duration.duration/90:void 0;return A.splice_time.time_specified_flag?{splice_command_type:l,pts:(_+A.splice_time.pts_time)%Math.pow(2,33),auto_return:R,duraiton:T,detail:E,data:e}:{splice_command_type:l,auto_return:R,duraiton:T,detail:E,data:e}}return{splice_command_type:l,auto_return:R=A.duration_flag?A.break_duration.auto_return:void 0,duraiton:T=A.duration_flag?A.break_duration.duration/90:void 0,detail:E,data:e}}if(l===N.kTimeSignal){var L=p;return L.splice_time.time_specified_flag?{splice_command_type:l,pts:(_+L.splice_time.pts_time)%Math.pow(2,33),detail:E,data:e}:{splice_command_type:l,detail:E,data:e}}return{splice_command_type:l,detail:E,data:e}};!function(e){e[e.kSliceIDR_W_RADL=19]="kSliceIDR_W_RADL",e[e.kSliceIDR_N_LP=20]="kSliceIDR_N_LP",e[e.kSliceCRA_NUT=21]="kSliceCRA_NUT",e[e.kSliceVPS=32]="kSliceVPS",e[e.kSliceSPS=33]="kSliceSPS",e[e.kSlicePPS=34]="kSlicePPS",e[e.kSliceAUD=35]="kSliceAUD"}(X||(X={}));var de=function(){},_e=function(e){var t=e.data.byteLength;this.type=e.type,this.data=new Uint8Array(4+t),new DataView(this.data.buffer).setUint32(0,t),this.data.set(e.data,4)},he=function(){function e(e){this.TAG="H265AnnexBParser",this.current_startcode_offset_=0,this.eof_flag_=!1,this.data_=e,this.current_startcode_offset_=this.findNextStartCodeOffset(0),this.eof_flag_&&r.a.e(this.TAG,"Could not find H265 startcode until payload end!")}return e.prototype.findNextStartCodeOffset=function(e){for(var t=e,i=this.data_;;){if(t+3>=i.byteLength)return this.eof_flag_=!0,i.byteLength;var n=i[t+0]<<24|i[t+1]<<16|i[t+2]<<8|i[t+3],a=i[t+0]<<16|i[t+1]<<8|i[t+2];if(1===n||1===a)return t;t++}},e.prototype.readNextNaluPayload=function(){for(var e=this.data_,t=null;null==t&&!this.eof_flag_;){var i=this.current_startcode_offset_,n=e[i+=1===(e[i]<<24|e[i+1]<<16|e[i+2]<<8|e[i+3])?4:3]>>1&63,a=(128&e[i])>>>7,r=this.findNextStartCodeOffset(i);if(this.current_startcode_offset_=r,0===a){var s=e.subarray(i,r);(t=new de).type=n,t.data=s}}return t},e}(),ce=function(){function e(e,t,i,n){var a=23+(5+e.byteLength)+(5+t.byteLength)+(5+i.byteLength),r=this.data=new Uint8Array(a);r[0]=1,r[1]=(3&n.general_profile_space)<<6|(n.general_tier_flag?1:0)<<5|31&n.general_profile_idc,r[2]=n.general_profile_compatibility_flags_1,r[3]=n.general_profile_compatibility_flags_2,r[4]=n.general_profile_compatibility_flags_3,r[5]=n.general_profile_compatibility_flags_4,r[6]=n.general_constraint_indicator_flags_1,r[7]=n.general_constraint_indicator_flags_2,r[8]=n.general_constraint_indicator_flags_3,r[9]=n.general_constraint_indicator_flags_4,r[10]=n.general_constraint_indicator_flags_5,r[11]=n.general_constraint_indicator_flags_6,r[12]=n.general_level_idc,r[13]=240|(3840&n.min_spatial_segmentation_idc)>>8,r[14]=255&n.min_spatial_segmentation_idc,r[15]=252|3&n.parallelismType,r[16]=252|3&n.chroma_format_idc,r[17]=248|7&n.bit_depth_luma_minus8,r[18]=248|7&n.bit_depth_chroma_minus8,r[19]=0,r[20]=0,r[21]=(3&n.constant_frame_rate)<<6|(7&n.num_temporal_layers)<<3|(n.temporal_id_nested?1:0)<<2|3,r[22]=3,r[23]=128|X.kSliceVPS,r[24]=0,r[25]=1,r[26]=(65280&e.byteLength)>>8,r[27]=(255&e.byteLength)>>0,r.set(e,28),r[23+(5+e.byteLength)+0]=128|X.kSliceSPS,r[23+(5+e.byteLength)+1]=0,r[23+(5+e.byteLength)+2]=1,r[23+(5+e.byteLength)+3]=(65280&t.byteLength)>>8,r[23+(5+e.byteLength)+4]=(255&t.byteLength)>>0,r.set(t,23+(5+e.byteLength)+5),r[23+(5+e.byteLength+5+t.byteLength)+0]=128|X.kSlicePPS,r[23+(5+e.byteLength+5+t.byteLength)+1]=0,r[23+(5+e.byteLength+5+t.byteLength)+2]=1,r[23+(5+e.byteLength+5+t.byteLength)+3]=(65280&i.byteLength)>>8,r[23+(5+e.byteLength+5+t.byteLength)+4]=(255&i.byteLength)>>0,r.set(i,23+(5+e.byteLength+5+t.byteLength)+5)}return e.prototype.getData=function(){return this.data},e}(),ue=function(){},le=function(){},fe=function(){},pe=[[64,64,80,80,96,96,112,112,128,128,160,160,192,192,224,224,256,256,320,320,384,384,448,448,512,512,640,640,768,768,896,896,1024,1024,1152,1152,1280,1280],[69,70,87,88,104,105,121,122,139,140,174,175,208,209,243,244,278,279,348,349,417,418,487,488,557,558,696,697,835,836,975,976,1114,1115,1253,1254,1393,1394],[96,96,120,120,144,144,168,168,192,192,240,240,288,288,336,336,384,384,480,480,576,576,672,672,768,768,960,960,1152,1152,1344,1344,1536,1536,1728,1728,1920,1920]],me=function(){function e(e){this.TAG="AC3Parser",this.data_=e,this.current_syncword_offset_=this.findNextSyncwordOffset(0),this.eof_flag_&&r.a.e(this.TAG,"Could not found AC3 syncword until payload end")}return e.prototype.findNextSyncwordOffset=function(e){for(var t=e,i=this.data_;;){if(t+7>=i.byteLength)return this.eof_flag_=!0,i.byteLength;if(2935===(i[t+0]<<8|i[t+1]<<0))return t;t++}},e.prototype.readNextAC3Frame=function(){for(var e=this.data_,t=null;null==t&&!this.eof_flag_;){var i=this.current_syncword_offset_,n=e[i+4]>>6,a=[48e3,44200,33e3][n],r=63&e[i+4],s=2*pe[n][r];if(i+s>this.data_.byteLength){this.eof_flag_=!0,this.has_last_incomplete_data=!0;break}var o=this.findNextSyncwordOffset(i+s);this.current_syncword_offset_=o;var d=e[i+5]>>3,_=7&e[i+5],h=e[i+6]>>5,c=0;0!=(1&h)&&1!==h&&(c+=2),0!=(4&h)&&(c+=2),2===h&&(c+=2);var u=(e[i+6]<<8|e[i+7]<<0)>>12-c&1,l=[2,1,2,3,3,4,4,5][h]+u;(t=new fe).sampling_frequency=a,t.channel_count=l,t.channel_mode=h,t.bit_stream_identification=d,t.low_frequency_effects_channel_on=u,t.bit_stream_mode=_,t.frame_size_code=r,t.data=e.subarray(i,i+s)}return t},e.prototype.hasIncompleteData=function(){return this.has_last_incomplete_data},e.prototype.getIncompleteData=function(){return this.has_last_incomplete_data?this.data_.subarray(this.current_syncword_offset_):null},e}(),ge=function(e){var t;t=[e.sampling_rate_code<<6|e.bit_stream_identification<<1|e.bit_stream_mode>>2,(3&e.bit_stream_mode)<<6|e.channel_mode<<3|e.low_frequency_effects_channel_on<<2|e.frame_size_code>>4,e.frame_size_code<<4&224],this.config=t,this.sampling_rate=e.sampling_frequency,this.bit_stream_identification=e.bit_stream_identification,this.bit_stream_mode=e.bit_stream_mode,this.low_frequency_effects_channel_on=e.low_frequency_effects_channel_on,this.channel_count=e.channel_count,this.channel_mode=e.channel_mode,this.codec_mimetype="ac-3",this.original_codec_mimetype="ac-3"},ve=function(){var e=function(t,i){return(e=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(e,t){e.__proto__=t}||function(e,t){for(var i in t)t.hasOwnProperty(i)&&(e[i]=t[i])})(t,i)};return function(t,i){function n(){this.constructor=t}e(t,i),t.prototype=null===i?Object.create(i):(n.prototype=i.prototype,new n)}}(),ye=function(){return(ye=Object.assign||function(e){for(var t,i=1,n=arguments.length;i<n;i++)for(var a in t=arguments[i])Object.prototype.hasOwnProperty.call(t,a)&&(e[a]=t[a]);return e}).apply(this,arguments)},be=function(e){function t(t,i){var n=e.call(this)||this;return n.TAG="TSDemuxer",n.first_parse_=!0,n.media_info_=new o.a,n.timescale_=90,n.duration_=0,n.current_pmt_pid_=-1,n.program_pmt_map_={},n.pes_slice_queues_={},n.section_slice_queues_={},n.video_metadata_={vps:void 0,sps:void 0,pps:void 0,details:void 0},n.audio_metadata_={codec:void 0,audio_object_type:void 0,sampling_freq_index:void 0,sampling_frequency:void 0,channel_config:void 0},n.aac_last_sample_pts_=void 0,n.aac_last_incomplete_data_=null,n.has_video_=!1,n.has_audio_=!1,n.video_init_segment_dispatched_=!1,n.audio_init_segment_dispatched_=!1,n.video_metadata_changed_=!1,n.audio_metadata_changed_=!1,n.loas_previous_frame=null,n.video_track_={type:"video",id:1,sequenceNumber:0,samples:[],length:0},n.audio_track_={type:"audio",id:2,sequenceNumber:0,samples:[],length:0},n.ts_packet_size_=t.ts_packet_size,n.sync_offset_=t.sync_offset,n.config_=i,n}return ve(t,e),t.prototype.destroy=function(){this.media_info_=null,this.pes_slice_queues_=null,this.section_slice_queues_=null,this.video_metadata_=null,this.audio_metadata_=null,this.aac_last_incomplete_data_=null,this.video_track_=null,this.audio_track_=null,e.prototype.destroy.call(this)},t.probe=function(e){var t=new Uint8Array(e),i=-1,n=188;if(t.byteLength<=3*n)return{needMoreData:!0};for(;-1===i;){for(var a=Math.min(1e3,t.byteLength-3*n),s=0;s<a;){if(71===t[s]&&71===t[s+n]&&71===t[s+2*n]){i=s;break}s++}if(-1===i)if(188===n)n=192;else{if(192!==n)break;n=204}}return-1===i?{match:!1}:(192===n&&i>=4?(r.a.v("TSDemuxer","ts_packet_size = 192, m2ts mode"),i-=4):204===n&&r.a.v("TSDemuxer","ts_packet_size = 204, RS encoded MPEG2-TS stream"),{match:!0,consumed:0,ts_packet_size:n,sync_offset:i})},t.prototype.bindDataSource=function(e){return e.onDataArrival=this.parseChunks.bind(this),this},t.prototype.resetMediaInfo=function(){this.media_info_=new o.a},t.prototype.parseChunks=function(e,t){if(!(this.onError&&this.onMediaInfo&&this.onTrackMetadata&&this.onDataAvailable))throw new c.a("onError & onMediaInfo & onTrackMetadata & onDataAvailable callback must be specified");var i=0;for(this.first_parse_&&(this.first_parse_=!1,i=this.sync_offset_);i+this.ts_packet_size_<=e.byteLength;){var n=t+i;192===this.ts_packet_size_&&(i+=4);var a=new Uint8Array(e,i,188),s=a[0];if(71!==s){r.a.e(this.TAG,"sync_byte = "+s+", not 0x47");break}var o=(64&a[1])>>>6,d=(a[1],(31&a[1])<<8|a[2]),_=(48&a[3])>>>4,h=15&a[3],u={},l=4;if(2==_||3==_){var f=a[4];if(5+f===188){i+=188,204===this.ts_packet_size_&&(i+=16);continue}f>0&&(u=this.parseAdaptationField(e,i+4,1+f)),l=5+f}if(1==_||3==_)if(0===d||d===this.current_pmt_pid_||null!=this.pmt_&&this.pmt_.pid_stream_type[d]===E.kSCTE35){var p=188-l;this.handleSectionSlice(e,i+l,p,{pid:d,file_position:n,payload_unit_start_indicator:o,continuity_conunter:h,random_access_indicator:u.random_access_indicator})}else if(null!=this.pmt_&&null!=this.pmt_.pid_stream_type[d]){p=188-l;var m=this.pmt_.pid_stream_type[d];d!==this.pmt_.common_pids.h264&&d!==this.pmt_.common_pids.h265&&d!==this.pmt_.common_pids.adts_aac&&d!==this.pmt_.common_pids.loas_aac&&d!==this.pmt_.common_pids.ac3&&d!==this.pmt_.common_pids.opus&&d!==this.pmt_.common_pids.mp3&&!0!==this.pmt_.pes_private_data_pids[d]&&!0!==this.pmt_.timed_id3_pids[d]||this.handlePESSlice(e,i+l,p,{pid:d,stream_type:m,file_position:n,payload_unit_start_indicator:o,continuity_conunter:h,random_access_indicator:u.random_access_indicator})}i+=188,204===this.ts_packet_size_&&(i+=16)}return this.dispatchAudioVideoMediaSegment(),i},t.prototype.parseAdaptationField=function(e,t,i){var n=new Uint8Array(e,t,i),a=n[0];return a>0?a>183?(r.a.w(this.TAG,"Illegal adaptation_field_length: "+a),{}):{discontinuity_indicator:(128&n[1])>>>7,random_access_indicator:(64&n[1])>>>6,elementary_stream_priority_indicator:(32&n[1])>>>5}:{}},t.prototype.handleSectionSlice=function(e,t,i,n){var a=new Uint8Array(e,t,i),r=this.section_slice_queues_[n.pid];if(n.payload_unit_start_indicator){var s=a[0];if(null!=r&&0!==r.total_length){var o=new Uint8Array(e,t+1,Math.min(i,s));r.slices.push(o),r.total_length+=o.byteLength,r.total_length===r.expected_length?this.emitSectionSlices(r,n):this.clearSlices(r,n)}for(var d=1+s;d<a.byteLength;){if(255===a[d+0])break;var _=(15&a[d+1])<<8|a[d+2];this.section_slice_queues_[n.pid]=new C,(r=this.section_slice_queues_[n.pid]).expected_length=_+3,r.file_position=n.file_position,r.random_access_indicator=n.random_access_indicator;o=new Uint8Array(e,t+d,Math.min(i-d,r.expected_length-r.total_length));r.slices.push(o),r.total_length+=o.byteLength,r.total_length===r.expected_length?this.emitSectionSlices(r,n):r.total_length>=r.expected_length&&this.clearSlices(r,n),d+=o.byteLength}}else if(null!=r&&0!==r.total_length){o=new Uint8Array(e,t,Math.min(i,r.expected_length-r.total_length));r.slices.push(o),r.total_length+=o.byteLength,r.total_length===r.expected_length?this.emitSectionSlices(r,n):r.total_length>=r.expected_length&&this.clearSlices(r,n)}},t.prototype.handlePESSlice=function(e,t,i,n){var a=new Uint8Array(e,t,i),s=a[0]<<16|a[1]<<8|a[2],o=(a[3],a[4]<<8|a[5]);if(n.payload_unit_start_indicator){if(1!==s)return void r.a.e(this.TAG,"handlePESSlice: packet_start_code_prefix should be 1 but with value "+s);var d=this.pes_slice_queues_[n.pid];d&&(0===d.expected_length||d.expected_length===d.total_length?this.emitPESSlices(d,n):this.clearSlices(d,n)),this.pes_slice_queues_[n.pid]=new C,this.pes_slice_queues_[n.pid].file_position=n.file_position,this.pes_slice_queues_[n.pid].random_access_indicator=n.random_access_indicator}if(null!=this.pes_slice_queues_[n.pid]){var _=this.pes_slice_queues_[n.pid];_.slices.push(a),n.payload_unit_start_indicator&&(_.expected_length=0===o?0:o+6),_.total_length+=a.byteLength,_.expected_length>0&&_.expected_length===_.total_length?this.emitPESSlices(_,n):_.expected_length>0&&_.expected_length<_.total_length&&this.clearSlices(_,n)}},t.prototype.emitSectionSlices=function(e,t){for(var i=new Uint8Array(e.total_length),n=0,a=0;n<e.slices.length;n++){var r=e.slices[n];i.set(r,a),a+=r.byteLength}e.slices=[],e.expected_length=-1,e.total_length=0;var s=new D;s.pid=t.pid,s.data=i,s.file_position=e.file_position,s.random_access_indicator=e.random_access_indicator,this.parseSection(s)},t.prototype.emitPESSlices=function(e,t){for(var i=new Uint8Array(e.total_length),n=0,a=0;n<e.slices.length;n++){var r=e.slices[n];i.set(r,a),a+=r.byteLength}e.slices=[],e.expected_length=-1,e.total_length=0;var s=new k;s.pid=t.pid,s.data=i,s.stream_type=t.stream_type,s.file_position=e.file_position,s.random_access_indicator=e.random_access_indicator,this.parsePES(s)},t.prototype.clearSlices=function(e,t){e.slices=[],e.expected_length=-1,e.total_length=0},t.prototype.parseSection=function(e){var t=e.data,i=e.pid;0===i?this.parsePAT(t):i===this.current_pmt_pid_?this.parsePMT(t):null!=this.pmt_&&this.pmt_.scte_35_pids[i]&&this.parseSCTE35(t)},t.prototype.parsePES=function(e){var t=e.data,i=t[0]<<16|t[1]<<8|t[2],n=t[3],a=t[4]<<8|t[5];if(1===i){if(188!==n&&190!==n&&191!==n&&240!==n&&241!==n&&255!==n&&242!==n&&248!==n){t[6];var s=(192&t[7])>>>6,o=t[8],d=void 0,_=void 0;2!==s&&3!==s||(d=536870912*(14&t[9])+4194304*(255&t[10])+16384*(254&t[11])+128*(255&t[12])+(254&t[13])/2,_=3===s?536870912*(14&t[14])+4194304*(255&t[15])+16384*(254&t[16])+128*(255&t[17])+(254&t[18])/2:d);var h=9+o,c=void 0;if(0!==a){if(a<3+o)return void r.a.v(this.TAG,"Malformed PES: PES_packet_length < 3 + PES_header_data_length");c=a-3-o}else c=t.byteLength-h;var u=t.subarray(h,h+c);switch(e.stream_type){case E.kMPEG1Audio:case E.kMPEG2Audio:this.parseMP3Payload(u,d);break;case E.kPESPrivateData:this.pmt_.common_pids.opus===e.pid?this.parseOpusPayload(u,d):this.pmt_.common_pids.ac3===e.pid?this.parseAC3Payload(u,d):this.pmt_.smpte2038_pids[e.pid]?this.parseSMPTE2038MetadataPayload(u,d,_,e.pid,n):this.parsePESPrivateDataPayload(u,d,_,e.pid,n);break;case E.kADTSAAC:this.parseADTSAACPayload(u,d);break;case E.kLOASAAC:this.parseLOASAACPayload(u,d);break;case E.kAC3:this.parseAC3Payload(u,d);break;case E.kID3:this.parseTimedID3MetadataPayload(u,d,_,e.pid,n);break;case E.kH264:this.parseH264Payload(u,d,_,e.file_position,e.random_access_indicator);break;case E.kH265:this.parseH265Payload(u,d,_,e.file_position,e.random_access_indicator)}}else if((188===n||191===n||240===n||241===n||255===n||242===n||248===n)&&e.stream_type===E.kPESPrivateData){h=6,c=void 0;c=0!==a?a:t.byteLength-h;u=t.subarray(h,h+c);this.parsePESPrivateDataPayload(u,void 0,void 0,e.pid,n)}}else r.a.e(this.TAG,"parsePES: packet_start_code_prefix should be 1 but with value "+i)},t.prototype.parsePAT=function(e){var t=e[0];if(0===t){var i=(15&e[1])<<8|e[2],n=(e[3],e[4],(62&e[5])>>>1),a=1&e[5],s=e[6],o=(e[7],null);if(1===a&&0===s)(o=new T).version_number=n;else if(null==(o=this.pat_))return;for(var d=i-5-4,_=-1,h=-1,c=8;c<8+d;c+=4){var u=e[c]<<8|e[c+1],l=(31&e[c+2])<<8|e[c+3];0===u?o.network_pid=l:(o.program_pmt_pid[u]=l,-1===_&&(_=u),-1===h&&(h=l))}1===a&&0===s&&(null==this.pat_&&r.a.v(this.TAG,"Parsed first PAT: "+JSON.stringify(o)),this.pat_=o,this.current_program_=_,this.current_pmt_pid_=h)}else r.a.e(this.TAG,"parsePAT: table_id "+t+" is not corresponded to PAT!")},t.prototype.parsePMT=function(e){var t=e[0];if(2===t){var i=(15&e[1])<<8|e[2],n=e[3]<<8|e[4],a=(62&e[5])>>>1,s=1&e[5],o=e[6],d=(e[7],null);if(1===s&&0===o)(d=new w).program_number=n,d.version_number=a,this.program_pmt_map_[n]=d;else if(null==(d=this.program_pmt_map_[n]))return;e[8],e[9];for(var _=(15&e[10])<<8|e[11],h=12+_,c=i-9-_-4,u=h;u<h+c;){var l=e[u],f=(31&e[u+1])<<8|e[u+2],p=(15&e[u+3])<<8|e[u+4];d.pid_stream_type[f]=l;var m=d.common_pids.h264||d.common_pids.h265,g=d.common_pids.adts_aac||d.common_pids.loas_aac||d.common_pids.ac3||d.common_pids.opus||d.common_pids.mp3;if(l!==E.kH264||m)if(l!==E.kH265||m)if(l!==E.kADTSAAC||g)if(l!==E.kLOASAAC||g)if(l!==E.kAC3||g)if(l!==E.kMPEG1Audio&&l!==E.kMPEG2Audio||g)if(l===E.kPESPrivateData){if(d.pes_private_data_pids[f]=!0,p>0){for(var v=u+5;v<u+5+p;){var y=e[v+0],b=e[v+1];if(5===y){var S=String.fromCharCode.apply(String,Array.from(e.subarray(v+2,v+2+b)));"VANC"===S?d.smpte2038_pids[f]=!0:"Opus"===S&&(d.common_pids.opus=f)}else if(127===y&&f===d.common_pids.opus){var A=null;if(128===e[v+2]&&(A=e[v+3]),null==A){r.a.e(this.TAG,"Not Supported Opus channel count.");continue}var R={codec:"opus",channel_count:0==(15&A)?2:15&A,channel_config_code:A,sample_rate:48e3},T={codec:"opus",meta:R};0==this.audio_init_segment_dispatched_?(this.audio_metadata_=R,this.dispatchAudioInitSegment(T)):this.detectAudioMetadataChange(T)&&(this.dispatchAudioMediaSegment(),this.dispatchAudioInitSegment(T))}v+=2+b}var L=e.subarray(u+5,u+5+p);this.dispatchPESPrivateDataDescriptor(f,l,L)}}else l===E.kID3?d.timed_id3_pids[f]=!0:l===E.kSCTE35&&(d.scte_35_pids[f]=!0);else d.common_pids.mp3=f;else d.common_pids.ac3=f;else d.common_pids.loas_aac=f;else d.common_pids.adts_aac=f;else d.common_pids.h265=f;else d.common_pids.h264=f;u+=5+p}n===this.current_program_&&(null==this.pmt_&&r.a.v(this.TAG,"Parsed first PMT: "+JSON.stringify(d)),this.pmt_=d,(d.common_pids.h264||d.common_pids.h265)&&(this.has_video_=!0),(d.common_pids.adts_aac||d.common_pids.loas_aac||d.common_pids.ac3||d.common_pids.opus||d.common_pids.mp3)&&(this.has_audio_=!0))}else r.a.e(this.TAG,"parsePMT: table_id "+t+" is not corresponded to PMT!")},t.prototype.parseSCTE35=function(e){var t=oe(e);if(null!=t.pts){var i=Math.floor(t.pts/this.timescale_);t.pts=i}else t.nearest_pts=this.aac_last_sample_pts_;this.onSCTE35Metadata&&this.onSCTE35Metadata(t)},t.prototype.parseH264Payload=function(e,t,i,n,a){for(var s=new M(e),o=null,d=[],_=0,h=!1;null!=(o=s.readNextNaluPayload());){var c=new P(o);if(c.type===L.kSliceSPS){var u=p.parseSPS(o.data);this.video_init_segment_dispatched_?!0===this.detectVideoMetadataChange(c,u)&&(r.a.v(this.TAG,"H264: Critical h264 metadata has been changed, attempt to re-generate InitSegment"),this.video_metadata_changed_=!0,this.video_metadata_={vps:void 0,sps:c,pps:void 0,details:u}):(this.video_metadata_.sps=c,this.video_metadata_.details=u)}else c.type===L.kSlicePPS?this.video_init_segment_dispatched_&&!this.video_metadata_changed_||(this.video_metadata_.pps=c,this.video_metadata_.sps&&this.video_metadata_.pps&&(this.video_metadata_changed_&&this.dispatchVideoMediaSegment(),this.dispatchVideoInitSegment())):(c.type===L.kSliceIDR||c.type===L.kSliceNonIDR&&1===a)&&(h=!0);this.video_init_segment_dispatched_&&(d.push(c),_+=c.data.byteLength)}var l=Math.floor(t/this.timescale_),f=Math.floor(i/this.timescale_);if(d.length){var m=this.video_track_,g={units:d,length:_,isKeyframe:h,dts:f,pts:l,cts:l-f,file_position:n};m.samples.push(g),m.length+=_}},t.prototype.parseH265Payload=function(e,t,i,n,a){for(var s=new he(e),o=null,d=[],_=0,h=!1;null!=(o=s.readNextNaluPayload());){var c=new _e(o);if(c.type===X.kSliceVPS){if(!this.video_init_segment_dispatched_){var u=g.parseVPS(o.data);this.video_metadata_.vps=c,this.video_metadata_.details=ye(ye({},this.video_metadata_.details),u)}}else if(c.type===X.kSliceSPS){u=g.parseSPS(o.data);this.video_init_segment_dispatched_?!0===this.detectVideoMetadataChange(c,u)&&(r.a.v(this.TAG,"H265: Critical h265 metadata has been changed, attempt to re-generate InitSegment"),this.video_metadata_changed_=!0,this.video_metadata_={vps:void 0,sps:c,pps:void 0,details:u}):(this.video_metadata_.sps=c,this.video_metadata_.details=ye(ye({},this.video_metadata_.details),u))}else if(c.type===X.kSlicePPS){if(!this.video_init_segment_dispatched_||this.video_metadata_changed_){u=g.parsePPS(o.data);this.video_metadata_.pps=c,this.video_metadata_.details=ye(ye({},this.video_metadata_.details),u),this.video_metadata_.vps&&this.video_metadata_.sps&&this.video_metadata_.pps&&(this.video_metadata_changed_&&this.dispatchVideoMediaSegment(),this.dispatchVideoInitSegment())}}else c.type!==X.kSliceIDR_W_RADL&&c.type!==X.kSliceIDR_N_LP&&c.type!==X.kSliceCRA_NUT||(h=!0);this.video_init_segment_dispatched_&&(d.push(c),_+=c.data.byteLength)}var l=Math.floor(t/this.timescale_),f=Math.floor(i/this.timescale_);if(d.length){var p=this.video_track_,m={units:d,length:_,isKeyframe:h,dts:f,pts:l,cts:l-f,file_position:n};p.samples.push(m),p.length+=_}},t.prototype.detectVideoMetadataChange=function(e,t){if(t.codec_mimetype!==this.video_metadata_.details.codec_mimetype)return r.a.v(this.TAG,"Video: Codec mimeType changed from "+this.video_metadata_.details.codec_mimetype+" to "+t.codec_mimetype),!0;if(t.codec_size.width!==this.video_metadata_.details.codec_size.width||t.codec_size.height!==this.video_metadata_.details.codec_size.height){var i=this.video_metadata_.details.codec_size,n=t.codec_size;return r.a.v(this.TAG,"Video: Coded Resolution changed from "+i.width+"x"+i.height+" to "+n.width+"x"+n.height),!0}return t.present_size.width!==this.video_metadata_.details.present_size.width&&(r.a.v(this.TAG,"Video: Present resolution width changed from "+this.video_metadata_.details.present_size.width+" to "+t.present_size.width),!0)},t.prototype.isInitSegmentDispatched=function(){return this.has_video_&&this.has_audio_?this.video_init_segment_dispatched_&&this.audio_init_segment_dispatched_:this.has_video_&&!this.has_audio_?this.video_init_segment_dispatched_:!(this.has_video_||!this.has_audio_)&&this.audio_init_segment_dispatched_},t.prototype.dispatchVideoInitSegment=function(){var e=this.video_metadata_.details,t={type:"video"};t.id=this.video_track_.id,t.timescale=1e3,t.duration=this.duration_,t.codecWidth=e.codec_size.width,t.codecHeight=e.codec_size.height,t.presentWidth=e.present_size.width,t.presentHeight=e.present_size.height,t.profile=e.profile_string,t.level=e.level_string,t.bitDepth=e.bit_depth,t.chromaFormat=e.chroma_format,t.sarRatio=e.sar_ratio,t.frameRate=e.frame_rate;var i=t.frameRate.fps_den,n=t.frameRate.fps_num;if(t.refSampleDuration=i/n*1e3,t.codec=e.codec_mimetype,this.video_metadata_.vps){var a=this.video_metadata_.vps.data.subarray(4),s=this.video_metadata_.sps.data.subarray(4),o=this.video_metadata_.pps.data.subarray(4),d=new ce(a,s,o,e);t.hvcc=d.getData(),0==this.video_init_segment_dispatched_&&r.a.v(this.TAG,"Generated first HEVCDecoderConfigurationRecord for mimeType: "+t.codec)}else{s=this.video_metadata_.sps.data.subarray(4),o=this.video_metadata_.pps.data.subarray(4);var _=new x(s,o,e);t.avcc=_.getData(),0==this.video_init_segment_dispatched_&&r.a.v(this.TAG,"Generated first AVCDecoderConfigurationRecord for mimeType: "+t.codec)}this.onTrackMetadata("video",t),this.video_init_segment_dispatched_=!0,this.video_metadata_changed_=!1;var h=this.media_info_;h.hasVideo=!0,h.width=t.codecWidth,h.height=t.codecHeight,h.fps=t.frameRate.fps,h.profile=t.profile,h.level=t.level,h.refFrames=e.ref_frames,h.chromaFormat=e.chroma_format_string,h.sarNum=t.sarRatio.width,h.sarDen=t.sarRatio.height,h.videoCodec=t.codec,h.hasAudio&&h.audioCodec?h.mimeType='video/mp2t; codecs="'+h.videoCodec+","+h.audioCodec+'"':h.mimeType='video/mp2t; codecs="'+h.videoCodec+'"',h.isComplete()&&this.onMediaInfo(h)},t.prototype.dispatchVideoMediaSegment=function(){this.isInitSegmentDispatched()&&this.video_track_.length&&this.onDataAvailable(null,this.video_track_)},t.prototype.dispatchAudioMediaSegment=function(){this.isInitSegmentDispatched()&&this.audio_track_.length&&this.onDataAvailable(this.audio_track_,null)},t.prototype.dispatchAudioVideoMediaSegment=function(){this.isInitSegmentDispatched()&&(this.audio_track_.length||this.video_track_.length)&&this.onDataAvailable(this.audio_track_,this.video_track_)},t.prototype.parseADTSAACPayload=function(e,t){if(!this.has_video_||this.video_init_segment_dispatched_){if(this.aac_last_incomplete_data_){var i=new Uint8Array(e.byteLength+this.aac_last_incomplete_data_.byteLength);i.set(this.aac_last_incomplete_data_,0),i.set(e,this.aac_last_incomplete_data_.byteLength),e=i}var n,a;if(null!=t&&(a=t/this.timescale_),"aac"===this.audio_metadata_.codec){if(null==t&&null!=this.aac_last_sample_pts_)n=1024/this.audio_metadata_.sampling_frequency*1e3,a=this.aac_last_sample_pts_+n;else if(null==t)return void r.a.w(this.TAG,"AAC: Unknown pts");if(this.aac_last_incomplete_data_&&this.aac_last_sample_pts_){n=1024/this.audio_metadata_.sampling_frequency*1e3;var s=this.aac_last_sample_pts_+n;Math.abs(s-a)>1&&(r.a.w(this.TAG,"AAC: Detected pts overlapped, expected: "+s+"ms, PES pts: "+a+"ms"),a=s)}}for(var o,d=new z(e),_=null,h=a;null!=(_=d.readNextAACFrame());){n=1024/_.sampling_frequency*1e3;var c={codec:"aac",data:_};0==this.audio_init_segment_dispatched_?(this.audio_metadata_={codec:"aac",audio_object_type:_.audio_object_type,sampling_freq_index:_.sampling_freq_index,sampling_frequency:_.sampling_frequency,channel_config:_.channel_config},this.dispatchAudioInitSegment(c)):this.detectAudioMetadataChange(c)&&(this.dispatchAudioMediaSegment(),this.dispatchAudioInitSegment(c)),o=h;var u=Math.floor(h),l={unit:_.data,length:_.data.byteLength,pts:u,dts:u};this.audio_track_.samples.push(l),this.audio_track_.length+=_.data.byteLength,h+=n}d.hasIncompleteData()&&(this.aac_last_incomplete_data_=d.getIncompleteData()),o&&(this.aac_last_sample_pts_=o)}},t.prototype.parseLOASAACPayload=function(e,t){var i;if(!this.has_video_||this.video_init_segment_dispatched_){if(this.aac_last_incomplete_data_){var n=new Uint8Array(e.byteLength+this.aac_last_incomplete_data_.byteLength);n.set(this.aac_last_incomplete_data_,0),n.set(e,this.aac_last_incomplete_data_.byteLength),e=n}var a,s;if(null!=t&&(s=t/this.timescale_),"aac"===this.audio_metadata_.codec){if(null==t&&null!=this.aac_last_sample_pts_)a=1024/this.audio_metadata_.sampling_frequency*1e3,s=this.aac_last_sample_pts_+a;else if(null==t)return void r.a.w(this.TAG,"AAC: Unknown pts");if(this.aac_last_incomplete_data_&&this.aac_last_sample_pts_){a=1024/this.audio_metadata_.sampling_frequency*1e3;var o=this.aac_last_sample_pts_+a;Math.abs(o-s)>1&&(r.a.w(this.TAG,"AAC: Detected pts overlapped, expected: "+o+"ms, PES pts: "+s+"ms"),s=o)}}for(var d,_=new H(e),h=null,c=s;null!=(h=_.readNextAACFrame(null!==(i=this.loas_previous_frame)&&void 0!==i?i:void 0));){this.loas_previous_frame=h,a=1024/h.sampling_frequency*1e3;var u={codec:"aac",data:h};0==this.audio_init_segment_dispatched_?(this.audio_metadata_={codec:"aac",audio_object_type:h.audio_object_type,sampling_freq_index:h.sampling_freq_index,sampling_frequency:h.sampling_frequency,channel_config:h.channel_config},this.dispatchAudioInitSegment(u)):this.detectAudioMetadataChange(u)&&(this.dispatchAudioMediaSegment(),this.dispatchAudioInitSegment(u)),d=c;var l=Math.floor(c),f={unit:h.data,length:h.data.byteLength,pts:l,dts:l};this.audio_track_.samples.push(f),this.audio_track_.length+=h.data.byteLength,c+=a}_.hasIncompleteData()&&(this.aac_last_incomplete_data_=_.getIncompleteData()),d&&(this.aac_last_sample_pts_=d)}},t.prototype.parseAC3Payload=function(e,t){if(!this.has_video_||this.video_init_segment_dispatched_){var i,n;if(null!=t&&(n=t/this.timescale_),"ac-3"===this.audio_metadata_.codec)if(null==t&&null!=this.aac_last_sample_pts_)i=1536/this.audio_metadata_.sampling_frequency*1e3,n=this.aac_last_sample_pts_+i;else if(null==t)return void r.a.w(this.TAG,"Opus: Unknown pts");for(var a,s=new me(e),o=null,d=n;null!=(o=s.readNextAC3Frame());){i=1536/o.sampling_frequency*1e3;var _={codec:"ac-3",data:o};0==this.audio_init_segment_dispatched_?(this.audio_metadata_={codec:"ac-3",sampling_frequency:o.sampling_frequency,bit_stream_identification:o.bit_stream_identification,bit_stream_mode:o.bit_stream_mode,low_frequency_effects_channel_on:o.low_frequency_effects_channel_on,channel_mode:o.channel_mode},console.log(JSON.stringify(this.audio_metadata_)),this.dispatchAudioInitSegment(_)):this.detectAudioMetadataChange(_)&&(this.dispatchAudioMediaSegment(),this.dispatchAudioInitSegment(_)),a=d;var h=Math.floor(d),c={unit:o.data,length:o.data.byteLength,pts:h,dts:h};this.audio_track_.samples.push(c),this.audio_track_.length+=o.data.byteLength,d+=i}a&&(this.aac_last_sample_pts_=a)}},t.prototype.parseOpusPayload=function(e,t){if(!this.has_video_||this.video_init_segment_dispatched_){var i,n;if(null!=t&&(n=t/this.timescale_),"opus"===this.audio_metadata_.codec)if(null==t&&null!=this.aac_last_sample_pts_)i=20,n=this.aac_last_sample_pts_+i;else if(null==t)return void r.a.w(this.TAG,"Opus: Unknown pts");for(var a,s=n,o=0;o<e.length;){i=20;for(var d=0!=(16&e[o+1]),_=0!=(8&e[o+1]),h=o+2,c=0;255===e[h];)c+=255,h+=1;c+=e[h],h+=1,h+=d?2:0,h+=_?2:0,a=s;var u=Math.floor(s),l=e.slice(h,h+c),f={unit:l,length:l.byteLength,pts:u,dts:u};this.audio_track_.samples.push(f),this.audio_track_.length+=l.byteLength,s+=i,o=h+c}a&&(this.aac_last_sample_pts_=a)}},t.prototype.parseMP3Payload=function(e,t){if(!this.has_video_||this.video_init_segment_dispatched_){var i=[0,32,64,96,128,160,192,224,256,288,320,352,384,416,448,-1],n=[0,32,48,56,64,80,96,112,128,160,192,224,256,320,384,-1],a=[0,32,40,48,56,64,80,96,112,128,160,192,224,256,320,-1],r=e[1]>>>3&3,s=(6&e[1])>>1,o=(240&e[2])>>>4,d=(12&e[2])>>>2,_=3!==(e[3]>>>6&3)?2:1,h=0,c=34;switch(r){case 0:h=[11025,12e3,8e3,0][d];break;case 2:h=[22050,24e3,16e3,0][d];break;case 3:h=[44100,48e3,32e3,0][d]}switch(s){case 1:c=34,o<a.length&&a[o];break;case 2:c=33,o<n.length&&n[o];break;case 3:c=32,o<i.length&&i[o]}var u=new le;u.object_type=c,u.sample_rate=h,u.channel_count=_,u.data=e;var l={codec:"mp3",data:u};0==this.audio_init_segment_dispatched_?(this.audio_metadata_={codec:"mp3",object_type:c,sample_rate:h,channel_count:_},this.dispatchAudioInitSegment(l)):this.detectAudioMetadataChange(l)&&(this.dispatchAudioMediaSegment(),this.dispatchAudioInitSegment(l));var f={unit:e,length:e.byteLength,pts:t/this.timescale_,dts:t/this.timescale_};this.audio_track_.samples.push(f),this.audio_track_.length+=e.byteLength}},t.prototype.detectAudioMetadataChange=function(e){if(e.codec!==this.audio_metadata_.codec)return r.a.v(this.TAG,"Audio: Audio Codecs changed from "+this.audio_metadata_.codec+" to "+e.codec),!0;if("aac"===e.codec&&"aac"===this.audio_metadata_.codec){if((t=e.data).audio_object_type!==this.audio_metadata_.audio_object_type)return r.a.v(this.TAG,"AAC: AudioObjectType changed from "+this.audio_metadata_.audio_object_type+" to "+t.audio_object_type),!0;if(t.sampling_freq_index!==this.audio_metadata_.sampling_freq_index)return r.a.v(this.TAG,"AAC: SamplingFrequencyIndex changed from "+this.audio_metadata_.sampling_freq_index+" to "+t.sampling_freq_index),!0;if(t.channel_config!==this.audio_metadata_.channel_config)return r.a.v(this.TAG,"AAC: Channel configuration changed from "+this.audio_metadata_.channel_config+" to "+t.channel_config),!0}else if("ac-3"===e.codec&&"ac-3"===this.audio_metadata_.codec){var t;if((t=e.data).sampling_frequency!==this.audio_metadata_.sampling_frequency)return r.a.v(this.TAG,"AC3: Sampling Frequency changed from "+this.audio_metadata_.sampling_frequency+" to "+t.sampling_frequency),!0;if(t.bit_stream_identification!==this.audio_metadata_.bit_stream_identification)return r.a.v(this.TAG,"AC3: Bit Stream Identification changed from "+this.audio_metadata_.bit_stream_identification+" to "+t.bit_stream_identification),!0;if(t.bit_stream_mode!==this.audio_metadata_.bit_stream_mode)return r.a.v(this.TAG,"AC3: BitStream Mode changed from "+this.audio_metadata_.bit_stream_mode+" to "+t.bit_stream_mode),!0;if(t.channel_mode!==this.audio_metadata_.channel_mode)return r.a.v(this.TAG,"AC3: Channel Mode changed from "+this.audio_metadata_.channel_mode+" to "+t.channel_mode),!0;if(t.low_frequency_effects_channel_on!==this.audio_metadata_.low_frequency_effects_channel_on)return r.a.v(this.TAG,"AC3: Low Frequency Effects Channel On changed from "+this.audio_metadata_.low_frequency_effects_channel_on+" to "+t.low_frequency_effects_channel_on),!0}else if("opus"===e.codec&&"opus"===this.audio_metadata_.codec){if((i=e.meta).sample_rate!==this.audio_metadata_.sample_rate)return r.a.v(this.TAG,"Opus: SamplingFrequencyIndex changed from "+this.audio_metadata_.sample_rate+" to "+i.sample_rate),!0;if(i.channel_count!==this.audio_metadata_.channel_count)return r.a.v(this.TAG,"Opus: Channel count changed from "+this.audio_metadata_.channel_count+" to "+i.channel_count),!0}else if("mp3"===e.codec&&"mp3"===this.audio_metadata_.codec){var i;if((i=e.data).object_type!==this.audio_metadata_.object_type)return r.a.v(this.TAG,"MP3: AudioObjectType changed from "+this.audio_metadata_.object_type+" to "+i.object_type),!0;if(i.sample_rate!==this.audio_metadata_.sample_rate)return r.a.v(this.TAG,"MP3: SamplingFrequencyIndex changed from "+this.audio_metadata_.sample_rate+" to "+i.sample_rate),!0;if(i.channel_count!==this.audio_metadata_.channel_count)return r.a.v(this.TAG,"MP3: Channel count changed from "+this.audio_metadata_.channel_count+" to "+i.channel_count),!0}return!1},t.prototype.dispatchAudioInitSegment=function(e){var t={type:"audio"};if(t.id=this.audio_track_.id,t.timescale=1e3,t.duration=this.duration_,"aac"===this.audio_metadata_.codec){var i="aac"===e.codec?e.data:null,n=new q(i);t.audioSampleRate=n.sampling_rate,t.channelCount=n.channel_count,t.codec=n.codec_mimetype,t.originalCodec=n.original_codec_mimetype,t.config=n.config,t.refSampleDuration=1024/t.audioSampleRate*t.timescale}else if("ac-3"===this.audio_metadata_.codec){var a="ac-3"===e.codec?e.data:null,s=new ge(a);t.audioSampleRate=s.sampling_rate,t.channelCount=s.channel_count,t.codec=s.codec_mimetype,t.originalCodec=s.original_codec_mimetype,t.config=s.config,t.refSampleDuration=1536/t.audioSampleRate*t.timescale}else"opus"===this.audio_metadata_.codec?(t.audioSampleRate=this.audio_metadata_.sample_rate,t.channelCount=this.audio_metadata_.channel_count,t.channelConfigCode=this.audio_metadata_.channel_config_code,t.codec="opus",t.originalCodec="opus",t.config=void 0,t.refSampleDuration=20):"mp3"===this.audio_metadata_.codec&&(t.audioSampleRate=this.audio_metadata_.sample_rate,t.channelCount=this.audio_metadata_.channel_count,t.codec="mp3",t.originalCodec="mp3",t.config=void 0);0==this.audio_init_segment_dispatched_&&r.a.v(this.TAG,"Generated first AudioSpecificConfig for mimeType: "+t.codec),this.onTrackMetadata("audio",t),this.audio_init_segment_dispatched_=!0,this.video_metadata_changed_=!1;var o=this.media_info_;o.hasAudio=!0,o.audioCodec=t.originalCodec,o.audioSampleRate=t.audioSampleRate,o.audioChannelCount=t.channelCount,o.hasVideo&&o.videoCodec?o.mimeType='video/mp2t; codecs="'+o.videoCodec+","+o.audioCodec+'"':o.mimeType='video/mp2t; codecs="'+o.audioCodec+'"',o.isComplete()&&this.onMediaInfo(o)},t.prototype.dispatchPESPrivateDataDescriptor=function(e,t,i){var n=new W;n.pid=e,n.stream_type=t,n.descriptor=i,this.onPESPrivateDataDescriptor&&this.onPESPrivateDataDescriptor(n)},t.prototype.parsePESPrivateDataPayload=function(e,t,i,n,a){var r=new K;if(r.pid=n,r.stream_id=a,r.len=e.byteLength,r.data=e,null!=t){var s=Math.floor(t/this.timescale_);r.pts=s}else r.nearest_pts=this.aac_last_sample_pts_;if(null!=i){var o=Math.floor(i/this.timescale_);r.dts=o}this.onPESPrivateData&&this.onPESPrivateData(r)},t.prototype.parseTimedID3MetadataPayload=function(e,t,i,n,a){var r=new K;if(r.pid=n,r.stream_id=a,r.len=e.byteLength,r.data=e,null!=t){var s=Math.floor(t/this.timescale_);r.pts=s}if(null!=i){var o=Math.floor(i/this.timescale_);r.dts=o}this.onTimedID3Metadata&&this.onTimedID3Metadata(r)},t.prototype.parseSMPTE2038MetadataPayload=function(e,t,i,n,a){var r=new ue;if(r.pid=n,r.stream_id=a,r.len=e.byteLength,r.data=e,null!=t){var s=Math.floor(t/this.timescale_);r.pts=s}if(r.nearest_pts=this.aac_last_sample_pts_,null!=i){var o=Math.floor(i/this.timescale_);r.dts=o}r.ancillaries=function(e){for(var t=new f(e),i=0,n=[];;){if(i+=6,0!==t.readBits(6))break;var a=t.readBool();i+=1;var r=t.readBits(11);i+=11;var s=t.readBits(12);i+=12;var o=255&t.readBits(10);i+=10;var d=255&t.readBits(10);i+=10;var _=255&t.readBits(10);i+=10;for(var h=new Uint8Array(_),c=0;c<_;c++){var u=255&t.readBits(10);i+=10,h[c]=u}t.readBits(10);i+=10;var l="User Defined";65===o?7===d&&(l="SCTE-104"):95===o?220===d?l="ARIB STD-B37 (1SEG)":221===d?l="ARIB STD-B37 (ANALOG)":222===d?l="ARIB STD-B37 (SD)":223===d&&(l="ARIB STD-B37 (HD)"):97===o&&(1===d?l="EIA-708":2===d&&(l="EIA-608")),n.push({yc_indicator:a,line_number:r,horizontal_offset:s,did:o,sdid:d,user_data:h,description:l,information:{}}),t.readBits(8-(i-Math.floor(i/8))%8),i+=(8-(i-Math.floor(i/8)))%8}return t.destroy(),t=null,n}(e),this.onSMPTE2038Metadata&&this.onSMPTE2038Metadata(r)},t}(R),Se=function(){for(var e=0,t=0,i=arguments.length;t<i;t++)e+=arguments[t].length;var n=Array(e),a=0;for(t=0;t<i;t++)for(var r=arguments[t],s=0,o=r.length;s<o;s++,a++)n[a]=r[s];return n},Ee=function(){function e(){}return e.init=function(){for(var t in e.types={avc1:[],avcC:[],btrt:[],dinf:[],dref:[],esds:[],ftyp:[],hdlr:[],hvc1:[],hvcC:[],mdat:[],mdhd:[],mdia:[],mfhd:[],minf:[],moof:[],moov:[],mp4a:[],mvex:[],mvhd:[],sdtp:[],stbl:[],stco:[],stsc:[],stsd:[],stsz:[],stts:[],tfdt:[],tfhd:[],traf:[],trak:[],trun:[],trex:[],tkhd:[],vmhd:[],smhd:[],".mp3":[],Opus:[],dOps:[],"ac-3":[],dac3:[]},e.types)e.types.hasOwnProperty(t)&&(e.types[t]=[t.charCodeAt(0),t.charCodeAt(1),t.charCodeAt(2),t.charCodeAt(3)]);var i=e.constants={};i.FTYP=new Uint8Array([105,115,111,109,0,0,0,1,105,115,111,109,97,118,99,49]),i.STSD_PREFIX=new Uint8Array([0,0,0,0,0,0,0,1]),i.STTS=new Uint8Array([0,0,0,0,0,0,0,0]),i.STSC=i.STCO=i.STTS,i.STSZ=new Uint8Array([0,0,0,0,0,0,0,0,0,0,0,0]),i.HDLR_VIDEO=new Uint8Array([0,0,0,0,0,0,0,0,118,105,100,101,0,0,0,0,0,0,0,0,0,0,0,0,86,105,100,101,111,72,97,110,100,108,101,114,0]),i.HDLR_AUDIO=new Uint8Array([0,0,0,0,0,0,0,0,115,111,117,110,0,0,0,0,0,0,0,0,0,0,0,0,83,111,117,110,100,72,97,110,100,108,101,114,0]),i.DREF=new Uint8Array([0,0,0,0,0,0,0,1,0,0,0,12,117,114,108,32,0,0,0,1]),i.SMHD=new Uint8Array([0,0,0,0,0,0,0,0]),i.VMHD=new Uint8Array([0,0,0,1,0,0,0,0,0,0,0,0])},e.box=function(e){for(var t=8,i=null,n=Array.prototype.slice.call(arguments,1),a=n.length,r=0;r<a;r++)t+=n[r].byteLength;(i=new Uint8Array(t))[0]=t>>>24&255,i[1]=t>>>16&255,i[2]=t>>>8&255,i[3]=255&t,i.set(e,4);var s=8;for(r=0;r<a;r++)i.set(n[r],s),s+=n[r].byteLength;return i},e.generateInitSegment=function(t){var i=e.box(e.types.ftyp,e.constants.FTYP),n=e.moov(t),a=new Uint8Array(i.byteLength+n.byteLength);return a.set(i,0),a.set(n,i.byteLength),a},e.moov=function(t){var i=e.mvhd(t.timescale,t.duration),n=e.trak(t),a=e.mvex(t);return e.box(e.types.moov,i,n,a)},e.mvhd=function(t,i){return e.box(e.types.mvhd,new Uint8Array([0,0,0,0,0,0,0,0,0,0,0,0,t>>>24&255,t>>>16&255,t>>>8&255,255&t,i>>>24&255,i>>>16&255,i>>>8&255,255&i,0,1,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,64,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,255,255,255,255]))},e.trak=function(t){return e.box(e.types.trak,e.tkhd(t),e.mdia(t))},e.tkhd=function(t){var i=t.id,n=t.duration,a=t.presentWidth,r=t.presentHeight;return e.box(e.types.tkhd,new Uint8Array([0,0,0,7,0,0,0,0,0,0,0,0,i>>>24&255,i>>>16&255,i>>>8&255,255&i,0,0,0,0,n>>>24&255,n>>>16&255,n>>>8&255,255&n,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,64,0,0,0,a>>>8&255,255&a,0,0,r>>>8&255,255&r,0,0]))},e.mdia=function(t){return e.box(e.types.mdia,e.mdhd(t),e.hdlr(t),e.minf(t))},e.mdhd=function(t){var i=t.timescale,n=t.duration;return e.box(e.types.mdhd,new Uint8Array([0,0,0,0,0,0,0,0,0,0,0,0,i>>>24&255,i>>>16&255,i>>>8&255,255&i,n>>>24&255,n>>>16&255,n>>>8&255,255&n,85,196,0,0]))},e.hdlr=function(t){var i=null;return i="audio"===t.type?e.constants.HDLR_AUDIO:e.constants.HDLR_VIDEO,e.box(e.types.hdlr,i)},e.minf=function(t){var i=null;return i="audio"===t.type?e.box(e.types.smhd,e.constants.SMHD):e.box(e.types.vmhd,e.constants.VMHD),e.box(e.types.minf,i,e.dinf(),e.stbl(t))},e.dinf=function(){return e.box(e.types.dinf,e.box(e.types.dref,e.constants.DREF))},e.stbl=function(t){return e.box(e.types.stbl,e.stsd(t),e.box(e.types.stts,e.constants.STTS),e.box(e.types.stsc,e.constants.STSC),e.box(e.types.stsz,e.constants.STSZ),e.box(e.types.stco,e.constants.STCO))},e.stsd=function(t){return"audio"===t.type?"mp3"===t.codec?e.box(e.types.stsd,e.constants.STSD_PREFIX,e.mp3(t)):"ac-3"===t.codec?e.box(e.types.stsd,e.constants.STSD_PREFIX,e.ac3(t)):"opus"===t.codec?e.box(e.types.stsd,e.constants.STSD_PREFIX,e.Opus(t)):e.box(e.types.stsd,e.constants.STSD_PREFIX,e.mp4a(t)):"video"===t.type&&t.codec.startsWith("hvc1")?e.box(e.types.stsd,e.constants.STSD_PREFIX,e.hvc1(t)):e.box(e.types.stsd,e.constants.STSD_PREFIX,e.avc1(t))},e.mp3=function(t){var i=t.channelCount,n=t.audioSampleRate,a=new Uint8Array([0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,i,0,16,0,0,0,0,n>>>8&255,255&n,0,0]);return e.box(e.types[".mp3"],a)},e.mp4a=function(t){var i=t.channelCount,n=t.audioSampleRate,a=new Uint8Array([0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,i,0,16,0,0,0,0,n>>>8&255,255&n,0,0]);return e.box(e.types.mp4a,a,e.esds(t))},e.ac3=function(t){var i=t.channelCount,n=t.audioSampleRate,a=new Uint8Array([0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,i,0,16,0,0,0,0,n>>>8&255,255&n,0,0]);return e.box(e.types["ac-3"],a,e.box(e.types.dac3,new Uint8Array(t.config)))},e.esds=function(t){var i=t.config||[],n=i.length,a=new Uint8Array([0,0,0,0,3,23+n,0,1,0,4,15+n,64,21,0,0,0,0,0,0,0,0,0,0,0,5].concat([n]).concat(i).concat([6,1,2]));return e.box(e.types.esds,a)},e.Opus=function(t){var i=t.channelCount,n=t.audioSampleRate,a=new Uint8Array([0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,i,0,16,0,0,0,0,n>>>8&255,255&n,0,0]);return e.box(e.types.Opus,a,e.dOps(t))},e.dOps=function(t){var i=t.channelCount,n=t.channelConfigCode,a=t.audioSampleRate;if(t.config)return e.box(e.types.dOps,s);var r=[];switch(n){case 1:case 2:r=[0];break;case 0:r=[255,1,1,0,1];break;case 128:r=[255,2,0,0,1];break;case 3:r=[1,2,1,0,2,1];break;case 4:r=[1,2,2,0,1,2,3];break;case 5:r=[1,3,2,0,4,1,2,3];break;case 6:r=[1,4,2,0,4,1,2,3,5];break;case 7:r=[1,4,2,0,4,1,2,3,5,6];break;case 8:r=[1,5,3,0,6,1,2,3,4,5,7];break;case 130:r=[1,1,2,0,1];break;case 131:r=[1,1,3,0,1,2];break;case 132:r=[1,1,4,0,1,2,3];break;case 133:r=[1,1,5,0,1,2,3,4];break;case 134:r=[1,1,6,0,1,2,3,4,5];break;case 135:r=[1,1,7,0,1,2,3,4,5,6];break;case 136:r=[1,1,8,0,1,2,3,4,5,6,7]}var s=new Uint8Array(Se([0,i,0,0,a>>>24&255,a>>>17&255,a>>>8&255,a>>>0&255,0,0],r));return e.box(e.types.dOps,s)},e.avc1=function(t){var i=t.avcc,n=t.codecWidth,a=t.codecHeight,r=new Uint8Array([0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,n>>>8&255,255&n,a>>>8&255,255&a,0,72,0,0,0,72,0,0,0,0,0,0,0,1,10,120,113,113,47,102,108,118,46,106,115,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,24,255,255]);return e.box(e.types.avc1,r,e.box(e.types.avcC,i))},e.hvc1=function(t){var i=t.hvcc,n=t.codecWidth,a=t.codecHeight,r=new Uint8Array([0,0,0,0,0,0,0,1,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,n>>>8&255,255&n,a>>>8&255,255&a,0,72,0,0,0,72,0,0,0,0,0,0,0,1,10,120,113,113,47,102,108,118,46,106,115,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,24,255,255]);return e.box(e.types.hvc1,r,e.box(e.types.hvcC,i))},e.mvex=function(t){return e.box(e.types.mvex,e.trex(t))},e.trex=function(t){var i=t.id,n=new Uint8Array([0,0,0,0,i>>>24&255,i>>>16&255,i>>>8&255,255&i,0,0,0,1,0,0,0,0,0,0,0,0,0,1,0,1]);return e.box(e.types.trex,n)},e.moof=function(t,i){return e.box(e.types.moof,e.mfhd(t.sequenceNumber),e.traf(t,i))},e.mfhd=function(t){var i=new Uint8Array([0,0,0,0,t>>>24&255,t>>>16&255,t>>>8&255,255&t]);return e.box(e.types.mfhd,i)},e.traf=function(t,i){var n=t.id,a=e.box(e.types.tfhd,new Uint8Array([0,0,0,0,n>>>24&255,n>>>16&255,n>>>8&255,255&n])),r=e.box(e.types.tfdt,new Uint8Array([0,0,0,0,i>>>24&255,i>>>16&255,i>>>8&255,255&i])),s=e.sdtp(t),o=e.trun(t,s.byteLength+16+16+8+16+8+8);return e.box(e.types.traf,a,r,o,s)},e.sdtp=function(t){for(var i=t.samples||[],n=i.length,a=new Uint8Array(4+n),r=0;r<n;r++){var s=i[r].flags;a[r+4]=s.isLeading<<6|s.dependsOn<<4|s.isDependedOn<<2|s.hasRedundancy}return e.box(e.types.sdtp,a)},e.trun=function(t,i){var n=t.samples||[],a=n.length,r=12+16*a,s=new Uint8Array(r);i+=8+r,s.set([0,0,15,1,a>>>24&255,a>>>16&255,a>>>8&255,255&a,i>>>24&255,i>>>16&255,i>>>8&255,255&i],0);for(var o=0;o<a;o++){var d=n[o].duration,_=n[o].size,h=n[o].flags,c=n[o].cts;s.set([d>>>24&255,d>>>16&255,d>>>8&255,255&d,_>>>24&255,_>>>16&255,_>>>8&255,255&_,h.isLeading<<2|h.dependsOn,h.isDependedOn<<6|h.hasRedundancy<<4|h.isNonSync,0,0,c>>>24&255,c>>>16&255,c>>>8&255,255&c],12+16*o)}return e.box(e.types.trun,s)},e.mdat=function(t){return e.box(e.types.mdat,t)},e}();Ee.init();var Ae=Ee,Re=function(){function e(){}return e.getSilentFrame=function(e,t){if("mp4a.40.2"===e){if(1===t)return new Uint8Array([0,200,0,128,35,128]);if(2===t)return new Uint8Array([33,0,73,144,2,25,0,35,128]);if(3===t)return new Uint8Array([0,200,0,128,32,132,1,38,64,8,100,0,142]);if(4===t)return new Uint8Array([0,200,0,128,32,132,1,38,64,8,100,0,128,44,128,8,2,56]);if(5===t)return new Uint8Array([0,200,0,128,32,132,1,38,64,8,100,0,130,48,4,153,0,33,144,2,56]);if(6===t)return new Uint8Array([0,200,0,128,32,132,1,38,64,8,100,0,130,48,4,153,0,33,144,2,0,178,0,32,8,224])}else{if(1===t)return new Uint8Array([1,64,34,128,163,78,230,128,186,8,0,0,0,28,6,241,193,10,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,94]);if(2===t)return new Uint8Array([1,64,34,128,163,94,230,128,186,8,0,0,0,0,149,0,6,241,161,10,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,94]);if(3===t)return new Uint8Array([1,64,34,128,163,94,230,128,186,8,0,0,0,0,149,0,6,241,161,10,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,90,94])}return null},e}(),Te=i(7),Le=function(){function e(e){this.TAG="MP4Remuxer",this._config=e,this._isLive=!0===e.isLive,this._dtsBase=-1,this._dtsBaseInited=!1,this._audioDtsBase=1/0,this._videoDtsBase=1/0,this._audioNextDts=void 0,this._videoNextDts=void 0,this._audioStashedLastSample=null,this._videoStashedLastSample=null,this._audioMeta=null,this._videoMeta=null,this._audioSegmentInfoList=new Te.c("audio"),this._videoSegmentInfoList=new Te.c("video"),this._onInitSegment=null,this._onMediaSegment=null,this._forceFirstIDR=!(!s.a.chrome||!(s.a.version.major<50||50===s.a.version.major&&s.a.version.build<2661)),this._fillSilentAfterSeek=s.a.msedge||s.a.msie,this._mp3UseMpegAudio=!s.a.firefox,this._fillAudioTimestampGap=this._config.fixAudioTimestampGap}return e.prototype.destroy=function(){this._dtsBase=-1,this._dtsBaseInited=!1,this._audioMeta=null,this._videoMeta=null,this._audioSegmentInfoList.clear(),this._audioSegmentInfoList=null,this._videoSegmentInfoList.clear(),this._videoSegmentInfoList=null,this._onInitSegment=null,this._onMediaSegment=null},e.prototype.bindDataSource=function(e){return e.onDataAvailable=this.remux.bind(this),e.onTrackMetadata=this._onTrackMetadataReceived.bind(this),this},Object.defineProperty(e.prototype,"onInitSegment",{get:function(){return this._onInitSegment},set:function(e){this._onInitSegment=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onMediaSegment",{get:function(){return this._onMediaSegment},set:function(e){this._onMediaSegment=e},enumerable:!1,configurable:!0}),e.prototype.insertDiscontinuity=function(){this._audioNextDts=this._videoNextDts=void 0},e.prototype.seek=function(e){this._audioStashedLastSample=null,this._videoStashedLastSample=null,this._videoSegmentInfoList.clear(),this._audioSegmentInfoList.clear()},e.prototype.remux=function(e,t){if(!this._onMediaSegment)throw new c.a("MP4Remuxer: onMediaSegment callback must be specificed!");this._dtsBaseInited||this._calculateDtsBase(e,t),t&&this._remuxVideo(t),e&&this._remuxAudio(e)},e.prototype._onTrackMetadataReceived=function(e,t){var i=null,n="mp4",a=t.codec;if("audio"===e)this._audioMeta=t,"mp3"===t.codec&&this._mp3UseMpegAudio?(n="mpeg",a="",i=new Uint8Array):i=Ae.generateInitSegment(t);else{if("video"!==e)return;this._videoMeta=t,i=Ae.generateInitSegment(t)}if(!this._onInitSegment)throw new c.a("MP4Remuxer: onInitSegment callback must be specified!");this._onInitSegment(e,{type:e,data:i.buffer,codec:a,container:e+"/"+n,mediaDuration:t.duration})},e.prototype._calculateDtsBase=function(e,t){this._dtsBaseInited||(e&&e.samples&&e.samples.length&&(this._audioDtsBase=e.samples[0].dts),t&&t.samples&&t.samples.length&&(this._videoDtsBase=t.samples[0].dts),this._dtsBase=Math.min(this._audioDtsBase,this._videoDtsBase),this._dtsBaseInited=!0)},e.prototype.getTimestampBase=function(){if(this._dtsBaseInited)return this._dtsBase},e.prototype.flushStashedSamples=function(){var e=this._videoStashedLastSample,t=this._audioStashedLastSample,i={type:"video",id:1,sequenceNumber:0,samples:[],length:0};null!=e&&(i.samples.push(e),i.length=e.length);var n={type:"audio",id:2,sequenceNumber:0,samples:[],length:0};null!=t&&(n.samples.push(t),n.length=t.length),this._videoStashedLastSample=null,this._audioStashedLastSample=null,this._remuxVideo(i,!0),this._remuxAudio(n,!0)},e.prototype._remuxAudio=function(e,t){if(null!=this._audioMeta){var i,n=e,a=n.samples,o=void 0,d=-1,_=this._audioMeta.refSampleDuration,h="mp3"===this._audioMeta.codec&&this._mp3UseMpegAudio,c=this._dtsBaseInited&&void 0===this._audioNextDts,u=!1;if(a&&0!==a.length&&(1!==a.length||t)){var l=0,f=null,p=0;h?(l=0,p=n.length):(l=8,p=8+n.length);var m=null;if(a.length>1&&(p-=(m=a.pop()).length),null!=this._audioStashedLastSample){var g=this._audioStashedLastSample;this._audioStashedLastSample=null,a.unshift(g),p+=g.length}null!=m&&(this._audioStashedLastSample=m);var v=a[0].dts-this._dtsBase;if(this._audioNextDts)o=v-this._audioNextDts;else if(this._audioSegmentInfoList.isEmpty())o=0,this._fillSilentAfterSeek&&!this._videoSegmentInfoList.isEmpty()&&"mp3"!==this._audioMeta.originalCodec&&(u=!0);else{var y=this._audioSegmentInfoList.getLastSampleBefore(v);if(null!=y){var b=v-(y.originalDts+y.duration);b<=3&&(b=0),o=v-(y.dts+y.duration+b)}else o=0}if(u){var S=v-o,E=this._videoSegmentInfoList.getLastSegmentBefore(v);if(null!=E&&E.beginDts<S){if(O=Re.getSilentFrame(this._audioMeta.originalCodec,this._audioMeta.channelCount)){var A=E.beginDts,R=S-E.beginDts;r.a.v(this.TAG,"InsertPrefixSilentAudio: dts: "+A+", duration: "+R),a.unshift({unit:O,dts:A,pts:A}),p+=O.byteLength}}else u=!1}for(var T=[],L=0;L<a.length;L++){var w=(g=a[L]).unit,k=g.dts-this._dtsBase,D=(A=k,!1),C=null,B=0;if(!(k<-.001)){if("mp3"!==this._audioMeta.codec){var I=k;if(this._audioNextDts&&(I=this._audioNextDts),(o=k-I)<=-3*_){r.a.w(this.TAG,"Dropping 1 audio frame (originalDts: "+k+" ms ,curRefDts: "+I+" ms)  due to dtsCorrection: "+o+" ms overlap.");continue}if(o>=3*_&&this._fillAudioTimestampGap&&!s.a.safari){D=!0;var O,P=Math.floor(o/_);r.a.w(this.TAG,"Large audio timestamp gap detected, may cause AV sync to drift. Silent frames will be generated to avoid unsync.\noriginalDts: "+k+" ms, curRefDts: "+I+" ms, dtsCorrection: "+Math.round(o)+" ms, generate: "+P+" frames"),A=Math.floor(I),B=Math.floor(I+_)-A,null==(O=Re.getSilentFrame(this._audioMeta.originalCodec,this._audioMeta.channelCount))&&(r.a.w(this.TAG,"Unable to generate silent frame for "+this._audioMeta.originalCodec+" with "+this._audioMeta.channelCount+" channels, repeat last frame"),O=w),C=[];for(var M=0;M<P;M++){I+=_;var x=Math.floor(I),U=Math.floor(I+_)-x,N={dts:x,pts:x,cts:0,unit:O,size:O.byteLength,duration:U,originalDts:k,flags:{isLeading:0,dependsOn:1,isDependedOn:0,hasRedundancy:0}};C.push(N),p+=N.size}this._audioNextDts=I+_}else A=Math.floor(I),B=Math.floor(I+_)-A,this._audioNextDts=I+_}else{if(A=k-o,L!==a.length-1)B=a[L+1].dts-this._dtsBase-o-A;else if(null!=m)B=m.dts-this._dtsBase-o-A;else B=T.length>=1?T[T.length-1].duration:Math.floor(_);this._audioNextDts=A+B}-1===d&&(d=A),T.push({dts:A,pts:A,cts:0,unit:g.unit,size:g.unit.byteLength,duration:B,originalDts:k,flags:{isLeading:0,dependsOn:1,isDependedOn:0,hasRedundancy:0}}),D&&T.push.apply(T,C)}}if(0===T.length)return n.samples=[],void(n.length=0);h?f=new Uint8Array(p):((f=new Uint8Array(p))[0]=p>>>24&255,f[1]=p>>>16&255,f[2]=p>>>8&255,f[3]=255&p,f.set(Ae.types.mdat,4));for(L=0;L<T.length;L++){w=T[L].unit;f.set(w,l),l+=w.byteLength}var G=T[T.length-1];i=G.dts+G.duration;var V=new Te.b;V.beginDts=d,V.endDts=i,V.beginPts=d,V.endPts=i,V.originalBeginDts=T[0].originalDts,V.originalEndDts=G.originalDts+G.duration,V.firstSample=new Te.d(T[0].dts,T[0].pts,T[0].duration,T[0].originalDts,!1),V.lastSample=new Te.d(G.dts,G.pts,G.duration,G.originalDts,!1),this._isLive||this._audioSegmentInfoList.append(V),n.samples=T,n.sequenceNumber++;var F=null;F=h?new Uint8Array:Ae.moof(n,d),n.samples=[],n.length=0;var j={type:"audio",data:this._mergeBoxes(F,f).buffer,sampleCount:T.length,info:V};h&&c&&(j.timestampOffset=d),this._onMediaSegment("audio",j)}}},e.prototype._remuxVideo=function(e,t){if(null!=this._videoMeta){var i,n,a=e,r=a.samples,s=void 0,o=-1,d=-1;if(r&&0!==r.length&&(1!==r.length||t)){var _=8,h=null,c=8+e.length,u=null;if(r.length>1&&(c-=(u=r.pop()).length),null!=this._videoStashedLastSample){var l=this._videoStashedLastSample;this._videoStashedLastSample=null,r.unshift(l),c+=l.length}null!=u&&(this._videoStashedLastSample=u);var f=r[0].dts-this._dtsBase;if(this._videoNextDts)s=f-this._videoNextDts;else if(this._videoSegmentInfoList.isEmpty())s=0;else{var p=this._videoSegmentInfoList.getLastSampleBefore(f);if(null!=p){var m=f-(p.originalDts+p.duration);m<=3&&(m=0),s=f-(p.dts+p.duration+m)}else s=0}for(var g=new Te.b,v=[],y=0;y<r.length;y++){var b=(l=r[y]).dts-this._dtsBase,S=l.isKeyframe,E=b-s,A=l.cts,R=E+A;-1===o&&(o=E,d=R);var T=0;if(y!==r.length-1)T=r[y+1].dts-this._dtsBase-s-E;else if(null!=u)T=u.dts-this._dtsBase-s-E;else T=v.length>=1?v[v.length-1].duration:Math.floor(this._videoMeta.refSampleDuration);if(S){var L=new Te.d(E,R,T,l.dts,!0);L.fileposition=l.fileposition,g.appendSyncPoint(L)}v.push({dts:E,pts:R,cts:A,units:l.units,size:l.length,isKeyframe:S,duration:T,originalDts:b,flags:{isLeading:0,dependsOn:S?2:1,isDependedOn:S?1:0,hasRedundancy:0,isNonSync:S?0:1}})}(h=new Uint8Array(c))[0]=c>>>24&255,h[1]=c>>>16&255,h[2]=c>>>8&255,h[3]=255&c,h.set(Ae.types.mdat,4);for(y=0;y<v.length;y++)for(var w=v[y].units;w.length;){var k=w.shift().data;h.set(k,_),_+=k.byteLength}var D=v[v.length-1];if(i=D.dts+D.duration,n=D.pts+D.duration,this._videoNextDts=i,g.beginDts=o,g.endDts=i,g.beginPts=d,g.endPts=n,g.originalBeginDts=v[0].originalDts,g.originalEndDts=D.originalDts+D.duration,g.firstSample=new Te.d(v[0].dts,v[0].pts,v[0].duration,v[0].originalDts,v[0].isKeyframe),g.lastSample=new Te.d(D.dts,D.pts,D.duration,D.originalDts,D.isKeyframe),this._isLive||this._videoSegmentInfoList.append(g),a.samples=v,a.sequenceNumber++,this._forceFirstIDR){var C=v[0].flags;C.dependsOn=2,C.isNonSync=0}var B=Ae.moof(a,o);a.samples=[],a.length=0,this._onMediaSegment("video",{type:"video",data:this._mergeBoxes(B,h).buffer,sampleCount:v.length,info:g})}}},e.prototype._mergeBoxes=function(e,t){var i=new Uint8Array(e.byteLength+t.byteLength);return i.set(e,0),i.set(t,e.byteLength),i},e}(),we=i(11),ke=i(1),De=function(){function e(e,t){this.TAG="TransmuxingController",this._emitter=new a.a,this._config=t,e.segments||(e.segments=[{duration:e.duration,filesize:e.filesize,url:e.url}]),"boolean"!=typeof e.cors&&(e.cors=!0),"boolean"!=typeof e.withCredentials&&(e.withCredentials=!1),this._mediaDataSource=e,this._currentSegmentIndex=0;var i=0;this._mediaDataSource.segments.forEach((function(n){n.timestampBase=i,i+=n.duration,n.cors=e.cors,n.withCredentials=e.withCredentials,t.referrerPolicy&&(n.referrerPolicy=t.referrerPolicy)})),isNaN(i)||this._mediaDataSource.duration===i||(this._mediaDataSource.duration=i),this._mediaInfo=null,this._demuxer=null,this._remuxer=null,this._ioctl=null,this._pendingSeekTime=null,this._pendingResolveSeekPoint=null,this._statisticsReporter=null}return e.prototype.destroy=function(){this._mediaInfo=null,this._mediaDataSource=null,this._statisticsReporter&&this._disableStatisticsReporter(),this._ioctl&&(this._ioctl.destroy(),this._ioctl=null),this._demuxer&&(this._demuxer.destroy(),this._demuxer=null),this._remuxer&&(this._remuxer.destroy(),this._remuxer=null),this._emitter.removeAllListeners(),this._emitter=null},e.prototype.on=function(e,t){this._emitter.addListener(e,t)},e.prototype.off=function(e,t){this._emitter.removeListener(e,t)},e.prototype.start=function(){this._loadSegment(0),this._enableStatisticsReporter()},e.prototype._loadSegment=function(e,t){this._currentSegmentIndex=e;var i=this._mediaDataSource.segments[e],n=this._ioctl=new we.a(i,this._config,e);n.onError=this._onIOException.bind(this),n.onSeeked=this._onIOSeeked.bind(this),n.onComplete=this._onIOComplete.bind(this),n.onRedirect=this._onIORedirect.bind(this),n.onRecoveredEarlyEof=this._onIORecoveredEarlyEof.bind(this),t?this._demuxer.bindDataSource(this._ioctl):n.onDataArrival=this._onInitChunkArrival.bind(this),n.open(t)},e.prototype.stop=function(){this._internalAbort(),this._disableStatisticsReporter()},e.prototype._internalAbort=function(){this._ioctl&&(this._ioctl.destroy(),this._ioctl=null)},e.prototype.pause=function(){this._ioctl&&this._ioctl.isWorking()&&(this._ioctl.pause(),this._disableStatisticsReporter())},e.prototype.resume=function(){this._ioctl&&this._ioctl.isPaused()&&(this._ioctl.resume(),this._enableStatisticsReporter())},e.prototype.seek=function(e){if(null!=this._mediaInfo&&this._mediaInfo.isSeekable()){var t=this._searchSegmentIndexContains(e);if(t===this._currentSegmentIndex){var i=this._mediaInfo.segments[t];if(null==i)this._pendingSeekTime=e;else{var n=i.getNearestKeyframe(e);this._remuxer.seek(n.milliseconds),this._ioctl.seek(n.fileposition),this._pendingResolveSeekPoint=n.milliseconds}}else{var a=this._mediaInfo.segments[t];if(null==a)this._pendingSeekTime=e,this._internalAbort(),this._remuxer.seek(),this._remuxer.insertDiscontinuity(),this._loadSegment(t);else{n=a.getNearestKeyframe(e);this._internalAbort(),this._remuxer.seek(e),this._remuxer.insertDiscontinuity(),this._demuxer.resetMediaInfo(),this._demuxer.timestampBase=this._mediaDataSource.segments[t].timestampBase,this._loadSegment(t,n.fileposition),this._pendingResolveSeekPoint=n.milliseconds,this._reportSegmentMediaInfo(t)}}this._enableStatisticsReporter()}},e.prototype._searchSegmentIndexContains=function(e){for(var t=this._mediaDataSource.segments,i=t.length-1,n=0;n<t.length;n++)if(e<t[n].timestampBase){i=n-1;break}return i},e.prototype._onInitChunkArrival=function(e,t){var i=this,n=0;if(t>0)this._demuxer.bindDataSource(this._ioctl),this._demuxer.timestampBase=this._mediaDataSource.segments[this._currentSegmentIndex].timestampBase,n=this._demuxer.parseChunks(e,t);else{var a=null;(a=A.probe(e)).match&&(this._setupFLVDemuxerRemuxer(a),n=this._demuxer.parseChunks(e,t)),a.match||a.needMoreData||(a=be.probe(e)).match&&(this._setupTSDemuxerRemuxer(a),n=this._demuxer.parseChunks(e,t)),a.match||a.needMoreData||(a=null,r.a.e(this.TAG,"Non MPEG-TS/FLV, Unsupported media type!"),Promise.resolve().then((function(){i._internalAbort()})),this._emitter.emit(ke.a.DEMUX_ERROR,m.a.FORMAT_UNSUPPORTED,"Non MPEG-TS/FLV, Unsupported media type!"))}return n},e.prototype._setupFLVDemuxerRemuxer=function(e){this._demuxer=new A(e,this._config),this._remuxer||(this._remuxer=new Le(this._config));var t=this._mediaDataSource;null==t.duration||isNaN(t.duration)||(this._demuxer.overridedDuration=t.duration),"boolean"==typeof t.hasAudio&&(this._demuxer.overridedHasAudio=t.hasAudio),"boolean"==typeof t.hasVideo&&(this._demuxer.overridedHasVideo=t.hasVideo),this._demuxer.timestampBase=t.segments[this._currentSegmentIndex].timestampBase,this._demuxer.onError=this._onDemuxException.bind(this),this._demuxer.onMediaInfo=this._onMediaInfo.bind(this),this._demuxer.onMetaDataArrived=this._onMetaDataArrived.bind(this),this._demuxer.onScriptDataArrived=this._onScriptDataArrived.bind(this),this._remuxer.bindDataSource(this._demuxer.bindDataSource(this._ioctl)),this._remuxer.onInitSegment=this._onRemuxerInitSegmentArrival.bind(this),this._remuxer.onMediaSegment=this._onRemuxerMediaSegmentArrival.bind(this)},e.prototype._setupTSDemuxerRemuxer=function(e){var t=this._demuxer=new be(e,this._config);this._remuxer||(this._remuxer=new Le(this._config)),t.onError=this._onDemuxException.bind(this),t.onMediaInfo=this._onMediaInfo.bind(this),t.onMetaDataArrived=this._onMetaDataArrived.bind(this),t.onTimedID3Metadata=this._onTimedID3Metadata.bind(this),t.onSMPTE2038Metadata=this._onSMPTE2038Metadata.bind(this),t.onSCTE35Metadata=this._onSCTE35Metadata.bind(this),t.onPESPrivateDataDescriptor=this._onPESPrivateDataDescriptor.bind(this),t.onPESPrivateData=this._onPESPrivateData.bind(this),this._remuxer.bindDataSource(this._demuxer),this._demuxer.bindDataSource(this._ioctl),this._remuxer.onInitSegment=this._onRemuxerInitSegmentArrival.bind(this),this._remuxer.onMediaSegment=this._onRemuxerMediaSegmentArrival.bind(this)},e.prototype._onMediaInfo=function(e){var t=this;null==this._mediaInfo&&(this._mediaInfo=Object.assign({},e),this._mediaInfo.keyframesIndex=null,this._mediaInfo.segments=[],this._mediaInfo.segmentCount=this._mediaDataSource.segments.length,Object.setPrototypeOf(this._mediaInfo,o.a.prototype));var i=Object.assign({},e);Object.setPrototypeOf(i,o.a.prototype),this._mediaInfo.segments[this._currentSegmentIndex]=i,this._reportSegmentMediaInfo(this._currentSegmentIndex),null!=this._pendingSeekTime&&Promise.resolve().then((function(){var e=t._pendingSeekTime;t._pendingSeekTime=null,t.seek(e)}))},e.prototype._onMetaDataArrived=function(e){this._emitter.emit(ke.a.METADATA_ARRIVED,e)},e.prototype._onScriptDataArrived=function(e){this._emitter.emit(ke.a.SCRIPTDATA_ARRIVED,e)},e.prototype._onTimedID3Metadata=function(e){var t=this._remuxer.getTimestampBase();null!=t&&(null!=e.pts&&(e.pts-=t),null!=e.dts&&(e.dts-=t),this._emitter.emit(ke.a.TIMED_ID3_METADATA_ARRIVED,e))},e.prototype._onSMPTE2038Metadata=function(e){var t=this._remuxer.getTimestampBase();null!=t&&(null!=e.pts&&(e.pts-=t),null!=e.dts&&(e.dts-=t),null!=e.nearest_pts&&(e.nearest_pts-=t),this._emitter.emit(ke.a.SMPTE2038_METADATA_ARRIVED,e))},e.prototype._onSCTE35Metadata=function(e){var t=this._remuxer.getTimestampBase();null!=t&&(null!=e.pts&&(e.pts-=t),null!=e.nearest_pts&&(e.nearest_pts-=t),this._emitter.emit(ke.a.SCTE35_METADATA_ARRIVED,e))},e.prototype._onPESPrivateDataDescriptor=function(e){this._emitter.emit(ke.a.PES_PRIVATE_DATA_DESCRIPTOR,e)},e.prototype._onPESPrivateData=function(e){var t=this._remuxer.getTimestampBase();null!=t&&(null!=e.pts&&(e.pts-=t),null!=e.nearest_pts&&(e.nearest_pts-=t),null!=e.dts&&(e.dts-=t),this._emitter.emit(ke.a.PES_PRIVATE_DATA_ARRIVED,e))},e.prototype._onIOSeeked=function(){this._remuxer.insertDiscontinuity()},e.prototype._onIOComplete=function(e){var t=e+1;t<this._mediaDataSource.segments.length?(this._internalAbort(),this._remuxer&&this._remuxer.flushStashedSamples(),this._loadSegment(t)):(this._remuxer&&this._remuxer.flushStashedSamples(),this._emitter.emit(ke.a.LOADING_COMPLETE),this._disableStatisticsReporter())},e.prototype._onIORedirect=function(e){var t=this._ioctl.extraData;this._mediaDataSource.segments[t].redirectedURL=e},e.prototype._onIORecoveredEarlyEof=function(){this._emitter.emit(ke.a.RECOVERED_EARLY_EOF)},e.prototype._onIOException=function(e,t){r.a.e(this.TAG,"IOException: type = "+e+", code = "+t.code+", msg = "+t.msg),this._emitter.emit(ke.a.IO_ERROR,e,t),this._disableStatisticsReporter()},e.prototype._onDemuxException=function(e,t){r.a.e(this.TAG,"DemuxException: type = "+e+", info = "+t),this._emitter.emit(ke.a.DEMUX_ERROR,e,t)},e.prototype._onRemuxerInitSegmentArrival=function(e,t){this._emitter.emit(ke.a.INIT_SEGMENT,e,t)},e.prototype._onRemuxerMediaSegmentArrival=function(e,t){if(null==this._pendingSeekTime&&(this._emitter.emit(ke.a.MEDIA_SEGMENT,e,t),null!=this._pendingResolveSeekPoint&&"video"===e)){var i=t.info.syncPoints,n=this._pendingResolveSeekPoint;this._pendingResolveSeekPoint=null,s.a.safari&&i.length>0&&i[0].originalDts===n&&(n=i[0].pts),this._emitter.emit(ke.a.RECOMMEND_SEEKPOINT,n)}},e.prototype._enableStatisticsReporter=function(){null==this._statisticsReporter&&(this._statisticsReporter=self.setInterval(this._reportStatisticsInfo.bind(this),this._config.statisticsInfoReportInterval))},e.prototype._disableStatisticsReporter=function(){this._statisticsReporter&&(self.clearInterval(this._statisticsReporter),this._statisticsReporter=null)},e.prototype._reportSegmentMediaInfo=function(e){var t=this._mediaInfo.segments[e],i=Object.assign({},t);i.duration=this._mediaInfo.duration,i.segmentCount=this._mediaInfo.segmentCount,delete i.segments,delete i.keyframesIndex,this._emitter.emit(ke.a.MEDIA_INFO,i)},e.prototype._reportStatisticsInfo=function(){var e={};e.url=this._ioctl.currentURL,e.hasRedirect=this._ioctl.hasRedirect,e.hasRedirect&&(e.redirectedURL=this._ioctl.currentRedirectedURL),e.speed=this._ioctl.currentSpeed,e.loaderType=this._ioctl.loaderType,e.currentSegmentIndex=this._currentSegmentIndex,e.totalSegmentCount=this._mediaDataSource.segments.length,this._emitter.emit(ke.a.STATISTICS_INFO,e)},e}();t.a=De},function(e,t,i){"use strict";var n,a=i(0),r=function(){function e(){this._firstCheckpoint=0,this._lastCheckpoint=0,this._intervalBytes=0,this._totalBytes=0,this._lastSecondBytes=0,self.performance&&self.performance.now?this._now=self.performance.now.bind(self.performance):this._now=Date.now}return e.prototype.reset=function(){this._firstCheckpoint=this._lastCheckpoint=0,this._totalBytes=this._intervalBytes=0,this._lastSecondBytes=0},e.prototype.addBytes=function(e){0===this._firstCheckpoint?(this._firstCheckpoint=this._now(),this._lastCheckpoint=this._firstCheckpoint,this._intervalBytes+=e,this._totalBytes+=e):this._now()-this._lastCheckpoint<1e3?(this._intervalBytes+=e,this._totalBytes+=e):(this._lastSecondBytes=this._intervalBytes,this._intervalBytes=e,this._totalBytes+=e,this._lastCheckpoint=this._now())},Object.defineProperty(e.prototype,"currentKBps",{get:function(){this.addBytes(0);var e=(this._now()-this._lastCheckpoint)/1e3;return 0==e&&(e=1),this._intervalBytes/e/1024},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"lastSecondKBps",{get:function(){return this.addBytes(0),0!==this._lastSecondBytes?this._lastSecondBytes/1024:this._now()-this._lastCheckpoint>=500?this.currentKBps:0},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"averageKBps",{get:function(){var e=(this._now()-this._firstCheckpoint)/1e3;return this._totalBytes/e/1024},enumerable:!1,configurable:!0}),e}(),s=i(2),o=i(4),d=i(3),_=(n=function(e,t){return(n=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(e,t){e.__proto__=t}||function(e,t){for(var i in t)t.hasOwnProperty(i)&&(e[i]=t[i])})(e,t)},function(e,t){function i(){this.constructor=e}n(e,t),e.prototype=null===t?Object.create(t):(i.prototype=t.prototype,new i)}),h=function(e){function t(t,i){var n=e.call(this,"fetch-stream-loader")||this;return n.TAG="FetchStreamLoader",n._seekHandler=t,n._config=i,n._needStash=!0,n._requestAbort=!1,n._abortController=null,n._contentLength=null,n._receivedLength=0,n}return _(t,e),t.isSupported=function(){try{var e=o.a.msedge&&o.a.version.minor>=15048,t=!o.a.msedge||e;return self.fetch&&self.ReadableStream&&t}catch(e){return!1}},t.prototype.destroy=function(){this.isWorking()&&this.abort(),e.prototype.destroy.call(this)},t.prototype.open=function(e,t){var i=this;this._dataSource=e,this._range=t;var n=e.url;this._config.reuseRedirectedURL&&null!=e.redirectedURL&&(n=e.redirectedURL);var a=this._seekHandler.getConfig(n,t),r=new self.Headers;if("object"==typeof a.headers){var o=a.headers;for(var _ in o)o.hasOwnProperty(_)&&r.append(_,o[_])}var h={method:"GET",headers:r,mode:"cors",cache:"default",referrerPolicy:"no-referrer-when-downgrade"};if("object"==typeof this._config.headers)for(var _ in this._config.headers)r.append(_,this._config.headers[_]);!1===e.cors&&(h.mode="same-origin"),e.withCredentials&&(h.credentials="include"),e.referrerPolicy&&(h.referrerPolicy=e.referrerPolicy),self.AbortController&&(this._abortController=new self.AbortController,h.signal=this._abortController.signal),this._status=s.c.kConnecting,self.fetch(a.url,h).then((function(e){if(i._requestAbort)return i._status=s.c.kIdle,void e.body.cancel();if(e.ok&&e.status>=200&&e.status<=299){if(e.url!==a.url&&i._onURLRedirect){var t=i._seekHandler.removeURLParameters(e.url);i._onURLRedirect(t)}var n=e.headers.get("Content-Length");return null!=n&&(i._contentLength=parseInt(n),0!==i._contentLength&&i._onContentLengthKnown&&i._onContentLengthKnown(i._contentLength)),i._pump.call(i,e.body.getReader())}if(i._status=s.c.kError,!i._onError)throw new d.d("FetchStreamLoader: Http code invalid, "+e.status+" "+e.statusText);i._onError(s.b.HTTP_STATUS_CODE_INVALID,{code:e.status,msg:e.statusText})})).catch((function(e){if(!i._abortController||!i._abortController.signal.aborted){if(i._status=s.c.kError,!i._onError)throw e;i._onError(s.b.EXCEPTION,{code:-1,msg:e.message})}}))},t.prototype.abort=function(){if(this._requestAbort=!0,(this._status!==s.c.kBuffering||!o.a.chrome)&&this._abortController)try{this._abortController.abort()}catch(e){}},t.prototype._pump=function(e){var t=this;return e.read().then((function(i){if(i.done)if(null!==t._contentLength&&t._receivedLength<t._contentLength){t._status=s.c.kError;var n=s.b.EARLY_EOF,a={code:-1,msg:"Fetch stream meet Early-EOF"};if(!t._onError)throw new d.d(a.msg);t._onError(n,a)}else t._status=s.c.kComplete,t._onComplete&&t._onComplete(t._range.from,t._range.from+t._receivedLength-1);else{if(t._abortController&&t._abortController.signal.aborted)return void(t._status=s.c.kComplete);if(!0===t._requestAbort)return t._status=s.c.kComplete,e.cancel();t._status=s.c.kBuffering;var r=i.value.buffer,o=t._range.from+t._receivedLength;t._receivedLength+=r.byteLength,t._onDataArrival&&t._onDataArrival(r,o,t._receivedLength),t._pump(e)}})).catch((function(e){if(t._abortController&&t._abortController.signal.aborted)t._status=s.c.kComplete;else if(11!==e.code||!o.a.msedge){t._status=s.c.kError;var i=0,n=null;if(19!==e.code&&"network error"!==e.message||!(null===t._contentLength||null!==t._contentLength&&t._receivedLength<t._contentLength)?(i=s.b.EXCEPTION,n={code:e.code,msg:e.message}):(i=s.b.EARLY_EOF,n={code:e.code,msg:"Fetch stream meet Early-EOF"}),!t._onError)throw new d.d(n.msg);t._onError(i,n)}}))},t}(s.a),c=function(){var e=function(t,i){return(e=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(e,t){e.__proto__=t}||function(e,t){for(var i in t)t.hasOwnProperty(i)&&(e[i]=t[i])})(t,i)};return function(t,i){function n(){this.constructor=t}e(t,i),t.prototype=null===i?Object.create(i):(n.prototype=i.prototype,new n)}}(),u=function(e){function t(t,i){var n=e.call(this,"xhr-moz-chunked-loader")||this;return n.TAG="MozChunkedLoader",n._seekHandler=t,n._config=i,n._needStash=!0,n._xhr=null,n._requestAbort=!1,n._contentLength=null,n._receivedLength=0,n}return c(t,e),t.isSupported=function(){try{var e=new XMLHttpRequest;return e.open("GET","https://example.com",!0),e.responseType="moz-chunked-arraybuffer","moz-chunked-arraybuffer"===e.responseType}catch(e){return a.a.w("MozChunkedLoader",e.message),!1}},t.prototype.destroy=function(){this.isWorking()&&this.abort(),this._xhr&&(this._xhr.onreadystatechange=null,this._xhr.onprogress=null,this._xhr.onloadend=null,this._xhr.onerror=null,this._xhr=null),e.prototype.destroy.call(this)},t.prototype.open=function(e,t){this._dataSource=e,this._range=t;var i=e.url;this._config.reuseRedirectedURL&&null!=e.redirectedURL&&(i=e.redirectedURL);var n=this._seekHandler.getConfig(i,t);this._requestURL=n.url;var a=this._xhr=new XMLHttpRequest;if(a.open("GET",n.url,!0),a.responseType="moz-chunked-arraybuffer",a.onreadystatechange=this._onReadyStateChange.bind(this),a.onprogress=this._onProgress.bind(this),a.onloadend=this._onLoadEnd.bind(this),a.onerror=this._onXhrError.bind(this),e.withCredentials&&(a.withCredentials=!0),"object"==typeof n.headers){var r=n.headers;for(var o in r)r.hasOwnProperty(o)&&a.setRequestHeader(o,r[o])}if("object"==typeof this._config.headers){r=this._config.headers;for(var o in r)r.hasOwnProperty(o)&&a.setRequestHeader(o,r[o])}this._status=s.c.kConnecting,a.send()},t.prototype.abort=function(){this._requestAbort=!0,this._xhr&&this._xhr.abort(),this._status=s.c.kComplete},t.prototype._onReadyStateChange=function(e){var t=e.target;if(2===t.readyState){if(null!=t.responseURL&&t.responseURL!==this._requestURL&&this._onURLRedirect){var i=this._seekHandler.removeURLParameters(t.responseURL);this._onURLRedirect(i)}if(0!==t.status&&(t.status<200||t.status>299)){if(this._status=s.c.kError,!this._onError)throw new d.d("MozChunkedLoader: Http code invalid, "+t.status+" "+t.statusText);this._onError(s.b.HTTP_STATUS_CODE_INVALID,{code:t.status,msg:t.statusText})}else this._status=s.c.kBuffering}},t.prototype._onProgress=function(e){if(this._status!==s.c.kError){null===this._contentLength&&null!==e.total&&0!==e.total&&(this._contentLength=e.total,this._onContentLengthKnown&&this._onContentLengthKnown(this._contentLength));var t=e.target.response,i=this._range.from+this._receivedLength;this._receivedLength+=t.byteLength,this._onDataArrival&&this._onDataArrival(t,i,this._receivedLength)}},t.prototype._onLoadEnd=function(e){!0!==this._requestAbort?this._status!==s.c.kError&&(this._status=s.c.kComplete,this._onComplete&&this._onComplete(this._range.from,this._range.from+this._receivedLength-1)):this._requestAbort=!1},t.prototype._onXhrError=function(e){this._status=s.c.kError;var t=0,i=null;if(this._contentLength&&e.loaded<this._contentLength?(t=s.b.EARLY_EOF,i={code:-1,msg:"Moz-Chunked stream meet Early-Eof"}):(t=s.b.EXCEPTION,i={code:-1,msg:e.constructor.name+" "+e.type}),!this._onError)throw new d.d(i.msg);this._onError(t,i)},t}(s.a),l=function(){var e=function(t,i){return(e=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(e,t){e.__proto__=t}||function(e,t){for(var i in t)t.hasOwnProperty(i)&&(e[i]=t[i])})(t,i)};return function(t,i){function n(){this.constructor=t}e(t,i),t.prototype=null===i?Object.create(i):(n.prototype=i.prototype,new n)}}(),f=function(e){function t(t,i){var n=e.call(this,"xhr-range-loader")||this;return n.TAG="RangeLoader",n._seekHandler=t,n._config=i,n._needStash=!1,n._chunkSizeKBList=[128,256,384,512,768,1024,1536,2048,3072,4096,5120,6144,7168,8192],n._currentChunkSizeKB=384,n._currentSpeedNormalized=0,n._zeroSpeedChunkCount=0,n._xhr=null,n._speedSampler=new r,n._requestAbort=!1,n._waitForTotalLength=!1,n._totalLengthReceived=!1,n._currentRequestURL=null,n._currentRedirectedURL=null,n._currentRequestRange=null,n._totalLength=null,n._contentLength=null,n._receivedLength=0,n._lastTimeLoaded=0,n}return l(t,e),t.isSupported=function(){try{var e=new XMLHttpRequest;return e.open("GET","https://example.com",!0),e.responseType="arraybuffer","arraybuffer"===e.responseType}catch(e){return a.a.w("RangeLoader",e.message),!1}},t.prototype.destroy=function(){this.isWorking()&&this.abort(),this._xhr&&(this._xhr.onreadystatechange=null,this._xhr.onprogress=null,this._xhr.onload=null,this._xhr.onerror=null,this._xhr=null),e.prototype.destroy.call(this)},Object.defineProperty(t.prototype,"currentSpeed",{get:function(){return this._speedSampler.lastSecondKBps},enumerable:!1,configurable:!0}),t.prototype.open=function(e,t){this._dataSource=e,this._range=t,this._status=s.c.kConnecting;var i=!1;null!=this._dataSource.filesize&&0!==this._dataSource.filesize&&(i=!0,this._totalLength=this._dataSource.filesize),this._totalLengthReceived||i?this._openSubRange():(this._waitForTotalLength=!0,this._internalOpen(this._dataSource,{from:0,to:-1}))},t.prototype._openSubRange=function(){var e=1024*this._currentChunkSizeKB,t=this._range.from+this._receivedLength,i=t+e;null!=this._contentLength&&i-this._range.from>=this._contentLength&&(i=this._range.from+this._contentLength-1),this._currentRequestRange={from:t,to:i},this._internalOpen(this._dataSource,this._currentRequestRange)},t.prototype._internalOpen=function(e,t){this._lastTimeLoaded=0;var i=e.url;this._config.reuseRedirectedURL&&(null!=this._currentRedirectedURL?i=this._currentRedirectedURL:null!=e.redirectedURL&&(i=e.redirectedURL));var n=this._seekHandler.getConfig(i,t);this._currentRequestURL=n.url;var a=this._xhr=new XMLHttpRequest;if(a.open("GET",n.url,!0),a.responseType="arraybuffer",a.onreadystatechange=this._onReadyStateChange.bind(this),a.onprogress=this._onProgress.bind(this),a.onload=this._onLoad.bind(this),a.onerror=this._onXhrError.bind(this),e.withCredentials&&(a.withCredentials=!0),"object"==typeof n.headers){var r=n.headers;for(var s in r)r.hasOwnProperty(s)&&a.setRequestHeader(s,r[s])}if("object"==typeof this._config.headers){r=this._config.headers;for(var s in r)r.hasOwnProperty(s)&&a.setRequestHeader(s,r[s])}a.send()},t.prototype.abort=function(){this._requestAbort=!0,this._internalAbort(),this._status=s.c.kComplete},t.prototype._internalAbort=function(){this._xhr&&(this._xhr.onreadystatechange=null,this._xhr.onprogress=null,this._xhr.onload=null,this._xhr.onerror=null,this._xhr.abort(),this._xhr=null)},t.prototype._onReadyStateChange=function(e){var t=e.target;if(2===t.readyState){if(null!=t.responseURL){var i=this._seekHandler.removeURLParameters(t.responseURL);t.responseURL!==this._currentRequestURL&&i!==this._currentRedirectedURL&&(this._currentRedirectedURL=i,this._onURLRedirect&&this._onURLRedirect(i))}if(t.status>=200&&t.status<=299){if(this._waitForTotalLength)return;this._status=s.c.kBuffering}else{if(this._status=s.c.kError,!this._onError)throw new d.d("RangeLoader: Http code invalid, "+t.status+" "+t.statusText);this._onError(s.b.HTTP_STATUS_CODE_INVALID,{code:t.status,msg:t.statusText})}}},t.prototype._onProgress=function(e){if(this._status!==s.c.kError){if(null===this._contentLength){var t=!1;if(this._waitForTotalLength){this._waitForTotalLength=!1,this._totalLengthReceived=!0,t=!0;var i=e.total;this._internalAbort(),null!=i&0!==i&&(this._totalLength=i)}if(-1===this._range.to?this._contentLength=this._totalLength-this._range.from:this._contentLength=this._range.to-this._range.from+1,t)return void this._openSubRange();this._onContentLengthKnown&&this._onContentLengthKnown(this._contentLength)}var n=e.loaded-this._lastTimeLoaded;this._lastTimeLoaded=e.loaded,this._speedSampler.addBytes(n)}},t.prototype._normalizeSpeed=function(e){var t=this._chunkSizeKBList,i=t.length-1,n=0,a=0,r=i;if(e<t[0])return t[0];for(;a<=r;){if((n=a+Math.floor((r-a)/2))===i||e>=t[n]&&e<t[n+1])return t[n];t[n]<e?a=n+1:r=n-1}},t.prototype._onLoad=function(e){if(this._status!==s.c.kError)if(this._waitForTotalLength)this._waitForTotalLength=!1;else{this._lastTimeLoaded=0;var t=this._speedSampler.lastSecondKBps;if(0===t&&(this._zeroSpeedChunkCount++,this._zeroSpeedChunkCount>=3&&(t=this._speedSampler.currentKBps)),0!==t){var i=this._normalizeSpeed(t);this._currentSpeedNormalized!==i&&(this._currentSpeedNormalized=i,this._currentChunkSizeKB=i)}var n=e.target.response,a=this._range.from+this._receivedLength;this._receivedLength+=n.byteLength;var r=!1;null!=this._contentLength&&this._receivedLength<this._contentLength?this._openSubRange():r=!0,this._onDataArrival&&this._onDataArrival(n,a,this._receivedLength),r&&(this._status=s.c.kComplete,this._onComplete&&this._onComplete(this._range.from,this._range.from+this._receivedLength-1))}},t.prototype._onXhrError=function(e){this._status=s.c.kError;var t=0,i=null;if(this._contentLength&&this._receivedLength>0&&this._receivedLength<this._contentLength?(t=s.b.EARLY_EOF,i={code:-1,msg:"RangeLoader meet Early-Eof"}):(t=s.b.EXCEPTION,i={code:-1,msg:e.constructor.name+" "+e.type}),!this._onError)throw new d.d(i.msg);this._onError(t,i)},t}(s.a),p=function(){var e=function(t,i){return(e=Object.setPrototypeOf||{__proto__:[]}instanceof Array&&function(e,t){e.__proto__=t}||function(e,t){for(var i in t)t.hasOwnProperty(i)&&(e[i]=t[i])})(t,i)};return function(t,i){function n(){this.constructor=t}e(t,i),t.prototype=null===i?Object.create(i):(n.prototype=i.prototype,new n)}}(),m=function(e){function t(){var t=e.call(this,"websocket-loader")||this;return t.TAG="WebSocketLoader",t._needStash=!0,t._ws=null,t._requestAbort=!1,t._receivedLength=0,t}return p(t,e),t.isSupported=function(){try{return void 0!==self.WebSocket}catch(e){return!1}},t.prototype.destroy=function(){this._ws&&this.abort(),e.prototype.destroy.call(this)},t.prototype.open=function(e){try{var t=this._ws=new self.WebSocket(e.url);t.binaryType="arraybuffer",t.onopen=this._onWebSocketOpen.bind(this),t.onclose=this._onWebSocketClose.bind(this),t.onmessage=this._onWebSocketMessage.bind(this),t.onerror=this._onWebSocketError.bind(this),this._status=s.c.kConnecting}catch(e){this._status=s.c.kError;var i={code:e.code,msg:e.message};if(!this._onError)throw new d.d(i.msg);this._onError(s.b.EXCEPTION,i)}},t.prototype.abort=function(){var e=this._ws;!e||0!==e.readyState&&1!==e.readyState||(this._requestAbort=!0,e.close()),this._ws=null,this._status=s.c.kComplete},t.prototype._onWebSocketOpen=function(e){this._status=s.c.kBuffering},t.prototype._onWebSocketClose=function(e){!0!==this._requestAbort?(this._status=s.c.kComplete,this._onComplete&&this._onComplete(0,this._receivedLength-1)):this._requestAbort=!1},t.prototype._onWebSocketMessage=function(e){var t=this;if(e.data instanceof ArrayBuffer)this._dispatchArrayBuffer(e.data);else if(e.data instanceof Blob){var i=new FileReader;i.onload=function(){t._dispatchArrayBuffer(i.result)},i.readAsArrayBuffer(e.data)}else{this._status=s.c.kError;var n={code:-1,msg:"Unsupported WebSocket message type: "+e.data.constructor.name};if(!this._onError)throw new d.d(n.msg);this._onError(s.b.EXCEPTION,n)}},t.prototype._dispatchArrayBuffer=function(e){var t=e,i=this._receivedLength;this._receivedLength+=t.byteLength,this._onDataArrival&&this._onDataArrival(t,i,this._receivedLength)},t.prototype._onWebSocketError=function(e){this._status=s.c.kError;var t={code:e.code,msg:e.message};if(!this._onError)throw new d.d(t.msg);this._onError(s.b.EXCEPTION,t)},t}(s.a),g=function(){function e(e){this._zeroStart=e||!1}return e.prototype.getConfig=function(e,t){var i={};if(0!==t.from||-1!==t.to){var n=void 0;n=-1!==t.to?"bytes="+t.from.toString()+"-"+t.to.toString():"bytes="+t.from.toString()+"-",i.Range=n}else this._zeroStart&&(i.Range="bytes=0-");return{url:e,headers:i}},e.prototype.removeURLParameters=function(e){return e},e}(),v=function(){function e(e,t){this._startName=e,this._endName=t}return e.prototype.getConfig=function(e,t){var i=e;if(0!==t.from||-1!==t.to){var n=!0;-1===i.indexOf("?")&&(i+="?",n=!1),n&&(i+="&"),i+=this._startName+"="+t.from.toString(),-1!==t.to&&(i+="&"+this._endName+"="+t.to.toString())}return{url:i,headers:{}}},e.prototype.removeURLParameters=function(e){var t=e.split("?")[0],i=void 0,n=e.indexOf("?");-1!==n&&(i=e.substring(n+1));var a="";if(null!=i&&i.length>0)for(var r=i.split("&"),s=0;s<r.length;s++){var o=r[s].split("="),d=s>0;o[0]!==this._startName&&o[0]!==this._endName&&(d&&(a+="&"),a+=r[s])}return 0===a.length?t:t+"?"+a},e}(),y=function(){function e(e,t,i){this.TAG="IOController",this._config=t,this._extraData=i,this._stashInitialSize=65536,null!=t.stashInitialSize&&t.stashInitialSize>0&&(this._stashInitialSize=t.stashInitialSize),this._stashUsed=0,this._stashSize=this._stashInitialSize,this._bufferSize=3145728,this._stashBuffer=new ArrayBuffer(this._bufferSize),this._stashByteStart=0,this._enableStash=!0,!1===t.enableStashBuffer&&(this._enableStash=!1),this._loader=null,this._loaderClass=null,this._seekHandler=null,this._dataSource=e,this._isWebSocketURL=/wss?:\/\/(.+?)/.test(e.url),this._refTotalLength=e.filesize?e.filesize:null,this._totalLength=this._refTotalLength,this._fullRequestFlag=!1,this._currentRange=null,this._redirectedURL=null,this._speedNormalized=0,this._speedSampler=new r,this._speedNormalizeList=[32,64,96,128,192,256,384,512,768,1024,1536,2048,3072,4096],this._isEarlyEofReconnecting=!1,this._paused=!1,this._resumeFrom=0,this._onDataArrival=null,this._onSeeked=null,this._onError=null,this._onComplete=null,this._onRedirect=null,this._onRecoveredEarlyEof=null,this._selectSeekHandler(),this._selectLoader(),this._createLoader()}return e.prototype.destroy=function(){this._loader.isWorking()&&this._loader.abort(),this._loader.destroy(),this._loader=null,this._loaderClass=null,this._dataSource=null,this._stashBuffer=null,this._stashUsed=this._stashSize=this._bufferSize=this._stashByteStart=0,this._currentRange=null,this._speedSampler=null,this._isEarlyEofReconnecting=!1,this._onDataArrival=null,this._onSeeked=null,this._onError=null,this._onComplete=null,this._onRedirect=null,this._onRecoveredEarlyEof=null,this._extraData=null},e.prototype.isWorking=function(){return this._loader&&this._loader.isWorking()&&!this._paused},e.prototype.isPaused=function(){return this._paused},Object.defineProperty(e.prototype,"status",{get:function(){return this._loader.status},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"extraData",{get:function(){return this._extraData},set:function(e){this._extraData=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onDataArrival",{get:function(){return this._onDataArrival},set:function(e){this._onDataArrival=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onSeeked",{get:function(){return this._onSeeked},set:function(e){this._onSeeked=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onError",{get:function(){return this._onError},set:function(e){this._onError=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onComplete",{get:function(){return this._onComplete},set:function(e){this._onComplete=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onRedirect",{get:function(){return this._onRedirect},set:function(e){this._onRedirect=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"onRecoveredEarlyEof",{get:function(){return this._onRecoveredEarlyEof},set:function(e){this._onRecoveredEarlyEof=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"currentURL",{get:function(){return this._dataSource.url},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"hasRedirect",{get:function(){return null!=this._redirectedURL||null!=this._dataSource.redirectedURL},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"currentRedirectedURL",{get:function(){return this._redirectedURL||this._dataSource.redirectedURL},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"currentSpeed",{get:function(){return this._loaderClass===f?this._loader.currentSpeed:this._speedSampler.lastSecondKBps},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"loaderType",{get:function(){return this._loader.type},enumerable:!1,configurable:!0}),e.prototype._selectSeekHandler=function(){var e=this._config;if("range"===e.seekType)this._seekHandler=new g(this._config.rangeLoadZeroStart);else if("param"===e.seekType){var t=e.seekParamStart||"bstart",i=e.seekParamEnd||"bend";this._seekHandler=new v(t,i)}else{if("custom"!==e.seekType)throw new d.b("Invalid seekType in config: "+e.seekType);if("function"!=typeof e.customSeekHandler)throw new d.b("Custom seekType specified in config but invalid customSeekHandler!");this._seekHandler=new e.customSeekHandler}},e.prototype._selectLoader=function(){if(null!=this._config.customLoader)this._loaderClass=this._config.customLoader;else if(this._isWebSocketURL)this._loaderClass=m;else if(h.isSupported())this._loaderClass=h;else if(u.isSupported())this._loaderClass=u;else{if(!f.isSupported())throw new d.d("Your browser doesn't support xhr with arraybuffer responseType!");this._loaderClass=f}},e.prototype._createLoader=function(){this._loader=new this._loaderClass(this._seekHandler,this._config),!1===this._loader.needStashBuffer&&(this._enableStash=!1),this._loader.onContentLengthKnown=this._onContentLengthKnown.bind(this),this._loader.onURLRedirect=this._onURLRedirect.bind(this),this._loader.onDataArrival=this._onLoaderChunkArrival.bind(this),this._loader.onComplete=this._onLoaderComplete.bind(this),this._loader.onError=this._onLoaderError.bind(this)},e.prototype.open=function(e){this._currentRange={from:0,to:-1},e&&(this._currentRange.from=e),this._speedSampler.reset(),e||(this._fullRequestFlag=!0),this._loader.open(this._dataSource,Object.assign({},this._currentRange))},e.prototype.abort=function(){this._loader.abort(),this._paused&&(this._paused=!1,this._resumeFrom=0)},e.prototype.pause=function(){this.isWorking()&&(this._loader.abort(),0!==this._stashUsed?(this._resumeFrom=this._stashByteStart,this._currentRange.to=this._stashByteStart-1):this._resumeFrom=this._currentRange.to+1,this._stashUsed=0,this._stashByteStart=0,this._paused=!0)},e.prototype.resume=function(){if(this._paused){this._paused=!1;var e=this._resumeFrom;this._resumeFrom=0,this._internalSeek(e,!0)}},e.prototype.seek=function(e){this._paused=!1,this._stashUsed=0,this._stashByteStart=0,this._internalSeek(e,!0)},e.prototype._internalSeek=function(e,t){this._loader.isWorking()&&this._loader.abort(),this._flushStashBuffer(t),this._loader.destroy(),this._loader=null;var i={from:e,to:-1};this._currentRange={from:i.from,to:-1},this._speedSampler.reset(),this._stashSize=this._stashInitialSize,this._createLoader(),this._loader.open(this._dataSource,i),this._onSeeked&&this._onSeeked()},e.prototype.updateUrl=function(e){if(!e||"string"!=typeof e||0===e.length)throw new d.b("Url must be a non-empty string!");this._dataSource.url=e},e.prototype._expandBuffer=function(e){for(var t=this._stashSize;t+1048576<e;)t*=2;if((t+=1048576)!==this._bufferSize){var i=new ArrayBuffer(t);if(this._stashUsed>0){var n=new Uint8Array(this._stashBuffer,0,this._stashUsed);new Uint8Array(i,0,t).set(n,0)}this._stashBuffer=i,this._bufferSize=t}},e.prototype._normalizeSpeed=function(e){var t=this._speedNormalizeList,i=t.length-1,n=0,a=0,r=i;if(e<t[0])return t[0];for(;a<=r;){if((n=a+Math.floor((r-a)/2))===i||e>=t[n]&&e<t[n+1])return t[n];t[n]<e?a=n+1:r=n-1}},e.prototype._adjustStashSize=function(e){var t=0;(t=this._config.isLive?e/8:e<512?e:e>=512&&e<=1024?Math.floor(1.5*e):2*e)>8192&&(t=8192);var i=1024*t+1048576;this._bufferSize<i&&this._expandBuffer(i),this._stashSize=1024*t},e.prototype._dispatchChunks=function(e,t){return this._currentRange.to=t+e.byteLength-1,this._onDataArrival(e,t)},e.prototype._onURLRedirect=function(e){this._redirectedURL=e,this._onRedirect&&this._onRedirect(e)},e.prototype._onContentLengthKnown=function(e){e&&this._fullRequestFlag&&(this._totalLength=e,this._fullRequestFlag=!1)},e.prototype._onLoaderChunkArrival=function(e,t,i){if(!this._onDataArrival)throw new d.a("IOController: No existing consumer (onDataArrival) callback!");if(!this._paused){this._isEarlyEofReconnecting&&(this._isEarlyEofReconnecting=!1,this._onRecoveredEarlyEof&&this._onRecoveredEarlyEof()),this._speedSampler.addBytes(e.byteLength);var n=this._speedSampler.lastSecondKBps;if(0!==n){var a=this._normalizeSpeed(n);this._speedNormalized!==a&&(this._speedNormalized=a,this._adjustStashSize(a))}if(this._enableStash)if(0===this._stashUsed&&0===this._stashByteStart&&(this._stashByteStart=t),this._stashUsed+e.byteLength<=this._stashSize){(o=new Uint8Array(this._stashBuffer,0,this._stashSize)).set(new Uint8Array(e),this._stashUsed),this._stashUsed+=e.byteLength}else{o=new Uint8Array(this._stashBuffer,0,this._bufferSize);if(this._stashUsed>0){var r=this._stashBuffer.slice(0,this._stashUsed);if((_=this._dispatchChunks(r,this._stashByteStart))<r.byteLength){if(_>0){h=new Uint8Array(r,_);o.set(h,0),this._stashUsed=h.byteLength,this._stashByteStart+=_}}else this._stashUsed=0,this._stashByteStart+=_;this._stashUsed+e.byteLength>this._bufferSize&&(this._expandBuffer(this._stashUsed+e.byteLength),o=new Uint8Array(this._stashBuffer,0,this._bufferSize)),o.set(new Uint8Array(e),this._stashUsed),this._stashUsed+=e.byteLength}else{if((_=this._dispatchChunks(e,t))<e.byteLength)(s=e.byteLength-_)>this._bufferSize&&(this._expandBuffer(s),o=new Uint8Array(this._stashBuffer,0,this._bufferSize)),o.set(new Uint8Array(e,_),0),this._stashUsed+=s,this._stashByteStart=t+_}}else if(0===this._stashUsed){var s;if((_=this._dispatchChunks(e,t))<e.byteLength)(s=e.byteLength-_)>this._bufferSize&&this._expandBuffer(s),(o=new Uint8Array(this._stashBuffer,0,this._bufferSize)).set(new Uint8Array(e,_),0),this._stashUsed+=s,this._stashByteStart=t+_}else{var o,_;if(this._stashUsed+e.byteLength>this._bufferSize&&this._expandBuffer(this._stashUsed+e.byteLength),(o=new Uint8Array(this._stashBuffer,0,this._bufferSize)).set(new Uint8Array(e),this._stashUsed),this._stashUsed+=e.byteLength,(_=this._dispatchChunks(this._stashBuffer.slice(0,this._stashUsed),this._stashByteStart))<this._stashUsed&&_>0){var h=new Uint8Array(this._stashBuffer,_);o.set(h,0)}this._stashUsed-=_,this._stashByteStart+=_}}},e.prototype._flushStashBuffer=function(e){if(this._stashUsed>0){var t=this._stashBuffer.slice(0,this._stashUsed),i=this._dispatchChunks(t,this._stashByteStart),n=t.byteLength-i;if(i<t.byteLength){if(!e){if(i>0){var r=new Uint8Array(this._stashBuffer,0,this._bufferSize),s=new Uint8Array(t,i);r.set(s,0),this._stashUsed=s.byteLength,this._stashByteStart+=i}return 0}a.a.w(this.TAG,n+" bytes unconsumed data remain when flush buffer, dropped")}return this._stashUsed=0,this._stashByteStart=0,n}return 0},e.prototype._onLoaderComplete=function(e,t){this._flushStashBuffer(!0),this._onComplete&&this._onComplete(this._extraData)},e.prototype._onLoaderError=function(e,t){switch(a.a.e(this.TAG,"Loader error, code = "+t.code+", msg = "+t.msg),this._flushStashBuffer(!1),this._isEarlyEofReconnecting&&(this._isEarlyEofReconnecting=!1,e=s.b.UNRECOVERABLE_EARLY_EOF),e){case s.b.EARLY_EOF:if(!this._config.isLive&&this._totalLength){var i=this._currentRange.to+1;return void(i<this._totalLength&&(a.a.w(this.TAG,"Connection lost, trying reconnect..."),this._isEarlyEofReconnecting=!0,this._internalSeek(i,!1)))}e=s.b.UNRECOVERABLE_EARLY_EOF;break;case s.b.UNRECOVERABLE_EARLY_EOF:case s.b.CONNECTING_TIMEOUT:case s.b.HTTP_STATUS_CODE_INVALID:case s.b.EXCEPTION:}if(!this._onError)throw new d.d("IOException: "+t.msg);this._onError(e,t)},e}();t.a=y},function(e,t,i){"use strict";var n=function(){function e(){}return e.install=function(){Object.setPrototypeOf=Object.setPrototypeOf||function(e,t){return e.__proto__=t,e},Object.assign=Object.assign||function(e){if(null==e)throw new TypeError("Cannot convert undefined or null to object");for(var t=Object(e),i=1;i<arguments.length;i++){var n=arguments[i];if(null!=n)for(var a in n)n.hasOwnProperty(a)&&(t[a]=n[a])}return t},"function"!=typeof self.Promise&&i(15).polyfill()},e}();n.install(),t.a=n},function(e,t,i){function n(e){var t={};function i(n){if(t[n])return t[n].exports;var a=t[n]={i:n,l:!1,exports:{}};return e[n].call(a.exports,a,a.exports,i),a.l=!0,a.exports}i.m=e,i.c=t,i.i=function(e){return e},i.d=function(e,t,n){i.o(e,t)||Object.defineProperty(e,t,{configurable:!1,enumerable:!0,get:n})},i.r=function(e){Object.defineProperty(e,"__esModule",{value:!0})},i.n=function(e){var t=e&&e.__esModule?function(){return e.default}:function(){return e};return i.d(t,"a",t),t},i.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)},i.p="/",i.oe=function(e){throw console.error(e),e};var n=i(i.s=ENTRY_MODULE);return n.default||n}function a(e){return(e+"").replace(/[.?*+^$[\]\\(){}|-]/g,"\\$&")}function r(e,t,n){var r={};r[n]=[];var s=t.toString(),o=s.match(/^function\s?\w*\(\w+,\s*\w+,\s*(\w+)\)/);if(!o)return r;for(var d,_=o[1],h=new RegExp("(\\\\n|\\W)"+a(_)+"\\(\\s*(/\\*.*?\\*/)?\\s*.*?([\\.|\\-|\\+|\\w|/|@]+).*?\\)","g");d=h.exec(s);)"dll-reference"!==d[3]&&r[n].push(d[3]);for(h=new RegExp("\\("+a(_)+'\\("(dll-reference\\s([\\.|\\-|\\+|\\w|/|@]+))"\\)\\)\\(\\s*(/\\*.*?\\*/)?\\s*.*?([\\.|\\-|\\+|\\w|/|@]+).*?\\)',"g");d=h.exec(s);)e[d[2]]||(r[n].push(d[1]),e[d[2]]=i(d[1]).m),r[d[2]]=r[d[2]]||[],r[d[2]].push(d[4]);for(var c,u=Object.keys(r),l=0;l<u.length;l++)for(var f=0;f<r[u[l]].length;f++)c=r[u[l]][f],isNaN(1*c)||(r[u[l]][f]=1*r[u[l]][f]);return r}function s(e){return Object.keys(e).reduce((function(t,i){return t||e[i].length>0}),!1)}e.exports=function(e,t){t=t||{};var a={main:i.m},o=t.all?{main:Object.keys(a.main)}:function(e,t){for(var i={main:[t]},n={main:[]},a={main:{}};s(i);)for(var o=Object.keys(i),d=0;d<o.length;d++){var _=o[d],h=i[_].pop();if(a[_]=a[_]||{},!a[_][h]&&e[_][h]){a[_][h]=!0,n[_]=n[_]||[],n[_].push(h);for(var c=r(e,e[_][h],_),u=Object.keys(c),l=0;l<u.length;l++)i[u[l]]=i[u[l]]||[],i[u[l]]=i[u[l]].concat(c[u[l]])}}return n}(a,e),d="";Object.keys(o).filter((function(e){return"main"!==e})).forEach((function(e){for(var t=0;o[e][t];)t++;o[e].push(t),a[e][t]="(function(module, exports, __webpack_require__) { module.exports = __webpack_require__; })",d=d+"var "+e+" = ("+n.toString().replace("ENTRY_MODULE",JSON.stringify(t))+")({"+o[e].map((function(t){return JSON.stringify(t)+": "+a[e][t].toString()})).join(",")+"});\n"})),d=d+"new (("+n.toString().replace("ENTRY_MODULE",JSON.stringify(e))+")({"+o.main.map((function(e){return JSON.stringify(e)+": "+a.main[e].toString()})).join(",")+"}))(self);";var _=new window.Blob([d],{type:"text/javascript"});if(t.bare)return _;var h=(window.URL||window.webkitURL||window.mozURL||window.msURL).createObjectURL(_),c=new window.Worker(h);return c.objectURL=h,c}},function(e,t,i){e.exports=i(19).default},function(e,t,i){(function(t,i){
/*!
 * @overview es6-promise - a tiny implementation of Promises/A+.
 * @copyright Copyright (c) 2014 Yehuda Katz, Tom Dale, Stefan Penner and contributors (Conversion to ES6 API by Jake Archibald)
 * @license   Licensed under MIT license
 *            See https://raw.githubusercontent.com/stefanpenner/es6-promise/master/LICENSE
 * @version   v4.2.8+1e68dce6
 */var n;n=function(){"use strict";function e(e){return"function"==typeof e}var n=Array.isArray?Array.isArray:function(e){return"[object Array]"===Object.prototype.toString.call(e)},a=0,r=void 0,s=void 0,o=function(e,t){f[a]=e,f[a+1]=t,2===(a+=2)&&(s?s(p):b())},d="undefined"!=typeof window?window:void 0,_=d||{},h=_.MutationObserver||_.WebKitMutationObserver,c="undefined"==typeof self&&void 0!==t&&"[object process]"==={}.toString.call(t),u="undefined"!=typeof Uint8ClampedArray&&"undefined"!=typeof importScripts&&"undefined"!=typeof MessageChannel;function l(){var e=setTimeout;return function(){return e(p,1)}}var f=new Array(1e3);function p(){for(var e=0;e<a;e+=2)(0,f[e])(f[e+1]),f[e]=void 0,f[e+1]=void 0;a=0}var m,g,v,y,b=void 0;function S(e,t){var i=this,n=new this.constructor(R);void 0===n[A]&&P(n);var a=i._state;if(a){var r=arguments[a-1];o((function(){return I(a,n,r,i._result)}))}else C(i,n,e,t);return n}function E(e){if(e&&"object"==typeof e&&e.constructor===this)return e;var t=new this(R);return L(t,e),t}c?b=function(){return t.nextTick(p)}:h?(g=0,v=new h(p),y=document.createTextNode(""),v.observe(y,{characterData:!0}),b=function(){y.data=g=++g%2}):u?((m=new MessageChannel).port1.onmessage=p,b=function(){return m.port2.postMessage(0)}):b=void 0===d?function(){try{var e=Function("return this")().require("vertx");return void 0!==(r=e.runOnLoop||e.runOnContext)?function(){r(p)}:l()}catch(e){return l()}}():l();var A=Math.random().toString(36).substring(2);function R(){}function T(t,i,n){i.constructor===t.constructor&&n===S&&i.constructor.resolve===E?function(e,t){1===t._state?k(e,t._result):2===t._state?D(e,t._result):C(t,void 0,(function(t){return L(e,t)}),(function(t){return D(e,t)}))}(t,i):void 0===n?k(t,i):e(n)?function(e,t,i){o((function(e){var n=!1,a=function(e,t,i,n){try{e.call(t,i,n)}catch(e){return e}}(i,t,(function(i){n||(n=!0,t!==i?L(e,i):k(e,i))}),(function(t){n||(n=!0,D(e,t))}),e._label);!n&&a&&(n=!0,D(e,a))}),e)}(t,i,n):k(t,i)}function L(e,t){if(e===t)D(e,new TypeError("You cannot resolve a promise with itself"));else if(a=typeof(n=t),null===n||"object"!==a&&"function"!==a)k(e,t);else{var i=void 0;try{i=t.then}catch(t){return void D(e,t)}T(e,t,i)}var n,a}function w(e){e._onerror&&e._onerror(e._result),B(e)}function k(e,t){void 0===e._state&&(e._result=t,e._state=1,0!==e._subscribers.length&&o(B,e))}function D(e,t){void 0===e._state&&(e._state=2,e._result=t,o(w,e))}function C(e,t,i,n){var a=e._subscribers,r=a.length;e._onerror=null,a[r]=t,a[r+1]=i,a[r+2]=n,0===r&&e._state&&o(B,e)}function B(e){var t=e._subscribers,i=e._state;if(0!==t.length){for(var n=void 0,a=void 0,r=e._result,s=0;s<t.length;s+=3)n=t[s],a=t[s+i],n?I(i,n,a,r):a(r);e._subscribers.length=0}}function I(t,i,n,a){var r=e(n),s=void 0,o=void 0,d=!0;if(r){try{s=n(a)}catch(e){d=!1,o=e}if(i===s)return void D(i,new TypeError("A promises callback cannot return that same promise."))}else s=a;void 0!==i._state||(r&&d?L(i,s):!1===d?D(i,o):1===t?k(i,s):2===t&&D(i,s))}var O=0;function P(e){e[A]=O++,e._state=void 0,e._result=void 0,e._subscribers=[]}var M=function(){function e(e,t){this._instanceConstructor=e,this.promise=new e(R),this.promise[A]||P(this.promise),n(t)?(this.length=t.length,this._remaining=t.length,this._result=new Array(this.length),0===this.length?k(this.promise,this._result):(this.length=this.length||0,this._enumerate(t),0===this._remaining&&k(this.promise,this._result))):D(this.promise,new Error("Array Methods must be provided an Array"))}return e.prototype._enumerate=function(e){for(var t=0;void 0===this._state&&t<e.length;t++)this._eachEntry(e[t],t)},e.prototype._eachEntry=function(e,t){var i=this._instanceConstructor,n=i.resolve;if(n===E){var a=void 0,r=void 0,s=!1;try{a=e.then}catch(e){s=!0,r=e}if(a===S&&void 0!==e._state)this._settledAt(e._state,t,e._result);else if("function"!=typeof a)this._remaining--,this._result[t]=e;else if(i===x){var o=new i(R);s?D(o,r):T(o,e,a),this._willSettleAt(o,t)}else this._willSettleAt(new i((function(t){return t(e)})),t)}else this._willSettleAt(n(e),t)},e.prototype._settledAt=function(e,t,i){var n=this.promise;void 0===n._state&&(this._remaining--,2===e?D(n,i):this._result[t]=i),0===this._remaining&&k(n,this._result)},e.prototype._willSettleAt=function(e,t){var i=this;C(e,void 0,(function(e){return i._settledAt(1,t,e)}),(function(e){return i._settledAt(2,t,e)}))},e}(),x=function(){function t(e){this[A]=O++,this._result=this._state=void 0,this._subscribers=[],R!==e&&("function"!=typeof e&&function(){throw new TypeError("You must pass a resolver function as the first argument to the promise constructor")}(),this instanceof t?function(e,t){try{t((function(t){L(e,t)}),(function(t){D(e,t)}))}catch(t){D(e,t)}}(this,e):function(){throw new TypeError("Failed to construct 'Promise': Please use the 'new' operator, this object constructor cannot be called as a function.")}())}return t.prototype.catch=function(e){return this.then(null,e)},t.prototype.finally=function(t){var i=this.constructor;return e(t)?this.then((function(e){return i.resolve(t()).then((function(){return e}))}),(function(e){return i.resolve(t()).then((function(){throw e}))})):this.then(t,t)},t}();return x.prototype.then=S,x.all=function(e){return new M(this,e).promise},x.race=function(e){var t=this;return n(e)?new t((function(i,n){for(var a=e.length,r=0;r<a;r++)t.resolve(e[r]).then(i,n)})):new t((function(e,t){return t(new TypeError("You must pass an array to race."))}))},x.resolve=E,x.reject=function(e){var t=new this(R);return D(t,e),t},x._setScheduler=function(e){s=e},x._setAsap=function(e){o=e},x._asap=o,x.polyfill=function(){var e=void 0;if(void 0!==i)e=i;else if("undefined"!=typeof self)e=self;else try{e=Function("return this")()}catch(e){throw new Error("polyfill failed because global object is unavailable in this environment")}var t=e.Promise;if(t){var n=null;try{n=Object.prototype.toString.call(t.resolve())}catch(e){}if("[object Promise]"===n&&!t.cast)return}e.Promise=x},x.Promise=x,x},e.exports=n()}).call(this,i(16),i(17))},function(e,t){var i,n,a=e.exports={};function r(){throw new Error("setTimeout has not been defined")}function s(){throw new Error("clearTimeout has not been defined")}function o(e){if(i===setTimeout)return setTimeout(e,0);if((i===r||!i)&&setTimeout)return i=setTimeout,setTimeout(e,0);try{return i(e,0)}catch(t){try{return i.call(null,e,0)}catch(t){return i.call(this,e,0)}}}!function(){try{i="function"==typeof setTimeout?setTimeout:r}catch(e){i=r}try{n="function"==typeof clearTimeout?clearTimeout:s}catch(e){n=s}}();var d,_=[],h=!1,c=-1;function u(){h&&d&&(h=!1,d.length?_=d.concat(_):c=-1,_.length&&l())}function l(){if(!h){var e=o(u);h=!0;for(var t=_.length;t;){for(d=_,_=[];++c<t;)d&&d[c].run();c=-1,t=_.length}d=null,h=!1,function(e){if(n===clearTimeout)return clearTimeout(e);if((n===s||!n)&&clearTimeout)return n=clearTimeout,clearTimeout(e);try{n(e)}catch(t){try{return n.call(null,e)}catch(t){return n.call(this,e)}}}(e)}}function f(e,t){this.fun=e,this.array=t}function p(){}a.nextTick=function(e){var t=new Array(arguments.length-1);if(arguments.length>1)for(var i=1;i<arguments.length;i++)t[i-1]=arguments[i];_.push(new f(e,t)),1!==_.length||h||o(l)},f.prototype.run=function(){this.fun.apply(null,this.array)},a.title="browser",a.browser=!0,a.env={},a.argv=[],a.version="",a.versions={},a.on=p,a.addListener=p,a.once=p,a.off=p,a.removeListener=p,a.removeAllListeners=p,a.emit=p,a.prependListener=p,a.prependOnceListener=p,a.listeners=function(e){return[]},a.binding=function(e){throw new Error("process.binding is not supported")},a.cwd=function(){return"/"},a.chdir=function(e){throw new Error("process.chdir is not supported")},a.umask=function(){return 0}},function(e,t){var i;i=function(){return this}();try{i=i||new Function("return this")()}catch(e){"object"==typeof window&&(i=window)}e.exports=i},function(e,t,i){"use strict";i.r(t);var n=i(9),a=i(12),r=i(10),s=i(1);t.default=function(e){var t=null,i=function(t,i){e.postMessage({msg:"logcat_callback",data:{type:t,logcat:i}})}.bind(this);function o(t,i){var n={msg:s.a.INIT_SEGMENT,data:{type:t,data:i}};e.postMessage(n,[i.data])}function d(t,i){var n={msg:s.a.MEDIA_SEGMENT,data:{type:t,data:i}};e.postMessage(n,[i.data])}function _(){var t={msg:s.a.LOADING_COMPLETE};e.postMessage(t)}function h(){var t={msg:s.a.RECOVERED_EARLY_EOF};e.postMessage(t)}function c(t){var i={msg:s.a.MEDIA_INFO,data:t};e.postMessage(i)}function u(t){var i={msg:s.a.METADATA_ARRIVED,data:t};e.postMessage(i)}function l(t){var i={msg:s.a.SCRIPTDATA_ARRIVED,data:t};e.postMessage(i)}function f(t){var i={msg:s.a.TIMED_ID3_METADATA_ARRIVED,data:t};e.postMessage(i)}function p(t){var i={msg:s.a.SMPTE2038_METADATA_ARRIVED,data:t};e.postMessage(i)}function m(t){var i={msg:s.a.SCTE35_METADATA_ARRIVED,data:t};e.postMessage(i)}function g(t){var i={msg:s.a.PES_PRIVATE_DATA_DESCRIPTOR,data:t};e.postMessage(i)}function v(t){var i={msg:s.a.PES_PRIVATE_DATA_ARRIVED,data:t};e.postMessage(i)}function y(t){var i={msg:s.a.STATISTICS_INFO,data:t};e.postMessage(i)}function b(t,i){e.postMessage({msg:s.a.IO_ERROR,data:{type:t,info:i}})}function S(t,i){e.postMessage({msg:s.a.DEMUX_ERROR,data:{type:t,info:i}})}function E(t){e.postMessage({msg:s.a.RECOMMEND_SEEKPOINT,data:t})}a.a.install(),e.addEventListener("message",(function(a){switch(a.data.cmd){case"init":(t=new r.a(a.data.param[0],a.data.param[1])).on(s.a.IO_ERROR,b.bind(this)),t.on(s.a.DEMUX_ERROR,S.bind(this)),t.on(s.a.INIT_SEGMENT,o.bind(this)),t.on(s.a.MEDIA_SEGMENT,d.bind(this)),t.on(s.a.LOADING_COMPLETE,_.bind(this)),t.on(s.a.RECOVERED_EARLY_EOF,h.bind(this)),t.on(s.a.MEDIA_INFO,c.bind(this)),t.on(s.a.METADATA_ARRIVED,u.bind(this)),t.on(s.a.SCRIPTDATA_ARRIVED,l.bind(this)),t.on(s.a.TIMED_ID3_METADATA_ARRIVED,f.bind(this)),t.on(s.a.SMPTE2038_METADATA_ARRIVED,p.bind(this)),t.on(s.a.SCTE35_METADATA_ARRIVED,m.bind(this)),t.on(s.a.PES_PRIVATE_DATA_DESCRIPTOR,g.bind(this)),t.on(s.a.PES_PRIVATE_DATA_ARRIVED,v.bind(this)),t.on(s.a.STATISTICS_INFO,y.bind(this)),t.on(s.a.RECOMMEND_SEEKPOINT,E.bind(this));break;case"destroy":t&&(t.destroy(),t=null),e.postMessage({msg:"destroyed"});break;case"start":t.start();break;case"stop":t.stop();break;case"seek":t.seek(a.data.param);break;case"pause":t.pause();break;case"resume":t.resume();break;case"logging_config":var A=a.data.param;n.a.applyConfig(A),!0===A.enableCallback?n.a.addLogListener(i):n.a.removeLogListener(i)}}))}},function(e,t,i){"use strict";i.r(t);var n=i(12),a=i(11),r={enableWorker:!1,enableStashBuffer:!0,stashInitialSize:void 0,isLive:!1,liveBufferLatencyChasing:!1,liveBufferLatencyMaxLatency:1.5,liveBufferLatencyMinRemain:.5,lazyLoad:!0,lazyLoadMaxDuration:180,lazyLoadRecoverDuration:30,deferLoadAfterSourceOpen:!0,autoCleanupMaxBackwardDuration:180,autoCleanupMinBackwardDuration:120,statisticsInfoReportInterval:600,fixAudioTimestampGap:!0,accurateSeek:!1,seekType:"range",seekParamStart:"bstart",seekParamEnd:"bend",rangeLoadZeroStart:!1,customSeekHandler:void 0,reuseRedirectedURL:!1,headers:void 0,customLoader:void 0};function s(){return Object.assign({},r)}var o=function(){function e(){}return e.supportMSEH264Playback=function(){return window.MediaSource&&window.MediaSource.isTypeSupported('video/mp4; codecs="avc1.42E01E,mp4a.40.2"')},e.supportMSEH265Playback=function(){return window.MediaSource&&window.MediaSource.isTypeSupported('video/mp4; codecs="hvc1.1.6.L93.B0"')},e.supportNetworkStreamIO=function(){var e=new a.a({},s()),t=e.loaderType;return e.destroy(),"fetch-stream-loader"==t||"xhr-moz-chunked-loader"==t},e.getNetworkLoaderTypeName=function(){var e=new a.a({},s()),t=e.loaderType;return e.destroy(),t},e.supportNativeMediaPlayback=function(t){null==e.videoElement&&(e.videoElement=window.document.createElement("video"));var i=e.videoElement.canPlayType(t);return"probably"===i||"maybe"==i},e.getFeatureList=function(){var t={msePlayback:!1,mseLivePlayback:!1,mseH265Playback:!1,networkStreamIO:!1,networkLoaderName:"",nativeMP4H264Playback:!1,nativeMP4H265Playback:!1,nativeWebmVP8Playback:!1,nativeWebmVP9Playback:!1};return t.msePlayback=e.supportMSEH264Playback(),t.networkStreamIO=e.supportNetworkStreamIO(),t.networkLoaderName=e.getNetworkLoaderTypeName(),t.mseLivePlayback=t.msePlayback&&t.networkStreamIO,t.mseH265Playback=e.supportMSEH265Playback(),t.nativeMP4H264Playback=e.supportNativeMediaPlayback('video/mp4; codecs="avc1.42001E, mp4a.40.2"'),t.nativeMP4H265Playback=e.supportNativeMediaPlayback('video/mp4; codecs="hvc1.1.6.L93.B0"'),t.nativeWebmVP8Playback=e.supportNativeMediaPlayback('video/webm; codecs="vp8.0, vorbis"'),t.nativeWebmVP9Playback=e.supportNativeMediaPlayback('video/webm; codecs="vp9"'),t},e}(),d=i(2),_=i(6),h=i.n(_),c=i(0),u=i(4),l={ERROR:"error",LOADING_COMPLETE:"loading_complete",RECOVERED_EARLY_EOF:"recovered_early_eof",MEDIA_INFO:"media_info",METADATA_ARRIVED:"metadata_arrived",SCRIPTDATA_ARRIVED:"scriptdata_arrived",TIMED_ID3_METADATA_ARRIVED:"timed_id3_metadata_arrived",SMPTE2038_METADATA_ARRIVED:"smpte2038_metadata_arrived",SCTE35_METADATA_ARRIVED:"scte35_metadata_arrived",PES_PRIVATE_DATA_DESCRIPTOR:"pes_private_data_descriptor",PES_PRIVATE_DATA_ARRIVED:"pes_private_data_arrived",STATISTICS_INFO:"statistics_info"},f=i(13),p=i.n(f),m=i(9),g=i(10),v=i(1),y=i(8),b=function(){function e(e,t){if(this.TAG="Transmuxer",this._emitter=new h.a,t.enableWorker&&"undefined"!=typeof Worker)try{this._worker=p()(18),this._workerDestroying=!1,this._worker.addEventListener("message",this._onWorkerMessage.bind(this)),this._worker.postMessage({cmd:"init",param:[e,t]}),this.e={onLoggingConfigChanged:this._onLoggingConfigChanged.bind(this)},m.a.registerListener(this.e.onLoggingConfigChanged),this._worker.postMessage({cmd:"logging_config",param:m.a.getConfig()})}catch(i){c.a.e(this.TAG,"Error while initialize transmuxing worker, fallback to inline transmuxing"),this._worker=null,this._controller=new g.a(e,t)}else this._controller=new g.a(e,t);if(this._controller){var i=this._controller;i.on(v.a.IO_ERROR,this._onIOError.bind(this)),i.on(v.a.DEMUX_ERROR,this._onDemuxError.bind(this)),i.on(v.a.INIT_SEGMENT,this._onInitSegment.bind(this)),i.on(v.a.MEDIA_SEGMENT,this._onMediaSegment.bind(this)),i.on(v.a.LOADING_COMPLETE,this._onLoadingComplete.bind(this)),i.on(v.a.RECOVERED_EARLY_EOF,this._onRecoveredEarlyEof.bind(this)),i.on(v.a.MEDIA_INFO,this._onMediaInfo.bind(this)),i.on(v.a.METADATA_ARRIVED,this._onMetaDataArrived.bind(this)),i.on(v.a.SCRIPTDATA_ARRIVED,this._onScriptDataArrived.bind(this)),i.on(v.a.TIMED_ID3_METADATA_ARRIVED,this._onTimedID3MetadataArrived.bind(this)),i.on(v.a.SMPTE2038_METADATA_ARRIVED,this._onSMPTE2038MetadataArrived.bind(this)),i.on(v.a.SCTE35_METADATA_ARRIVED,this._onSCTE35MetadataArrived.bind(this)),i.on(v.a.PES_PRIVATE_DATA_DESCRIPTOR,this._onPESPrivateDataDescriptor.bind(this)),i.on(v.a.PES_PRIVATE_DATA_ARRIVED,this._onPESPrivateDataArrived.bind(this)),i.on(v.a.STATISTICS_INFO,this._onStatisticsInfo.bind(this)),i.on(v.a.RECOMMEND_SEEKPOINT,this._onRecommendSeekpoint.bind(this))}}return e.prototype.destroy=function(){this._worker?this._workerDestroying||(this._workerDestroying=!0,this._worker.postMessage({cmd:"destroy"}),m.a.removeListener(this.e.onLoggingConfigChanged),this.e=null):(this._controller.destroy(),this._controller=null),this._emitter.removeAllListeners(),this._emitter=null},e.prototype.on=function(e,t){this._emitter.addListener(e,t)},e.prototype.off=function(e,t){this._emitter.removeListener(e,t)},e.prototype.hasWorker=function(){return null!=this._worker},e.prototype.open=function(){this._worker?this._worker.postMessage({cmd:"start"}):this._controller.start()},e.prototype.close=function(){this._worker?this._worker.postMessage({cmd:"stop"}):this._controller.stop()},e.prototype.seek=function(e){this._worker?this._worker.postMessage({cmd:"seek",param:e}):this._controller.seek(e)},e.prototype.pause=function(){this._worker?this._worker.postMessage({cmd:"pause"}):this._controller.pause()},e.prototype.resume=function(){this._worker?this._worker.postMessage({cmd:"resume"}):this._controller.resume()},e.prototype._onInitSegment=function(e,t){var i=this;Promise.resolve().then((function(){i._emitter.emit(v.a.INIT_SEGMENT,e,t)}))},e.prototype._onMediaSegment=function(e,t){var i=this;Promise.resolve().then((function(){i._emitter.emit(v.a.MEDIA_SEGMENT,e,t)}))},e.prototype._onLoadingComplete=function(){var e=this;Promise.resolve().then((function(){e._emitter.emit(v.a.LOADING_COMPLETE)}))},e.prototype._onRecoveredEarlyEof=function(){var e=this;Promise.resolve().then((function(){e._emitter.emit(v.a.RECOVERED_EARLY_EOF)}))},e.prototype._onMediaInfo=function(e){var t=this;Promise.resolve().then((function(){t._emitter.emit(v.a.MEDIA_INFO,e)}))},e.prototype._onMetaDataArrived=function(e){var t=this;Promise.resolve().then((function(){t._emitter.emit(v.a.METADATA_ARRIVED,e)}))},e.prototype._onScriptDataArrived=function(e){var t=this;Promise.resolve().then((function(){t._emitter.emit(v.a.SCRIPTDATA_ARRIVED,e)}))},e.prototype._onTimedID3MetadataArrived=function(e){var t=this;Promise.resolve().then((function(){t._emitter.emit(v.a.TIMED_ID3_METADATA_ARRIVED,e)}))},e.prototype._onSMPTE2038MetadataArrived=function(e){var t=this;Promise.resolve().then((function(){t._emitter.emit(v.a.SMPTE2038_METADATA_ARRIVED,e)}))},e.prototype._onSCTE35MetadataArrived=function(e){var t=this;Promise.resolve().then((function(){t._emitter.emit(v.a.SCTE35_METADATA_ARRIVED,e)}))},e.prototype._onPESPrivateDataDescriptor=function(e){var t=this;Promise.resolve().then((function(){t._emitter.emit(v.a.PES_PRIVATE_DATA_DESCRIPTOR,e)}))},e.prototype._onPESPrivateDataArrived=function(e){var t=this;Promise.resolve().then((function(){t._emitter.emit(v.a.PES_PRIVATE_DATA_ARRIVED,e)}))},e.prototype._onStatisticsInfo=function(e){var t=this;Promise.resolve().then((function(){t._emitter.emit(v.a.STATISTICS_INFO,e)}))},e.prototype._onIOError=function(e,t){var i=this;Promise.resolve().then((function(){i._emitter.emit(v.a.IO_ERROR,e,t)}))},e.prototype._onDemuxError=function(e,t){var i=this;Promise.resolve().then((function(){i._emitter.emit(v.a.DEMUX_ERROR,e,t)}))},e.prototype._onRecommendSeekpoint=function(e){var t=this;Promise.resolve().then((function(){t._emitter.emit(v.a.RECOMMEND_SEEKPOINT,e)}))},e.prototype._onLoggingConfigChanged=function(e){this._worker&&this._worker.postMessage({cmd:"logging_config",param:e})},e.prototype._onWorkerMessage=function(e){var t=e.data,i=t.data;if("destroyed"===t.msg||this._workerDestroying)return this._workerDestroying=!1,this._worker.terminate(),void(this._worker=null);switch(t.msg){case v.a.INIT_SEGMENT:case v.a.MEDIA_SEGMENT:this._emitter.emit(t.msg,i.type,i.data);break;case v.a.LOADING_COMPLETE:case v.a.RECOVERED_EARLY_EOF:this._emitter.emit(t.msg);break;case v.a.MEDIA_INFO:Object.setPrototypeOf(i,y.a.prototype),this._emitter.emit(t.msg,i);break;case v.a.METADATA_ARRIVED:case v.a.SCRIPTDATA_ARRIVED:case v.a.TIMED_ID3_METADATA_ARRIVED:case v.a.SMPTE2038_METADATA_ARRIVED:case v.a.SCTE35_METADATA_ARRIVED:case v.a.PES_PRIVATE_DATA_DESCRIPTOR:case v.a.PES_PRIVATE_DATA_ARRIVED:case v.a.STATISTICS_INFO:this._emitter.emit(t.msg,i);break;case v.a.IO_ERROR:case v.a.DEMUX_ERROR:this._emitter.emit(t.msg,i.type,i.info);break;case v.a.RECOMMEND_SEEKPOINT:this._emitter.emit(t.msg,i);break;case"logcat_callback":c.a.emitter.emit("log",i.type,i.logcat)}},e}(),S={ERROR:"error",SOURCE_OPEN:"source_open",UPDATE_END:"update_end",BUFFER_FULL:"buffer_full"},E=i(7),A=i(3),R=function(){function e(e){this.TAG="MSEController",this._config=e,this._emitter=new h.a,this._config.isLive&&null==this._config.autoCleanupSourceBuffer&&(this._config.autoCleanupSourceBuffer=!0),this.e={onSourceOpen:this._onSourceOpen.bind(this),onSourceEnded:this._onSourceEnded.bind(this),onSourceClose:this._onSourceClose.bind(this),onSourceBufferError:this._onSourceBufferError.bind(this),onSourceBufferUpdateEnd:this._onSourceBufferUpdateEnd.bind(this)},this._mediaSource=null,this._mediaSourceObjectURL=null,this._mediaElement=null,this._isBufferFull=!1,this._hasPendingEos=!1,this._requireSetMediaDuration=!1,this._pendingMediaDuration=0,this._pendingSourceBufferInit=[],this._mimeTypes={video:null,audio:null},this._sourceBuffers={video:null,audio:null},this._lastInitSegments={video:null,audio:null},this._pendingSegments={video:[],audio:[]},this._pendingRemoveRanges={video:[],audio:[]},this._idrList=new E.a}return e.prototype.destroy=function(){(this._mediaElement||this._mediaSource)&&this.detachMediaElement(),this.e=null,this._emitter.removeAllListeners(),this._emitter=null},e.prototype.on=function(e,t){this._emitter.addListener(e,t)},e.prototype.off=function(e,t){this._emitter.removeListener(e,t)},e.prototype.attachMediaElement=function(e){if(this._mediaSource)throw new A.a("MediaSource has been attached to an HTMLMediaElement!");var t=this._mediaSource=new window.MediaSource;t.addEventListener("sourceopen",this.e.onSourceOpen),t.addEventListener("sourceended",this.e.onSourceEnded),t.addEventListener("sourceclose",this.e.onSourceClose),this._mediaElement=e,this._mediaSourceObjectURL=window.URL.createObjectURL(this._mediaSource),e.src=this._mediaSourceObjectURL},e.prototype.detachMediaElement=function(){if(this._mediaSource){var e=this._mediaSource;for(var t in this._sourceBuffers){var i=this._pendingSegments[t];i.splice(0,i.length),this._pendingSegments[t]=null,this._pendingRemoveRanges[t]=null,this._lastInitSegments[t]=null;var n=this._sourceBuffers[t];if(n){if("closed"!==e.readyState){try{e.removeSourceBuffer(n)}catch(e){c.a.e(this.TAG,e.message)}n.removeEventListener("error",this.e.onSourceBufferError),n.removeEventListener("updateend",this.e.onSourceBufferUpdateEnd)}this._mimeTypes[t]=null,this._sourceBuffers[t]=null}}if("open"===e.readyState)try{e.endOfStream()}catch(e){c.a.e(this.TAG,e.message)}e.removeEventListener("sourceopen",this.e.onSourceOpen),e.removeEventListener("sourceended",this.e.onSourceEnded),e.removeEventListener("sourceclose",this.e.onSourceClose),this._pendingSourceBufferInit=[],this._isBufferFull=!1,this._idrList.clear(),this._mediaSource=null}this._mediaElement&&(this._mediaElement.src="",this._mediaElement.removeAttribute("src"),this._mediaElement=null),this._mediaSourceObjectURL&&(window.URL.revokeObjectURL(this._mediaSourceObjectURL),this._mediaSourceObjectURL=null)},e.prototype.appendInitSegment=function(e,t){if(!this._mediaSource||"open"!==this._mediaSource.readyState)return this._pendingSourceBufferInit.push(e),void this._pendingSegments[e.type].push(e);var i=e,n=""+i.container;i.codec&&i.codec.length>0&&(n+=";codecs="+i.codec);var a=!1;if(c.a.v(this.TAG,"Received Initialization Segment, mimeType: "+n),this._lastInitSegments[i.type]=i,n!==this._mimeTypes[i.type]){if(this._mimeTypes[i.type])c.a.v(this.TAG,"Notice: "+i.type+" mimeType changed, origin: "+this._mimeTypes[i.type]+", target: "+n);else{a=!0;try{var r=this._sourceBuffers[i.type]=this._mediaSource.addSourceBuffer(n);r.addEventListener("error",this.e.onSourceBufferError),r.addEventListener("updateend",this.e.onSourceBufferUpdateEnd)}catch(e){return c.a.e(this.TAG,e.message),void this._emitter.emit(S.ERROR,{code:e.code,msg:e.message})}}this._mimeTypes[i.type]=n}t||this._pendingSegments[i.type].push(i),a||this._sourceBuffers[i.type]&&!this._sourceBuffers[i.type].updating&&this._doAppendSegments(),u.a.safari&&"audio/mpeg"===i.container&&i.mediaDuration>0&&(this._requireSetMediaDuration=!0,this._pendingMediaDuration=i.mediaDuration/1e3,this._updateMediaSourceDuration())},e.prototype.appendMediaSegment=function(e){var t=e;this._pendingSegments[t.type].push(t),this._config.autoCleanupSourceBuffer&&this._needCleanupSourceBuffer()&&this._doCleanupSourceBuffer();var i=this._sourceBuffers[t.type];!i||i.updating||this._hasPendingRemoveRanges()||this._doAppendSegments()},e.prototype.seek=function(e){for(var t in this._sourceBuffers)if(this._sourceBuffers[t]){var i=this._sourceBuffers[t];if("open"===this._mediaSource.readyState)try{i.abort()}catch(e){c.a.e(this.TAG,e.message)}this._idrList.clear();var n=this._pendingSegments[t];if(n.splice(0,n.length),"closed"!==this._mediaSource.readyState){for(var a=0;a<i.buffered.length;a++){var r=i.buffered.start(a),s=i.buffered.end(a);this._pendingRemoveRanges[t].push({start:r,end:s})}if(i.updating||this._doRemoveRanges(),u.a.safari){var o=this._lastInitSegments[t];o&&(this._pendingSegments[t].push(o),i.updating||this._doAppendSegments())}}}},e.prototype.endOfStream=function(){var e=this._mediaSource,t=this._sourceBuffers;e&&"open"===e.readyState?t.video&&t.video.updating||t.audio&&t.audio.updating?this._hasPendingEos=!0:(this._hasPendingEos=!1,e.endOfStream()):e&&"closed"===e.readyState&&this._hasPendingSegments()&&(this._hasPendingEos=!0)},e.prototype.getNearestKeyframe=function(e){return this._idrList.getLastSyncPointBeforeDts(e)},e.prototype._needCleanupSourceBuffer=function(){if(!this._config.autoCleanupSourceBuffer)return!1;var e=this._mediaElement.currentTime;for(var t in this._sourceBuffers){var i=this._sourceBuffers[t];if(i){var n=i.buffered;if(n.length>=1&&e-n.start(0)>=this._config.autoCleanupMaxBackwardDuration)return!0}}return!1},e.prototype._doCleanupSourceBuffer=function(){var e=this._mediaElement.currentTime;for(var t in this._sourceBuffers){var i=this._sourceBuffers[t];if(i){for(var n=i.buffered,a=!1,r=0;r<n.length;r++){var s=n.start(r),o=n.end(r);if(s<=e&&e<o+3){if(e-s>=this._config.autoCleanupMaxBackwardDuration){a=!0;var d=e-this._config.autoCleanupMinBackwardDuration;this._pendingRemoveRanges[t].push({start:s,end:d})}}else o<e&&(a=!0,this._pendingRemoveRanges[t].push({start:s,end:o}))}a&&!i.updating&&this._doRemoveRanges()}}},e.prototype._updateMediaSourceDuration=function(){var e=this._sourceBuffers;if(0!==this._mediaElement.readyState&&"open"===this._mediaSource.readyState&&!(e.video&&e.video.updating||e.audio&&e.audio.updating)){var t=this._mediaSource.duration,i=this._pendingMediaDuration;i>0&&(isNaN(t)||i>t)&&(c.a.v(this.TAG,"Update MediaSource duration from "+t+" to "+i),this._mediaSource.duration=i),this._requireSetMediaDuration=!1,this._pendingMediaDuration=0}},e.prototype._doRemoveRanges=function(){for(var e in this._pendingRemoveRanges)if(this._sourceBuffers[e]&&!this._sourceBuffers[e].updating)for(var t=this._sourceBuffers[e],i=this._pendingRemoveRanges[e];i.length&&!t.updating;){var n=i.shift();t.remove(n.start,n.end)}},e.prototype._doAppendSegments=function(){var e=this._pendingSegments;for(var t in e)if(this._sourceBuffers[t]&&!this._sourceBuffers[t].updating&&e[t].length>0){var i=e[t].shift();if(i.timestampOffset){var n=this._sourceBuffers[t].timestampOffset,a=i.timestampOffset/1e3;Math.abs(n-a)>.1&&(c.a.v(this.TAG,"Update MPEG audio timestampOffset from "+n+" to "+a),this._sourceBuffers[t].timestampOffset=a),delete i.timestampOffset}if(!i.data||0===i.data.byteLength)continue;try{this._sourceBuffers[t].appendBuffer(i.data),this._isBufferFull=!1,"video"===t&&i.hasOwnProperty("info")&&this._idrList.appendArray(i.info.syncPoints)}catch(e){this._pendingSegments[t].unshift(i),22===e.code?(this._isBufferFull||this._emitter.emit(S.BUFFER_FULL),this._isBufferFull=!0):(c.a.e(this.TAG,e.message),this._emitter.emit(S.ERROR,{code:e.code,msg:e.message}))}}},e.prototype._onSourceOpen=function(){if(c.a.v(this.TAG,"MediaSource onSourceOpen"),this._mediaSource.removeEventListener("sourceopen",this.e.onSourceOpen),this._pendingSourceBufferInit.length>0)for(var e=this._pendingSourceBufferInit;e.length;){var t=e.shift();this.appendInitSegment(t,!0)}this._hasPendingSegments()&&this._doAppendSegments(),this._emitter.emit(S.SOURCE_OPEN)},e.prototype._onSourceEnded=function(){c.a.v(this.TAG,"MediaSource onSourceEnded")},e.prototype._onSourceClose=function(){c.a.v(this.TAG,"MediaSource onSourceClose"),this._mediaSource&&null!=this.e&&(this._mediaSource.removeEventListener("sourceopen",this.e.onSourceOpen),this._mediaSource.removeEventListener("sourceended",this.e.onSourceEnded),this._mediaSource.removeEventListener("sourceclose",this.e.onSourceClose))},e.prototype._hasPendingSegments=function(){var e=this._pendingSegments;return e.video.length>0||e.audio.length>0},e.prototype._hasPendingRemoveRanges=function(){var e=this._pendingRemoveRanges;return e.video.length>0||e.audio.length>0},e.prototype._onSourceBufferUpdateEnd=function(){this._requireSetMediaDuration?this._updateMediaSourceDuration():this._hasPendingRemoveRanges()?this._doRemoveRanges():this._hasPendingSegments()?this._doAppendSegments():this._hasPendingEos&&this.endOfStream(),this._emitter.emit(S.UPDATE_END)},e.prototype._onSourceBufferError=function(e){c.a.e(this.TAG,"SourceBuffer Error: "+e)},e}(),T=i(5),L={NETWORK_ERROR:"NetworkError",MEDIA_ERROR:"MediaError",OTHER_ERROR:"OtherError"},w={NETWORK_EXCEPTION:d.b.EXCEPTION,NETWORK_STATUS_CODE_INVALID:d.b.HTTP_STATUS_CODE_INVALID,NETWORK_TIMEOUT:d.b.CONNECTING_TIMEOUT,NETWORK_UNRECOVERABLE_EARLY_EOF:d.b.UNRECOVERABLE_EARLY_EOF,MEDIA_MSE_ERROR:"MediaMSEError",MEDIA_FORMAT_ERROR:T.a.FORMAT_ERROR,MEDIA_FORMAT_UNSUPPORTED:T.a.FORMAT_UNSUPPORTED,MEDIA_CODEC_UNSUPPORTED:T.a.CODEC_UNSUPPORTED},k=function(){function e(e,t){this.TAG="MSEPlayer",this._type="MSEPlayer",this._emitter=new h.a,this._config=s(),"object"==typeof t&&Object.assign(this._config,t);var i=e.type.toLowerCase();if("mse"!==i&&"mpegts"!==i&&"m2ts"!==i&&"flv"!==i)throw new A.b("MSEPlayer requires an mpegts/m2ts/flv MediaDataSource input!");!0===e.isLive&&(this._config.isLive=!0),this.e={onvLoadedMetadata:this._onvLoadedMetadata.bind(this),onvSeeking:this._onvSeeking.bind(this),onvCanPlay:this._onvCanPlay.bind(this),onvStalled:this._onvStalled.bind(this),onvProgress:this._onvProgress.bind(this)},self.performance&&self.performance.now?this._now=self.performance.now.bind(self.performance):this._now=Date.now,this._pendingSeekTime=null,this._requestSetTime=!1,this._seekpointRecord=null,this._progressChecker=null,this._mediaDataSource=e,this._mediaElement=null,this._msectl=null,this._transmuxer=null,this._mseSourceOpened=!1,this._hasPendingLoad=!1,this._receivedCanPlay=!1,this._mediaInfo=null,this._statisticsInfo=null;var n=u.a.chrome&&(u.a.version.major<50||50===u.a.version.major&&u.a.version.build<2661);this._alwaysSeekKeyframe=!!(n||u.a.msedge||u.a.msie),this._alwaysSeekKeyframe&&(this._config.accurateSeek=!1)}return e.prototype.destroy=function(){null!=this._progressChecker&&(window.clearInterval(this._progressChecker),this._progressChecker=null),this._transmuxer&&this.unload(),this._mediaElement&&this.detachMediaElement(),this.e=null,this._mediaDataSource=null,this._emitter.removeAllListeners(),this._emitter=null},e.prototype.on=function(e,t){var i=this;e===l.MEDIA_INFO?null!=this._mediaInfo&&Promise.resolve().then((function(){i._emitter.emit(l.MEDIA_INFO,i.mediaInfo)})):e===l.STATISTICS_INFO&&null!=this._statisticsInfo&&Promise.resolve().then((function(){i._emitter.emit(l.STATISTICS_INFO,i.statisticsInfo)})),this._emitter.addListener(e,t)},e.prototype.off=function(e,t){this._emitter.removeListener(e,t)},e.prototype.attachMediaElement=function(e){var t=this;if(this._mediaElement=e,e.addEventListener("loadedmetadata",this.e.onvLoadedMetadata),e.addEventListener("seeking",this.e.onvSeeking),e.addEventListener("canplay",this.e.onvCanPlay),e.addEventListener("stalled",this.e.onvStalled),e.addEventListener("progress",this.e.onvProgress),this._msectl=new R(this._config),this._msectl.on(S.UPDATE_END,this._onmseUpdateEnd.bind(this)),this._msectl.on(S.BUFFER_FULL,this._onmseBufferFull.bind(this)),this._msectl.on(S.SOURCE_OPEN,(function(){t._mseSourceOpened=!0,t._hasPendingLoad&&(t._hasPendingLoad=!1,t.load())})),this._msectl.on(S.ERROR,(function(e){t._emitter.emit(l.ERROR,L.MEDIA_ERROR,w.MEDIA_MSE_ERROR,e)})),this._msectl.attachMediaElement(e),null!=this._pendingSeekTime)try{e.currentTime=this._pendingSeekTime,this._pendingSeekTime=null}catch(e){}},e.prototype.detachMediaElement=function(){this._mediaElement&&(this._msectl.detachMediaElement(),this._mediaElement.removeEventListener("loadedmetadata",this.e.onvLoadedMetadata),this._mediaElement.removeEventListener("seeking",this.e.onvSeeking),this._mediaElement.removeEventListener("canplay",this.e.onvCanPlay),this._mediaElement.removeEventListener("stalled",this.e.onvStalled),this._mediaElement.removeEventListener("progress",this.e.onvProgress),this._mediaElement=null),this._msectl&&(this._msectl.destroy(),this._msectl=null)},e.prototype.load=function(){var e=this;if(!this._mediaElement)throw new A.a("HTMLMediaElement must be attached before load()!");if(this._transmuxer)throw new A.a("MSEPlayer.load() has been called, please call unload() first!");this._hasPendingLoad||(this._config.deferLoadAfterSourceOpen&&!1===this._mseSourceOpened?this._hasPendingLoad=!0:(this._mediaElement.readyState>0&&(this._requestSetTime=!0,this._mediaElement.currentTime=0),this._transmuxer=new b(this._mediaDataSource,this._config),this._transmuxer.on(v.a.INIT_SEGMENT,(function(t,i){e._msectl.appendInitSegment(i)})),this._transmuxer.on(v.a.MEDIA_SEGMENT,(function(t,i){if(e._msectl.appendMediaSegment(i),e._config.lazyLoad&&!e._config.isLive){var n=e._mediaElement.currentTime;i.info.endDts>=1e3*(n+e._config.lazyLoadMaxDuration)&&null==e._progressChecker&&(c.a.v(e.TAG,"Maximum buffering duration exceeded, suspend transmuxing task"),e._suspendTransmuxer())}})),this._transmuxer.on(v.a.LOADING_COMPLETE,(function(){e._msectl.endOfStream(),e._emitter.emit(l.LOADING_COMPLETE)})),this._transmuxer.on(v.a.RECOVERED_EARLY_EOF,(function(){e._emitter.emit(l.RECOVERED_EARLY_EOF)})),this._transmuxer.on(v.a.IO_ERROR,(function(t,i){e._emitter.emit(l.ERROR,L.NETWORK_ERROR,t,i)})),this._transmuxer.on(v.a.DEMUX_ERROR,(function(t,i){e._emitter.emit(l.ERROR,L.MEDIA_ERROR,t,{code:-1,msg:i})})),this._transmuxer.on(v.a.MEDIA_INFO,(function(t){e._mediaInfo=t,e._emitter.emit(l.MEDIA_INFO,Object.assign({},t))})),this._transmuxer.on(v.a.METADATA_ARRIVED,(function(t){e._emitter.emit(l.METADATA_ARRIVED,t)})),this._transmuxer.on(v.a.SCRIPTDATA_ARRIVED,(function(t){e._emitter.emit(l.SCRIPTDATA_ARRIVED,t)})),this._transmuxer.on(v.a.TIMED_ID3_METADATA_ARRIVED,(function(t){e._emitter.emit(l.TIMED_ID3_METADATA_ARRIVED,t)})),this._transmuxer.on(v.a.SMPTE2038_METADATA_ARRIVED,(function(t){e._emitter.emit(l.SMPTE2038_METADATA_ARRIVED,t)})),this._transmuxer.on(v.a.SCTE35_METADATA_ARRIVED,(function(t){e._emitter.emit(l.SCTE35_METADATA_ARRIVED,t)})),this._transmuxer.on(v.a.PES_PRIVATE_DATA_DESCRIPTOR,(function(t){e._emitter.emit(l.PES_PRIVATE_DATA_DESCRIPTOR,t)})),this._transmuxer.on(v.a.PES_PRIVATE_DATA_ARRIVED,(function(t){e._emitter.emit(l.PES_PRIVATE_DATA_ARRIVED,t)})),this._transmuxer.on(v.a.STATISTICS_INFO,(function(t){e._statisticsInfo=e._fillStatisticsInfo(t),e._emitter.emit(l.STATISTICS_INFO,Object.assign({},e._statisticsInfo))})),this._transmuxer.on(v.a.RECOMMEND_SEEKPOINT,(function(t){e._mediaElement&&!e._config.accurateSeek&&(e._requestSetTime=!0,e._mediaElement.currentTime=t/1e3)})),this._transmuxer.open()))},e.prototype.unload=function(){this._mediaElement&&this._mediaElement.pause(),this._msectl&&this._msectl.seek(0),this._transmuxer&&(this._transmuxer.close(),this._transmuxer.destroy(),this._transmuxer=null)},e.prototype.play=function(){return this._mediaElement.play()},e.prototype.pause=function(){this._mediaElement.pause()},Object.defineProperty(e.prototype,"type",{get:function(){return this._type},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"buffered",{get:function(){return this._mediaElement.buffered},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"duration",{get:function(){return this._mediaElement.duration},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"volume",{get:function(){return this._mediaElement.volume},set:function(e){this._mediaElement.volume=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"muted",{get:function(){return this._mediaElement.muted},set:function(e){this._mediaElement.muted=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"currentTime",{get:function(){return this._mediaElement?this._mediaElement.currentTime:0},set:function(e){this._mediaElement?this._internalSeek(e):this._pendingSeekTime=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"mediaInfo",{get:function(){return Object.assign({},this._mediaInfo)},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"statisticsInfo",{get:function(){return null==this._statisticsInfo&&(this._statisticsInfo={}),this._statisticsInfo=this._fillStatisticsInfo(this._statisticsInfo),Object.assign({},this._statisticsInfo)},enumerable:!1,configurable:!0}),e.prototype._fillStatisticsInfo=function(e){if(e.playerType=this._type,!(this._mediaElement instanceof HTMLVideoElement))return e;var t=!0,i=0,n=0;if(this._mediaElement.getVideoPlaybackQuality){var a=this._mediaElement.getVideoPlaybackQuality();i=a.totalVideoFrames,n=a.droppedVideoFrames}else null!=this._mediaElement.webkitDecodedFrameCount?(i=this._mediaElement.webkitDecodedFrameCount,n=this._mediaElement.webkitDroppedFrameCount):t=!1;return t&&(e.decodedFrames=i,e.droppedFrames=n),e},e.prototype._onmseUpdateEnd=function(){var e=this._mediaElement.buffered,t=this._mediaElement.currentTime;if(this._config.isLive&&this._config.liveBufferLatencyChasing&&e.length>0&&!this._mediaElement.paused){var i=e.end(e.length-1);if(i>this._config.liveBufferLatencyMaxLatency&&i-t>this._config.liveBufferLatencyMaxLatency){var n=i-this._config.liveBufferLatencyMinRemain;this.currentTime=n}}if(this._config.lazyLoad&&!this._config.isLive){for(var a=0,r=0;r<e.length;r++){var s=e.start(r),o=e.end(r);if(s<=t&&t<o){s,a=o;break}}a>=t+this._config.lazyLoadMaxDuration&&null==this._progressChecker&&(c.a.v(this.TAG,"Maximum buffering duration exceeded, suspend transmuxing task"),this._suspendTransmuxer())}},e.prototype._onmseBufferFull=function(){c.a.v(this.TAG,"MSE SourceBuffer is full, suspend transmuxing task"),null==this._progressChecker&&this._suspendTransmuxer()},e.prototype._suspendTransmuxer=function(){this._transmuxer&&(this._transmuxer.pause(),null==this._progressChecker&&(this._progressChecker=window.setInterval(this._checkProgressAndResume.bind(this),1e3)))},e.prototype._checkProgressAndResume=function(){for(var e=this._mediaElement.currentTime,t=this._mediaElement.buffered,i=!1,n=0;n<t.length;n++){var a=t.start(n),r=t.end(n);if(e>=a&&e<r){e>=r-this._config.lazyLoadRecoverDuration&&(i=!0);break}}i&&(window.clearInterval(this._progressChecker),this._progressChecker=null,i&&(c.a.v(this.TAG,"Continue loading from paused position"),this._transmuxer.resume()))},e.prototype._isTimepointBuffered=function(e){for(var t=this._mediaElement.buffered,i=0;i<t.length;i++){var n=t.start(i),a=t.end(i);if(e>=n&&e<a)return!0}return!1},e.prototype._internalSeek=function(e){var t=this._isTimepointBuffered(e),i=!1,n=0;if(e<1&&this._mediaElement.buffered.length>0){var a=this._mediaElement.buffered.start(0);(a<1&&e<a||u.a.safari)&&(i=!0,n=u.a.safari?.1:a)}if(i)this._requestSetTime=!0,this._mediaElement.currentTime=n;else if(t){if(this._alwaysSeekKeyframe){var r=this._msectl.getNearestKeyframe(Math.floor(1e3*e));this._requestSetTime=!0,this._mediaElement.currentTime=null!=r?r.dts/1e3:e}else this._requestSetTime=!0,this._mediaElement.currentTime=e;null!=this._progressChecker&&this._checkProgressAndResume()}else null!=this._progressChecker&&(window.clearInterval(this._progressChecker),this._progressChecker=null),this._msectl.seek(e),this._transmuxer.seek(Math.floor(1e3*e)),this._config.accurateSeek&&(this._requestSetTime=!0,this._mediaElement.currentTime=e)},e.prototype._checkAndApplyUnbufferedSeekpoint=function(){if(this._seekpointRecord)if(this._seekpointRecord.recordTime<=this._now()-100){var e=this._mediaElement.currentTime;this._seekpointRecord=null,this._isTimepointBuffered(e)||(null!=this._progressChecker&&(window.clearTimeout(this._progressChecker),this._progressChecker=null),this._msectl.seek(e),this._transmuxer.seek(Math.floor(1e3*e)),this._config.accurateSeek&&(this._requestSetTime=!0,this._mediaElement.currentTime=e))}else window.setTimeout(this._checkAndApplyUnbufferedSeekpoint.bind(this),50)},e.prototype._checkAndResumeStuckPlayback=function(e){var t=this._mediaElement;if(e||!this._receivedCanPlay||t.readyState<2){var i=t.buffered;i.length>0&&t.currentTime<i.start(0)&&(c.a.w(this.TAG,"Playback seems stuck at "+t.currentTime+", seek to "+i.start(0)),this._requestSetTime=!0,this._mediaElement.currentTime=i.start(0),this._mediaElement.removeEventListener("progress",this.e.onvProgress))}else this._mediaElement.removeEventListener("progress",this.e.onvProgress)},e.prototype._onvLoadedMetadata=function(e){null!=this._pendingSeekTime&&(this._mediaElement.currentTime=this._pendingSeekTime,this._pendingSeekTime=null)},e.prototype._onvSeeking=function(e){var t=this._mediaElement.currentTime,i=this._mediaElement.buffered;if(this._requestSetTime)this._requestSetTime=!1;else{if(t<1&&i.length>0){var n=i.start(0);if(n<1&&t<n||u.a.safari)return this._requestSetTime=!0,void(this._mediaElement.currentTime=u.a.safari?.1:n)}if(this._isTimepointBuffered(t)){if(this._alwaysSeekKeyframe){var a=this._msectl.getNearestKeyframe(Math.floor(1e3*t));null!=a&&(this._requestSetTime=!0,this._mediaElement.currentTime=a.dts/1e3)}null!=this._progressChecker&&this._checkProgressAndResume()}else this._seekpointRecord={seekPoint:t,recordTime:this._now()},window.setTimeout(this._checkAndApplyUnbufferedSeekpoint.bind(this),50)}},e.prototype._onvCanPlay=function(e){this._receivedCanPlay=!0,this._mediaElement.removeEventListener("canplay",this.e.onvCanPlay)},e.prototype._onvStalled=function(e){this._checkAndResumeStuckPlayback(!0)},e.prototype._onvProgress=function(e){this._checkAndResumeStuckPlayback()},e}(),D=function(){function e(e,t){this.TAG="NativePlayer",this._type="NativePlayer",this._emitter=new h.a,this._config=s(),"object"==typeof t&&Object.assign(this._config,t);var i=e.type.toLowerCase();if("mse"===i||"mpegts"===i||"m2ts"===i||"flv"===i)throw new A.b("NativePlayer does't support mse/mpegts/m2ts/flv MediaDataSource input!");if(e.hasOwnProperty("segments"))throw new A.b("NativePlayer("+e.type+") doesn't support multipart playback!");this.e={onvLoadedMetadata:this._onvLoadedMetadata.bind(this)},this._pendingSeekTime=null,this._statisticsReporter=null,this._mediaDataSource=e,this._mediaElement=null}return e.prototype.destroy=function(){this._mediaElement&&(this.unload(),this.detachMediaElement()),this.e=null,this._mediaDataSource=null,this._emitter.removeAllListeners(),this._emitter=null},e.prototype.on=function(e,t){var i=this;e===l.MEDIA_INFO?null!=this._mediaElement&&0!==this._mediaElement.readyState&&Promise.resolve().then((function(){i._emitter.emit(l.MEDIA_INFO,i.mediaInfo)})):e===l.STATISTICS_INFO&&null!=this._mediaElement&&0!==this._mediaElement.readyState&&Promise.resolve().then((function(){i._emitter.emit(l.STATISTICS_INFO,i.statisticsInfo)})),this._emitter.addListener(e,t)},e.prototype.off=function(e,t){this._emitter.removeListener(e,t)},e.prototype.attachMediaElement=function(e){if(this._mediaElement=e,e.addEventListener("loadedmetadata",this.e.onvLoadedMetadata),null!=this._pendingSeekTime)try{e.currentTime=this._pendingSeekTime,this._pendingSeekTime=null}catch(e){}},e.prototype.detachMediaElement=function(){this._mediaElement&&(this._mediaElement.src="",this._mediaElement.removeAttribute("src"),this._mediaElement.removeEventListener("loadedmetadata",this.e.onvLoadedMetadata),this._mediaElement=null),null!=this._statisticsReporter&&(window.clearInterval(this._statisticsReporter),this._statisticsReporter=null)},e.prototype.load=function(){if(!this._mediaElement)throw new A.a("HTMLMediaElement must be attached before load()!");this._mediaElement.src=this._mediaDataSource.url,this._mediaElement.readyState>0&&(this._mediaElement.currentTime=0),this._mediaElement.preload="auto",this._mediaElement.load(),this._statisticsReporter=window.setInterval(this._reportStatisticsInfo.bind(this),this._config.statisticsInfoReportInterval)},e.prototype.unload=function(){this._mediaElement&&(this._mediaElement.src="",this._mediaElement.removeAttribute("src")),null!=this._statisticsReporter&&(window.clearInterval(this._statisticsReporter),this._statisticsReporter=null)},e.prototype.play=function(){return this._mediaElement.play()},e.prototype.pause=function(){this._mediaElement.pause()},Object.defineProperty(e.prototype,"type",{get:function(){return this._type},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"buffered",{get:function(){return this._mediaElement.buffered},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"duration",{get:function(){return this._mediaElement.duration},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"volume",{get:function(){return this._mediaElement.volume},set:function(e){this._mediaElement.volume=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"muted",{get:function(){return this._mediaElement.muted},set:function(e){this._mediaElement.muted=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"currentTime",{get:function(){return this._mediaElement?this._mediaElement.currentTime:0},set:function(e){this._mediaElement?this._mediaElement.currentTime=e:this._pendingSeekTime=e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"mediaInfo",{get:function(){var e={mimeType:(this._mediaElement instanceof HTMLAudioElement?"audio/":"video/")+this._mediaDataSource.type};return this._mediaElement&&(e.duration=Math.floor(1e3*this._mediaElement.duration),this._mediaElement instanceof HTMLVideoElement&&(e.width=this._mediaElement.videoWidth,e.height=this._mediaElement.videoHeight)),e},enumerable:!1,configurable:!0}),Object.defineProperty(e.prototype,"statisticsInfo",{get:function(){var e={playerType:this._type,url:this._mediaDataSource.url};if(!(this._mediaElement instanceof HTMLVideoElement))return e;var t=!0,i=0,n=0;if(this._mediaElement.getVideoPlaybackQuality){var a=this._mediaElement.getVideoPlaybackQuality();i=a.totalVideoFrames,n=a.droppedVideoFrames}else null!=this._mediaElement.webkitDecodedFrameCount?(i=this._mediaElement.webkitDecodedFrameCount,n=this._mediaElement.webkitDroppedFrameCount):t=!1;return t&&(e.decodedFrames=i,e.droppedFrames=n),e},enumerable:!1,configurable:!0}),e.prototype._onvLoadedMetadata=function(e){null!=this._pendingSeekTime&&(this._mediaElement.currentTime=this._pendingSeekTime,this._pendingSeekTime=null),this._emitter.emit(l.MEDIA_INFO,this.mediaInfo)},e.prototype._reportStatisticsInfo=function(){this._emitter.emit(l.STATISTICS_INFO,this.statisticsInfo)},e}();n.a.install();var C={createPlayer:function(e,t){var i=e;if(null==i||"object"!=typeof i)throw new A.b("MediaDataSource must be an javascript object!");if(!i.hasOwnProperty("type"))throw new A.b("MediaDataSource must has type field to indicate video file type!");switch(i.type){case"mse":case"mpegts":case"m2ts":case"flv":return new k(i,t);default:return new D(i,t)}},isSupported:function(){return o.supportMSEH264Playback()},getFeatureList:function(){return o.getFeatureList()}};C.BaseLoader=d.a,C.LoaderStatus=d.c,C.LoaderErrors=d.b,C.Events=l,C.ErrorTypes=L,C.ErrorDetails=w,C.MSEPlayer=k,C.NativePlayer=D,C.LoggingControl=m.a,Object.defineProperty(C,"version",{enumerable:!0,get:function(){return"1.7.3"}});t.default=C}])}));
//# sourceMappingURL=mpegts.js.map

/***/ }),

/***/ "./node_modules/object-inspect/index.js":
/*!**********************************************!*\
  !*** ./node_modules/object-inspect/index.js ***!
  \**********************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

var hasMap = typeof Map === 'function' && Map.prototype;
var mapSizeDescriptor = Object.getOwnPropertyDescriptor && hasMap ? Object.getOwnPropertyDescriptor(Map.prototype, 'size') : null;
var mapSize = hasMap && mapSizeDescriptor && typeof mapSizeDescriptor.get === 'function' ? mapSizeDescriptor.get : null;
var mapForEach = hasMap && Map.prototype.forEach;
var hasSet = typeof Set === 'function' && Set.prototype;
var setSizeDescriptor = Object.getOwnPropertyDescriptor && hasSet ? Object.getOwnPropertyDescriptor(Set.prototype, 'size') : null;
var setSize = hasSet && setSizeDescriptor && typeof setSizeDescriptor.get === 'function' ? setSizeDescriptor.get : null;
var setForEach = hasSet && Set.prototype.forEach;
var hasWeakMap = typeof WeakMap === 'function' && WeakMap.prototype;
var weakMapHas = hasWeakMap ? WeakMap.prototype.has : null;
var hasWeakSet = typeof WeakSet === 'function' && WeakSet.prototype;
var weakSetHas = hasWeakSet ? WeakSet.prototype.has : null;
var hasWeakRef = typeof WeakRef === 'function' && WeakRef.prototype;
var weakRefDeref = hasWeakRef ? WeakRef.prototype.deref : null;
var booleanValueOf = Boolean.prototype.valueOf;
var objectToString = Object.prototype.toString;
var functionToString = Function.prototype.toString;
var $match = String.prototype.match;
var $slice = String.prototype.slice;
var $replace = String.prototype.replace;
var $toUpperCase = String.prototype.toUpperCase;
var $toLowerCase = String.prototype.toLowerCase;
var $test = RegExp.prototype.test;
var $concat = Array.prototype.concat;
var $join = Array.prototype.join;
var $arrSlice = Array.prototype.slice;
var $floor = Math.floor;
var bigIntValueOf = typeof BigInt === 'function' ? BigInt.prototype.valueOf : null;
var gOPS = Object.getOwnPropertySymbols;
var symToString = typeof Symbol === 'function' && typeof Symbol.iterator === 'symbol' ? Symbol.prototype.toString : null;
var hasShammedSymbols = typeof Symbol === 'function' && typeof Symbol.iterator === 'object';
// ie, `has-tostringtag/shams
var toStringTag = typeof Symbol === 'function' && Symbol.toStringTag && (typeof Symbol.toStringTag === hasShammedSymbols ? 'object' : 'symbol')
    ? Symbol.toStringTag
    : null;
var isEnumerable = Object.prototype.propertyIsEnumerable;

var gPO = (typeof Reflect === 'function' ? Reflect.getPrototypeOf : Object.getPrototypeOf) || (
    [].__proto__ === Array.prototype // eslint-disable-line no-proto
        ? function (O) {
            return O.__proto__; // eslint-disable-line no-proto
        }
        : null
);

function addNumericSeparator(num, str) {
    if (
        num === Infinity
        || num === -Infinity
        || num !== num
        || (num && num > -1000 && num < 1000)
        || $test.call(/e/, str)
    ) {
        return str;
    }
    var sepRegex = /[0-9](?=(?:[0-9]{3})+(?![0-9]))/g;
    if (typeof num === 'number') {
        var int = num < 0 ? -$floor(-num) : $floor(num); // trunc(num)
        if (int !== num) {
            var intStr = String(int);
            var dec = $slice.call(str, intStr.length + 1);
            return $replace.call(intStr, sepRegex, '$&_') + '.' + $replace.call($replace.call(dec, /([0-9]{3})/g, '$&_'), /_$/, '');
        }
    }
    return $replace.call(str, sepRegex, '$&_');
}

var utilInspect = __webpack_require__(/*! ./util.inspect */ "?4f7e");
var inspectCustom = utilInspect.custom;
var inspectSymbol = isSymbol(inspectCustom) ? inspectCustom : null;

module.exports = function inspect_(obj, options, depth, seen) {
    var opts = options || {};

    if (has(opts, 'quoteStyle') && (opts.quoteStyle !== 'single' && opts.quoteStyle !== 'double')) {
        throw new TypeError('option "quoteStyle" must be "single" or "double"');
    }
    if (
        has(opts, 'maxStringLength') && (typeof opts.maxStringLength === 'number'
            ? opts.maxStringLength < 0 && opts.maxStringLength !== Infinity
            : opts.maxStringLength !== null
        )
    ) {
        throw new TypeError('option "maxStringLength", if provided, must be a positive integer, Infinity, or `null`');
    }
    var customInspect = has(opts, 'customInspect') ? opts.customInspect : true;
    if (typeof customInspect !== 'boolean' && customInspect !== 'symbol') {
        throw new TypeError('option "customInspect", if provided, must be `true`, `false`, or `\'symbol\'`');
    }

    if (
        has(opts, 'indent')
        && opts.indent !== null
        && opts.indent !== '\t'
        && !(parseInt(opts.indent, 10) === opts.indent && opts.indent > 0)
    ) {
        throw new TypeError('option "indent" must be "\\t", an integer > 0, or `null`');
    }
    if (has(opts, 'numericSeparator') && typeof opts.numericSeparator !== 'boolean') {
        throw new TypeError('option "numericSeparator", if provided, must be `true` or `false`');
    }
    var numericSeparator = opts.numericSeparator;

    if (typeof obj === 'undefined') {
        return 'undefined';
    }
    if (obj === null) {
        return 'null';
    }
    if (typeof obj === 'boolean') {
        return obj ? 'true' : 'false';
    }

    if (typeof obj === 'string') {
        return inspectString(obj, opts);
    }
    if (typeof obj === 'number') {
        if (obj === 0) {
            return Infinity / obj > 0 ? '0' : '-0';
        }
        var str = String(obj);
        return numericSeparator ? addNumericSeparator(obj, str) : str;
    }
    if (typeof obj === 'bigint') {
        var bigIntStr = String(obj) + 'n';
        return numericSeparator ? addNumericSeparator(obj, bigIntStr) : bigIntStr;
    }

    var maxDepth = typeof opts.depth === 'undefined' ? 5 : opts.depth;
    if (typeof depth === 'undefined') { depth = 0; }
    if (depth >= maxDepth && maxDepth > 0 && typeof obj === 'object') {
        return isArray(obj) ? '[Array]' : '[Object]';
    }

    var indent = getIndent(opts, depth);

    if (typeof seen === 'undefined') {
        seen = [];
    } else if (indexOf(seen, obj) >= 0) {
        return '[Circular]';
    }

    function inspect(value, from, noIndent) {
        if (from) {
            seen = $arrSlice.call(seen);
            seen.push(from);
        }
        if (noIndent) {
            var newOpts = {
                depth: opts.depth
            };
            if (has(opts, 'quoteStyle')) {
                newOpts.quoteStyle = opts.quoteStyle;
            }
            return inspect_(value, newOpts, depth + 1, seen);
        }
        return inspect_(value, opts, depth + 1, seen);
    }

    if (typeof obj === 'function' && !isRegExp(obj)) { // in older engines, regexes are callable
        var name = nameOf(obj);
        var keys = arrObjKeys(obj, inspect);
        return '[Function' + (name ? ': ' + name : ' (anonymous)') + ']' + (keys.length > 0 ? ' { ' + $join.call(keys, ', ') + ' }' : '');
    }
    if (isSymbol(obj)) {
        var symString = hasShammedSymbols ? $replace.call(String(obj), /^(Symbol\(.*\))_[^)]*$/, '$1') : symToString.call(obj);
        return typeof obj === 'object' && !hasShammedSymbols ? markBoxed(symString) : symString;
    }
    if (isElement(obj)) {
        var s = '<' + $toLowerCase.call(String(obj.nodeName));
        var attrs = obj.attributes || [];
        for (var i = 0; i < attrs.length; i++) {
            s += ' ' + attrs[i].name + '=' + wrapQuotes(quote(attrs[i].value), 'double', opts);
        }
        s += '>';
        if (obj.childNodes && obj.childNodes.length) { s += '...'; }
        s += '</' + $toLowerCase.call(String(obj.nodeName)) + '>';
        return s;
    }
    if (isArray(obj)) {
        if (obj.length === 0) { return '[]'; }
        var xs = arrObjKeys(obj, inspect);
        if (indent && !singleLineValues(xs)) {
            return '[' + indentedJoin(xs, indent) + ']';
        }
        return '[ ' + $join.call(xs, ', ') + ' ]';
    }
    if (isError(obj)) {
        var parts = arrObjKeys(obj, inspect);
        if (!('cause' in Error.prototype) && 'cause' in obj && !isEnumerable.call(obj, 'cause')) {
            return '{ [' + String(obj) + '] ' + $join.call($concat.call('[cause]: ' + inspect(obj.cause), parts), ', ') + ' }';
        }
        if (parts.length === 0) { return '[' + String(obj) + ']'; }
        return '{ [' + String(obj) + '] ' + $join.call(parts, ', ') + ' }';
    }
    if (typeof obj === 'object' && customInspect) {
        if (inspectSymbol && typeof obj[inspectSymbol] === 'function' && utilInspect) {
            return utilInspect(obj, { depth: maxDepth - depth });
        } else if (customInspect !== 'symbol' && typeof obj.inspect === 'function') {
            return obj.inspect();
        }
    }
    if (isMap(obj)) {
        var mapParts = [];
        if (mapForEach) {
            mapForEach.call(obj, function (value, key) {
                mapParts.push(inspect(key, obj, true) + ' => ' + inspect(value, obj));
            });
        }
        return collectionOf('Map', mapSize.call(obj), mapParts, indent);
    }
    if (isSet(obj)) {
        var setParts = [];
        if (setForEach) {
            setForEach.call(obj, function (value) {
                setParts.push(inspect(value, obj));
            });
        }
        return collectionOf('Set', setSize.call(obj), setParts, indent);
    }
    if (isWeakMap(obj)) {
        return weakCollectionOf('WeakMap');
    }
    if (isWeakSet(obj)) {
        return weakCollectionOf('WeakSet');
    }
    if (isWeakRef(obj)) {
        return weakCollectionOf('WeakRef');
    }
    if (isNumber(obj)) {
        return markBoxed(inspect(Number(obj)));
    }
    if (isBigInt(obj)) {
        return markBoxed(inspect(bigIntValueOf.call(obj)));
    }
    if (isBoolean(obj)) {
        return markBoxed(booleanValueOf.call(obj));
    }
    if (isString(obj)) {
        return markBoxed(inspect(String(obj)));
    }
    // note: in IE 8, sometimes `global !== window` but both are the prototypes of each other
    /* eslint-env browser */
    if (typeof window !== 'undefined' && obj === window) {
        return '{ [object Window] }';
    }
    if (obj === __webpack_require__.g) {
        return '{ [object globalThis] }';
    }
    if (!isDate(obj) && !isRegExp(obj)) {
        var ys = arrObjKeys(obj, inspect);
        var isPlainObject = gPO ? gPO(obj) === Object.prototype : obj instanceof Object || obj.constructor === Object;
        var protoTag = obj instanceof Object ? '' : 'null prototype';
        var stringTag = !isPlainObject && toStringTag && Object(obj) === obj && toStringTag in obj ? $slice.call(toStr(obj), 8, -1) : protoTag ? 'Object' : '';
        var constructorTag = isPlainObject || typeof obj.constructor !== 'function' ? '' : obj.constructor.name ? obj.constructor.name + ' ' : '';
        var tag = constructorTag + (stringTag || protoTag ? '[' + $join.call($concat.call([], stringTag || [], protoTag || []), ': ') + '] ' : '');
        if (ys.length === 0) { return tag + '{}'; }
        if (indent) {
            return tag + '{' + indentedJoin(ys, indent) + '}';
        }
        return tag + '{ ' + $join.call(ys, ', ') + ' }';
    }
    return String(obj);
};

function wrapQuotes(s, defaultStyle, opts) {
    var quoteChar = (opts.quoteStyle || defaultStyle) === 'double' ? '"' : "'";
    return quoteChar + s + quoteChar;
}

function quote(s) {
    return $replace.call(String(s), /"/g, '&quot;');
}

function isArray(obj) { return toStr(obj) === '[object Array]' && (!toStringTag || !(typeof obj === 'object' && toStringTag in obj)); }
function isDate(obj) { return toStr(obj) === '[object Date]' && (!toStringTag || !(typeof obj === 'object' && toStringTag in obj)); }
function isRegExp(obj) { return toStr(obj) === '[object RegExp]' && (!toStringTag || !(typeof obj === 'object' && toStringTag in obj)); }
function isError(obj) { return toStr(obj) === '[object Error]' && (!toStringTag || !(typeof obj === 'object' && toStringTag in obj)); }
function isString(obj) { return toStr(obj) === '[object String]' && (!toStringTag || !(typeof obj === 'object' && toStringTag in obj)); }
function isNumber(obj) { return toStr(obj) === '[object Number]' && (!toStringTag || !(typeof obj === 'object' && toStringTag in obj)); }
function isBoolean(obj) { return toStr(obj) === '[object Boolean]' && (!toStringTag || !(typeof obj === 'object' && toStringTag in obj)); }

// Symbol and BigInt do have Symbol.toStringTag by spec, so that can't be used to eliminate false positives
function isSymbol(obj) {
    if (hasShammedSymbols) {
        return obj && typeof obj === 'object' && obj instanceof Symbol;
    }
    if (typeof obj === 'symbol') {
        return true;
    }
    if (!obj || typeof obj !== 'object' || !symToString) {
        return false;
    }
    try {
        symToString.call(obj);
        return true;
    } catch (e) {}
    return false;
}

function isBigInt(obj) {
    if (!obj || typeof obj !== 'object' || !bigIntValueOf) {
        return false;
    }
    try {
        bigIntValueOf.call(obj);
        return true;
    } catch (e) {}
    return false;
}

var hasOwn = Object.prototype.hasOwnProperty || function (key) { return key in this; };
function has(obj, key) {
    return hasOwn.call(obj, key);
}

function toStr(obj) {
    return objectToString.call(obj);
}

function nameOf(f) {
    if (f.name) { return f.name; }
    var m = $match.call(functionToString.call(f), /^function\s*([\w$]+)/);
    if (m) { return m[1]; }
    return null;
}

function indexOf(xs, x) {
    if (xs.indexOf) { return xs.indexOf(x); }
    for (var i = 0, l = xs.length; i < l; i++) {
        if (xs[i] === x) { return i; }
    }
    return -1;
}

function isMap(x) {
    if (!mapSize || !x || typeof x !== 'object') {
        return false;
    }
    try {
        mapSize.call(x);
        try {
            setSize.call(x);
        } catch (s) {
            return true;
        }
        return x instanceof Map; // core-js workaround, pre-v2.5.0
    } catch (e) {}
    return false;
}

function isWeakMap(x) {
    if (!weakMapHas || !x || typeof x !== 'object') {
        return false;
    }
    try {
        weakMapHas.call(x, weakMapHas);
        try {
            weakSetHas.call(x, weakSetHas);
        } catch (s) {
            return true;
        }
        return x instanceof WeakMap; // core-js workaround, pre-v2.5.0
    } catch (e) {}
    return false;
}

function isWeakRef(x) {
    if (!weakRefDeref || !x || typeof x !== 'object') {
        return false;
    }
    try {
        weakRefDeref.call(x);
        return true;
    } catch (e) {}
    return false;
}

function isSet(x) {
    if (!setSize || !x || typeof x !== 'object') {
        return false;
    }
    try {
        setSize.call(x);
        try {
            mapSize.call(x);
        } catch (m) {
            return true;
        }
        return x instanceof Set; // core-js workaround, pre-v2.5.0
    } catch (e) {}
    return false;
}

function isWeakSet(x) {
    if (!weakSetHas || !x || typeof x !== 'object') {
        return false;
    }
    try {
        weakSetHas.call(x, weakSetHas);
        try {
            weakMapHas.call(x, weakMapHas);
        } catch (s) {
            return true;
        }
        return x instanceof WeakSet; // core-js workaround, pre-v2.5.0
    } catch (e) {}
    return false;
}

function isElement(x) {
    if (!x || typeof x !== 'object') { return false; }
    if (typeof HTMLElement !== 'undefined' && x instanceof HTMLElement) {
        return true;
    }
    return typeof x.nodeName === 'string' && typeof x.getAttribute === 'function';
}

function inspectString(str, opts) {
    if (str.length > opts.maxStringLength) {
        var remaining = str.length - opts.maxStringLength;
        var trailer = '... ' + remaining + ' more character' + (remaining > 1 ? 's' : '');
        return inspectString($slice.call(str, 0, opts.maxStringLength), opts) + trailer;
    }
    // eslint-disable-next-line no-control-regex
    var s = $replace.call($replace.call(str, /(['\\])/g, '\\$1'), /[\x00-\x1f]/g, lowbyte);
    return wrapQuotes(s, 'single', opts);
}

function lowbyte(c) {
    var n = c.charCodeAt(0);
    var x = {
        8: 'b',
        9: 't',
        10: 'n',
        12: 'f',
        13: 'r'
    }[n];
    if (x) { return '\\' + x; }
    return '\\x' + (n < 0x10 ? '0' : '') + $toUpperCase.call(n.toString(16));
}

function markBoxed(str) {
    return 'Object(' + str + ')';
}

function weakCollectionOf(type) {
    return type + ' { ? }';
}

function collectionOf(type, size, entries, indent) {
    var joinedEntries = indent ? indentedJoin(entries, indent) : $join.call(entries, ', ');
    return type + ' (' + size + ') {' + joinedEntries + '}';
}

function singleLineValues(xs) {
    for (var i = 0; i < xs.length; i++) {
        if (indexOf(xs[i], '\n') >= 0) {
            return false;
        }
    }
    return true;
}

function getIndent(opts, depth) {
    var baseIndent;
    if (opts.indent === '\t') {
        baseIndent = '\t';
    } else if (typeof opts.indent === 'number' && opts.indent > 0) {
        baseIndent = $join.call(Array(opts.indent + 1), ' ');
    } else {
        return null;
    }
    return {
        base: baseIndent,
        prev: $join.call(Array(depth + 1), baseIndent)
    };
}

function indentedJoin(xs, indent) {
    if (xs.length === 0) { return ''; }
    var lineJoiner = '\n' + indent.prev + indent.base;
    return lineJoiner + $join.call(xs, ',' + lineJoiner) + '\n' + indent.prev;
}

function arrObjKeys(obj, inspect) {
    var isArr = isArray(obj);
    var xs = [];
    if (isArr) {
        xs.length = obj.length;
        for (var i = 0; i < obj.length; i++) {
            xs[i] = has(obj, i) ? inspect(obj[i], obj) : '';
        }
    }
    var syms = typeof gOPS === 'function' ? gOPS(obj) : [];
    var symMap;
    if (hasShammedSymbols) {
        symMap = {};
        for (var k = 0; k < syms.length; k++) {
            symMap['$' + syms[k]] = syms[k];
        }
    }

    for (var key in obj) { // eslint-disable-line no-restricted-syntax
        if (!has(obj, key)) { continue; } // eslint-disable-line no-restricted-syntax, no-continue
        if (isArr && String(Number(key)) === key && key < obj.length) { continue; } // eslint-disable-line no-restricted-syntax, no-continue
        if (hasShammedSymbols && symMap['$' + key] instanceof Symbol) {
            // this is to prevent shammed Symbols, which are stored as strings, from being included in the string key section
            continue; // eslint-disable-line no-restricted-syntax, no-continue
        } else if ($test.call(/[^\w$]/, key)) {
            xs.push(inspect(key, obj) + ': ' + inspect(obj[key], obj));
        } else {
            xs.push(key + ': ' + inspect(obj[key], obj));
        }
    }
    if (typeof gOPS === 'function') {
        for (var j = 0; j < syms.length; j++) {
            if (isEnumerable.call(obj, syms[j])) {
                xs.push('[' + inspect(syms[j]) + ']: ' + inspect(obj[syms[j]], obj));
            }
        }
    }
    return xs;
}


/***/ }),

/***/ "./node_modules/process/browser.js":
/*!*****************************************!*\
  !*** ./node_modules/process/browser.js ***!
  \*****************************************/
/***/ ((module) => {

// shim for using process in browser
var process = module.exports = {};

// cached from whatever global is present so that test runners that stub it
// don't break things.  But we need to wrap it in a try catch in case it is
// wrapped in strict mode code which doesn't define any globals.  It's inside a
// function because try/catches deoptimize in certain engines.

var cachedSetTimeout;
var cachedClearTimeout;

function defaultSetTimout() {
    throw new Error('setTimeout has not been defined');
}
function defaultClearTimeout () {
    throw new Error('clearTimeout has not been defined');
}
(function () {
    try {
        if (typeof setTimeout === 'function') {
            cachedSetTimeout = setTimeout;
        } else {
            cachedSetTimeout = defaultSetTimout;
        }
    } catch (e) {
        cachedSetTimeout = defaultSetTimout;
    }
    try {
        if (typeof clearTimeout === 'function') {
            cachedClearTimeout = clearTimeout;
        } else {
            cachedClearTimeout = defaultClearTimeout;
        }
    } catch (e) {
        cachedClearTimeout = defaultClearTimeout;
    }
} ())
function runTimeout(fun) {
    if (cachedSetTimeout === setTimeout) {
        //normal enviroments in sane situations
        return setTimeout(fun, 0);
    }
    // if setTimeout wasn't available but was latter defined
    if ((cachedSetTimeout === defaultSetTimout || !cachedSetTimeout) && setTimeout) {
        cachedSetTimeout = setTimeout;
        return setTimeout(fun, 0);
    }
    try {
        // when when somebody has screwed with setTimeout but no I.E. maddness
        return cachedSetTimeout(fun, 0);
    } catch(e){
        try {
            // When we are in I.E. but the script has been evaled so I.E. doesn't trust the global object when called normally
            return cachedSetTimeout.call(null, fun, 0);
        } catch(e){
            // same as above but when it's a version of I.E. that must have the global object for 'this', hopfully our context correct otherwise it will throw a global error
            return cachedSetTimeout.call(this, fun, 0);
        }
    }


}
function runClearTimeout(marker) {
    if (cachedClearTimeout === clearTimeout) {
        //normal enviroments in sane situations
        return clearTimeout(marker);
    }
    // if clearTimeout wasn't available but was latter defined
    if ((cachedClearTimeout === defaultClearTimeout || !cachedClearTimeout) && clearTimeout) {
        cachedClearTimeout = clearTimeout;
        return clearTimeout(marker);
    }
    try {
        // when when somebody has screwed with setTimeout but no I.E. maddness
        return cachedClearTimeout(marker);
    } catch (e){
        try {
            // When we are in I.E. but the script has been evaled so I.E. doesn't  trust the global object when called normally
            return cachedClearTimeout.call(null, marker);
        } catch (e){
            // same as above but when it's a version of I.E. that must have the global object for 'this', hopfully our context correct otherwise it will throw a global error.
            // Some versions of I.E. have different rules for clearTimeout vs setTimeout
            return cachedClearTimeout.call(this, marker);
        }
    }



}
var queue = [];
var draining = false;
var currentQueue;
var queueIndex = -1;

function cleanUpNextTick() {
    if (!draining || !currentQueue) {
        return;
    }
    draining = false;
    if (currentQueue.length) {
        queue = currentQueue.concat(queue);
    } else {
        queueIndex = -1;
    }
    if (queue.length) {
        drainQueue();
    }
}

function drainQueue() {
    if (draining) {
        return;
    }
    var timeout = runTimeout(cleanUpNextTick);
    draining = true;

    var len = queue.length;
    while(len) {
        currentQueue = queue;
        queue = [];
        while (++queueIndex < len) {
            if (currentQueue) {
                currentQueue[queueIndex].run();
            }
        }
        queueIndex = -1;
        len = queue.length;
    }
    currentQueue = null;
    draining = false;
    runClearTimeout(timeout);
}

process.nextTick = function (fun) {
    var args = new Array(arguments.length - 1);
    if (arguments.length > 1) {
        for (var i = 1; i < arguments.length; i++) {
            args[i - 1] = arguments[i];
        }
    }
    queue.push(new Item(fun, args));
    if (queue.length === 1 && !draining) {
        runTimeout(drainQueue);
    }
};

// v8 likes predictible objects
function Item(fun, array) {
    this.fun = fun;
    this.array = array;
}
Item.prototype.run = function () {
    this.fun.apply(null, this.array);
};
process.title = 'browser';
process.browser = true;
process.env = {};
process.argv = [];
process.version = ''; // empty string to avoid regexp issues
process.versions = {};

function noop() {}

process.on = noop;
process.addListener = noop;
process.once = noop;
process.off = noop;
process.removeListener = noop;
process.removeAllListeners = noop;
process.emit = noop;
process.prependListener = noop;
process.prependOnceListener = noop;

process.listeners = function (name) { return [] }

process.binding = function (name) {
    throw new Error('process.binding is not supported');
};

process.cwd = function () { return '/' };
process.chdir = function (dir) {
    throw new Error('process.chdir is not supported');
};
process.umask = function() { return 0; };


/***/ }),

/***/ "./node_modules/punycode/punycode.js":
/*!*******************************************!*\
  !*** ./node_modules/punycode/punycode.js ***!
  \*******************************************/
/***/ (function(module, exports, __webpack_require__) {

/* module decorator */ module = __webpack_require__.nmd(module);
var __WEBPACK_AMD_DEFINE_RESULT__;/*! https://mths.be/punycode v1.4.1 by @mathias */
;(function(root) {

	/** Detect free variables */
	var freeExports =  true && exports &&
		!exports.nodeType && exports;
	var freeModule =  true && module &&
		!module.nodeType && module;
	var freeGlobal = typeof __webpack_require__.g == 'object' && __webpack_require__.g;
	if (
		freeGlobal.global === freeGlobal ||
		freeGlobal.window === freeGlobal ||
		freeGlobal.self === freeGlobal
	) {
		root = freeGlobal;
	}

	/**
	 * The `punycode` object.
	 * @name punycode
	 * @type Object
	 */
	var punycode,

	/** Highest positive signed 32-bit float value */
	maxInt = 2147483647, // aka. 0x7FFFFFFF or 2^31-1

	/** Bootstring parameters */
	base = 36,
	tMin = 1,
	tMax = 26,
	skew = 38,
	damp = 700,
	initialBias = 72,
	initialN = 128, // 0x80
	delimiter = '-', // '\x2D'

	/** Regular expressions */
	regexPunycode = /^xn--/,
	regexNonASCII = /[^\x20-\x7E]/, // unprintable ASCII chars + non-ASCII chars
	regexSeparators = /[\x2E\u3002\uFF0E\uFF61]/g, // RFC 3490 separators

	/** Error messages */
	errors = {
		'overflow': 'Overflow: input needs wider integers to process',
		'not-basic': 'Illegal input >= 0x80 (not a basic code point)',
		'invalid-input': 'Invalid input'
	},

	/** Convenience shortcuts */
	baseMinusTMin = base - tMin,
	floor = Math.floor,
	stringFromCharCode = String.fromCharCode,

	/** Temporary variable */
	key;

	/*--------------------------------------------------------------------------*/

	/**
	 * A generic error utility function.
	 * @private
	 * @param {String} type The error type.
	 * @returns {Error} Throws a `RangeError` with the applicable error message.
	 */
	function error(type) {
		throw new RangeError(errors[type]);
	}

	/**
	 * A generic `Array#map` utility function.
	 * @private
	 * @param {Array} array The array to iterate over.
	 * @param {Function} callback The function that gets called for every array
	 * item.
	 * @returns {Array} A new array of values returned by the callback function.
	 */
	function map(array, fn) {
		var length = array.length;
		var result = [];
		while (length--) {
			result[length] = fn(array[length]);
		}
		return result;
	}

	/**
	 * A simple `Array#map`-like wrapper to work with domain name strings or email
	 * addresses.
	 * @private
	 * @param {String} domain The domain name or email address.
	 * @param {Function} callback The function that gets called for every
	 * character.
	 * @returns {Array} A new string of characters returned by the callback
	 * function.
	 */
	function mapDomain(string, fn) {
		var parts = string.split('@');
		var result = '';
		if (parts.length > 1) {
			// In email addresses, only the domain name should be punycoded. Leave
			// the local part (i.e. everything up to `@`) intact.
			result = parts[0] + '@';
			string = parts[1];
		}
		// Avoid `split(regex)` for IE8 compatibility. See #17.
		string = string.replace(regexSeparators, '\x2E');
		var labels = string.split('.');
		var encoded = map(labels, fn).join('.');
		return result + encoded;
	}

	/**
	 * Creates an array containing the numeric code points of each Unicode
	 * character in the string. While JavaScript uses UCS-2 internally,
	 * this function will convert a pair of surrogate halves (each of which
	 * UCS-2 exposes as separate characters) into a single code point,
	 * matching UTF-16.
	 * @see `punycode.ucs2.encode`
	 * @see <https://mathiasbynens.be/notes/javascript-encoding>
	 * @memberOf punycode.ucs2
	 * @name decode
	 * @param {String} string The Unicode input string (UCS-2).
	 * @returns {Array} The new array of code points.
	 */
	function ucs2decode(string) {
		var output = [],
		    counter = 0,
		    length = string.length,
		    value,
		    extra;
		while (counter < length) {
			value = string.charCodeAt(counter++);
			if (value >= 0xD800 && value <= 0xDBFF && counter < length) {
				// high surrogate, and there is a next character
				extra = string.charCodeAt(counter++);
				if ((extra & 0xFC00) == 0xDC00) { // low surrogate
					output.push(((value & 0x3FF) << 10) + (extra & 0x3FF) + 0x10000);
				} else {
					// unmatched surrogate; only append this code unit, in case the next
					// code unit is the high surrogate of a surrogate pair
					output.push(value);
					counter--;
				}
			} else {
				output.push(value);
			}
		}
		return output;
	}

	/**
	 * Creates a string based on an array of numeric code points.
	 * @see `punycode.ucs2.decode`
	 * @memberOf punycode.ucs2
	 * @name encode
	 * @param {Array} codePoints The array of numeric code points.
	 * @returns {String} The new Unicode string (UCS-2).
	 */
	function ucs2encode(array) {
		return map(array, function(value) {
			var output = '';
			if (value > 0xFFFF) {
				value -= 0x10000;
				output += stringFromCharCode(value >>> 10 & 0x3FF | 0xD800);
				value = 0xDC00 | value & 0x3FF;
			}
			output += stringFromCharCode(value);
			return output;
		}).join('');
	}

	/**
	 * Converts a basic code point into a digit/integer.
	 * @see `digitToBasic()`
	 * @private
	 * @param {Number} codePoint The basic numeric code point value.
	 * @returns {Number} The numeric value of a basic code point (for use in
	 * representing integers) in the range `0` to `base - 1`, or `base` if
	 * the code point does not represent a value.
	 */
	function basicToDigit(codePoint) {
		if (codePoint - 48 < 10) {
			return codePoint - 22;
		}
		if (codePoint - 65 < 26) {
			return codePoint - 65;
		}
		if (codePoint - 97 < 26) {
			return codePoint - 97;
		}
		return base;
	}

	/**
	 * Converts a digit/integer into a basic code point.
	 * @see `basicToDigit()`
	 * @private
	 * @param {Number} digit The numeric value of a basic code point.
	 * @returns {Number} The basic code point whose value (when used for
	 * representing integers) is `digit`, which needs to be in the range
	 * `0` to `base - 1`. If `flag` is non-zero, the uppercase form is
	 * used; else, the lowercase form is used. The behavior is undefined
	 * if `flag` is non-zero and `digit` has no uppercase form.
	 */
	function digitToBasic(digit, flag) {
		//  0..25 map to ASCII a..z or A..Z
		// 26..35 map to ASCII 0..9
		return digit + 22 + 75 * (digit < 26) - ((flag != 0) << 5);
	}

	/**
	 * Bias adaptation function as per section 3.4 of RFC 3492.
	 * https://tools.ietf.org/html/rfc3492#section-3.4
	 * @private
	 */
	function adapt(delta, numPoints, firstTime) {
		var k = 0;
		delta = firstTime ? floor(delta / damp) : delta >> 1;
		delta += floor(delta / numPoints);
		for (/* no initialization */; delta > baseMinusTMin * tMax >> 1; k += base) {
			delta = floor(delta / baseMinusTMin);
		}
		return floor(k + (baseMinusTMin + 1) * delta / (delta + skew));
	}

	/**
	 * Converts a Punycode string of ASCII-only symbols to a string of Unicode
	 * symbols.
	 * @memberOf punycode
	 * @param {String} input The Punycode string of ASCII-only symbols.
	 * @returns {String} The resulting string of Unicode symbols.
	 */
	function decode(input) {
		// Don't use UCS-2
		var output = [],
		    inputLength = input.length,
		    out,
		    i = 0,
		    n = initialN,
		    bias = initialBias,
		    basic,
		    j,
		    index,
		    oldi,
		    w,
		    k,
		    digit,
		    t,
		    /** Cached calculation results */
		    baseMinusT;

		// Handle the basic code points: let `basic` be the number of input code
		// points before the last delimiter, or `0` if there is none, then copy
		// the first basic code points to the output.

		basic = input.lastIndexOf(delimiter);
		if (basic < 0) {
			basic = 0;
		}

		for (j = 0; j < basic; ++j) {
			// if it's not a basic code point
			if (input.charCodeAt(j) >= 0x80) {
				error('not-basic');
			}
			output.push(input.charCodeAt(j));
		}

		// Main decoding loop: start just after the last delimiter if any basic code
		// points were copied; start at the beginning otherwise.

		for (index = basic > 0 ? basic + 1 : 0; index < inputLength; /* no final expression */) {

			// `index` is the index of the next character to be consumed.
			// Decode a generalized variable-length integer into `delta`,
			// which gets added to `i`. The overflow checking is easier
			// if we increase `i` as we go, then subtract off its starting
			// value at the end to obtain `delta`.
			for (oldi = i, w = 1, k = base; /* no condition */; k += base) {

				if (index >= inputLength) {
					error('invalid-input');
				}

				digit = basicToDigit(input.charCodeAt(index++));

				if (digit >= base || digit > floor((maxInt - i) / w)) {
					error('overflow');
				}

				i += digit * w;
				t = k <= bias ? tMin : (k >= bias + tMax ? tMax : k - bias);

				if (digit < t) {
					break;
				}

				baseMinusT = base - t;
				if (w > floor(maxInt / baseMinusT)) {
					error('overflow');
				}

				w *= baseMinusT;

			}

			out = output.length + 1;
			bias = adapt(i - oldi, out, oldi == 0);

			// `i` was supposed to wrap around from `out` to `0`,
			// incrementing `n` each time, so we'll fix that now:
			if (floor(i / out) > maxInt - n) {
				error('overflow');
			}

			n += floor(i / out);
			i %= out;

			// Insert `n` at position `i` of the output
			output.splice(i++, 0, n);

		}

		return ucs2encode(output);
	}

	/**
	 * Converts a string of Unicode symbols (e.g. a domain name label) to a
	 * Punycode string of ASCII-only symbols.
	 * @memberOf punycode
	 * @param {String} input The string of Unicode symbols.
	 * @returns {String} The resulting Punycode string of ASCII-only symbols.
	 */
	function encode(input) {
		var n,
		    delta,
		    handledCPCount,
		    basicLength,
		    bias,
		    j,
		    m,
		    q,
		    k,
		    t,
		    currentValue,
		    output = [],
		    /** `inputLength` will hold the number of code points in `input`. */
		    inputLength,
		    /** Cached calculation results */
		    handledCPCountPlusOne,
		    baseMinusT,
		    qMinusT;

		// Convert the input in UCS-2 to Unicode
		input = ucs2decode(input);

		// Cache the length
		inputLength = input.length;

		// Initialize the state
		n = initialN;
		delta = 0;
		bias = initialBias;

		// Handle the basic code points
		for (j = 0; j < inputLength; ++j) {
			currentValue = input[j];
			if (currentValue < 0x80) {
				output.push(stringFromCharCode(currentValue));
			}
		}

		handledCPCount = basicLength = output.length;

		// `handledCPCount` is the number of code points that have been handled;
		// `basicLength` is the number of basic code points.

		// Finish the basic string - if it is not empty - with a delimiter
		if (basicLength) {
			output.push(delimiter);
		}

		// Main encoding loop:
		while (handledCPCount < inputLength) {

			// All non-basic code points < n have been handled already. Find the next
			// larger one:
			for (m = maxInt, j = 0; j < inputLength; ++j) {
				currentValue = input[j];
				if (currentValue >= n && currentValue < m) {
					m = currentValue;
				}
			}

			// Increase `delta` enough to advance the decoder's <n,i> state to <m,0>,
			// but guard against overflow
			handledCPCountPlusOne = handledCPCount + 1;
			if (m - n > floor((maxInt - delta) / handledCPCountPlusOne)) {
				error('overflow');
			}

			delta += (m - n) * handledCPCountPlusOne;
			n = m;

			for (j = 0; j < inputLength; ++j) {
				currentValue = input[j];

				if (currentValue < n && ++delta > maxInt) {
					error('overflow');
				}

				if (currentValue == n) {
					// Represent delta as a generalized variable-length integer
					for (q = delta, k = base; /* no condition */; k += base) {
						t = k <= bias ? tMin : (k >= bias + tMax ? tMax : k - bias);
						if (q < t) {
							break;
						}
						qMinusT = q - t;
						baseMinusT = base - t;
						output.push(
							stringFromCharCode(digitToBasic(t + qMinusT % baseMinusT, 0))
						);
						q = floor(qMinusT / baseMinusT);
					}

					output.push(stringFromCharCode(digitToBasic(q, 0)));
					bias = adapt(delta, handledCPCountPlusOne, handledCPCount == basicLength);
					delta = 0;
					++handledCPCount;
				}
			}

			++delta;
			++n;

		}
		return output.join('');
	}

	/**
	 * Converts a Punycode string representing a domain name or an email address
	 * to Unicode. Only the Punycoded parts of the input will be converted, i.e.
	 * it doesn't matter if you call it on a string that has already been
	 * converted to Unicode.
	 * @memberOf punycode
	 * @param {String} input The Punycoded domain name or email address to
	 * convert to Unicode.
	 * @returns {String} The Unicode representation of the given Punycode
	 * string.
	 */
	function toUnicode(input) {
		return mapDomain(input, function(string) {
			return regexPunycode.test(string)
				? decode(string.slice(4).toLowerCase())
				: string;
		});
	}

	/**
	 * Converts a Unicode string representing a domain name or an email address to
	 * Punycode. Only the non-ASCII parts of the domain name will be converted,
	 * i.e. it doesn't matter if you call it with a domain that's already in
	 * ASCII.
	 * @memberOf punycode
	 * @param {String} input The domain name or email address to convert, as a
	 * Unicode string.
	 * @returns {String} The Punycode representation of the given domain name or
	 * email address.
	 */
	function toASCII(input) {
		return mapDomain(input, function(string) {
			return regexNonASCII.test(string)
				? 'xn--' + encode(string)
				: string;
		});
	}

	/*--------------------------------------------------------------------------*/

	/** Define the public API */
	punycode = {
		/**
		 * A string representing the current Punycode.js version number.
		 * @memberOf punycode
		 * @type String
		 */
		'version': '1.4.1',
		/**
		 * An object of methods to convert from JavaScript's internal character
		 * representation (UCS-2) to Unicode code points, and back.
		 * @see <https://mathiasbynens.be/notes/javascript-encoding>
		 * @memberOf punycode
		 * @type Object
		 */
		'ucs2': {
			'decode': ucs2decode,
			'encode': ucs2encode
		},
		'decode': decode,
		'encode': encode,
		'toASCII': toASCII,
		'toUnicode': toUnicode
	};

	/** Expose `punycode` */
	// Some AMD build optimizers, like r.js, check for specific condition patterns
	// like the following:
	if (
		true
	) {
		!(__WEBPACK_AMD_DEFINE_RESULT__ = (function() {
			return punycode;
		}).call(exports, __webpack_require__, exports, module),
		__WEBPACK_AMD_DEFINE_RESULT__ !== undefined && (module.exports = __WEBPACK_AMD_DEFINE_RESULT__));
	} else {}

}(this));


/***/ }),

/***/ "./node_modules/qs/lib/formats.js":
/*!****************************************!*\
  !*** ./node_modules/qs/lib/formats.js ***!
  \****************************************/
/***/ ((module) => {

"use strict";


var replace = String.prototype.replace;
var percentTwenties = /%20/g;

var Format = {
    RFC1738: 'RFC1738',
    RFC3986: 'RFC3986'
};

module.exports = {
    'default': Format.RFC3986,
    formatters: {
        RFC1738: function (value) {
            return replace.call(value, percentTwenties, '+');
        },
        RFC3986: function (value) {
            return String(value);
        }
    },
    RFC1738: Format.RFC1738,
    RFC3986: Format.RFC3986
};


/***/ }),

/***/ "./node_modules/qs/lib/index.js":
/*!**************************************!*\
  !*** ./node_modules/qs/lib/index.js ***!
  \**************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var stringify = __webpack_require__(/*! ./stringify */ "./node_modules/qs/lib/stringify.js");
var parse = __webpack_require__(/*! ./parse */ "./node_modules/qs/lib/parse.js");
var formats = __webpack_require__(/*! ./formats */ "./node_modules/qs/lib/formats.js");

module.exports = {
    formats: formats,
    parse: parse,
    stringify: stringify
};


/***/ }),

/***/ "./node_modules/qs/lib/parse.js":
/*!**************************************!*\
  !*** ./node_modules/qs/lib/parse.js ***!
  \**************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var utils = __webpack_require__(/*! ./utils */ "./node_modules/qs/lib/utils.js");

var has = Object.prototype.hasOwnProperty;
var isArray = Array.isArray;

var defaults = {
    allowDots: false,
    allowPrototypes: false,
    allowSparse: false,
    arrayLimit: 20,
    charset: 'utf-8',
    charsetSentinel: false,
    comma: false,
    decoder: utils.decode,
    delimiter: '&',
    depth: 5,
    ignoreQueryPrefix: false,
    interpretNumericEntities: false,
    parameterLimit: 1000,
    parseArrays: true,
    plainObjects: false,
    strictNullHandling: false
};

var interpretNumericEntities = function (str) {
    return str.replace(/&#(\d+);/g, function ($0, numberStr) {
        return String.fromCharCode(parseInt(numberStr, 10));
    });
};

var parseArrayValue = function (val, options) {
    if (val && typeof val === 'string' && options.comma && val.indexOf(',') > -1) {
        return val.split(',');
    }

    return val;
};

// This is what browsers will submit when the ✓ character occurs in an
// application/x-www-form-urlencoded body and the encoding of the page containing
// the form is iso-8859-1, or when the submitted form has an accept-charset
// attribute of iso-8859-1. Presumably also with other charsets that do not contain
// the ✓ character, such as us-ascii.
var isoSentinel = 'utf8=%26%2310003%3B'; // encodeURIComponent('&#10003;')

// These are the percent-encoded utf-8 octets representing a checkmark, indicating that the request actually is utf-8 encoded.
var charsetSentinel = 'utf8=%E2%9C%93'; // encodeURIComponent('✓')

var parseValues = function parseQueryStringValues(str, options) {
    var obj = { __proto__: null };

    var cleanStr = options.ignoreQueryPrefix ? str.replace(/^\?/, '') : str;
    var limit = options.parameterLimit === Infinity ? undefined : options.parameterLimit;
    var parts = cleanStr.split(options.delimiter, limit);
    var skipIndex = -1; // Keep track of where the utf8 sentinel was found
    var i;

    var charset = options.charset;
    if (options.charsetSentinel) {
        for (i = 0; i < parts.length; ++i) {
            if (parts[i].indexOf('utf8=') === 0) {
                if (parts[i] === charsetSentinel) {
                    charset = 'utf-8';
                } else if (parts[i] === isoSentinel) {
                    charset = 'iso-8859-1';
                }
                skipIndex = i;
                i = parts.length; // The eslint settings do not allow break;
            }
        }
    }

    for (i = 0; i < parts.length; ++i) {
        if (i === skipIndex) {
            continue;
        }
        var part = parts[i];

        var bracketEqualsPos = part.indexOf(']=');
        var pos = bracketEqualsPos === -1 ? part.indexOf('=') : bracketEqualsPos + 1;

        var key, val;
        if (pos === -1) {
            key = options.decoder(part, defaults.decoder, charset, 'key');
            val = options.strictNullHandling ? null : '';
        } else {
            key = options.decoder(part.slice(0, pos), defaults.decoder, charset, 'key');
            val = utils.maybeMap(
                parseArrayValue(part.slice(pos + 1), options),
                function (encodedVal) {
                    return options.decoder(encodedVal, defaults.decoder, charset, 'value');
                }
            );
        }

        if (val && options.interpretNumericEntities && charset === 'iso-8859-1') {
            val = interpretNumericEntities(val);
        }

        if (part.indexOf('[]=') > -1) {
            val = isArray(val) ? [val] : val;
        }

        if (has.call(obj, key)) {
            obj[key] = utils.combine(obj[key], val);
        } else {
            obj[key] = val;
        }
    }

    return obj;
};

var parseObject = function (chain, val, options, valuesParsed) {
    var leaf = valuesParsed ? val : parseArrayValue(val, options);

    for (var i = chain.length - 1; i >= 0; --i) {
        var obj;
        var root = chain[i];

        if (root === '[]' && options.parseArrays) {
            obj = [].concat(leaf);
        } else {
            obj = options.plainObjects ? Object.create(null) : {};
            var cleanRoot = root.charAt(0) === '[' && root.charAt(root.length - 1) === ']' ? root.slice(1, -1) : root;
            var index = parseInt(cleanRoot, 10);
            if (!options.parseArrays && cleanRoot === '') {
                obj = { 0: leaf };
            } else if (
                !isNaN(index)
                && root !== cleanRoot
                && String(index) === cleanRoot
                && index >= 0
                && (options.parseArrays && index <= options.arrayLimit)
            ) {
                obj = [];
                obj[index] = leaf;
            } else if (cleanRoot !== '__proto__') {
                obj[cleanRoot] = leaf;
            }
        }

        leaf = obj;
    }

    return leaf;
};

var parseKeys = function parseQueryStringKeys(givenKey, val, options, valuesParsed) {
    if (!givenKey) {
        return;
    }

    // Transform dot notation to bracket notation
    var key = options.allowDots ? givenKey.replace(/\.([^.[]+)/g, '[$1]') : givenKey;

    // The regex chunks

    var brackets = /(\[[^[\]]*])/;
    var child = /(\[[^[\]]*])/g;

    // Get the parent

    var segment = options.depth > 0 && brackets.exec(key);
    var parent = segment ? key.slice(0, segment.index) : key;

    // Stash the parent if it exists

    var keys = [];
    if (parent) {
        // If we aren't using plain objects, optionally prefix keys that would overwrite object prototype properties
        if (!options.plainObjects && has.call(Object.prototype, parent)) {
            if (!options.allowPrototypes) {
                return;
            }
        }

        keys.push(parent);
    }

    // Loop through children appending to the array until we hit depth

    var i = 0;
    while (options.depth > 0 && (segment = child.exec(key)) !== null && i < options.depth) {
        i += 1;
        if (!options.plainObjects && has.call(Object.prototype, segment[1].slice(1, -1))) {
            if (!options.allowPrototypes) {
                return;
            }
        }
        keys.push(segment[1]);
    }

    // If there's a remainder, just add whatever is left

    if (segment) {
        keys.push('[' + key.slice(segment.index) + ']');
    }

    return parseObject(keys, val, options, valuesParsed);
};

var normalizeParseOptions = function normalizeParseOptions(opts) {
    if (!opts) {
        return defaults;
    }

    if (opts.decoder !== null && opts.decoder !== undefined && typeof opts.decoder !== 'function') {
        throw new TypeError('Decoder has to be a function.');
    }

    if (typeof opts.charset !== 'undefined' && opts.charset !== 'utf-8' && opts.charset !== 'iso-8859-1') {
        throw new TypeError('The charset option must be either utf-8, iso-8859-1, or undefined');
    }
    var charset = typeof opts.charset === 'undefined' ? defaults.charset : opts.charset;

    return {
        allowDots: typeof opts.allowDots === 'undefined' ? defaults.allowDots : !!opts.allowDots,
        allowPrototypes: typeof opts.allowPrototypes === 'boolean' ? opts.allowPrototypes : defaults.allowPrototypes,
        allowSparse: typeof opts.allowSparse === 'boolean' ? opts.allowSparse : defaults.allowSparse,
        arrayLimit: typeof opts.arrayLimit === 'number' ? opts.arrayLimit : defaults.arrayLimit,
        charset: charset,
        charsetSentinel: typeof opts.charsetSentinel === 'boolean' ? opts.charsetSentinel : defaults.charsetSentinel,
        comma: typeof opts.comma === 'boolean' ? opts.comma : defaults.comma,
        decoder: typeof opts.decoder === 'function' ? opts.decoder : defaults.decoder,
        delimiter: typeof opts.delimiter === 'string' || utils.isRegExp(opts.delimiter) ? opts.delimiter : defaults.delimiter,
        // eslint-disable-next-line no-implicit-coercion, no-extra-parens
        depth: (typeof opts.depth === 'number' || opts.depth === false) ? +opts.depth : defaults.depth,
        ignoreQueryPrefix: opts.ignoreQueryPrefix === true,
        interpretNumericEntities: typeof opts.interpretNumericEntities === 'boolean' ? opts.interpretNumericEntities : defaults.interpretNumericEntities,
        parameterLimit: typeof opts.parameterLimit === 'number' ? opts.parameterLimit : defaults.parameterLimit,
        parseArrays: opts.parseArrays !== false,
        plainObjects: typeof opts.plainObjects === 'boolean' ? opts.plainObjects : defaults.plainObjects,
        strictNullHandling: typeof opts.strictNullHandling === 'boolean' ? opts.strictNullHandling : defaults.strictNullHandling
    };
};

module.exports = function (str, opts) {
    var options = normalizeParseOptions(opts);

    if (str === '' || str === null || typeof str === 'undefined') {
        return options.plainObjects ? Object.create(null) : {};
    }

    var tempObj = typeof str === 'string' ? parseValues(str, options) : str;
    var obj = options.plainObjects ? Object.create(null) : {};

    // Iterate over the keys and setup the new object

    var keys = Object.keys(tempObj);
    for (var i = 0; i < keys.length; ++i) {
        var key = keys[i];
        var newObj = parseKeys(key, tempObj[key], options, typeof str === 'string');
        obj = utils.merge(obj, newObj, options);
    }

    if (options.allowSparse === true) {
        return obj;
    }

    return utils.compact(obj);
};


/***/ }),

/***/ "./node_modules/qs/lib/stringify.js":
/*!******************************************!*\
  !*** ./node_modules/qs/lib/stringify.js ***!
  \******************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var getSideChannel = __webpack_require__(/*! side-channel */ "./node_modules/side-channel/index.js");
var utils = __webpack_require__(/*! ./utils */ "./node_modules/qs/lib/utils.js");
var formats = __webpack_require__(/*! ./formats */ "./node_modules/qs/lib/formats.js");
var has = Object.prototype.hasOwnProperty;

var arrayPrefixGenerators = {
    brackets: function brackets(prefix) {
        return prefix + '[]';
    },
    comma: 'comma',
    indices: function indices(prefix, key) {
        return prefix + '[' + key + ']';
    },
    repeat: function repeat(prefix) {
        return prefix;
    }
};

var isArray = Array.isArray;
var push = Array.prototype.push;
var pushToArray = function (arr, valueOrArray) {
    push.apply(arr, isArray(valueOrArray) ? valueOrArray : [valueOrArray]);
};

var toISO = Date.prototype.toISOString;

var defaultFormat = formats['default'];
var defaults = {
    addQueryPrefix: false,
    allowDots: false,
    charset: 'utf-8',
    charsetSentinel: false,
    delimiter: '&',
    encode: true,
    encoder: utils.encode,
    encodeValuesOnly: false,
    format: defaultFormat,
    formatter: formats.formatters[defaultFormat],
    // deprecated
    indices: false,
    serializeDate: function serializeDate(date) {
        return toISO.call(date);
    },
    skipNulls: false,
    strictNullHandling: false
};

var isNonNullishPrimitive = function isNonNullishPrimitive(v) {
    return typeof v === 'string'
        || typeof v === 'number'
        || typeof v === 'boolean'
        || typeof v === 'symbol'
        || typeof v === 'bigint';
};

var sentinel = {};

var stringify = function stringify(
    object,
    prefix,
    generateArrayPrefix,
    commaRoundTrip,
    strictNullHandling,
    skipNulls,
    encoder,
    filter,
    sort,
    allowDots,
    serializeDate,
    format,
    formatter,
    encodeValuesOnly,
    charset,
    sideChannel
) {
    var obj = object;

    var tmpSc = sideChannel;
    var step = 0;
    var findFlag = false;
    while ((tmpSc = tmpSc.get(sentinel)) !== void undefined && !findFlag) {
        // Where object last appeared in the ref tree
        var pos = tmpSc.get(object);
        step += 1;
        if (typeof pos !== 'undefined') {
            if (pos === step) {
                throw new RangeError('Cyclic object value');
            } else {
                findFlag = true; // Break while
            }
        }
        if (typeof tmpSc.get(sentinel) === 'undefined') {
            step = 0;
        }
    }

    if (typeof filter === 'function') {
        obj = filter(prefix, obj);
    } else if (obj instanceof Date) {
        obj = serializeDate(obj);
    } else if (generateArrayPrefix === 'comma' && isArray(obj)) {
        obj = utils.maybeMap(obj, function (value) {
            if (value instanceof Date) {
                return serializeDate(value);
            }
            return value;
        });
    }

    if (obj === null) {
        if (strictNullHandling) {
            return encoder && !encodeValuesOnly ? encoder(prefix, defaults.encoder, charset, 'key', format) : prefix;
        }

        obj = '';
    }

    if (isNonNullishPrimitive(obj) || utils.isBuffer(obj)) {
        if (encoder) {
            var keyValue = encodeValuesOnly ? prefix : encoder(prefix, defaults.encoder, charset, 'key', format);
            return [formatter(keyValue) + '=' + formatter(encoder(obj, defaults.encoder, charset, 'value', format))];
        }
        return [formatter(prefix) + '=' + formatter(String(obj))];
    }

    var values = [];

    if (typeof obj === 'undefined') {
        return values;
    }

    var objKeys;
    if (generateArrayPrefix === 'comma' && isArray(obj)) {
        // we need to join elements in
        if (encodeValuesOnly && encoder) {
            obj = utils.maybeMap(obj, encoder);
        }
        objKeys = [{ value: obj.length > 0 ? obj.join(',') || null : void undefined }];
    } else if (isArray(filter)) {
        objKeys = filter;
    } else {
        var keys = Object.keys(obj);
        objKeys = sort ? keys.sort(sort) : keys;
    }

    var adjustedPrefix = commaRoundTrip && isArray(obj) && obj.length === 1 ? prefix + '[]' : prefix;

    for (var j = 0; j < objKeys.length; ++j) {
        var key = objKeys[j];
        var value = typeof key === 'object' && typeof key.value !== 'undefined' ? key.value : obj[key];

        if (skipNulls && value === null) {
            continue;
        }

        var keyPrefix = isArray(obj)
            ? typeof generateArrayPrefix === 'function' ? generateArrayPrefix(adjustedPrefix, key) : adjustedPrefix
            : adjustedPrefix + (allowDots ? '.' + key : '[' + key + ']');

        sideChannel.set(object, step);
        var valueSideChannel = getSideChannel();
        valueSideChannel.set(sentinel, sideChannel);
        pushToArray(values, stringify(
            value,
            keyPrefix,
            generateArrayPrefix,
            commaRoundTrip,
            strictNullHandling,
            skipNulls,
            generateArrayPrefix === 'comma' && encodeValuesOnly && isArray(obj) ? null : encoder,
            filter,
            sort,
            allowDots,
            serializeDate,
            format,
            formatter,
            encodeValuesOnly,
            charset,
            valueSideChannel
        ));
    }

    return values;
};

var normalizeStringifyOptions = function normalizeStringifyOptions(opts) {
    if (!opts) {
        return defaults;
    }

    if (opts.encoder !== null && typeof opts.encoder !== 'undefined' && typeof opts.encoder !== 'function') {
        throw new TypeError('Encoder has to be a function.');
    }

    var charset = opts.charset || defaults.charset;
    if (typeof opts.charset !== 'undefined' && opts.charset !== 'utf-8' && opts.charset !== 'iso-8859-1') {
        throw new TypeError('The charset option must be either utf-8, iso-8859-1, or undefined');
    }

    var format = formats['default'];
    if (typeof opts.format !== 'undefined') {
        if (!has.call(formats.formatters, opts.format)) {
            throw new TypeError('Unknown format option provided.');
        }
        format = opts.format;
    }
    var formatter = formats.formatters[format];

    var filter = defaults.filter;
    if (typeof opts.filter === 'function' || isArray(opts.filter)) {
        filter = opts.filter;
    }

    return {
        addQueryPrefix: typeof opts.addQueryPrefix === 'boolean' ? opts.addQueryPrefix : defaults.addQueryPrefix,
        allowDots: typeof opts.allowDots === 'undefined' ? defaults.allowDots : !!opts.allowDots,
        charset: charset,
        charsetSentinel: typeof opts.charsetSentinel === 'boolean' ? opts.charsetSentinel : defaults.charsetSentinel,
        delimiter: typeof opts.delimiter === 'undefined' ? defaults.delimiter : opts.delimiter,
        encode: typeof opts.encode === 'boolean' ? opts.encode : defaults.encode,
        encoder: typeof opts.encoder === 'function' ? opts.encoder : defaults.encoder,
        encodeValuesOnly: typeof opts.encodeValuesOnly === 'boolean' ? opts.encodeValuesOnly : defaults.encodeValuesOnly,
        filter: filter,
        format: format,
        formatter: formatter,
        serializeDate: typeof opts.serializeDate === 'function' ? opts.serializeDate : defaults.serializeDate,
        skipNulls: typeof opts.skipNulls === 'boolean' ? opts.skipNulls : defaults.skipNulls,
        sort: typeof opts.sort === 'function' ? opts.sort : null,
        strictNullHandling: typeof opts.strictNullHandling === 'boolean' ? opts.strictNullHandling : defaults.strictNullHandling
    };
};

module.exports = function (object, opts) {
    var obj = object;
    var options = normalizeStringifyOptions(opts);

    var objKeys;
    var filter;

    if (typeof options.filter === 'function') {
        filter = options.filter;
        obj = filter('', obj);
    } else if (isArray(options.filter)) {
        filter = options.filter;
        objKeys = filter;
    }

    var keys = [];

    if (typeof obj !== 'object' || obj === null) {
        return '';
    }

    var arrayFormat;
    if (opts && opts.arrayFormat in arrayPrefixGenerators) {
        arrayFormat = opts.arrayFormat;
    } else if (opts && 'indices' in opts) {
        arrayFormat = opts.indices ? 'indices' : 'repeat';
    } else {
        arrayFormat = 'indices';
    }

    var generateArrayPrefix = arrayPrefixGenerators[arrayFormat];
    if (opts && 'commaRoundTrip' in opts && typeof opts.commaRoundTrip !== 'boolean') {
        throw new TypeError('`commaRoundTrip` must be a boolean, or absent');
    }
    var commaRoundTrip = generateArrayPrefix === 'comma' && opts && opts.commaRoundTrip;

    if (!objKeys) {
        objKeys = Object.keys(obj);
    }

    if (options.sort) {
        objKeys.sort(options.sort);
    }

    var sideChannel = getSideChannel();
    for (var i = 0; i < objKeys.length; ++i) {
        var key = objKeys[i];

        if (options.skipNulls && obj[key] === null) {
            continue;
        }
        pushToArray(keys, stringify(
            obj[key],
            key,
            generateArrayPrefix,
            commaRoundTrip,
            options.strictNullHandling,
            options.skipNulls,
            options.encode ? options.encoder : null,
            options.filter,
            options.sort,
            options.allowDots,
            options.serializeDate,
            options.format,
            options.formatter,
            options.encodeValuesOnly,
            options.charset,
            sideChannel
        ));
    }

    var joined = keys.join(options.delimiter);
    var prefix = options.addQueryPrefix === true ? '?' : '';

    if (options.charsetSentinel) {
        if (options.charset === 'iso-8859-1') {
            // encodeURIComponent('&#10003;'), the "numeric entity" representation of a checkmark
            prefix += 'utf8=%26%2310003%3B&';
        } else {
            // encodeURIComponent('✓')
            prefix += 'utf8=%E2%9C%93&';
        }
    }

    return joined.length > 0 ? prefix + joined : '';
};


/***/ }),

/***/ "./node_modules/qs/lib/utils.js":
/*!**************************************!*\
  !*** ./node_modules/qs/lib/utils.js ***!
  \**************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var formats = __webpack_require__(/*! ./formats */ "./node_modules/qs/lib/formats.js");

var has = Object.prototype.hasOwnProperty;
var isArray = Array.isArray;

var hexTable = (function () {
    var array = [];
    for (var i = 0; i < 256; ++i) {
        array.push('%' + ((i < 16 ? '0' : '') + i.toString(16)).toUpperCase());
    }

    return array;
}());

var compactQueue = function compactQueue(queue) {
    while (queue.length > 1) {
        var item = queue.pop();
        var obj = item.obj[item.prop];

        if (isArray(obj)) {
            var compacted = [];

            for (var j = 0; j < obj.length; ++j) {
                if (typeof obj[j] !== 'undefined') {
                    compacted.push(obj[j]);
                }
            }

            item.obj[item.prop] = compacted;
        }
    }
};

var arrayToObject = function arrayToObject(source, options) {
    var obj = options && options.plainObjects ? Object.create(null) : {};
    for (var i = 0; i < source.length; ++i) {
        if (typeof source[i] !== 'undefined') {
            obj[i] = source[i];
        }
    }

    return obj;
};

var merge = function merge(target, source, options) {
    /* eslint no-param-reassign: 0 */
    if (!source) {
        return target;
    }

    if (typeof source !== 'object') {
        if (isArray(target)) {
            target.push(source);
        } else if (target && typeof target === 'object') {
            if ((options && (options.plainObjects || options.allowPrototypes)) || !has.call(Object.prototype, source)) {
                target[source] = true;
            }
        } else {
            return [target, source];
        }

        return target;
    }

    if (!target || typeof target !== 'object') {
        return [target].concat(source);
    }

    var mergeTarget = target;
    if (isArray(target) && !isArray(source)) {
        mergeTarget = arrayToObject(target, options);
    }

    if (isArray(target) && isArray(source)) {
        source.forEach(function (item, i) {
            if (has.call(target, i)) {
                var targetItem = target[i];
                if (targetItem && typeof targetItem === 'object' && item && typeof item === 'object') {
                    target[i] = merge(targetItem, item, options);
                } else {
                    target.push(item);
                }
            } else {
                target[i] = item;
            }
        });
        return target;
    }

    return Object.keys(source).reduce(function (acc, key) {
        var value = source[key];

        if (has.call(acc, key)) {
            acc[key] = merge(acc[key], value, options);
        } else {
            acc[key] = value;
        }
        return acc;
    }, mergeTarget);
};

var assign = function assignSingleSource(target, source) {
    return Object.keys(source).reduce(function (acc, key) {
        acc[key] = source[key];
        return acc;
    }, target);
};

var decode = function (str, decoder, charset) {
    var strWithoutPlus = str.replace(/\+/g, ' ');
    if (charset === 'iso-8859-1') {
        // unescape never throws, no try...catch needed:
        return strWithoutPlus.replace(/%[0-9a-f]{2}/gi, unescape);
    }
    // utf-8
    try {
        return decodeURIComponent(strWithoutPlus);
    } catch (e) {
        return strWithoutPlus;
    }
};

var encode = function encode(str, defaultEncoder, charset, kind, format) {
    // This code was originally written by Brian White (mscdex) for the io.js core querystring library.
    // It has been adapted here for stricter adherence to RFC 3986
    if (str.length === 0) {
        return str;
    }

    var string = str;
    if (typeof str === 'symbol') {
        string = Symbol.prototype.toString.call(str);
    } else if (typeof str !== 'string') {
        string = String(str);
    }

    if (charset === 'iso-8859-1') {
        return escape(string).replace(/%u[0-9a-f]{4}/gi, function ($0) {
            return '%26%23' + parseInt($0.slice(2), 16) + '%3B';
        });
    }

    var out = '';
    for (var i = 0; i < string.length; ++i) {
        var c = string.charCodeAt(i);

        if (
            c === 0x2D // -
            || c === 0x2E // .
            || c === 0x5F // _
            || c === 0x7E // ~
            || (c >= 0x30 && c <= 0x39) // 0-9
            || (c >= 0x41 && c <= 0x5A) // a-z
            || (c >= 0x61 && c <= 0x7A) // A-Z
            || (format === formats.RFC1738 && (c === 0x28 || c === 0x29)) // ( )
        ) {
            out += string.charAt(i);
            continue;
        }

        if (c < 0x80) {
            out = out + hexTable[c];
            continue;
        }

        if (c < 0x800) {
            out = out + (hexTable[0xC0 | (c >> 6)] + hexTable[0x80 | (c & 0x3F)]);
            continue;
        }

        if (c < 0xD800 || c >= 0xE000) {
            out = out + (hexTable[0xE0 | (c >> 12)] + hexTable[0x80 | ((c >> 6) & 0x3F)] + hexTable[0x80 | (c & 0x3F)]);
            continue;
        }

        i += 1;
        c = 0x10000 + (((c & 0x3FF) << 10) | (string.charCodeAt(i) & 0x3FF));
        /* eslint operator-linebreak: [2, "before"] */
        out += hexTable[0xF0 | (c >> 18)]
            + hexTable[0x80 | ((c >> 12) & 0x3F)]
            + hexTable[0x80 | ((c >> 6) & 0x3F)]
            + hexTable[0x80 | (c & 0x3F)];
    }

    return out;
};

var compact = function compact(value) {
    var queue = [{ obj: { o: value }, prop: 'o' }];
    var refs = [];

    for (var i = 0; i < queue.length; ++i) {
        var item = queue[i];
        var obj = item.obj[item.prop];

        var keys = Object.keys(obj);
        for (var j = 0; j < keys.length; ++j) {
            var key = keys[j];
            var val = obj[key];
            if (typeof val === 'object' && val !== null && refs.indexOf(val) === -1) {
                queue.push({ obj: obj, prop: key });
                refs.push(val);
            }
        }
    }

    compactQueue(queue);

    return value;
};

var isRegExp = function isRegExp(obj) {
    return Object.prototype.toString.call(obj) === '[object RegExp]';
};

var isBuffer = function isBuffer(obj) {
    if (!obj || typeof obj !== 'object') {
        return false;
    }

    return !!(obj.constructor && obj.constructor.isBuffer && obj.constructor.isBuffer(obj));
};

var combine = function combine(a, b) {
    return [].concat(a, b);
};

var maybeMap = function maybeMap(val, fn) {
    if (isArray(val)) {
        var mapped = [];
        for (var i = 0; i < val.length; i += 1) {
            mapped.push(fn(val[i]));
        }
        return mapped;
    }
    return fn(val);
};

module.exports = {
    arrayToObject: arrayToObject,
    assign: assign,
    combine: combine,
    compact: compact,
    decode: decode,
    encode: encode,
    isBuffer: isBuffer,
    isRegExp: isRegExp,
    maybeMap: maybeMap,
    merge: merge
};


/***/ }),

/***/ "./node_modules/readable-stream/errors-browser.js":
/*!********************************************************!*\
  !*** ./node_modules/readable-stream/errors-browser.js ***!
  \********************************************************/
/***/ ((module) => {

"use strict";


function _inheritsLoose(subClass, superClass) { subClass.prototype = Object.create(superClass.prototype); subClass.prototype.constructor = subClass; subClass.__proto__ = superClass; }

var codes = {};

function createErrorType(code, message, Base) {
  if (!Base) {
    Base = Error;
  }

  function getMessage(arg1, arg2, arg3) {
    if (typeof message === 'string') {
      return message;
    } else {
      return message(arg1, arg2, arg3);
    }
  }

  var NodeError =
  /*#__PURE__*/
  function (_Base) {
    _inheritsLoose(NodeError, _Base);

    function NodeError(arg1, arg2, arg3) {
      return _Base.call(this, getMessage(arg1, arg2, arg3)) || this;
    }

    return NodeError;
  }(Base);

  NodeError.prototype.name = Base.name;
  NodeError.prototype.code = code;
  codes[code] = NodeError;
} // https://github.com/nodejs/node/blob/v10.8.0/lib/internal/errors.js


function oneOf(expected, thing) {
  if (Array.isArray(expected)) {
    var len = expected.length;
    expected = expected.map(function (i) {
      return String(i);
    });

    if (len > 2) {
      return "one of ".concat(thing, " ").concat(expected.slice(0, len - 1).join(', '), ", or ") + expected[len - 1];
    } else if (len === 2) {
      return "one of ".concat(thing, " ").concat(expected[0], " or ").concat(expected[1]);
    } else {
      return "of ".concat(thing, " ").concat(expected[0]);
    }
  } else {
    return "of ".concat(thing, " ").concat(String(expected));
  }
} // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/startsWith


function startsWith(str, search, pos) {
  return str.substr(!pos || pos < 0 ? 0 : +pos, search.length) === search;
} // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/endsWith


function endsWith(str, search, this_len) {
  if (this_len === undefined || this_len > str.length) {
    this_len = str.length;
  }

  return str.substring(this_len - search.length, this_len) === search;
} // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/includes


function includes(str, search, start) {
  if (typeof start !== 'number') {
    start = 0;
  }

  if (start + search.length > str.length) {
    return false;
  } else {
    return str.indexOf(search, start) !== -1;
  }
}

createErrorType('ERR_INVALID_OPT_VALUE', function (name, value) {
  return 'The value "' + value + '" is invalid for option "' + name + '"';
}, TypeError);
createErrorType('ERR_INVALID_ARG_TYPE', function (name, expected, actual) {
  // determiner: 'must be' or 'must not be'
  var determiner;

  if (typeof expected === 'string' && startsWith(expected, 'not ')) {
    determiner = 'must not be';
    expected = expected.replace(/^not /, '');
  } else {
    determiner = 'must be';
  }

  var msg;

  if (endsWith(name, ' argument')) {
    // For cases like 'first argument'
    msg = "The ".concat(name, " ").concat(determiner, " ").concat(oneOf(expected, 'type'));
  } else {
    var type = includes(name, '.') ? 'property' : 'argument';
    msg = "The \"".concat(name, "\" ").concat(type, " ").concat(determiner, " ").concat(oneOf(expected, 'type'));
  }

  msg += ". Received type ".concat(typeof actual);
  return msg;
}, TypeError);
createErrorType('ERR_STREAM_PUSH_AFTER_EOF', 'stream.push() after EOF');
createErrorType('ERR_METHOD_NOT_IMPLEMENTED', function (name) {
  return 'The ' + name + ' method is not implemented';
});
createErrorType('ERR_STREAM_PREMATURE_CLOSE', 'Premature close');
createErrorType('ERR_STREAM_DESTROYED', function (name) {
  return 'Cannot call ' + name + ' after a stream was destroyed';
});
createErrorType('ERR_MULTIPLE_CALLBACK', 'Callback called multiple times');
createErrorType('ERR_STREAM_CANNOT_PIPE', 'Cannot pipe, not readable');
createErrorType('ERR_STREAM_WRITE_AFTER_END', 'write after end');
createErrorType('ERR_STREAM_NULL_VALUES', 'May not write null values to stream', TypeError);
createErrorType('ERR_UNKNOWN_ENCODING', function (arg) {
  return 'Unknown encoding: ' + arg;
}, TypeError);
createErrorType('ERR_STREAM_UNSHIFT_AFTER_END_EVENT', 'stream.unshift() after end event');
module.exports.codes = codes;


/***/ }),

/***/ "./node_modules/readable-stream/lib/_stream_duplex.js":
/*!************************************************************!*\
  !*** ./node_modules/readable-stream/lib/_stream_duplex.js ***!
  \************************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";
/* provided dependency */ var process = __webpack_require__(/*! process/browser */ "./node_modules/process/browser.js");
// Copyright Joyent, Inc. and other Node contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to permit
// persons to whom the Software is furnished to do so, subject to the
// following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN
// NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
// USE OR OTHER DEALINGS IN THE SOFTWARE.

// a duplex stream is just a stream that is both readable and writable.
// Since JS doesn't have multiple prototypal inheritance, this class
// prototypally inherits from Readable, and then parasitically from
// Writable.



/*<replacement>*/
var objectKeys = Object.keys || function (obj) {
  var keys = [];
  for (var key in obj) keys.push(key);
  return keys;
};
/*</replacement>*/

module.exports = Duplex;
var Readable = __webpack_require__(/*! ./_stream_readable */ "./node_modules/readable-stream/lib/_stream_readable.js");
var Writable = __webpack_require__(/*! ./_stream_writable */ "./node_modules/readable-stream/lib/_stream_writable.js");
__webpack_require__(/*! inherits */ "./node_modules/inherits/inherits_browser.js")(Duplex, Readable);
{
  // Allow the keys array to be GC'ed.
  var keys = objectKeys(Writable.prototype);
  for (var v = 0; v < keys.length; v++) {
    var method = keys[v];
    if (!Duplex.prototype[method]) Duplex.prototype[method] = Writable.prototype[method];
  }
}
function Duplex(options) {
  if (!(this instanceof Duplex)) return new Duplex(options);
  Readable.call(this, options);
  Writable.call(this, options);
  this.allowHalfOpen = true;
  if (options) {
    if (options.readable === false) this.readable = false;
    if (options.writable === false) this.writable = false;
    if (options.allowHalfOpen === false) {
      this.allowHalfOpen = false;
      this.once('end', onend);
    }
  }
}
Object.defineProperty(Duplex.prototype, 'writableHighWaterMark', {
  // making it explicit this property is not enumerable
  // because otherwise some prototype manipulation in
  // userland will fail
  enumerable: false,
  get: function get() {
    return this._writableState.highWaterMark;
  }
});
Object.defineProperty(Duplex.prototype, 'writableBuffer', {
  // making it explicit this property is not enumerable
  // because otherwise some prototype manipulation in
  // userland will fail
  enumerable: false,
  get: function get() {
    return this._writableState && this._writableState.getBuffer();
  }
});
Object.defineProperty(Duplex.prototype, 'writableLength', {
  // making it explicit this property is not enumerable
  // because otherwise some prototype manipulation in
  // userland will fail
  enumerable: false,
  get: function get() {
    return this._writableState.length;
  }
});

// the no-half-open enforcer
function onend() {
  // If the writable side ended, then we're ok.
  if (this._writableState.ended) return;

  // no more data can be written.
  // But allow more writes to happen in this tick.
  process.nextTick(onEndNT, this);
}
function onEndNT(self) {
  self.end();
}
Object.defineProperty(Duplex.prototype, 'destroyed', {
  // making it explicit this property is not enumerable
  // because otherwise some prototype manipulation in
  // userland will fail
  enumerable: false,
  get: function get() {
    if (this._readableState === undefined || this._writableState === undefined) {
      return false;
    }
    return this._readableState.destroyed && this._writableState.destroyed;
  },
  set: function set(value) {
    // we ignore the value if the stream
    // has not been initialized yet
    if (this._readableState === undefined || this._writableState === undefined) {
      return;
    }

    // backward compatibility, the user is explicitly
    // managing destroyed
    this._readableState.destroyed = value;
    this._writableState.destroyed = value;
  }
});

/***/ }),

/***/ "./node_modules/readable-stream/lib/_stream_passthrough.js":
/*!*****************************************************************!*\
  !*** ./node_modules/readable-stream/lib/_stream_passthrough.js ***!
  \*****************************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";
// Copyright Joyent, Inc. and other Node contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to permit
// persons to whom the Software is furnished to do so, subject to the
// following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN
// NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
// USE OR OTHER DEALINGS IN THE SOFTWARE.

// a passthrough stream.
// basically just the most minimal sort of Transform stream.
// Every written chunk gets output as-is.



module.exports = PassThrough;
var Transform = __webpack_require__(/*! ./_stream_transform */ "./node_modules/readable-stream/lib/_stream_transform.js");
__webpack_require__(/*! inherits */ "./node_modules/inherits/inherits_browser.js")(PassThrough, Transform);
function PassThrough(options) {
  if (!(this instanceof PassThrough)) return new PassThrough(options);
  Transform.call(this, options);
}
PassThrough.prototype._transform = function (chunk, encoding, cb) {
  cb(null, chunk);
};

/***/ }),

/***/ "./node_modules/readable-stream/lib/_stream_readable.js":
/*!**************************************************************!*\
  !*** ./node_modules/readable-stream/lib/_stream_readable.js ***!
  \**************************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";
/* provided dependency */ var process = __webpack_require__(/*! process/browser */ "./node_modules/process/browser.js");
// Copyright Joyent, Inc. and other Node contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to permit
// persons to whom the Software is furnished to do so, subject to the
// following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN
// NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
// USE OR OTHER DEALINGS IN THE SOFTWARE.



module.exports = Readable;

/*<replacement>*/
var Duplex;
/*</replacement>*/

Readable.ReadableState = ReadableState;

/*<replacement>*/
var EE = (__webpack_require__(/*! events */ "./node_modules/events/events.js").EventEmitter);
var EElistenerCount = function EElistenerCount(emitter, type) {
  return emitter.listeners(type).length;
};
/*</replacement>*/

/*<replacement>*/
var Stream = __webpack_require__(/*! ./internal/streams/stream */ "./node_modules/readable-stream/lib/internal/streams/stream-browser.js");
/*</replacement>*/

var Buffer = (__webpack_require__(/*! buffer */ "./node_modules/buffer/index.js").Buffer);
var OurUint8Array = (typeof __webpack_require__.g !== 'undefined' ? __webpack_require__.g : typeof window !== 'undefined' ? window : typeof self !== 'undefined' ? self : {}).Uint8Array || function () {};
function _uint8ArrayToBuffer(chunk) {
  return Buffer.from(chunk);
}
function _isUint8Array(obj) {
  return Buffer.isBuffer(obj) || obj instanceof OurUint8Array;
}

/*<replacement>*/
var debugUtil = __webpack_require__(/*! util */ "?d17e");
var debug;
if (debugUtil && debugUtil.debuglog) {
  debug = debugUtil.debuglog('stream');
} else {
  debug = function debug() {};
}
/*</replacement>*/

var BufferList = __webpack_require__(/*! ./internal/streams/buffer_list */ "./node_modules/readable-stream/lib/internal/streams/buffer_list.js");
var destroyImpl = __webpack_require__(/*! ./internal/streams/destroy */ "./node_modules/readable-stream/lib/internal/streams/destroy.js");
var _require = __webpack_require__(/*! ./internal/streams/state */ "./node_modules/readable-stream/lib/internal/streams/state.js"),
  getHighWaterMark = _require.getHighWaterMark;
var _require$codes = (__webpack_require__(/*! ../errors */ "./node_modules/readable-stream/errors-browser.js").codes),
  ERR_INVALID_ARG_TYPE = _require$codes.ERR_INVALID_ARG_TYPE,
  ERR_STREAM_PUSH_AFTER_EOF = _require$codes.ERR_STREAM_PUSH_AFTER_EOF,
  ERR_METHOD_NOT_IMPLEMENTED = _require$codes.ERR_METHOD_NOT_IMPLEMENTED,
  ERR_STREAM_UNSHIFT_AFTER_END_EVENT = _require$codes.ERR_STREAM_UNSHIFT_AFTER_END_EVENT;

// Lazy loaded to improve the startup performance.
var StringDecoder;
var createReadableStreamAsyncIterator;
var from;
__webpack_require__(/*! inherits */ "./node_modules/inherits/inherits_browser.js")(Readable, Stream);
var errorOrDestroy = destroyImpl.errorOrDestroy;
var kProxyEvents = ['error', 'close', 'destroy', 'pause', 'resume'];
function prependListener(emitter, event, fn) {
  // Sadly this is not cacheable as some libraries bundle their own
  // event emitter implementation with them.
  if (typeof emitter.prependListener === 'function') return emitter.prependListener(event, fn);

  // This is a hack to make sure that our error handler is attached before any
  // userland ones.  NEVER DO THIS. This is here only because this code needs
  // to continue to work with older versions of Node.js that do not include
  // the prependListener() method. The goal is to eventually remove this hack.
  if (!emitter._events || !emitter._events[event]) emitter.on(event, fn);else if (Array.isArray(emitter._events[event])) emitter._events[event].unshift(fn);else emitter._events[event] = [fn, emitter._events[event]];
}
function ReadableState(options, stream, isDuplex) {
  Duplex = Duplex || __webpack_require__(/*! ./_stream_duplex */ "./node_modules/readable-stream/lib/_stream_duplex.js");
  options = options || {};

  // Duplex streams are both readable and writable, but share
  // the same options object.
  // However, some cases require setting options to different
  // values for the readable and the writable sides of the duplex stream.
  // These options can be provided separately as readableXXX and writableXXX.
  if (typeof isDuplex !== 'boolean') isDuplex = stream instanceof Duplex;

  // object stream flag. Used to make read(n) ignore n and to
  // make all the buffer merging and length checks go away
  this.objectMode = !!options.objectMode;
  if (isDuplex) this.objectMode = this.objectMode || !!options.readableObjectMode;

  // the point at which it stops calling _read() to fill the buffer
  // Note: 0 is a valid value, means "don't call _read preemptively ever"
  this.highWaterMark = getHighWaterMark(this, options, 'readableHighWaterMark', isDuplex);

  // A linked list is used to store data chunks instead of an array because the
  // linked list can remove elements from the beginning faster than
  // array.shift()
  this.buffer = new BufferList();
  this.length = 0;
  this.pipes = null;
  this.pipesCount = 0;
  this.flowing = null;
  this.ended = false;
  this.endEmitted = false;
  this.reading = false;

  // a flag to be able to tell if the event 'readable'/'data' is emitted
  // immediately, or on a later tick.  We set this to true at first, because
  // any actions that shouldn't happen until "later" should generally also
  // not happen before the first read call.
  this.sync = true;

  // whenever we return null, then we set a flag to say
  // that we're awaiting a 'readable' event emission.
  this.needReadable = false;
  this.emittedReadable = false;
  this.readableListening = false;
  this.resumeScheduled = false;
  this.paused = true;

  // Should close be emitted on destroy. Defaults to true.
  this.emitClose = options.emitClose !== false;

  // Should .destroy() be called after 'end' (and potentially 'finish')
  this.autoDestroy = !!options.autoDestroy;

  // has it been destroyed
  this.destroyed = false;

  // Crypto is kind of old and crusty.  Historically, its default string
  // encoding is 'binary' so we have to make this configurable.
  // Everything else in the universe uses 'utf8', though.
  this.defaultEncoding = options.defaultEncoding || 'utf8';

  // the number of writers that are awaiting a drain event in .pipe()s
  this.awaitDrain = 0;

  // if true, a maybeReadMore has been scheduled
  this.readingMore = false;
  this.decoder = null;
  this.encoding = null;
  if (options.encoding) {
    if (!StringDecoder) StringDecoder = (__webpack_require__(/*! string_decoder/ */ "./node_modules/string_decoder/lib/string_decoder.js").StringDecoder);
    this.decoder = new StringDecoder(options.encoding);
    this.encoding = options.encoding;
  }
}
function Readable(options) {
  Duplex = Duplex || __webpack_require__(/*! ./_stream_duplex */ "./node_modules/readable-stream/lib/_stream_duplex.js");
  if (!(this instanceof Readable)) return new Readable(options);

  // Checking for a Stream.Duplex instance is faster here instead of inside
  // the ReadableState constructor, at least with V8 6.5
  var isDuplex = this instanceof Duplex;
  this._readableState = new ReadableState(options, this, isDuplex);

  // legacy
  this.readable = true;
  if (options) {
    if (typeof options.read === 'function') this._read = options.read;
    if (typeof options.destroy === 'function') this._destroy = options.destroy;
  }
  Stream.call(this);
}
Object.defineProperty(Readable.prototype, 'destroyed', {
  // making it explicit this property is not enumerable
  // because otherwise some prototype manipulation in
  // userland will fail
  enumerable: false,
  get: function get() {
    if (this._readableState === undefined) {
      return false;
    }
    return this._readableState.destroyed;
  },
  set: function set(value) {
    // we ignore the value if the stream
    // has not been initialized yet
    if (!this._readableState) {
      return;
    }

    // backward compatibility, the user is explicitly
    // managing destroyed
    this._readableState.destroyed = value;
  }
});
Readable.prototype.destroy = destroyImpl.destroy;
Readable.prototype._undestroy = destroyImpl.undestroy;
Readable.prototype._destroy = function (err, cb) {
  cb(err);
};

// Manually shove something into the read() buffer.
// This returns true if the highWaterMark has not been hit yet,
// similar to how Writable.write() returns true if you should
// write() some more.
Readable.prototype.push = function (chunk, encoding) {
  var state = this._readableState;
  var skipChunkCheck;
  if (!state.objectMode) {
    if (typeof chunk === 'string') {
      encoding = encoding || state.defaultEncoding;
      if (encoding !== state.encoding) {
        chunk = Buffer.from(chunk, encoding);
        encoding = '';
      }
      skipChunkCheck = true;
    }
  } else {
    skipChunkCheck = true;
  }
  return readableAddChunk(this, chunk, encoding, false, skipChunkCheck);
};

// Unshift should *always* be something directly out of read()
Readable.prototype.unshift = function (chunk) {
  return readableAddChunk(this, chunk, null, true, false);
};
function readableAddChunk(stream, chunk, encoding, addToFront, skipChunkCheck) {
  debug('readableAddChunk', chunk);
  var state = stream._readableState;
  if (chunk === null) {
    state.reading = false;
    onEofChunk(stream, state);
  } else {
    var er;
    if (!skipChunkCheck) er = chunkInvalid(state, chunk);
    if (er) {
      errorOrDestroy(stream, er);
    } else if (state.objectMode || chunk && chunk.length > 0) {
      if (typeof chunk !== 'string' && !state.objectMode && Object.getPrototypeOf(chunk) !== Buffer.prototype) {
        chunk = _uint8ArrayToBuffer(chunk);
      }
      if (addToFront) {
        if (state.endEmitted) errorOrDestroy(stream, new ERR_STREAM_UNSHIFT_AFTER_END_EVENT());else addChunk(stream, state, chunk, true);
      } else if (state.ended) {
        errorOrDestroy(stream, new ERR_STREAM_PUSH_AFTER_EOF());
      } else if (state.destroyed) {
        return false;
      } else {
        state.reading = false;
        if (state.decoder && !encoding) {
          chunk = state.decoder.write(chunk);
          if (state.objectMode || chunk.length !== 0) addChunk(stream, state, chunk, false);else maybeReadMore(stream, state);
        } else {
          addChunk(stream, state, chunk, false);
        }
      }
    } else if (!addToFront) {
      state.reading = false;
      maybeReadMore(stream, state);
    }
  }

  // We can push more data if we are below the highWaterMark.
  // Also, if we have no data yet, we can stand some more bytes.
  // This is to work around cases where hwm=0, such as the repl.
  return !state.ended && (state.length < state.highWaterMark || state.length === 0);
}
function addChunk(stream, state, chunk, addToFront) {
  if (state.flowing && state.length === 0 && !state.sync) {
    state.awaitDrain = 0;
    stream.emit('data', chunk);
  } else {
    // update the buffer info.
    state.length += state.objectMode ? 1 : chunk.length;
    if (addToFront) state.buffer.unshift(chunk);else state.buffer.push(chunk);
    if (state.needReadable) emitReadable(stream);
  }
  maybeReadMore(stream, state);
}
function chunkInvalid(state, chunk) {
  var er;
  if (!_isUint8Array(chunk) && typeof chunk !== 'string' && chunk !== undefined && !state.objectMode) {
    er = new ERR_INVALID_ARG_TYPE('chunk', ['string', 'Buffer', 'Uint8Array'], chunk);
  }
  return er;
}
Readable.prototype.isPaused = function () {
  return this._readableState.flowing === false;
};

// backwards compatibility.
Readable.prototype.setEncoding = function (enc) {
  if (!StringDecoder) StringDecoder = (__webpack_require__(/*! string_decoder/ */ "./node_modules/string_decoder/lib/string_decoder.js").StringDecoder);
  var decoder = new StringDecoder(enc);
  this._readableState.decoder = decoder;
  // If setEncoding(null), decoder.encoding equals utf8
  this._readableState.encoding = this._readableState.decoder.encoding;

  // Iterate over current buffer to convert already stored Buffers:
  var p = this._readableState.buffer.head;
  var content = '';
  while (p !== null) {
    content += decoder.write(p.data);
    p = p.next;
  }
  this._readableState.buffer.clear();
  if (content !== '') this._readableState.buffer.push(content);
  this._readableState.length = content.length;
  return this;
};

// Don't raise the hwm > 1GB
var MAX_HWM = 0x40000000;
function computeNewHighWaterMark(n) {
  if (n >= MAX_HWM) {
    // TODO(ronag): Throw ERR_VALUE_OUT_OF_RANGE.
    n = MAX_HWM;
  } else {
    // Get the next highest power of 2 to prevent increasing hwm excessively in
    // tiny amounts
    n--;
    n |= n >>> 1;
    n |= n >>> 2;
    n |= n >>> 4;
    n |= n >>> 8;
    n |= n >>> 16;
    n++;
  }
  return n;
}

// This function is designed to be inlinable, so please take care when making
// changes to the function body.
function howMuchToRead(n, state) {
  if (n <= 0 || state.length === 0 && state.ended) return 0;
  if (state.objectMode) return 1;
  if (n !== n) {
    // Only flow one buffer at a time
    if (state.flowing && state.length) return state.buffer.head.data.length;else return state.length;
  }
  // If we're asking for more than the current hwm, then raise the hwm.
  if (n > state.highWaterMark) state.highWaterMark = computeNewHighWaterMark(n);
  if (n <= state.length) return n;
  // Don't have enough
  if (!state.ended) {
    state.needReadable = true;
    return 0;
  }
  return state.length;
}

// you can override either this method, or the async _read(n) below.
Readable.prototype.read = function (n) {
  debug('read', n);
  n = parseInt(n, 10);
  var state = this._readableState;
  var nOrig = n;
  if (n !== 0) state.emittedReadable = false;

  // if we're doing read(0) to trigger a readable event, but we
  // already have a bunch of data in the buffer, then just trigger
  // the 'readable' event and move on.
  if (n === 0 && state.needReadable && ((state.highWaterMark !== 0 ? state.length >= state.highWaterMark : state.length > 0) || state.ended)) {
    debug('read: emitReadable', state.length, state.ended);
    if (state.length === 0 && state.ended) endReadable(this);else emitReadable(this);
    return null;
  }
  n = howMuchToRead(n, state);

  // if we've ended, and we're now clear, then finish it up.
  if (n === 0 && state.ended) {
    if (state.length === 0) endReadable(this);
    return null;
  }

  // All the actual chunk generation logic needs to be
  // *below* the call to _read.  The reason is that in certain
  // synthetic stream cases, such as passthrough streams, _read
  // may be a completely synchronous operation which may change
  // the state of the read buffer, providing enough data when
  // before there was *not* enough.
  //
  // So, the steps are:
  // 1. Figure out what the state of things will be after we do
  // a read from the buffer.
  //
  // 2. If that resulting state will trigger a _read, then call _read.
  // Note that this may be asynchronous, or synchronous.  Yes, it is
  // deeply ugly to write APIs this way, but that still doesn't mean
  // that the Readable class should behave improperly, as streams are
  // designed to be sync/async agnostic.
  // Take note if the _read call is sync or async (ie, if the read call
  // has returned yet), so that we know whether or not it's safe to emit
  // 'readable' etc.
  //
  // 3. Actually pull the requested chunks out of the buffer and return.

  // if we need a readable event, then we need to do some reading.
  var doRead = state.needReadable;
  debug('need readable', doRead);

  // if we currently have less than the highWaterMark, then also read some
  if (state.length === 0 || state.length - n < state.highWaterMark) {
    doRead = true;
    debug('length less than watermark', doRead);
  }

  // however, if we've ended, then there's no point, and if we're already
  // reading, then it's unnecessary.
  if (state.ended || state.reading) {
    doRead = false;
    debug('reading or ended', doRead);
  } else if (doRead) {
    debug('do read');
    state.reading = true;
    state.sync = true;
    // if the length is currently zero, then we *need* a readable event.
    if (state.length === 0) state.needReadable = true;
    // call internal read method
    this._read(state.highWaterMark);
    state.sync = false;
    // If _read pushed data synchronously, then `reading` will be false,
    // and we need to re-evaluate how much data we can return to the user.
    if (!state.reading) n = howMuchToRead(nOrig, state);
  }
  var ret;
  if (n > 0) ret = fromList(n, state);else ret = null;
  if (ret === null) {
    state.needReadable = state.length <= state.highWaterMark;
    n = 0;
  } else {
    state.length -= n;
    state.awaitDrain = 0;
  }
  if (state.length === 0) {
    // If we have nothing in the buffer, then we want to know
    // as soon as we *do* get something into the buffer.
    if (!state.ended) state.needReadable = true;

    // If we tried to read() past the EOF, then emit end on the next tick.
    if (nOrig !== n && state.ended) endReadable(this);
  }
  if (ret !== null) this.emit('data', ret);
  return ret;
};
function onEofChunk(stream, state) {
  debug('onEofChunk');
  if (state.ended) return;
  if (state.decoder) {
    var chunk = state.decoder.end();
    if (chunk && chunk.length) {
      state.buffer.push(chunk);
      state.length += state.objectMode ? 1 : chunk.length;
    }
  }
  state.ended = true;
  if (state.sync) {
    // if we are sync, wait until next tick to emit the data.
    // Otherwise we risk emitting data in the flow()
    // the readable code triggers during a read() call
    emitReadable(stream);
  } else {
    // emit 'readable' now to make sure it gets picked up.
    state.needReadable = false;
    if (!state.emittedReadable) {
      state.emittedReadable = true;
      emitReadable_(stream);
    }
  }
}

// Don't emit readable right away in sync mode, because this can trigger
// another read() call => stack overflow.  This way, it might trigger
// a nextTick recursion warning, but that's not so bad.
function emitReadable(stream) {
  var state = stream._readableState;
  debug('emitReadable', state.needReadable, state.emittedReadable);
  state.needReadable = false;
  if (!state.emittedReadable) {
    debug('emitReadable', state.flowing);
    state.emittedReadable = true;
    process.nextTick(emitReadable_, stream);
  }
}
function emitReadable_(stream) {
  var state = stream._readableState;
  debug('emitReadable_', state.destroyed, state.length, state.ended);
  if (!state.destroyed && (state.length || state.ended)) {
    stream.emit('readable');
    state.emittedReadable = false;
  }

  // The stream needs another readable event if
  // 1. It is not flowing, as the flow mechanism will take
  //    care of it.
  // 2. It is not ended.
  // 3. It is below the highWaterMark, so we can schedule
  //    another readable later.
  state.needReadable = !state.flowing && !state.ended && state.length <= state.highWaterMark;
  flow(stream);
}

// at this point, the user has presumably seen the 'readable' event,
// and called read() to consume some data.  that may have triggered
// in turn another _read(n) call, in which case reading = true if
// it's in progress.
// However, if we're not ended, or reading, and the length < hwm,
// then go ahead and try to read some more preemptively.
function maybeReadMore(stream, state) {
  if (!state.readingMore) {
    state.readingMore = true;
    process.nextTick(maybeReadMore_, stream, state);
  }
}
function maybeReadMore_(stream, state) {
  // Attempt to read more data if we should.
  //
  // The conditions for reading more data are (one of):
  // - Not enough data buffered (state.length < state.highWaterMark). The loop
  //   is responsible for filling the buffer with enough data if such data
  //   is available. If highWaterMark is 0 and we are not in the flowing mode
  //   we should _not_ attempt to buffer any extra data. We'll get more data
  //   when the stream consumer calls read() instead.
  // - No data in the buffer, and the stream is in flowing mode. In this mode
  //   the loop below is responsible for ensuring read() is called. Failing to
  //   call read here would abort the flow and there's no other mechanism for
  //   continuing the flow if the stream consumer has just subscribed to the
  //   'data' event.
  //
  // In addition to the above conditions to keep reading data, the following
  // conditions prevent the data from being read:
  // - The stream has ended (state.ended).
  // - There is already a pending 'read' operation (state.reading). This is a
  //   case where the the stream has called the implementation defined _read()
  //   method, but they are processing the call asynchronously and have _not_
  //   called push() with new data. In this case we skip performing more
  //   read()s. The execution ends in this method again after the _read() ends
  //   up calling push() with more data.
  while (!state.reading && !state.ended && (state.length < state.highWaterMark || state.flowing && state.length === 0)) {
    var len = state.length;
    debug('maybeReadMore read 0');
    stream.read(0);
    if (len === state.length)
      // didn't get any data, stop spinning.
      break;
  }
  state.readingMore = false;
}

// abstract method.  to be overridden in specific implementation classes.
// call cb(er, data) where data is <= n in length.
// for virtual (non-string, non-buffer) streams, "length" is somewhat
// arbitrary, and perhaps not very meaningful.
Readable.prototype._read = function (n) {
  errorOrDestroy(this, new ERR_METHOD_NOT_IMPLEMENTED('_read()'));
};
Readable.prototype.pipe = function (dest, pipeOpts) {
  var src = this;
  var state = this._readableState;
  switch (state.pipesCount) {
    case 0:
      state.pipes = dest;
      break;
    case 1:
      state.pipes = [state.pipes, dest];
      break;
    default:
      state.pipes.push(dest);
      break;
  }
  state.pipesCount += 1;
  debug('pipe count=%d opts=%j', state.pipesCount, pipeOpts);
  var doEnd = (!pipeOpts || pipeOpts.end !== false) && dest !== process.stdout && dest !== process.stderr;
  var endFn = doEnd ? onend : unpipe;
  if (state.endEmitted) process.nextTick(endFn);else src.once('end', endFn);
  dest.on('unpipe', onunpipe);
  function onunpipe(readable, unpipeInfo) {
    debug('onunpipe');
    if (readable === src) {
      if (unpipeInfo && unpipeInfo.hasUnpiped === false) {
        unpipeInfo.hasUnpiped = true;
        cleanup();
      }
    }
  }
  function onend() {
    debug('onend');
    dest.end();
  }

  // when the dest drains, it reduces the awaitDrain counter
  // on the source.  This would be more elegant with a .once()
  // handler in flow(), but adding and removing repeatedly is
  // too slow.
  var ondrain = pipeOnDrain(src);
  dest.on('drain', ondrain);
  var cleanedUp = false;
  function cleanup() {
    debug('cleanup');
    // cleanup event handlers once the pipe is broken
    dest.removeListener('close', onclose);
    dest.removeListener('finish', onfinish);
    dest.removeListener('drain', ondrain);
    dest.removeListener('error', onerror);
    dest.removeListener('unpipe', onunpipe);
    src.removeListener('end', onend);
    src.removeListener('end', unpipe);
    src.removeListener('data', ondata);
    cleanedUp = true;

    // if the reader is waiting for a drain event from this
    // specific writer, then it would cause it to never start
    // flowing again.
    // So, if this is awaiting a drain, then we just call it now.
    // If we don't know, then assume that we are waiting for one.
    if (state.awaitDrain && (!dest._writableState || dest._writableState.needDrain)) ondrain();
  }
  src.on('data', ondata);
  function ondata(chunk) {
    debug('ondata');
    var ret = dest.write(chunk);
    debug('dest.write', ret);
    if (ret === false) {
      // If the user unpiped during `dest.write()`, it is possible
      // to get stuck in a permanently paused state if that write
      // also returned false.
      // => Check whether `dest` is still a piping destination.
      if ((state.pipesCount === 1 && state.pipes === dest || state.pipesCount > 1 && indexOf(state.pipes, dest) !== -1) && !cleanedUp) {
        debug('false write response, pause', state.awaitDrain);
        state.awaitDrain++;
      }
      src.pause();
    }
  }

  // if the dest has an error, then stop piping into it.
  // however, don't suppress the throwing behavior for this.
  function onerror(er) {
    debug('onerror', er);
    unpipe();
    dest.removeListener('error', onerror);
    if (EElistenerCount(dest, 'error') === 0) errorOrDestroy(dest, er);
  }

  // Make sure our error handler is attached before userland ones.
  prependListener(dest, 'error', onerror);

  // Both close and finish should trigger unpipe, but only once.
  function onclose() {
    dest.removeListener('finish', onfinish);
    unpipe();
  }
  dest.once('close', onclose);
  function onfinish() {
    debug('onfinish');
    dest.removeListener('close', onclose);
    unpipe();
  }
  dest.once('finish', onfinish);
  function unpipe() {
    debug('unpipe');
    src.unpipe(dest);
  }

  // tell the dest that it's being piped to
  dest.emit('pipe', src);

  // start the flow if it hasn't been started already.
  if (!state.flowing) {
    debug('pipe resume');
    src.resume();
  }
  return dest;
};
function pipeOnDrain(src) {
  return function pipeOnDrainFunctionResult() {
    var state = src._readableState;
    debug('pipeOnDrain', state.awaitDrain);
    if (state.awaitDrain) state.awaitDrain--;
    if (state.awaitDrain === 0 && EElistenerCount(src, 'data')) {
      state.flowing = true;
      flow(src);
    }
  };
}
Readable.prototype.unpipe = function (dest) {
  var state = this._readableState;
  var unpipeInfo = {
    hasUnpiped: false
  };

  // if we're not piping anywhere, then do nothing.
  if (state.pipesCount === 0) return this;

  // just one destination.  most common case.
  if (state.pipesCount === 1) {
    // passed in one, but it's not the right one.
    if (dest && dest !== state.pipes) return this;
    if (!dest) dest = state.pipes;

    // got a match.
    state.pipes = null;
    state.pipesCount = 0;
    state.flowing = false;
    if (dest) dest.emit('unpipe', this, unpipeInfo);
    return this;
  }

  // slow case. multiple pipe destinations.

  if (!dest) {
    // remove all.
    var dests = state.pipes;
    var len = state.pipesCount;
    state.pipes = null;
    state.pipesCount = 0;
    state.flowing = false;
    for (var i = 0; i < len; i++) dests[i].emit('unpipe', this, {
      hasUnpiped: false
    });
    return this;
  }

  // try to find the right one.
  var index = indexOf(state.pipes, dest);
  if (index === -1) return this;
  state.pipes.splice(index, 1);
  state.pipesCount -= 1;
  if (state.pipesCount === 1) state.pipes = state.pipes[0];
  dest.emit('unpipe', this, unpipeInfo);
  return this;
};

// set up data events if they are asked for
// Ensure readable listeners eventually get something
Readable.prototype.on = function (ev, fn) {
  var res = Stream.prototype.on.call(this, ev, fn);
  var state = this._readableState;
  if (ev === 'data') {
    // update readableListening so that resume() may be a no-op
    // a few lines down. This is needed to support once('readable').
    state.readableListening = this.listenerCount('readable') > 0;

    // Try start flowing on next tick if stream isn't explicitly paused
    if (state.flowing !== false) this.resume();
  } else if (ev === 'readable') {
    if (!state.endEmitted && !state.readableListening) {
      state.readableListening = state.needReadable = true;
      state.flowing = false;
      state.emittedReadable = false;
      debug('on readable', state.length, state.reading);
      if (state.length) {
        emitReadable(this);
      } else if (!state.reading) {
        process.nextTick(nReadingNextTick, this);
      }
    }
  }
  return res;
};
Readable.prototype.addListener = Readable.prototype.on;
Readable.prototype.removeListener = function (ev, fn) {
  var res = Stream.prototype.removeListener.call(this, ev, fn);
  if (ev === 'readable') {
    // We need to check if there is someone still listening to
    // readable and reset the state. However this needs to happen
    // after readable has been emitted but before I/O (nextTick) to
    // support once('readable', fn) cycles. This means that calling
    // resume within the same tick will have no
    // effect.
    process.nextTick(updateReadableListening, this);
  }
  return res;
};
Readable.prototype.removeAllListeners = function (ev) {
  var res = Stream.prototype.removeAllListeners.apply(this, arguments);
  if (ev === 'readable' || ev === undefined) {
    // We need to check if there is someone still listening to
    // readable and reset the state. However this needs to happen
    // after readable has been emitted but before I/O (nextTick) to
    // support once('readable', fn) cycles. This means that calling
    // resume within the same tick will have no
    // effect.
    process.nextTick(updateReadableListening, this);
  }
  return res;
};
function updateReadableListening(self) {
  var state = self._readableState;
  state.readableListening = self.listenerCount('readable') > 0;
  if (state.resumeScheduled && !state.paused) {
    // flowing needs to be set to true now, otherwise
    // the upcoming resume will not flow.
    state.flowing = true;

    // crude way to check if we should resume
  } else if (self.listenerCount('data') > 0) {
    self.resume();
  }
}
function nReadingNextTick(self) {
  debug('readable nexttick read 0');
  self.read(0);
}

// pause() and resume() are remnants of the legacy readable stream API
// If the user uses them, then switch into old mode.
Readable.prototype.resume = function () {
  var state = this._readableState;
  if (!state.flowing) {
    debug('resume');
    // we flow only if there is no one listening
    // for readable, but we still have to call
    // resume()
    state.flowing = !state.readableListening;
    resume(this, state);
  }
  state.paused = false;
  return this;
};
function resume(stream, state) {
  if (!state.resumeScheduled) {
    state.resumeScheduled = true;
    process.nextTick(resume_, stream, state);
  }
}
function resume_(stream, state) {
  debug('resume', state.reading);
  if (!state.reading) {
    stream.read(0);
  }
  state.resumeScheduled = false;
  stream.emit('resume');
  flow(stream);
  if (state.flowing && !state.reading) stream.read(0);
}
Readable.prototype.pause = function () {
  debug('call pause flowing=%j', this._readableState.flowing);
  if (this._readableState.flowing !== false) {
    debug('pause');
    this._readableState.flowing = false;
    this.emit('pause');
  }
  this._readableState.paused = true;
  return this;
};
function flow(stream) {
  var state = stream._readableState;
  debug('flow', state.flowing);
  while (state.flowing && stream.read() !== null);
}

// wrap an old-style stream as the async data source.
// This is *not* part of the readable stream interface.
// It is an ugly unfortunate mess of history.
Readable.prototype.wrap = function (stream) {
  var _this = this;
  var state = this._readableState;
  var paused = false;
  stream.on('end', function () {
    debug('wrapped end');
    if (state.decoder && !state.ended) {
      var chunk = state.decoder.end();
      if (chunk && chunk.length) _this.push(chunk);
    }
    _this.push(null);
  });
  stream.on('data', function (chunk) {
    debug('wrapped data');
    if (state.decoder) chunk = state.decoder.write(chunk);

    // don't skip over falsy values in objectMode
    if (state.objectMode && (chunk === null || chunk === undefined)) return;else if (!state.objectMode && (!chunk || !chunk.length)) return;
    var ret = _this.push(chunk);
    if (!ret) {
      paused = true;
      stream.pause();
    }
  });

  // proxy all the other methods.
  // important when wrapping filters and duplexes.
  for (var i in stream) {
    if (this[i] === undefined && typeof stream[i] === 'function') {
      this[i] = function methodWrap(method) {
        return function methodWrapReturnFunction() {
          return stream[method].apply(stream, arguments);
        };
      }(i);
    }
  }

  // proxy certain important events.
  for (var n = 0; n < kProxyEvents.length; n++) {
    stream.on(kProxyEvents[n], this.emit.bind(this, kProxyEvents[n]));
  }

  // when we try to consume some more bytes, simply unpause the
  // underlying stream.
  this._read = function (n) {
    debug('wrapped _read', n);
    if (paused) {
      paused = false;
      stream.resume();
    }
  };
  return this;
};
if (typeof Symbol === 'function') {
  Readable.prototype[Symbol.asyncIterator] = function () {
    if (createReadableStreamAsyncIterator === undefined) {
      createReadableStreamAsyncIterator = __webpack_require__(/*! ./internal/streams/async_iterator */ "./node_modules/readable-stream/lib/internal/streams/async_iterator.js");
    }
    return createReadableStreamAsyncIterator(this);
  };
}
Object.defineProperty(Readable.prototype, 'readableHighWaterMark', {
  // making it explicit this property is not enumerable
  // because otherwise some prototype manipulation in
  // userland will fail
  enumerable: false,
  get: function get() {
    return this._readableState.highWaterMark;
  }
});
Object.defineProperty(Readable.prototype, 'readableBuffer', {
  // making it explicit this property is not enumerable
  // because otherwise some prototype manipulation in
  // userland will fail
  enumerable: false,
  get: function get() {
    return this._readableState && this._readableState.buffer;
  }
});
Object.defineProperty(Readable.prototype, 'readableFlowing', {
  // making it explicit this property is not enumerable
  // because otherwise some prototype manipulation in
  // userland will fail
  enumerable: false,
  get: function get() {
    return this._readableState.flowing;
  },
  set: function set(state) {
    if (this._readableState) {
      this._readableState.flowing = state;
    }
  }
});

// exposed for testing purposes only.
Readable._fromList = fromList;
Object.defineProperty(Readable.prototype, 'readableLength', {
  // making it explicit this property is not enumerable
  // because otherwise some prototype manipulation in
  // userland will fail
  enumerable: false,
  get: function get() {
    return this._readableState.length;
  }
});

// Pluck off n bytes from an array of buffers.
// Length is the combined lengths of all the buffers in the list.
// This function is designed to be inlinable, so please take care when making
// changes to the function body.
function fromList(n, state) {
  // nothing buffered
  if (state.length === 0) return null;
  var ret;
  if (state.objectMode) ret = state.buffer.shift();else if (!n || n >= state.length) {
    // read it all, truncate the list
    if (state.decoder) ret = state.buffer.join('');else if (state.buffer.length === 1) ret = state.buffer.first();else ret = state.buffer.concat(state.length);
    state.buffer.clear();
  } else {
    // read part of list
    ret = state.buffer.consume(n, state.decoder);
  }
  return ret;
}
function endReadable(stream) {
  var state = stream._readableState;
  debug('endReadable', state.endEmitted);
  if (!state.endEmitted) {
    state.ended = true;
    process.nextTick(endReadableNT, state, stream);
  }
}
function endReadableNT(state, stream) {
  debug('endReadableNT', state.endEmitted, state.length);

  // Check that we didn't get one last unshift.
  if (!state.endEmitted && state.length === 0) {
    state.endEmitted = true;
    stream.readable = false;
    stream.emit('end');
    if (state.autoDestroy) {
      // In case of duplex streams we need a way to detect
      // if the writable side is ready for autoDestroy as well
      var wState = stream._writableState;
      if (!wState || wState.autoDestroy && wState.finished) {
        stream.destroy();
      }
    }
  }
}
if (typeof Symbol === 'function') {
  Readable.from = function (iterable, opts) {
    if (from === undefined) {
      from = __webpack_require__(/*! ./internal/streams/from */ "./node_modules/readable-stream/lib/internal/streams/from-browser.js");
    }
    return from(Readable, iterable, opts);
  };
}
function indexOf(xs, x) {
  for (var i = 0, l = xs.length; i < l; i++) {
    if (xs[i] === x) return i;
  }
  return -1;
}

/***/ }),

/***/ "./node_modules/readable-stream/lib/_stream_transform.js":
/*!***************************************************************!*\
  !*** ./node_modules/readable-stream/lib/_stream_transform.js ***!
  \***************************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";
// Copyright Joyent, Inc. and other Node contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to permit
// persons to whom the Software is furnished to do so, subject to the
// following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN
// NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
// USE OR OTHER DEALINGS IN THE SOFTWARE.

// a transform stream is a readable/writable stream where you do
// something with the data.  Sometimes it's called a "filter",
// but that's not a great name for it, since that implies a thing where
// some bits pass through, and others are simply ignored.  (That would
// be a valid example of a transform, of course.)
//
// While the output is causally related to the input, it's not a
// necessarily symmetric or synchronous transformation.  For example,
// a zlib stream might take multiple plain-text writes(), and then
// emit a single compressed chunk some time in the future.
//
// Here's how this works:
//
// The Transform stream has all the aspects of the readable and writable
// stream classes.  When you write(chunk), that calls _write(chunk,cb)
// internally, and returns false if there's a lot of pending writes
// buffered up.  When you call read(), that calls _read(n) until
// there's enough pending readable data buffered up.
//
// In a transform stream, the written data is placed in a buffer.  When
// _read(n) is called, it transforms the queued up data, calling the
// buffered _write cb's as it consumes chunks.  If consuming a single
// written chunk would result in multiple output chunks, then the first
// outputted bit calls the readcb, and subsequent chunks just go into
// the read buffer, and will cause it to emit 'readable' if necessary.
//
// This way, back-pressure is actually determined by the reading side,
// since _read has to be called to start processing a new chunk.  However,
// a pathological inflate type of transform can cause excessive buffering
// here.  For example, imagine a stream where every byte of input is
// interpreted as an integer from 0-255, and then results in that many
// bytes of output.  Writing the 4 bytes {ff,ff,ff,ff} would result in
// 1kb of data being output.  In this case, you could write a very small
// amount of input, and end up with a very large amount of output.  In
// such a pathological inflating mechanism, there'd be no way to tell
// the system to stop doing the transform.  A single 4MB write could
// cause the system to run out of memory.
//
// However, even in such a pathological case, only a single written chunk
// would be consumed, and then the rest would wait (un-transformed) until
// the results of the previous transformed chunk were consumed.



module.exports = Transform;
var _require$codes = (__webpack_require__(/*! ../errors */ "./node_modules/readable-stream/errors-browser.js").codes),
  ERR_METHOD_NOT_IMPLEMENTED = _require$codes.ERR_METHOD_NOT_IMPLEMENTED,
  ERR_MULTIPLE_CALLBACK = _require$codes.ERR_MULTIPLE_CALLBACK,
  ERR_TRANSFORM_ALREADY_TRANSFORMING = _require$codes.ERR_TRANSFORM_ALREADY_TRANSFORMING,
  ERR_TRANSFORM_WITH_LENGTH_0 = _require$codes.ERR_TRANSFORM_WITH_LENGTH_0;
var Duplex = __webpack_require__(/*! ./_stream_duplex */ "./node_modules/readable-stream/lib/_stream_duplex.js");
__webpack_require__(/*! inherits */ "./node_modules/inherits/inherits_browser.js")(Transform, Duplex);
function afterTransform(er, data) {
  var ts = this._transformState;
  ts.transforming = false;
  var cb = ts.writecb;
  if (cb === null) {
    return this.emit('error', new ERR_MULTIPLE_CALLBACK());
  }
  ts.writechunk = null;
  ts.writecb = null;
  if (data != null)
    // single equals check for both `null` and `undefined`
    this.push(data);
  cb(er);
  var rs = this._readableState;
  rs.reading = false;
  if (rs.needReadable || rs.length < rs.highWaterMark) {
    this._read(rs.highWaterMark);
  }
}
function Transform(options) {
  if (!(this instanceof Transform)) return new Transform(options);
  Duplex.call(this, options);
  this._transformState = {
    afterTransform: afterTransform.bind(this),
    needTransform: false,
    transforming: false,
    writecb: null,
    writechunk: null,
    writeencoding: null
  };

  // start out asking for a readable event once data is transformed.
  this._readableState.needReadable = true;

  // we have implemented the _read method, and done the other things
  // that Readable wants before the first _read call, so unset the
  // sync guard flag.
  this._readableState.sync = false;
  if (options) {
    if (typeof options.transform === 'function') this._transform = options.transform;
    if (typeof options.flush === 'function') this._flush = options.flush;
  }

  // When the writable side finishes, then flush out anything remaining.
  this.on('prefinish', prefinish);
}
function prefinish() {
  var _this = this;
  if (typeof this._flush === 'function' && !this._readableState.destroyed) {
    this._flush(function (er, data) {
      done(_this, er, data);
    });
  } else {
    done(this, null, null);
  }
}
Transform.prototype.push = function (chunk, encoding) {
  this._transformState.needTransform = false;
  return Duplex.prototype.push.call(this, chunk, encoding);
};

// This is the part where you do stuff!
// override this function in implementation classes.
// 'chunk' is an input chunk.
//
// Call `push(newChunk)` to pass along transformed output
// to the readable side.  You may call 'push' zero or more times.
//
// Call `cb(err)` when you are done with this chunk.  If you pass
// an error, then that'll put the hurt on the whole operation.  If you
// never call cb(), then you'll never get another chunk.
Transform.prototype._transform = function (chunk, encoding, cb) {
  cb(new ERR_METHOD_NOT_IMPLEMENTED('_transform()'));
};
Transform.prototype._write = function (chunk, encoding, cb) {
  var ts = this._transformState;
  ts.writecb = cb;
  ts.writechunk = chunk;
  ts.writeencoding = encoding;
  if (!ts.transforming) {
    var rs = this._readableState;
    if (ts.needTransform || rs.needReadable || rs.length < rs.highWaterMark) this._read(rs.highWaterMark);
  }
};

// Doesn't matter what the args are here.
// _transform does all the work.
// That we got here means that the readable side wants more data.
Transform.prototype._read = function (n) {
  var ts = this._transformState;
  if (ts.writechunk !== null && !ts.transforming) {
    ts.transforming = true;
    this._transform(ts.writechunk, ts.writeencoding, ts.afterTransform);
  } else {
    // mark that we need a transform, so that any data that comes in
    // will get processed, now that we've asked for it.
    ts.needTransform = true;
  }
};
Transform.prototype._destroy = function (err, cb) {
  Duplex.prototype._destroy.call(this, err, function (err2) {
    cb(err2);
  });
};
function done(stream, er, data) {
  if (er) return stream.emit('error', er);
  if (data != null)
    // single equals check for both `null` and `undefined`
    stream.push(data);

  // TODO(BridgeAR): Write a test for these two error cases
  // if there's nothing in the write buffer, then that means
  // that nothing more will ever be provided
  if (stream._writableState.length) throw new ERR_TRANSFORM_WITH_LENGTH_0();
  if (stream._transformState.transforming) throw new ERR_TRANSFORM_ALREADY_TRANSFORMING();
  return stream.push(null);
}

/***/ }),

/***/ "./node_modules/readable-stream/lib/_stream_writable.js":
/*!**************************************************************!*\
  !*** ./node_modules/readable-stream/lib/_stream_writable.js ***!
  \**************************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";
/* provided dependency */ var process = __webpack_require__(/*! process/browser */ "./node_modules/process/browser.js");
// Copyright Joyent, Inc. and other Node contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to permit
// persons to whom the Software is furnished to do so, subject to the
// following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN
// NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
// USE OR OTHER DEALINGS IN THE SOFTWARE.

// A bit simpler than readable streams.
// Implement an async ._write(chunk, encoding, cb), and it'll handle all
// the drain event emission and buffering.



module.exports = Writable;

/* <replacement> */
function WriteReq(chunk, encoding, cb) {
  this.chunk = chunk;
  this.encoding = encoding;
  this.callback = cb;
  this.next = null;
}

// It seems a linked list but it is not
// there will be only 2 of these for each stream
function CorkedRequest(state) {
  var _this = this;
  this.next = null;
  this.entry = null;
  this.finish = function () {
    onCorkedFinish(_this, state);
  };
}
/* </replacement> */

/*<replacement>*/
var Duplex;
/*</replacement>*/

Writable.WritableState = WritableState;

/*<replacement>*/
var internalUtil = {
  deprecate: __webpack_require__(/*! util-deprecate */ "./node_modules/util-deprecate/browser.js")
};
/*</replacement>*/

/*<replacement>*/
var Stream = __webpack_require__(/*! ./internal/streams/stream */ "./node_modules/readable-stream/lib/internal/streams/stream-browser.js");
/*</replacement>*/

var Buffer = (__webpack_require__(/*! buffer */ "./node_modules/buffer/index.js").Buffer);
var OurUint8Array = (typeof __webpack_require__.g !== 'undefined' ? __webpack_require__.g : typeof window !== 'undefined' ? window : typeof self !== 'undefined' ? self : {}).Uint8Array || function () {};
function _uint8ArrayToBuffer(chunk) {
  return Buffer.from(chunk);
}
function _isUint8Array(obj) {
  return Buffer.isBuffer(obj) || obj instanceof OurUint8Array;
}
var destroyImpl = __webpack_require__(/*! ./internal/streams/destroy */ "./node_modules/readable-stream/lib/internal/streams/destroy.js");
var _require = __webpack_require__(/*! ./internal/streams/state */ "./node_modules/readable-stream/lib/internal/streams/state.js"),
  getHighWaterMark = _require.getHighWaterMark;
var _require$codes = (__webpack_require__(/*! ../errors */ "./node_modules/readable-stream/errors-browser.js").codes),
  ERR_INVALID_ARG_TYPE = _require$codes.ERR_INVALID_ARG_TYPE,
  ERR_METHOD_NOT_IMPLEMENTED = _require$codes.ERR_METHOD_NOT_IMPLEMENTED,
  ERR_MULTIPLE_CALLBACK = _require$codes.ERR_MULTIPLE_CALLBACK,
  ERR_STREAM_CANNOT_PIPE = _require$codes.ERR_STREAM_CANNOT_PIPE,
  ERR_STREAM_DESTROYED = _require$codes.ERR_STREAM_DESTROYED,
  ERR_STREAM_NULL_VALUES = _require$codes.ERR_STREAM_NULL_VALUES,
  ERR_STREAM_WRITE_AFTER_END = _require$codes.ERR_STREAM_WRITE_AFTER_END,
  ERR_UNKNOWN_ENCODING = _require$codes.ERR_UNKNOWN_ENCODING;
var errorOrDestroy = destroyImpl.errorOrDestroy;
__webpack_require__(/*! inherits */ "./node_modules/inherits/inherits_browser.js")(Writable, Stream);
function nop() {}
function WritableState(options, stream, isDuplex) {
  Duplex = Duplex || __webpack_require__(/*! ./_stream_duplex */ "./node_modules/readable-stream/lib/_stream_duplex.js");
  options = options || {};

  // Duplex streams are both readable and writable, but share
  // the same options object.
  // However, some cases require setting options to different
  // values for the readable and the writable sides of the duplex stream,
  // e.g. options.readableObjectMode vs. options.writableObjectMode, etc.
  if (typeof isDuplex !== 'boolean') isDuplex = stream instanceof Duplex;

  // object stream flag to indicate whether or not this stream
  // contains buffers or objects.
  this.objectMode = !!options.objectMode;
  if (isDuplex) this.objectMode = this.objectMode || !!options.writableObjectMode;

  // the point at which write() starts returning false
  // Note: 0 is a valid value, means that we always return false if
  // the entire buffer is not flushed immediately on write()
  this.highWaterMark = getHighWaterMark(this, options, 'writableHighWaterMark', isDuplex);

  // if _final has been called
  this.finalCalled = false;

  // drain event flag.
  this.needDrain = false;
  // at the start of calling end()
  this.ending = false;
  // when end() has been called, and returned
  this.ended = false;
  // when 'finish' is emitted
  this.finished = false;

  // has it been destroyed
  this.destroyed = false;

  // should we decode strings into buffers before passing to _write?
  // this is here so that some node-core streams can optimize string
  // handling at a lower level.
  var noDecode = options.decodeStrings === false;
  this.decodeStrings = !noDecode;

  // Crypto is kind of old and crusty.  Historically, its default string
  // encoding is 'binary' so we have to make this configurable.
  // Everything else in the universe uses 'utf8', though.
  this.defaultEncoding = options.defaultEncoding || 'utf8';

  // not an actual buffer we keep track of, but a measurement
  // of how much we're waiting to get pushed to some underlying
  // socket or file.
  this.length = 0;

  // a flag to see when we're in the middle of a write.
  this.writing = false;

  // when true all writes will be buffered until .uncork() call
  this.corked = 0;

  // a flag to be able to tell if the onwrite cb is called immediately,
  // or on a later tick.  We set this to true at first, because any
  // actions that shouldn't happen until "later" should generally also
  // not happen before the first write call.
  this.sync = true;

  // a flag to know if we're processing previously buffered items, which
  // may call the _write() callback in the same tick, so that we don't
  // end up in an overlapped onwrite situation.
  this.bufferProcessing = false;

  // the callback that's passed to _write(chunk,cb)
  this.onwrite = function (er) {
    onwrite(stream, er);
  };

  // the callback that the user supplies to write(chunk,encoding,cb)
  this.writecb = null;

  // the amount that is being written when _write is called.
  this.writelen = 0;
  this.bufferedRequest = null;
  this.lastBufferedRequest = null;

  // number of pending user-supplied write callbacks
  // this must be 0 before 'finish' can be emitted
  this.pendingcb = 0;

  // emit prefinish if the only thing we're waiting for is _write cbs
  // This is relevant for synchronous Transform streams
  this.prefinished = false;

  // True if the error was already emitted and should not be thrown again
  this.errorEmitted = false;

  // Should close be emitted on destroy. Defaults to true.
  this.emitClose = options.emitClose !== false;

  // Should .destroy() be called after 'finish' (and potentially 'end')
  this.autoDestroy = !!options.autoDestroy;

  // count buffered requests
  this.bufferedRequestCount = 0;

  // allocate the first CorkedRequest, there is always
  // one allocated and free to use, and we maintain at most two
  this.corkedRequestsFree = new CorkedRequest(this);
}
WritableState.prototype.getBuffer = function getBuffer() {
  var current = this.bufferedRequest;
  var out = [];
  while (current) {
    out.push(current);
    current = current.next;
  }
  return out;
};
(function () {
  try {
    Object.defineProperty(WritableState.prototype, 'buffer', {
      get: internalUtil.deprecate(function writableStateBufferGetter() {
        return this.getBuffer();
      }, '_writableState.buffer is deprecated. Use _writableState.getBuffer ' + 'instead.', 'DEP0003')
    });
  } catch (_) {}
})();

// Test _writableState for inheritance to account for Duplex streams,
// whose prototype chain only points to Readable.
var realHasInstance;
if (typeof Symbol === 'function' && Symbol.hasInstance && typeof Function.prototype[Symbol.hasInstance] === 'function') {
  realHasInstance = Function.prototype[Symbol.hasInstance];
  Object.defineProperty(Writable, Symbol.hasInstance, {
    value: function value(object) {
      if (realHasInstance.call(this, object)) return true;
      if (this !== Writable) return false;
      return object && object._writableState instanceof WritableState;
    }
  });
} else {
  realHasInstance = function realHasInstance(object) {
    return object instanceof this;
  };
}
function Writable(options) {
  Duplex = Duplex || __webpack_require__(/*! ./_stream_duplex */ "./node_modules/readable-stream/lib/_stream_duplex.js");

  // Writable ctor is applied to Duplexes, too.
  // `realHasInstance` is necessary because using plain `instanceof`
  // would return false, as no `_writableState` property is attached.

  // Trying to use the custom `instanceof` for Writable here will also break the
  // Node.js LazyTransform implementation, which has a non-trivial getter for
  // `_writableState` that would lead to infinite recursion.

  // Checking for a Stream.Duplex instance is faster here instead of inside
  // the WritableState constructor, at least with V8 6.5
  var isDuplex = this instanceof Duplex;
  if (!isDuplex && !realHasInstance.call(Writable, this)) return new Writable(options);
  this._writableState = new WritableState(options, this, isDuplex);

  // legacy.
  this.writable = true;
  if (options) {
    if (typeof options.write === 'function') this._write = options.write;
    if (typeof options.writev === 'function') this._writev = options.writev;
    if (typeof options.destroy === 'function') this._destroy = options.destroy;
    if (typeof options.final === 'function') this._final = options.final;
  }
  Stream.call(this);
}

// Otherwise people can pipe Writable streams, which is just wrong.
Writable.prototype.pipe = function () {
  errorOrDestroy(this, new ERR_STREAM_CANNOT_PIPE());
};
function writeAfterEnd(stream, cb) {
  var er = new ERR_STREAM_WRITE_AFTER_END();
  // TODO: defer error events consistently everywhere, not just the cb
  errorOrDestroy(stream, er);
  process.nextTick(cb, er);
}

// Checks that a user-supplied chunk is valid, especially for the particular
// mode the stream is in. Currently this means that `null` is never accepted
// and undefined/non-string values are only allowed in object mode.
function validChunk(stream, state, chunk, cb) {
  var er;
  if (chunk === null) {
    er = new ERR_STREAM_NULL_VALUES();
  } else if (typeof chunk !== 'string' && !state.objectMode) {
    er = new ERR_INVALID_ARG_TYPE('chunk', ['string', 'Buffer'], chunk);
  }
  if (er) {
    errorOrDestroy(stream, er);
    process.nextTick(cb, er);
    return false;
  }
  return true;
}
Writable.prototype.write = function (chunk, encoding, cb) {
  var state = this._writableState;
  var ret = false;
  var isBuf = !state.objectMode && _isUint8Array(chunk);
  if (isBuf && !Buffer.isBuffer(chunk)) {
    chunk = _uint8ArrayToBuffer(chunk);
  }
  if (typeof encoding === 'function') {
    cb = encoding;
    encoding = null;
  }
  if (isBuf) encoding = 'buffer';else if (!encoding) encoding = state.defaultEncoding;
  if (typeof cb !== 'function') cb = nop;
  if (state.ending) writeAfterEnd(this, cb);else if (isBuf || validChunk(this, state, chunk, cb)) {
    state.pendingcb++;
    ret = writeOrBuffer(this, state, isBuf, chunk, encoding, cb);
  }
  return ret;
};
Writable.prototype.cork = function () {
  this._writableState.corked++;
};
Writable.prototype.uncork = function () {
  var state = this._writableState;
  if (state.corked) {
    state.corked--;
    if (!state.writing && !state.corked && !state.bufferProcessing && state.bufferedRequest) clearBuffer(this, state);
  }
};
Writable.prototype.setDefaultEncoding = function setDefaultEncoding(encoding) {
  // node::ParseEncoding() requires lower case.
  if (typeof encoding === 'string') encoding = encoding.toLowerCase();
  if (!(['hex', 'utf8', 'utf-8', 'ascii', 'binary', 'base64', 'ucs2', 'ucs-2', 'utf16le', 'utf-16le', 'raw'].indexOf((encoding + '').toLowerCase()) > -1)) throw new ERR_UNKNOWN_ENCODING(encoding);
  this._writableState.defaultEncoding = encoding;
  return this;
};
Object.defineProperty(Writable.prototype, 'writableBuffer', {
  // making it explicit this property is not enumerable
  // because otherwise some prototype manipulation in
  // userland will fail
  enumerable: false,
  get: function get() {
    return this._writableState && this._writableState.getBuffer();
  }
});
function decodeChunk(state, chunk, encoding) {
  if (!state.objectMode && state.decodeStrings !== false && typeof chunk === 'string') {
    chunk = Buffer.from(chunk, encoding);
  }
  return chunk;
}
Object.defineProperty(Writable.prototype, 'writableHighWaterMark', {
  // making it explicit this property is not enumerable
  // because otherwise some prototype manipulation in
  // userland will fail
  enumerable: false,
  get: function get() {
    return this._writableState.highWaterMark;
  }
});

// if we're already writing something, then just put this
// in the queue, and wait our turn.  Otherwise, call _write
// If we return false, then we need a drain event, so set that flag.
function writeOrBuffer(stream, state, isBuf, chunk, encoding, cb) {
  if (!isBuf) {
    var newChunk = decodeChunk(state, chunk, encoding);
    if (chunk !== newChunk) {
      isBuf = true;
      encoding = 'buffer';
      chunk = newChunk;
    }
  }
  var len = state.objectMode ? 1 : chunk.length;
  state.length += len;
  var ret = state.length < state.highWaterMark;
  // we must ensure that previous needDrain will not be reset to false.
  if (!ret) state.needDrain = true;
  if (state.writing || state.corked) {
    var last = state.lastBufferedRequest;
    state.lastBufferedRequest = {
      chunk: chunk,
      encoding: encoding,
      isBuf: isBuf,
      callback: cb,
      next: null
    };
    if (last) {
      last.next = state.lastBufferedRequest;
    } else {
      state.bufferedRequest = state.lastBufferedRequest;
    }
    state.bufferedRequestCount += 1;
  } else {
    doWrite(stream, state, false, len, chunk, encoding, cb);
  }
  return ret;
}
function doWrite(stream, state, writev, len, chunk, encoding, cb) {
  state.writelen = len;
  state.writecb = cb;
  state.writing = true;
  state.sync = true;
  if (state.destroyed) state.onwrite(new ERR_STREAM_DESTROYED('write'));else if (writev) stream._writev(chunk, state.onwrite);else stream._write(chunk, encoding, state.onwrite);
  state.sync = false;
}
function onwriteError(stream, state, sync, er, cb) {
  --state.pendingcb;
  if (sync) {
    // defer the callback if we are being called synchronously
    // to avoid piling up things on the stack
    process.nextTick(cb, er);
    // this can emit finish, and it will always happen
    // after error
    process.nextTick(finishMaybe, stream, state);
    stream._writableState.errorEmitted = true;
    errorOrDestroy(stream, er);
  } else {
    // the caller expect this to happen before if
    // it is async
    cb(er);
    stream._writableState.errorEmitted = true;
    errorOrDestroy(stream, er);
    // this can emit finish, but finish must
    // always follow error
    finishMaybe(stream, state);
  }
}
function onwriteStateUpdate(state) {
  state.writing = false;
  state.writecb = null;
  state.length -= state.writelen;
  state.writelen = 0;
}
function onwrite(stream, er) {
  var state = stream._writableState;
  var sync = state.sync;
  var cb = state.writecb;
  if (typeof cb !== 'function') throw new ERR_MULTIPLE_CALLBACK();
  onwriteStateUpdate(state);
  if (er) onwriteError(stream, state, sync, er, cb);else {
    // Check if we're actually ready to finish, but don't emit yet
    var finished = needFinish(state) || stream.destroyed;
    if (!finished && !state.corked && !state.bufferProcessing && state.bufferedRequest) {
      clearBuffer(stream, state);
    }
    if (sync) {
      process.nextTick(afterWrite, stream, state, finished, cb);
    } else {
      afterWrite(stream, state, finished, cb);
    }
  }
}
function afterWrite(stream, state, finished, cb) {
  if (!finished) onwriteDrain(stream, state);
  state.pendingcb--;
  cb();
  finishMaybe(stream, state);
}

// Must force callback to be called on nextTick, so that we don't
// emit 'drain' before the write() consumer gets the 'false' return
// value, and has a chance to attach a 'drain' listener.
function onwriteDrain(stream, state) {
  if (state.length === 0 && state.needDrain) {
    state.needDrain = false;
    stream.emit('drain');
  }
}

// if there's something in the buffer waiting, then process it
function clearBuffer(stream, state) {
  state.bufferProcessing = true;
  var entry = state.bufferedRequest;
  if (stream._writev && entry && entry.next) {
    // Fast case, write everything using _writev()
    var l = state.bufferedRequestCount;
    var buffer = new Array(l);
    var holder = state.corkedRequestsFree;
    holder.entry = entry;
    var count = 0;
    var allBuffers = true;
    while (entry) {
      buffer[count] = entry;
      if (!entry.isBuf) allBuffers = false;
      entry = entry.next;
      count += 1;
    }
    buffer.allBuffers = allBuffers;
    doWrite(stream, state, true, state.length, buffer, '', holder.finish);

    // doWrite is almost always async, defer these to save a bit of time
    // as the hot path ends with doWrite
    state.pendingcb++;
    state.lastBufferedRequest = null;
    if (holder.next) {
      state.corkedRequestsFree = holder.next;
      holder.next = null;
    } else {
      state.corkedRequestsFree = new CorkedRequest(state);
    }
    state.bufferedRequestCount = 0;
  } else {
    // Slow case, write chunks one-by-one
    while (entry) {
      var chunk = entry.chunk;
      var encoding = entry.encoding;
      var cb = entry.callback;
      var len = state.objectMode ? 1 : chunk.length;
      doWrite(stream, state, false, len, chunk, encoding, cb);
      entry = entry.next;
      state.bufferedRequestCount--;
      // if we didn't call the onwrite immediately, then
      // it means that we need to wait until it does.
      // also, that means that the chunk and cb are currently
      // being processed, so move the buffer counter past them.
      if (state.writing) {
        break;
      }
    }
    if (entry === null) state.lastBufferedRequest = null;
  }
  state.bufferedRequest = entry;
  state.bufferProcessing = false;
}
Writable.prototype._write = function (chunk, encoding, cb) {
  cb(new ERR_METHOD_NOT_IMPLEMENTED('_write()'));
};
Writable.prototype._writev = null;
Writable.prototype.end = function (chunk, encoding, cb) {
  var state = this._writableState;
  if (typeof chunk === 'function') {
    cb = chunk;
    chunk = null;
    encoding = null;
  } else if (typeof encoding === 'function') {
    cb = encoding;
    encoding = null;
  }
  if (chunk !== null && chunk !== undefined) this.write(chunk, encoding);

  // .end() fully uncorks
  if (state.corked) {
    state.corked = 1;
    this.uncork();
  }

  // ignore unnecessary end() calls.
  if (!state.ending) endWritable(this, state, cb);
  return this;
};
Object.defineProperty(Writable.prototype, 'writableLength', {
  // making it explicit this property is not enumerable
  // because otherwise some prototype manipulation in
  // userland will fail
  enumerable: false,
  get: function get() {
    return this._writableState.length;
  }
});
function needFinish(state) {
  return state.ending && state.length === 0 && state.bufferedRequest === null && !state.finished && !state.writing;
}
function callFinal(stream, state) {
  stream._final(function (err) {
    state.pendingcb--;
    if (err) {
      errorOrDestroy(stream, err);
    }
    state.prefinished = true;
    stream.emit('prefinish');
    finishMaybe(stream, state);
  });
}
function prefinish(stream, state) {
  if (!state.prefinished && !state.finalCalled) {
    if (typeof stream._final === 'function' && !state.destroyed) {
      state.pendingcb++;
      state.finalCalled = true;
      process.nextTick(callFinal, stream, state);
    } else {
      state.prefinished = true;
      stream.emit('prefinish');
    }
  }
}
function finishMaybe(stream, state) {
  var need = needFinish(state);
  if (need) {
    prefinish(stream, state);
    if (state.pendingcb === 0) {
      state.finished = true;
      stream.emit('finish');
      if (state.autoDestroy) {
        // In case of duplex streams we need a way to detect
        // if the readable side is ready for autoDestroy as well
        var rState = stream._readableState;
        if (!rState || rState.autoDestroy && rState.endEmitted) {
          stream.destroy();
        }
      }
    }
  }
  return need;
}
function endWritable(stream, state, cb) {
  state.ending = true;
  finishMaybe(stream, state);
  if (cb) {
    if (state.finished) process.nextTick(cb);else stream.once('finish', cb);
  }
  state.ended = true;
  stream.writable = false;
}
function onCorkedFinish(corkReq, state, err) {
  var entry = corkReq.entry;
  corkReq.entry = null;
  while (entry) {
    var cb = entry.callback;
    state.pendingcb--;
    cb(err);
    entry = entry.next;
  }

  // reuse the free corkReq.
  state.corkedRequestsFree.next = corkReq;
}
Object.defineProperty(Writable.prototype, 'destroyed', {
  // making it explicit this property is not enumerable
  // because otherwise some prototype manipulation in
  // userland will fail
  enumerable: false,
  get: function get() {
    if (this._writableState === undefined) {
      return false;
    }
    return this._writableState.destroyed;
  },
  set: function set(value) {
    // we ignore the value if the stream
    // has not been initialized yet
    if (!this._writableState) {
      return;
    }

    // backward compatibility, the user is explicitly
    // managing destroyed
    this._writableState.destroyed = value;
  }
});
Writable.prototype.destroy = destroyImpl.destroy;
Writable.prototype._undestroy = destroyImpl.undestroy;
Writable.prototype._destroy = function (err, cb) {
  cb(err);
};

/***/ }),

/***/ "./node_modules/readable-stream/lib/internal/streams/async_iterator.js":
/*!*****************************************************************************!*\
  !*** ./node_modules/readable-stream/lib/internal/streams/async_iterator.js ***!
  \*****************************************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";
/* provided dependency */ var process = __webpack_require__(/*! process/browser */ "./node_modules/process/browser.js");


var _Object$setPrototypeO;
function _defineProperty(obj, key, value) { key = _toPropertyKey(key); if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }
function _toPropertyKey(arg) { var key = _toPrimitive(arg, "string"); return typeof key === "symbol" ? key : String(key); }
function _toPrimitive(input, hint) { if (typeof input !== "object" || input === null) return input; var prim = input[Symbol.toPrimitive]; if (prim !== undefined) { var res = prim.call(input, hint || "default"); if (typeof res !== "object") return res; throw new TypeError("@@toPrimitive must return a primitive value."); } return (hint === "string" ? String : Number)(input); }
var finished = __webpack_require__(/*! ./end-of-stream */ "./node_modules/readable-stream/lib/internal/streams/end-of-stream.js");
var kLastResolve = Symbol('lastResolve');
var kLastReject = Symbol('lastReject');
var kError = Symbol('error');
var kEnded = Symbol('ended');
var kLastPromise = Symbol('lastPromise');
var kHandlePromise = Symbol('handlePromise');
var kStream = Symbol('stream');
function createIterResult(value, done) {
  return {
    value: value,
    done: done
  };
}
function readAndResolve(iter) {
  var resolve = iter[kLastResolve];
  if (resolve !== null) {
    var data = iter[kStream].read();
    // we defer if data is null
    // we can be expecting either 'end' or
    // 'error'
    if (data !== null) {
      iter[kLastPromise] = null;
      iter[kLastResolve] = null;
      iter[kLastReject] = null;
      resolve(createIterResult(data, false));
    }
  }
}
function onReadable(iter) {
  // we wait for the next tick, because it might
  // emit an error with process.nextTick
  process.nextTick(readAndResolve, iter);
}
function wrapForNext(lastPromise, iter) {
  return function (resolve, reject) {
    lastPromise.then(function () {
      if (iter[kEnded]) {
        resolve(createIterResult(undefined, true));
        return;
      }
      iter[kHandlePromise](resolve, reject);
    }, reject);
  };
}
var AsyncIteratorPrototype = Object.getPrototypeOf(function () {});
var ReadableStreamAsyncIteratorPrototype = Object.setPrototypeOf((_Object$setPrototypeO = {
  get stream() {
    return this[kStream];
  },
  next: function next() {
    var _this = this;
    // if we have detected an error in the meanwhile
    // reject straight away
    var error = this[kError];
    if (error !== null) {
      return Promise.reject(error);
    }
    if (this[kEnded]) {
      return Promise.resolve(createIterResult(undefined, true));
    }
    if (this[kStream].destroyed) {
      // We need to defer via nextTick because if .destroy(err) is
      // called, the error will be emitted via nextTick, and
      // we cannot guarantee that there is no error lingering around
      // waiting to be emitted.
      return new Promise(function (resolve, reject) {
        process.nextTick(function () {
          if (_this[kError]) {
            reject(_this[kError]);
          } else {
            resolve(createIterResult(undefined, true));
          }
        });
      });
    }

    // if we have multiple next() calls
    // we will wait for the previous Promise to finish
    // this logic is optimized to support for await loops,
    // where next() is only called once at a time
    var lastPromise = this[kLastPromise];
    var promise;
    if (lastPromise) {
      promise = new Promise(wrapForNext(lastPromise, this));
    } else {
      // fast path needed to support multiple this.push()
      // without triggering the next() queue
      var data = this[kStream].read();
      if (data !== null) {
        return Promise.resolve(createIterResult(data, false));
      }
      promise = new Promise(this[kHandlePromise]);
    }
    this[kLastPromise] = promise;
    return promise;
  }
}, _defineProperty(_Object$setPrototypeO, Symbol.asyncIterator, function () {
  return this;
}), _defineProperty(_Object$setPrototypeO, "return", function _return() {
  var _this2 = this;
  // destroy(err, cb) is a private API
  // we can guarantee we have that here, because we control the
  // Readable class this is attached to
  return new Promise(function (resolve, reject) {
    _this2[kStream].destroy(null, function (err) {
      if (err) {
        reject(err);
        return;
      }
      resolve(createIterResult(undefined, true));
    });
  });
}), _Object$setPrototypeO), AsyncIteratorPrototype);
var createReadableStreamAsyncIterator = function createReadableStreamAsyncIterator(stream) {
  var _Object$create;
  var iterator = Object.create(ReadableStreamAsyncIteratorPrototype, (_Object$create = {}, _defineProperty(_Object$create, kStream, {
    value: stream,
    writable: true
  }), _defineProperty(_Object$create, kLastResolve, {
    value: null,
    writable: true
  }), _defineProperty(_Object$create, kLastReject, {
    value: null,
    writable: true
  }), _defineProperty(_Object$create, kError, {
    value: null,
    writable: true
  }), _defineProperty(_Object$create, kEnded, {
    value: stream._readableState.endEmitted,
    writable: true
  }), _defineProperty(_Object$create, kHandlePromise, {
    value: function value(resolve, reject) {
      var data = iterator[kStream].read();
      if (data) {
        iterator[kLastPromise] = null;
        iterator[kLastResolve] = null;
        iterator[kLastReject] = null;
        resolve(createIterResult(data, false));
      } else {
        iterator[kLastResolve] = resolve;
        iterator[kLastReject] = reject;
      }
    },
    writable: true
  }), _Object$create));
  iterator[kLastPromise] = null;
  finished(stream, function (err) {
    if (err && err.code !== 'ERR_STREAM_PREMATURE_CLOSE') {
      var reject = iterator[kLastReject];
      // reject if we are waiting for data in the Promise
      // returned by next() and store the error
      if (reject !== null) {
        iterator[kLastPromise] = null;
        iterator[kLastResolve] = null;
        iterator[kLastReject] = null;
        reject(err);
      }
      iterator[kError] = err;
      return;
    }
    var resolve = iterator[kLastResolve];
    if (resolve !== null) {
      iterator[kLastPromise] = null;
      iterator[kLastResolve] = null;
      iterator[kLastReject] = null;
      resolve(createIterResult(undefined, true));
    }
    iterator[kEnded] = true;
  });
  stream.on('readable', onReadable.bind(null, iterator));
  return iterator;
};
module.exports = createReadableStreamAsyncIterator;

/***/ }),

/***/ "./node_modules/readable-stream/lib/internal/streams/buffer_list.js":
/*!**************************************************************************!*\
  !*** ./node_modules/readable-stream/lib/internal/streams/buffer_list.js ***!
  \**************************************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


function ownKeys(object, enumerableOnly) { var keys = Object.keys(object); if (Object.getOwnPropertySymbols) { var symbols = Object.getOwnPropertySymbols(object); enumerableOnly && (symbols = symbols.filter(function (sym) { return Object.getOwnPropertyDescriptor(object, sym).enumerable; })), keys.push.apply(keys, symbols); } return keys; }
function _objectSpread(target) { for (var i = 1; i < arguments.length; i++) { var source = null != arguments[i] ? arguments[i] : {}; i % 2 ? ownKeys(Object(source), !0).forEach(function (key) { _defineProperty(target, key, source[key]); }) : Object.getOwnPropertyDescriptors ? Object.defineProperties(target, Object.getOwnPropertyDescriptors(source)) : ownKeys(Object(source)).forEach(function (key) { Object.defineProperty(target, key, Object.getOwnPropertyDescriptor(source, key)); }); } return target; }
function _defineProperty(obj, key, value) { key = _toPropertyKey(key); if (key in obj) { Object.defineProperty(obj, key, { value: value, enumerable: true, configurable: true, writable: true }); } else { obj[key] = value; } return obj; }
function _classCallCheck(instance, Constructor) { if (!(instance instanceof Constructor)) { throw new TypeError("Cannot call a class as a function"); } }
function _defineProperties(target, props) { for (var i = 0; i < props.length; i++) { var descriptor = props[i]; descriptor.enumerable = descriptor.enumerable || false; descriptor.configurable = true; if ("value" in descriptor) descriptor.writable = true; Object.defineProperty(target, _toPropertyKey(descriptor.key), descriptor); } }
function _createClass(Constructor, protoProps, staticProps) { if (protoProps) _defineProperties(Constructor.prototype, protoProps); if (staticProps) _defineProperties(Constructor, staticProps); Object.defineProperty(Constructor, "prototype", { writable: false }); return Constructor; }
function _toPropertyKey(arg) { var key = _toPrimitive(arg, "string"); return typeof key === "symbol" ? key : String(key); }
function _toPrimitive(input, hint) { if (typeof input !== "object" || input === null) return input; var prim = input[Symbol.toPrimitive]; if (prim !== undefined) { var res = prim.call(input, hint || "default"); if (typeof res !== "object") return res; throw new TypeError("@@toPrimitive must return a primitive value."); } return (hint === "string" ? String : Number)(input); }
var _require = __webpack_require__(/*! buffer */ "./node_modules/buffer/index.js"),
  Buffer = _require.Buffer;
var _require2 = __webpack_require__(/*! util */ "?ed1b"),
  inspect = _require2.inspect;
var custom = inspect && inspect.custom || 'inspect';
function copyBuffer(src, target, offset) {
  Buffer.prototype.copy.call(src, target, offset);
}
module.exports = /*#__PURE__*/function () {
  function BufferList() {
    _classCallCheck(this, BufferList);
    this.head = null;
    this.tail = null;
    this.length = 0;
  }
  _createClass(BufferList, [{
    key: "push",
    value: function push(v) {
      var entry = {
        data: v,
        next: null
      };
      if (this.length > 0) this.tail.next = entry;else this.head = entry;
      this.tail = entry;
      ++this.length;
    }
  }, {
    key: "unshift",
    value: function unshift(v) {
      var entry = {
        data: v,
        next: this.head
      };
      if (this.length === 0) this.tail = entry;
      this.head = entry;
      ++this.length;
    }
  }, {
    key: "shift",
    value: function shift() {
      if (this.length === 0) return;
      var ret = this.head.data;
      if (this.length === 1) this.head = this.tail = null;else this.head = this.head.next;
      --this.length;
      return ret;
    }
  }, {
    key: "clear",
    value: function clear() {
      this.head = this.tail = null;
      this.length = 0;
    }
  }, {
    key: "join",
    value: function join(s) {
      if (this.length === 0) return '';
      var p = this.head;
      var ret = '' + p.data;
      while (p = p.next) ret += s + p.data;
      return ret;
    }
  }, {
    key: "concat",
    value: function concat(n) {
      if (this.length === 0) return Buffer.alloc(0);
      var ret = Buffer.allocUnsafe(n >>> 0);
      var p = this.head;
      var i = 0;
      while (p) {
        copyBuffer(p.data, ret, i);
        i += p.data.length;
        p = p.next;
      }
      return ret;
    }

    // Consumes a specified amount of bytes or characters from the buffered data.
  }, {
    key: "consume",
    value: function consume(n, hasStrings) {
      var ret;
      if (n < this.head.data.length) {
        // `slice` is the same for buffers and strings.
        ret = this.head.data.slice(0, n);
        this.head.data = this.head.data.slice(n);
      } else if (n === this.head.data.length) {
        // First chunk is a perfect match.
        ret = this.shift();
      } else {
        // Result spans more than one buffer.
        ret = hasStrings ? this._getString(n) : this._getBuffer(n);
      }
      return ret;
    }
  }, {
    key: "first",
    value: function first() {
      return this.head.data;
    }

    // Consumes a specified amount of characters from the buffered data.
  }, {
    key: "_getString",
    value: function _getString(n) {
      var p = this.head;
      var c = 1;
      var ret = p.data;
      n -= ret.length;
      while (p = p.next) {
        var str = p.data;
        var nb = n > str.length ? str.length : n;
        if (nb === str.length) ret += str;else ret += str.slice(0, n);
        n -= nb;
        if (n === 0) {
          if (nb === str.length) {
            ++c;
            if (p.next) this.head = p.next;else this.head = this.tail = null;
          } else {
            this.head = p;
            p.data = str.slice(nb);
          }
          break;
        }
        ++c;
      }
      this.length -= c;
      return ret;
    }

    // Consumes a specified amount of bytes from the buffered data.
  }, {
    key: "_getBuffer",
    value: function _getBuffer(n) {
      var ret = Buffer.allocUnsafe(n);
      var p = this.head;
      var c = 1;
      p.data.copy(ret);
      n -= p.data.length;
      while (p = p.next) {
        var buf = p.data;
        var nb = n > buf.length ? buf.length : n;
        buf.copy(ret, ret.length - n, 0, nb);
        n -= nb;
        if (n === 0) {
          if (nb === buf.length) {
            ++c;
            if (p.next) this.head = p.next;else this.head = this.tail = null;
          } else {
            this.head = p;
            p.data = buf.slice(nb);
          }
          break;
        }
        ++c;
      }
      this.length -= c;
      return ret;
    }

    // Make sure the linked list only shows the minimal necessary information.
  }, {
    key: custom,
    value: function value(_, options) {
      return inspect(this, _objectSpread(_objectSpread({}, options), {}, {
        // Only inspect one level.
        depth: 0,
        // It should not recurse.
        customInspect: false
      }));
    }
  }]);
  return BufferList;
}();

/***/ }),

/***/ "./node_modules/readable-stream/lib/internal/streams/destroy.js":
/*!**********************************************************************!*\
  !*** ./node_modules/readable-stream/lib/internal/streams/destroy.js ***!
  \**********************************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";
/* provided dependency */ var process = __webpack_require__(/*! process/browser */ "./node_modules/process/browser.js");


// undocumented cb() API, needed for core, not for public API
function destroy(err, cb) {
  var _this = this;
  var readableDestroyed = this._readableState && this._readableState.destroyed;
  var writableDestroyed = this._writableState && this._writableState.destroyed;
  if (readableDestroyed || writableDestroyed) {
    if (cb) {
      cb(err);
    } else if (err) {
      if (!this._writableState) {
        process.nextTick(emitErrorNT, this, err);
      } else if (!this._writableState.errorEmitted) {
        this._writableState.errorEmitted = true;
        process.nextTick(emitErrorNT, this, err);
      }
    }
    return this;
  }

  // we set destroyed to true before firing error callbacks in order
  // to make it re-entrance safe in case destroy() is called within callbacks

  if (this._readableState) {
    this._readableState.destroyed = true;
  }

  // if this is a duplex stream mark the writable part as destroyed as well
  if (this._writableState) {
    this._writableState.destroyed = true;
  }
  this._destroy(err || null, function (err) {
    if (!cb && err) {
      if (!_this._writableState) {
        process.nextTick(emitErrorAndCloseNT, _this, err);
      } else if (!_this._writableState.errorEmitted) {
        _this._writableState.errorEmitted = true;
        process.nextTick(emitErrorAndCloseNT, _this, err);
      } else {
        process.nextTick(emitCloseNT, _this);
      }
    } else if (cb) {
      process.nextTick(emitCloseNT, _this);
      cb(err);
    } else {
      process.nextTick(emitCloseNT, _this);
    }
  });
  return this;
}
function emitErrorAndCloseNT(self, err) {
  emitErrorNT(self, err);
  emitCloseNT(self);
}
function emitCloseNT(self) {
  if (self._writableState && !self._writableState.emitClose) return;
  if (self._readableState && !self._readableState.emitClose) return;
  self.emit('close');
}
function undestroy() {
  if (this._readableState) {
    this._readableState.destroyed = false;
    this._readableState.reading = false;
    this._readableState.ended = false;
    this._readableState.endEmitted = false;
  }
  if (this._writableState) {
    this._writableState.destroyed = false;
    this._writableState.ended = false;
    this._writableState.ending = false;
    this._writableState.finalCalled = false;
    this._writableState.prefinished = false;
    this._writableState.finished = false;
    this._writableState.errorEmitted = false;
  }
}
function emitErrorNT(self, err) {
  self.emit('error', err);
}
function errorOrDestroy(stream, err) {
  // We have tests that rely on errors being emitted
  // in the same tick, so changing this is semver major.
  // For now when you opt-in to autoDestroy we allow
  // the error to be emitted nextTick. In a future
  // semver major update we should change the default to this.

  var rState = stream._readableState;
  var wState = stream._writableState;
  if (rState && rState.autoDestroy || wState && wState.autoDestroy) stream.destroy(err);else stream.emit('error', err);
}
module.exports = {
  destroy: destroy,
  undestroy: undestroy,
  errorOrDestroy: errorOrDestroy
};

/***/ }),

/***/ "./node_modules/readable-stream/lib/internal/streams/end-of-stream.js":
/*!****************************************************************************!*\
  !*** ./node_modules/readable-stream/lib/internal/streams/end-of-stream.js ***!
  \****************************************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";
// Ported from https://github.com/mafintosh/end-of-stream with
// permission from the author, Mathias Buus (@mafintosh).



var ERR_STREAM_PREMATURE_CLOSE = (__webpack_require__(/*! ../../../errors */ "./node_modules/readable-stream/errors-browser.js").codes).ERR_STREAM_PREMATURE_CLOSE;
function once(callback) {
  var called = false;
  return function () {
    if (called) return;
    called = true;
    for (var _len = arguments.length, args = new Array(_len), _key = 0; _key < _len; _key++) {
      args[_key] = arguments[_key];
    }
    callback.apply(this, args);
  };
}
function noop() {}
function isRequest(stream) {
  return stream.setHeader && typeof stream.abort === 'function';
}
function eos(stream, opts, callback) {
  if (typeof opts === 'function') return eos(stream, null, opts);
  if (!opts) opts = {};
  callback = once(callback || noop);
  var readable = opts.readable || opts.readable !== false && stream.readable;
  var writable = opts.writable || opts.writable !== false && stream.writable;
  var onlegacyfinish = function onlegacyfinish() {
    if (!stream.writable) onfinish();
  };
  var writableEnded = stream._writableState && stream._writableState.finished;
  var onfinish = function onfinish() {
    writable = false;
    writableEnded = true;
    if (!readable) callback.call(stream);
  };
  var readableEnded = stream._readableState && stream._readableState.endEmitted;
  var onend = function onend() {
    readable = false;
    readableEnded = true;
    if (!writable) callback.call(stream);
  };
  var onerror = function onerror(err) {
    callback.call(stream, err);
  };
  var onclose = function onclose() {
    var err;
    if (readable && !readableEnded) {
      if (!stream._readableState || !stream._readableState.ended) err = new ERR_STREAM_PREMATURE_CLOSE();
      return callback.call(stream, err);
    }
    if (writable && !writableEnded) {
      if (!stream._writableState || !stream._writableState.ended) err = new ERR_STREAM_PREMATURE_CLOSE();
      return callback.call(stream, err);
    }
  };
  var onrequest = function onrequest() {
    stream.req.on('finish', onfinish);
  };
  if (isRequest(stream)) {
    stream.on('complete', onfinish);
    stream.on('abort', onclose);
    if (stream.req) onrequest();else stream.on('request', onrequest);
  } else if (writable && !stream._writableState) {
    // legacy streams
    stream.on('end', onlegacyfinish);
    stream.on('close', onlegacyfinish);
  }
  stream.on('end', onend);
  stream.on('finish', onfinish);
  if (opts.error !== false) stream.on('error', onerror);
  stream.on('close', onclose);
  return function () {
    stream.removeListener('complete', onfinish);
    stream.removeListener('abort', onclose);
    stream.removeListener('request', onrequest);
    if (stream.req) stream.req.removeListener('finish', onfinish);
    stream.removeListener('end', onlegacyfinish);
    stream.removeListener('close', onlegacyfinish);
    stream.removeListener('finish', onfinish);
    stream.removeListener('end', onend);
    stream.removeListener('error', onerror);
    stream.removeListener('close', onclose);
  };
}
module.exports = eos;

/***/ }),

/***/ "./node_modules/readable-stream/lib/internal/streams/from-browser.js":
/*!***************************************************************************!*\
  !*** ./node_modules/readable-stream/lib/internal/streams/from-browser.js ***!
  \***************************************************************************/
/***/ ((module) => {

module.exports = function () {
  throw new Error('Readable.from is not available in the browser')
};


/***/ }),

/***/ "./node_modules/readable-stream/lib/internal/streams/pipeline.js":
/*!***********************************************************************!*\
  !*** ./node_modules/readable-stream/lib/internal/streams/pipeline.js ***!
  \***********************************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";
// Ported from https://github.com/mafintosh/pump with
// permission from the author, Mathias Buus (@mafintosh).



var eos;
function once(callback) {
  var called = false;
  return function () {
    if (called) return;
    called = true;
    callback.apply(void 0, arguments);
  };
}
var _require$codes = (__webpack_require__(/*! ../../../errors */ "./node_modules/readable-stream/errors-browser.js").codes),
  ERR_MISSING_ARGS = _require$codes.ERR_MISSING_ARGS,
  ERR_STREAM_DESTROYED = _require$codes.ERR_STREAM_DESTROYED;
function noop(err) {
  // Rethrow the error if it exists to avoid swallowing it
  if (err) throw err;
}
function isRequest(stream) {
  return stream.setHeader && typeof stream.abort === 'function';
}
function destroyer(stream, reading, writing, callback) {
  callback = once(callback);
  var closed = false;
  stream.on('close', function () {
    closed = true;
  });
  if (eos === undefined) eos = __webpack_require__(/*! ./end-of-stream */ "./node_modules/readable-stream/lib/internal/streams/end-of-stream.js");
  eos(stream, {
    readable: reading,
    writable: writing
  }, function (err) {
    if (err) return callback(err);
    closed = true;
    callback();
  });
  var destroyed = false;
  return function (err) {
    if (closed) return;
    if (destroyed) return;
    destroyed = true;

    // request.destroy just do .end - .abort is what we want
    if (isRequest(stream)) return stream.abort();
    if (typeof stream.destroy === 'function') return stream.destroy();
    callback(err || new ERR_STREAM_DESTROYED('pipe'));
  };
}
function call(fn) {
  fn();
}
function pipe(from, to) {
  return from.pipe(to);
}
function popCallback(streams) {
  if (!streams.length) return noop;
  if (typeof streams[streams.length - 1] !== 'function') return noop;
  return streams.pop();
}
function pipeline() {
  for (var _len = arguments.length, streams = new Array(_len), _key = 0; _key < _len; _key++) {
    streams[_key] = arguments[_key];
  }
  var callback = popCallback(streams);
  if (Array.isArray(streams[0])) streams = streams[0];
  if (streams.length < 2) {
    throw new ERR_MISSING_ARGS('streams');
  }
  var error;
  var destroys = streams.map(function (stream, i) {
    var reading = i < streams.length - 1;
    var writing = i > 0;
    return destroyer(stream, reading, writing, function (err) {
      if (!error) error = err;
      if (err) destroys.forEach(call);
      if (reading) return;
      destroys.forEach(call);
      callback(error);
    });
  });
  return streams.reduce(pipe);
}
module.exports = pipeline;

/***/ }),

/***/ "./node_modules/readable-stream/lib/internal/streams/state.js":
/*!********************************************************************!*\
  !*** ./node_modules/readable-stream/lib/internal/streams/state.js ***!
  \********************************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var ERR_INVALID_OPT_VALUE = (__webpack_require__(/*! ../../../errors */ "./node_modules/readable-stream/errors-browser.js").codes).ERR_INVALID_OPT_VALUE;
function highWaterMarkFrom(options, isDuplex, duplexKey) {
  return options.highWaterMark != null ? options.highWaterMark : isDuplex ? options[duplexKey] : null;
}
function getHighWaterMark(state, options, duplexKey, isDuplex) {
  var hwm = highWaterMarkFrom(options, isDuplex, duplexKey);
  if (hwm != null) {
    if (!(isFinite(hwm) && Math.floor(hwm) === hwm) || hwm < 0) {
      var name = isDuplex ? duplexKey : 'highWaterMark';
      throw new ERR_INVALID_OPT_VALUE(name, hwm);
    }
    return Math.floor(hwm);
  }

  // Default value
  return state.objectMode ? 16 : 16 * 1024;
}
module.exports = {
  getHighWaterMark: getHighWaterMark
};

/***/ }),

/***/ "./node_modules/readable-stream/lib/internal/streams/stream-browser.js":
/*!*****************************************************************************!*\
  !*** ./node_modules/readable-stream/lib/internal/streams/stream-browser.js ***!
  \*****************************************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

module.exports = __webpack_require__(/*! events */ "./node_modules/events/events.js").EventEmitter;


/***/ }),

/***/ "./node_modules/readable-stream/readable-browser.js":
/*!**********************************************************!*\
  !*** ./node_modules/readable-stream/readable-browser.js ***!
  \**********************************************************/
/***/ ((module, exports, __webpack_require__) => {

exports = module.exports = __webpack_require__(/*! ./lib/_stream_readable.js */ "./node_modules/readable-stream/lib/_stream_readable.js");
exports.Stream = exports;
exports.Readable = exports;
exports.Writable = __webpack_require__(/*! ./lib/_stream_writable.js */ "./node_modules/readable-stream/lib/_stream_writable.js");
exports.Duplex = __webpack_require__(/*! ./lib/_stream_duplex.js */ "./node_modules/readable-stream/lib/_stream_duplex.js");
exports.Transform = __webpack_require__(/*! ./lib/_stream_transform.js */ "./node_modules/readable-stream/lib/_stream_transform.js");
exports.PassThrough = __webpack_require__(/*! ./lib/_stream_passthrough.js */ "./node_modules/readable-stream/lib/_stream_passthrough.js");
exports.finished = __webpack_require__(/*! ./lib/internal/streams/end-of-stream.js */ "./node_modules/readable-stream/lib/internal/streams/end-of-stream.js");
exports.pipeline = __webpack_require__(/*! ./lib/internal/streams/pipeline.js */ "./node_modules/readable-stream/lib/internal/streams/pipeline.js");


/***/ }),

/***/ "./node_modules/safe-buffer/index.js":
/*!*******************************************!*\
  !*** ./node_modules/safe-buffer/index.js ***!
  \*******************************************/
/***/ ((module, exports, __webpack_require__) => {

/*! safe-buffer. MIT License. Feross Aboukhadijeh <https://feross.org/opensource> */
/* eslint-disable node/no-deprecated-api */
var buffer = __webpack_require__(/*! buffer */ "./node_modules/buffer/index.js")
var Buffer = buffer.Buffer

// alternative to using Object.keys for old browsers
function copyProps (src, dst) {
  for (var key in src) {
    dst[key] = src[key]
  }
}
if (Buffer.from && Buffer.alloc && Buffer.allocUnsafe && Buffer.allocUnsafeSlow) {
  module.exports = buffer
} else {
  // Copy properties from require('buffer')
  copyProps(buffer, exports)
  exports.Buffer = SafeBuffer
}

function SafeBuffer (arg, encodingOrOffset, length) {
  return Buffer(arg, encodingOrOffset, length)
}

SafeBuffer.prototype = Object.create(Buffer.prototype)

// Copy static methods from Buffer
copyProps(Buffer, SafeBuffer)

SafeBuffer.from = function (arg, encodingOrOffset, length) {
  if (typeof arg === 'number') {
    throw new TypeError('Argument must not be a number')
  }
  return Buffer(arg, encodingOrOffset, length)
}

SafeBuffer.alloc = function (size, fill, encoding) {
  if (typeof size !== 'number') {
    throw new TypeError('Argument must be a number')
  }
  var buf = Buffer(size)
  if (fill !== undefined) {
    if (typeof encoding === 'string') {
      buf.fill(fill, encoding)
    } else {
      buf.fill(fill)
    }
  } else {
    buf.fill(0)
  }
  return buf
}

SafeBuffer.allocUnsafe = function (size) {
  if (typeof size !== 'number') {
    throw new TypeError('Argument must be a number')
  }
  return Buffer(size)
}

SafeBuffer.allocUnsafeSlow = function (size) {
  if (typeof size !== 'number') {
    throw new TypeError('Argument must be a number')
  }
  return buffer.SlowBuffer(size)
}


/***/ }),

/***/ "./node_modules/set-function-length/index.js":
/*!***************************************************!*\
  !*** ./node_modules/set-function-length/index.js ***!
  \***************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var GetIntrinsic = __webpack_require__(/*! get-intrinsic */ "./node_modules/get-intrinsic/index.js");
var define = __webpack_require__(/*! define-data-property */ "./node_modules/define-data-property/index.js");
var hasDescriptors = __webpack_require__(/*! has-property-descriptors */ "./node_modules/has-property-descriptors/index.js")();
var gOPD = __webpack_require__(/*! gopd */ "./node_modules/gopd/index.js");

var $TypeError = GetIntrinsic('%TypeError%');
var $floor = GetIntrinsic('%Math.floor%');

module.exports = function setFunctionLength(fn, length) {
	if (typeof fn !== 'function') {
		throw new $TypeError('`fn` is not a function');
	}
	if (typeof length !== 'number' || length < 0 || length > 0xFFFFFFFF || $floor(length) !== length) {
		throw new $TypeError('`length` must be a positive 32-bit integer');
	}

	var loose = arguments.length > 2 && !!arguments[2];

	var functionLengthIsConfigurable = true;
	var functionLengthIsWritable = true;
	if ('length' in fn && gOPD) {
		var desc = gOPD(fn, 'length');
		if (desc && !desc.configurable) {
			functionLengthIsConfigurable = false;
		}
		if (desc && !desc.writable) {
			functionLengthIsWritable = false;
		}
	}

	if (functionLengthIsConfigurable || functionLengthIsWritable || !loose) {
		if (hasDescriptors) {
			define(fn, 'length', length, true, true);
		} else {
			define(fn, 'length', length);
		}
	}
	return fn;
};


/***/ }),

/***/ "./node_modules/side-channel/index.js":
/*!********************************************!*\
  !*** ./node_modules/side-channel/index.js ***!
  \********************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

"use strict";


var GetIntrinsic = __webpack_require__(/*! get-intrinsic */ "./node_modules/get-intrinsic/index.js");
var callBound = __webpack_require__(/*! call-bind/callBound */ "./node_modules/call-bind/callBound.js");
var inspect = __webpack_require__(/*! object-inspect */ "./node_modules/object-inspect/index.js");

var $TypeError = GetIntrinsic('%TypeError%');
var $WeakMap = GetIntrinsic('%WeakMap%', true);
var $Map = GetIntrinsic('%Map%', true);

var $weakMapGet = callBound('WeakMap.prototype.get', true);
var $weakMapSet = callBound('WeakMap.prototype.set', true);
var $weakMapHas = callBound('WeakMap.prototype.has', true);
var $mapGet = callBound('Map.prototype.get', true);
var $mapSet = callBound('Map.prototype.set', true);
var $mapHas = callBound('Map.prototype.has', true);

/*
 * This function traverses the list returning the node corresponding to the
 * given key.
 *
 * That node is also moved to the head of the list, so that if it's accessed
 * again we don't need to traverse the whole list. By doing so, all the recently
 * used nodes can be accessed relatively quickly.
 */
var listGetNode = function (list, key) { // eslint-disable-line consistent-return
	for (var prev = list, curr; (curr = prev.next) !== null; prev = curr) {
		if (curr.key === key) {
			prev.next = curr.next;
			curr.next = list.next;
			list.next = curr; // eslint-disable-line no-param-reassign
			return curr;
		}
	}
};

var listGet = function (objects, key) {
	var node = listGetNode(objects, key);
	return node && node.value;
};
var listSet = function (objects, key, value) {
	var node = listGetNode(objects, key);
	if (node) {
		node.value = value;
	} else {
		// Prepend the new node to the beginning of the list
		objects.next = { // eslint-disable-line no-param-reassign
			key: key,
			next: objects.next,
			value: value
		};
	}
};
var listHas = function (objects, key) {
	return !!listGetNode(objects, key);
};

module.exports = function getSideChannel() {
	var $wm;
	var $m;
	var $o;
	var channel = {
		assert: function (key) {
			if (!channel.has(key)) {
				throw new $TypeError('Side channel does not contain ' + inspect(key));
			}
		},
		get: function (key) { // eslint-disable-line consistent-return
			if ($WeakMap && key && (typeof key === 'object' || typeof key === 'function')) {
				if ($wm) {
					return $weakMapGet($wm, key);
				}
			} else if ($Map) {
				if ($m) {
					return $mapGet($m, key);
				}
			} else {
				if ($o) { // eslint-disable-line no-lonely-if
					return listGet($o, key);
				}
			}
		},
		has: function (key) {
			if ($WeakMap && key && (typeof key === 'object' || typeof key === 'function')) {
				if ($wm) {
					return $weakMapHas($wm, key);
				}
			} else if ($Map) {
				if ($m) {
					return $mapHas($m, key);
				}
			} else {
				if ($o) { // eslint-disable-line no-lonely-if
					return listHas($o, key);
				}
			}
			return false;
		},
		set: function (key, value) {
			if ($WeakMap && key && (typeof key === 'object' || typeof key === 'function')) {
				if (!$wm) {
					$wm = new $WeakMap();
				}
				$weakMapSet($wm, key, value);
			} else if ($Map) {
				if (!$m) {
					$m = new $Map();
				}
				$mapSet($m, key, value);
			} else {
				if (!$o) {
					/*
					 * Initialize the linked list as an empty node, so that we don't have
					 * to special-case handling of the first node: we can always refer to
					 * it as (previous node).next, instead of something like (list).head
					 */
					$o = { key: {}, next: null };
				}
				listSet($o, key, value);
			}
		}
	};
	return channel;
};


/***/ }),

/***/ "./node_modules/stream-http/index.js":
/*!*******************************************!*\
  !*** ./node_modules/stream-http/index.js ***!
  \*******************************************/
/***/ ((__unused_webpack_module, exports, __webpack_require__) => {

var ClientRequest = __webpack_require__(/*! ./lib/request */ "./node_modules/stream-http/lib/request.js")
var response = __webpack_require__(/*! ./lib/response */ "./node_modules/stream-http/lib/response.js")
var extend = __webpack_require__(/*! xtend */ "./node_modules/xtend/immutable.js")
var statusCodes = __webpack_require__(/*! builtin-status-codes */ "./node_modules/builtin-status-codes/browser.js")
var url = __webpack_require__(/*! url */ "./node_modules/url/url.js")

var http = exports

http.request = function (opts, cb) {
	if (typeof opts === 'string')
		opts = url.parse(opts)
	else
		opts = extend(opts)

	// Normally, the page is loaded from http or https, so not specifying a protocol
	// will result in a (valid) protocol-relative url. However, this won't work if
	// the protocol is something else, like 'file:'
	var defaultProtocol = __webpack_require__.g.location.protocol.search(/^https?:$/) === -1 ? 'http:' : ''

	var protocol = opts.protocol || defaultProtocol
	var host = opts.hostname || opts.host
	var port = opts.port
	var path = opts.path || '/'

	// Necessary for IPv6 addresses
	if (host && host.indexOf(':') !== -1)
		host = '[' + host + ']'

	// This may be a relative url. The browser should always be able to interpret it correctly.
	opts.url = (host ? (protocol + '//' + host) : '') + (port ? ':' + port : '') + path
	opts.method = (opts.method || 'GET').toUpperCase()
	opts.headers = opts.headers || {}

	// Also valid opts.auth, opts.mode

	var req = new ClientRequest(opts)
	if (cb)
		req.on('response', cb)
	return req
}

http.get = function get (opts, cb) {
	var req = http.request(opts, cb)
	req.end()
	return req
}

http.ClientRequest = ClientRequest
http.IncomingMessage = response.IncomingMessage

http.Agent = function () {}
http.Agent.defaultMaxSockets = 4

http.globalAgent = new http.Agent()

http.STATUS_CODES = statusCodes

http.METHODS = [
	'CHECKOUT',
	'CONNECT',
	'COPY',
	'DELETE',
	'GET',
	'HEAD',
	'LOCK',
	'M-SEARCH',
	'MERGE',
	'MKACTIVITY',
	'MKCOL',
	'MOVE',
	'NOTIFY',
	'OPTIONS',
	'PATCH',
	'POST',
	'PROPFIND',
	'PROPPATCH',
	'PURGE',
	'PUT',
	'REPORT',
	'SEARCH',
	'SUBSCRIBE',
	'TRACE',
	'UNLOCK',
	'UNSUBSCRIBE'
]

/***/ }),

/***/ "./node_modules/stream-http/lib/capability.js":
/*!****************************************************!*\
  !*** ./node_modules/stream-http/lib/capability.js ***!
  \****************************************************/
/***/ ((__unused_webpack_module, exports, __webpack_require__) => {

exports.fetch = isFunction(__webpack_require__.g.fetch) && isFunction(__webpack_require__.g.ReadableStream)

exports.writableStream = isFunction(__webpack_require__.g.WritableStream)

exports.abortController = isFunction(__webpack_require__.g.AbortController)

// The xhr request to example.com may violate some restrictive CSP configurations,
// so if we're running in a browser that supports `fetch`, avoid calling getXHR()
// and assume support for certain features below.
var xhr
function getXHR () {
	// Cache the xhr value
	if (xhr !== undefined) return xhr

	if (__webpack_require__.g.XMLHttpRequest) {
		xhr = new __webpack_require__.g.XMLHttpRequest()
		// If XDomainRequest is available (ie only, where xhr might not work
		// cross domain), use the page location. Otherwise use example.com
		// Note: this doesn't actually make an http request.
		try {
			xhr.open('GET', __webpack_require__.g.XDomainRequest ? '/' : 'https://example.com')
		} catch(e) {
			xhr = null
		}
	} else {
		// Service workers don't have XHR
		xhr = null
	}
	return xhr
}

function checkTypeSupport (type) {
	var xhr = getXHR()
	if (!xhr) return false
	try {
		xhr.responseType = type
		return xhr.responseType === type
	} catch (e) {}
	return false
}

// If fetch is supported, then arraybuffer will be supported too. Skip calling
// checkTypeSupport(), since that calls getXHR().
exports.arraybuffer = exports.fetch || checkTypeSupport('arraybuffer')

// These next two tests unavoidably show warnings in Chrome. Since fetch will always
// be used if it's available, just return false for these to avoid the warnings.
exports.msstream = !exports.fetch && checkTypeSupport('ms-stream')
exports.mozchunkedarraybuffer = !exports.fetch && checkTypeSupport('moz-chunked-arraybuffer')

// If fetch is supported, then overrideMimeType will be supported too. Skip calling
// getXHR().
exports.overrideMimeType = exports.fetch || (getXHR() ? isFunction(getXHR().overrideMimeType) : false)

function isFunction (value) {
	return typeof value === 'function'
}

xhr = null // Help gc


/***/ }),

/***/ "./node_modules/stream-http/lib/request.js":
/*!*************************************************!*\
  !*** ./node_modules/stream-http/lib/request.js ***!
  \*************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {

/* provided dependency */ var Buffer = __webpack_require__(/*! buffer */ "./node_modules/buffer/index.js")["Buffer"];
/* provided dependency */ var process = __webpack_require__(/*! process/browser */ "./node_modules/process/browser.js");
var capability = __webpack_require__(/*! ./capability */ "./node_modules/stream-http/lib/capability.js")
var inherits = __webpack_require__(/*! inherits */ "./node_modules/inherits/inherits_browser.js")
var response = __webpack_require__(/*! ./response */ "./node_modules/stream-http/lib/response.js")
var stream = __webpack_require__(/*! readable-stream */ "./node_modules/readable-stream/readable-browser.js")

var IncomingMessage = response.IncomingMessage
var rStates = response.readyStates

function decideMode (preferBinary, useFetch) {
	if (capability.fetch && useFetch) {
		return 'fetch'
	} else if (capability.mozchunkedarraybuffer) {
		return 'moz-chunked-arraybuffer'
	} else if (capability.msstream) {
		return 'ms-stream'
	} else if (capability.arraybuffer && preferBinary) {
		return 'arraybuffer'
	} else {
		return 'text'
	}
}

var ClientRequest = module.exports = function (opts) {
	var self = this
	stream.Writable.call(self)

	self._opts = opts
	self._body = []
	self._headers = {}
	if (opts.auth)
		self.setHeader('Authorization', 'Basic ' + Buffer.from(opts.auth).toString('base64'))
	Object.keys(opts.headers).forEach(function (name) {
		self.setHeader(name, opts.headers[name])
	})

	var preferBinary
	var useFetch = true
	if (opts.mode === 'disable-fetch' || ('requestTimeout' in opts && !capability.abortController)) {
		// If the use of XHR should be preferred. Not typically needed.
		useFetch = false
		preferBinary = true
	} else if (opts.mode === 'prefer-streaming') {
		// If streaming is a high priority but binary compatibility and
		// the accuracy of the 'content-type' header aren't
		preferBinary = false
	} else if (opts.mode === 'allow-wrong-content-type') {
		// If streaming is more important than preserving the 'content-type' header
		preferBinary = !capability.overrideMimeType
	} else if (!opts.mode || opts.mode === 'default' || opts.mode === 'prefer-fast') {
		// Use binary if text streaming may corrupt data or the content-type header, or for speed
		preferBinary = true
	} else {
		throw new Error('Invalid value for opts.mode')
	}
	self._mode = decideMode(preferBinary, useFetch)
	self._fetchTimer = null
	self._socketTimeout = null
	self._socketTimer = null

	self.on('finish', function () {
		self._onFinish()
	})
}

inherits(ClientRequest, stream.Writable)

ClientRequest.prototype.setHeader = function (name, value) {
	var self = this
	var lowerName = name.toLowerCase()
	// This check is not necessary, but it prevents warnings from browsers about setting unsafe
	// headers. To be honest I'm not entirely sure hiding these warnings is a good thing, but
	// http-browserify did it, so I will too.
	if (unsafeHeaders.indexOf(lowerName) !== -1)
		return

	self._headers[lowerName] = {
		name: name,
		value: value
	}
}

ClientRequest.prototype.getHeader = function (name) {
	var header = this._headers[name.toLowerCase()]
	if (header)
		return header.value
	return null
}

ClientRequest.prototype.removeHeader = function (name) {
	var self = this
	delete self._headers[name.toLowerCase()]
}

ClientRequest.prototype._onFinish = function () {
	var self = this

	if (self._destroyed)
		return
	var opts = self._opts

	if ('timeout' in opts && opts.timeout !== 0) {
		self.setTimeout(opts.timeout)
	}

	var headersObj = self._headers
	var body = null
	if (opts.method !== 'GET' && opts.method !== 'HEAD') {
        body = new Blob(self._body, {
            type: (headersObj['content-type'] || {}).value || ''
        });
    }

	// create flattened list of headers
	var headersList = []
	Object.keys(headersObj).forEach(function (keyName) {
		var name = headersObj[keyName].name
		var value = headersObj[keyName].value
		if (Array.isArray(value)) {
			value.forEach(function (v) {
				headersList.push([name, v])
			})
		} else {
			headersList.push([name, value])
		}
	})

	if (self._mode === 'fetch') {
		var signal = null
		if (capability.abortController) {
			var controller = new AbortController()
			signal = controller.signal
			self._fetchAbortController = controller

			if ('requestTimeout' in opts && opts.requestTimeout !== 0) {
				self._fetchTimer = __webpack_require__.g.setTimeout(function () {
					self.emit('requestTimeout')
					if (self._fetchAbortController)
						self._fetchAbortController.abort()
				}, opts.requestTimeout)
			}
		}

		__webpack_require__.g.fetch(self._opts.url, {
			method: self._opts.method,
			headers: headersList,
			body: body || undefined,
			mode: 'cors',
			credentials: opts.withCredentials ? 'include' : 'same-origin',
			signal: signal
		}).then(function (response) {
			self._fetchResponse = response
			self._resetTimers(false)
			self._connect()
		}, function (reason) {
			self._resetTimers(true)
			if (!self._destroyed)
				self.emit('error', reason)
		})
	} else {
		var xhr = self._xhr = new __webpack_require__.g.XMLHttpRequest()
		try {
			xhr.open(self._opts.method, self._opts.url, true)
		} catch (err) {
			process.nextTick(function () {
				self.emit('error', err)
			})
			return
		}

		// Can't set responseType on really old browsers
		if ('responseType' in xhr)
			xhr.responseType = self._mode

		if ('withCredentials' in xhr)
			xhr.withCredentials = !!opts.withCredentials

		if (self._mode === 'text' && 'overrideMimeType' in xhr)
			xhr.overrideMimeType('text/plain; charset=x-user-defined')

		if ('requestTimeout' in opts) {
			xhr.timeout = opts.requestTimeout
			xhr.ontimeout = function () {
				self.emit('requestTimeout')
			}
		}

		headersList.forEach(function (header) {
			xhr.setRequestHeader(header[0], header[1])
		})

		self._response = null
		xhr.onreadystatechange = function () {
			switch (xhr.readyState) {
				case rStates.LOADING:
				case rStates.DONE:
					self._onXHRProgress()
					break
			}
		}
		// Necessary for streaming in Firefox, since xhr.response is ONLY defined
		// in onprogress, not in onreadystatechange with xhr.readyState = 3
		if (self._mode === 'moz-chunked-arraybuffer') {
			xhr.onprogress = function () {
				self._onXHRProgress()
			}
		}

		xhr.onerror = function () {
			if (self._destroyed)
				return
			self._resetTimers(true)
			self.emit('error', new Error('XHR error'))
		}

		try {
			xhr.send(body)
		} catch (err) {
			process.nextTick(function () {
				self.emit('error', err)
			})
			return
		}
	}
}

/**
 * Checks if xhr.status is readable and non-zero, indicating no error.
 * Even though the spec says it should be available in readyState 3,
 * accessing it throws an exception in IE8
 */
function statusValid (xhr) {
	try {
		var status = xhr.status
		return (status !== null && status !== 0)
	} catch (e) {
		return false
	}
}

ClientRequest.prototype._onXHRProgress = function () {
	var self = this

	self._resetTimers(false)

	if (!statusValid(self._xhr) || self._destroyed)
		return

	if (!self._response)
		self._connect()

	self._response._onXHRProgress(self._resetTimers.bind(self))
}

ClientRequest.prototype._connect = function () {
	var self = this

	if (self._destroyed)
		return

	self._response = new IncomingMessage(self._xhr, self._fetchResponse, self._mode, self._resetTimers.bind(self))
	self._response.on('error', function(err) {
		self.emit('error', err)
	})

	self.emit('response', self._response)
}

ClientRequest.prototype._write = function (chunk, encoding, cb) {
	var self = this

	self._body.push(chunk)
	cb()
}

ClientRequest.prototype._resetTimers = function (done) {
	var self = this

	__webpack_require__.g.clearTimeout(self._socketTimer)
	self._socketTimer = null

	if (done) {
		__webpack_require__.g.clearTimeout(self._fetchTimer)
		self._fetchTimer = null
	} else if (self._socketTimeout) {
		self._socketTimer = __webpack_require__.g.setTimeout(function () {
			self.emit('timeout')
		}, self._socketTimeout)
	}
}

ClientRequest.prototype.abort = ClientRequest.prototype.destroy = function (err) {
	var self = this
	self._destroyed = true
	self._resetTimers(true)
	if (self._response)
		self._response._destroyed = true
	if (self._xhr)
		self._xhr.abort()
	else if (self._fetchAbortController)
		self._fetchAbortController.abort()

	if (err)
		self.emit('error', err)
}

ClientRequest.prototype.end = function (data, encoding, cb) {
	var self = this
	if (typeof data === 'function') {
		cb = data
		data = undefined
	}

	stream.Writable.prototype.end.call(self, data, encoding, cb)
}

ClientRequest.prototype.setTimeout = function (timeout, cb) {
	var self = this

	if (cb)
		self.once('timeout', cb)

	self._socketTimeout = timeout
	self._resetTimers(false)
}

ClientRequest.prototype.flushHeaders = function () {}
ClientRequest.prototype.setNoDelay = function () {}
ClientRequest.prototype.setSocketKeepAlive = function () {}

// Taken from http://www.w3.org/TR/XMLHttpRequest/#the-setrequestheader%28%29-method
var unsafeHeaders = [
	'accept-charset',
	'accept-encoding',
	'access-control-request-headers',
	'access-control-request-method',
	'connection',
	'content-length',
	'cookie',
	'cookie2',
	'date',
	'dnt',
	'expect',
	'host',
	'keep-alive',
	'origin',
	'referer',
	'te',
	'trailer',
	'transfer-encoding',
	'upgrade',
	'via'
]


/***/ }),

/***/ "./node_modules/stream-http/lib/response.js":
/*!**************************************************!*\
  !*** ./node_modules/stream-http/lib/response.js ***!
  \**************************************************/
/***/ ((__unused_webpack_module, exports, __webpack_require__) => {

/* provided dependency */ var process = __webpack_require__(/*! process/browser */ "./node_modules/process/browser.js");
/* provided dependency */ var Buffer = __webpack_require__(/*! buffer */ "./node_modules/buffer/index.js")["Buffer"];
var capability = __webpack_require__(/*! ./capability */ "./node_modules/stream-http/lib/capability.js")
var inherits = __webpack_require__(/*! inherits */ "./node_modules/inherits/inherits_browser.js")
var stream = __webpack_require__(/*! readable-stream */ "./node_modules/readable-stream/readable-browser.js")

var rStates = exports.readyStates = {
	UNSENT: 0,
	OPENED: 1,
	HEADERS_RECEIVED: 2,
	LOADING: 3,
	DONE: 4
}

var IncomingMessage = exports.IncomingMessage = function (xhr, response, mode, resetTimers) {
	var self = this
	stream.Readable.call(self)

	self._mode = mode
	self.headers = {}
	self.rawHeaders = []
	self.trailers = {}
	self.rawTrailers = []

	// Fake the 'close' event, but only once 'end' fires
	self.on('end', function () {
		// The nextTick is necessary to prevent the 'request' module from causing an infinite loop
		process.nextTick(function () {
			self.emit('close')
		})
	})

	if (mode === 'fetch') {
		self._fetchResponse = response

		self.url = response.url
		self.statusCode = response.status
		self.statusMessage = response.statusText
		
		response.headers.forEach(function (header, key){
			self.headers[key.toLowerCase()] = header
			self.rawHeaders.push(key, header)
		})

		if (capability.writableStream) {
			var writable = new WritableStream({
				write: function (chunk) {
					resetTimers(false)
					return new Promise(function (resolve, reject) {
						if (self._destroyed) {
							reject()
						} else if(self.push(Buffer.from(chunk))) {
							resolve()
						} else {
							self._resumeFetch = resolve
						}
					})
				},
				close: function () {
					resetTimers(true)
					if (!self._destroyed)
						self.push(null)
				},
				abort: function (err) {
					resetTimers(true)
					if (!self._destroyed)
						self.emit('error', err)
				}
			})

			try {
				response.body.pipeTo(writable).catch(function (err) {
					resetTimers(true)
					if (!self._destroyed)
						self.emit('error', err)
				})
				return
			} catch (e) {} // pipeTo method isn't defined. Can't find a better way to feature test this
		}
		// fallback for when writableStream or pipeTo aren't available
		var reader = response.body.getReader()
		function read () {
			reader.read().then(function (result) {
				if (self._destroyed)
					return
				resetTimers(result.done)
				if (result.done) {
					self.push(null)
					return
				}
				self.push(Buffer.from(result.value))
				read()
			}).catch(function (err) {
				resetTimers(true)
				if (!self._destroyed)
					self.emit('error', err)
			})
		}
		read()
	} else {
		self._xhr = xhr
		self._pos = 0

		self.url = xhr.responseURL
		self.statusCode = xhr.status
		self.statusMessage = xhr.statusText
		var headers = xhr.getAllResponseHeaders().split(/\r?\n/)
		headers.forEach(function (header) {
			var matches = header.match(/^([^:]+):\s*(.*)/)
			if (matches) {
				var key = matches[1].toLowerCase()
				if (key === 'set-cookie') {
					if (self.headers[key] === undefined) {
						self.headers[key] = []
					}
					self.headers[key].push(matches[2])
				} else if (self.headers[key] !== undefined) {
					self.headers[key] += ', ' + matches[2]
				} else {
					self.headers[key] = matches[2]
				}
				self.rawHeaders.push(matches[1], matches[2])
			}
		})

		self._charset = 'x-user-defined'
		if (!capability.overrideMimeType) {
			var mimeType = self.rawHeaders['mime-type']
			if (mimeType) {
				var charsetMatch = mimeType.match(/;\s*charset=([^;])(;|$)/)
				if (charsetMatch) {
					self._charset = charsetMatch[1].toLowerCase()
				}
			}
			if (!self._charset)
				self._charset = 'utf-8' // best guess
		}
	}
}

inherits(IncomingMessage, stream.Readable)

IncomingMessage.prototype._read = function () {
	var self = this

	var resolve = self._resumeFetch
	if (resolve) {
		self._resumeFetch = null
		resolve()
	}
}

IncomingMessage.prototype._onXHRProgress = function (resetTimers) {
	var self = this

	var xhr = self._xhr

	var response = null
	switch (self._mode) {
		case 'text':
			response = xhr.responseText
			if (response.length > self._pos) {
				var newData = response.substr(self._pos)
				if (self._charset === 'x-user-defined') {
					var buffer = Buffer.alloc(newData.length)
					for (var i = 0; i < newData.length; i++)
						buffer[i] = newData.charCodeAt(i) & 0xff

					self.push(buffer)
				} else {
					self.push(newData, self._charset)
				}
				self._pos = response.length
			}
			break
		case 'arraybuffer':
			if (xhr.readyState !== rStates.DONE || !xhr.response)
				break
			response = xhr.response
			self.push(Buffer.from(new Uint8Array(response)))
			break
		case 'moz-chunked-arraybuffer': // take whole
			response = xhr.response
			if (xhr.readyState !== rStates.LOADING || !response)
				break
			self.push(Buffer.from(new Uint8Array(response)))
			break
		case 'ms-stream':
			response = xhr.response
			if (xhr.readyState !== rStates.LOADING)
				break
			var reader = new __webpack_require__.g.MSStreamReader()
			reader.onprogress = function () {
				if (reader.result.byteLength > self._pos) {
					self.push(Buffer.from(new Uint8Array(reader.result.slice(self._pos))))
					self._pos = reader.result.byteLength
				}
			}
			reader.onload = function () {
				resetTimers(true)
				self.push(null)
			}
			// reader.onerror = ??? // TODO: this
			reader.readAsArrayBuffer(response)
			break
	}

	// The ms-stream case handles end separately in reader.onload()
	if (self._xhr.readyState === rStates.DONE && self._mode !== 'ms-stream') {
		resetTimers(true)
		self.push(null)
	}
}


/***/ }),

/***/ "./node_modules/string_decoder/lib/string_decoder.js":
/*!***********************************************************!*\
  !*** ./node_modules/string_decoder/lib/string_decoder.js ***!
  \***********************************************************/
/***/ ((__unused_webpack_module, exports, __webpack_require__) => {

"use strict";
// Copyright Joyent, Inc. and other Node contributors.
//
// Permission is hereby granted, free of charge, to any person obtaining a
// copy of this software and associated documentation files (the
// "Software"), to deal in the Software without restriction, including
// without limitation the rights to use, copy, modify, merge, publish,
// distribute, sublicense, and/or sell copies of the Software, and to permit
// persons to whom the Software is furnished to do so, subject to the
// following conditions:
//
// The above copyright notice and this permission notice shall be included
// in all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
// OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
// MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN
// NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
// DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
// USE OR OTHER DEALINGS IN THE SOFTWARE.



/*<replacement>*/

var Buffer = (__webpack_require__(/*! safe-buffer */ "./node_modules/safe-buffer/index.js").Buffer);
/*</replacement>*/

var isEncoding = Buffer.isEncoding || function (encoding) {
  encoding = '' + encoding;
  switch (encoding && encoding.toLowerCase()) {
    case 'hex':case 'utf8':case 'utf-8':case 'ascii':case 'binary':case 'base64':case 'ucs2':case 'ucs-2':case 'utf16le':case 'utf-16le':case 'raw':
      return true;
    default:
      return false;
  }
};

function _normalizeEncoding(enc) {
  if (!enc) return 'utf8';
  var retried;
  while (true) {
    switch (enc) {
      case 'utf8':
      case 'utf-8':
        return 'utf8';
      case 'ucs2':
      case 'ucs-2':
      case 'utf16le':
      case 'utf-16le':
        return 'utf16le';
      case 'latin1':
      case 'binary':
        return 'latin1';
      case 'base64':
      case 'ascii':
      case 'hex':
        return enc;
      default:
        if (retried) return; // undefined
        enc = ('' + enc).toLowerCase();
        retried = true;
    }
  }
};

// Do not cache `Buffer.isEncoding` when checking encoding names as some
// modules monkey-patch it to support additional encodings
function normalizeEncoding(enc) {
  var nenc = _normalizeEncoding(enc);
  if (typeof nenc !== 'string' && (Buffer.isEncoding === isEncoding || !isEncoding(enc))) throw new Error('Unknown encoding: ' + enc);
  return nenc || enc;
}

// StringDecoder provides an interface for efficiently splitting a series of
// buffers into a series of JS strings without breaking apart multi-byte
// characters.
exports.StringDecoder = StringDecoder;
function StringDecoder(encoding) {
  this.encoding = normalizeEncoding(encoding);
  var nb;
  switch (this.encoding) {
    case 'utf16le':
      this.text = utf16Text;
      this.end = utf16End;
      nb = 4;
      break;
    case 'utf8':
      this.fillLast = utf8FillLast;
      nb = 4;
      break;
    case 'base64':
      this.text = base64Text;
      this.end = base64End;
      nb = 3;
      break;
    default:
      this.write = simpleWrite;
      this.end = simpleEnd;
      return;
  }
  this.lastNeed = 0;
  this.lastTotal = 0;
  this.lastChar = Buffer.allocUnsafe(nb);
}

StringDecoder.prototype.write = function (buf) {
  if (buf.length === 0) return '';
  var r;
  var i;
  if (this.lastNeed) {
    r = this.fillLast(buf);
    if (r === undefined) return '';
    i = this.lastNeed;
    this.lastNeed = 0;
  } else {
    i = 0;
  }
  if (i < buf.length) return r ? r + this.text(buf, i) : this.text(buf, i);
  return r || '';
};

StringDecoder.prototype.end = utf8End;

// Returns only complete characters in a Buffer
StringDecoder.prototype.text = utf8Text;

// Attempts to complete a partial non-UTF-8 character using bytes from a Buffer
StringDecoder.prototype.fillLast = function (buf) {
  if (this.lastNeed <= buf.length) {
    buf.copy(this.lastChar, this.lastTotal - this.lastNeed, 0, this.lastNeed);
    return this.lastChar.toString(this.encoding, 0, this.lastTotal);
  }
  buf.copy(this.lastChar, this.lastTotal - this.lastNeed, 0, buf.length);
  this.lastNeed -= buf.length;
};

// Checks the type of a UTF-8 byte, whether it's ASCII, a leading byte, or a
// continuation byte. If an invalid byte is detected, -2 is returned.
function utf8CheckByte(byte) {
  if (byte <= 0x7F) return 0;else if (byte >> 5 === 0x06) return 2;else if (byte >> 4 === 0x0E) return 3;else if (byte >> 3 === 0x1E) return 4;
  return byte >> 6 === 0x02 ? -1 : -2;
}

// Checks at most 3 bytes at the end of a Buffer in order to detect an
// incomplete multi-byte UTF-8 character. The total number of bytes (2, 3, or 4)
// needed to complete the UTF-8 character (if applicable) are returned.
function utf8CheckIncomplete(self, buf, i) {
  var j = buf.length - 1;
  if (j < i) return 0;
  var nb = utf8CheckByte(buf[j]);
  if (nb >= 0) {
    if (nb > 0) self.lastNeed = nb - 1;
    return nb;
  }
  if (--j < i || nb === -2) return 0;
  nb = utf8CheckByte(buf[j]);
  if (nb >= 0) {
    if (nb > 0) self.lastNeed = nb - 2;
    return nb;
  }
  if (--j < i || nb === -2) return 0;
  nb = utf8CheckByte(buf[j]);
  if (nb >= 0) {
    if (nb > 0) {
      if (nb === 2) nb = 0;else self.lastNeed = nb - 3;
    }
    return nb;
  }
  return 0;
}

// Validates as many continuation bytes for a multi-byte UTF-8 character as
// needed or are available. If we see a non-continuation byte where we expect
// one, we "replace" the validated continuation bytes we've seen so far with
// a single UTF-8 replacement character ('\ufffd'), to match v8's UTF-8 decoding
// behavior. The continuation byte check is included three times in the case
// where all of the continuation bytes for a character exist in the same buffer.
// It is also done this way as a slight performance increase instead of using a
// loop.
function utf8CheckExtraBytes(self, buf, p) {
  if ((buf[0] & 0xC0) !== 0x80) {
    self.lastNeed = 0;
    return '\ufffd';
  }
  if (self.lastNeed > 1 && buf.length > 1) {
    if ((buf[1] & 0xC0) !== 0x80) {
      self.lastNeed = 1;
      return '\ufffd';
    }
    if (self.lastNeed > 2 && buf.length > 2) {
      if ((buf[2] & 0xC0) !== 0x80) {
        self.lastNeed = 2;
        return '\ufffd';
      }
    }
  }
}

// Attempts to complete a multi-byte UTF-8 character using bytes from a Buffer.
function utf8FillLast(buf) {
  var p = this.lastTotal - this.lastNeed;
  var r = utf8CheckExtraBytes(this, buf, p);
  if (r !== undefined) return r;
  if (this.lastNeed <= buf.length) {
    buf.copy(this.lastChar, p, 0, this.lastNeed);
    return this.lastChar.toString(this.encoding, 0, this.lastTotal);
  }
  buf.copy(this.lastChar, p, 0, buf.length);
  this.lastNeed -= buf.length;
}

// Returns all complete UTF-8 characters in a Buffer. If the Buffer ended on a
// partial character, the character's bytes are buffered until the required
// number of bytes are available.
function utf8Text(buf, i) {
  var total = utf8CheckIncomplete(this, buf, i);
  if (!this.lastNeed) return buf.toString('utf8', i);
  this.lastTotal = total;
  var end = buf.length - (total - this.lastNeed);
  buf.copy(this.lastChar, 0, end);
  return buf.toString('utf8', i, end);
}

// For UTF-8, a replacement character is added when ending on a partial
// character.
function utf8End(buf) {
  var r = buf && buf.length ? this.write(buf) : '';
  if (this.lastNeed) return r + '\ufffd';
  return r;
}

// UTF-16LE typically needs two bytes per character, but even if we have an even
// number of bytes available, we need to check if we end on a leading/high
// surrogate. In that case, we need to wait for the next two bytes in order to
// decode the last character properly.
function utf16Text(buf, i) {
  if ((buf.length - i) % 2 === 0) {
    var r = buf.toString('utf16le', i);
    if (r) {
      var c = r.charCodeAt(r.length - 1);
      if (c >= 0xD800 && c <= 0xDBFF) {
        this.lastNeed = 2;
        this.lastTotal = 4;
        this.lastChar[0] = buf[buf.length - 2];
        this.lastChar[1] = buf[buf.length - 1];
        return r.slice(0, -1);
      }
    }
    return r;
  }
  this.lastNeed = 1;
  this.lastTotal = 2;
  this.lastChar[0] = buf[buf.length - 1];
  return buf.toString('utf16le', i, buf.length - 1);
}

// For UTF-16LE we do not explicitly append special replacement characters if we
// end on a partial character, we simply let v8 handle that.
function utf16End(buf) {
  var r = buf && buf.length ? this.write(buf) : '';
  if (this.lastNeed) {
    var end = this.lastTotal - this.lastNeed;
    return r + this.lastChar.toString('utf16le', 0, end);
  }
  return r;
}

function base64Text(buf, i) {
  var n = (buf.length - i) % 3;
  if (n === 0) return buf.toString('base64', i);
  this.lastNeed = 3 - n;
  this.lastTotal = 3;
  if (n === 1) {
    this.lastChar[0] = buf[buf.length - 1];
  } else {
    this.lastChar[0] = buf[buf.length - 2];
    this.lastChar[1] = buf[buf.length - 1];
  }
  return buf.toString('base64', i, buf.length - n);
}

function base64End(buf) {
  var r = buf && buf.length ? this.write(buf) : '';
  if (this.lastNeed) return r + this.lastChar.toString('base64', 0, 3 - this.lastNeed);
  return r;
}

// Pass bytes on through for single-byte encodings (e.g. ascii, latin1, hex)
function simpleWrite(buf) {
  return buf.toString(this.encoding);
}

function simpleEnd(buf) {
  return buf && buf.length ? this.write(buf) : '';
}

/***/ }),

/***/ "./img/filp.svg":
/*!**********************!*\
  !*** ./img/filp.svg ***!
  \**********************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = ("data:image/svg+xml;base64,PD94bWwgdmVyc2lvbj0iMS4wIiBzdGFuZGFsb25lPSJubyI/PjwhRE9DVFlQRSBzdmcgUFVCTElDICItLy9XM0MvL0RURCBTVkcgMS4xLy9FTiIgImh0dHA6Ly93d3cudzMub3JnL0dyYXBoaWNzL1NWRy8xLjEvRFREL3N2ZzExLmR0ZCI+PHN2ZyB0PSIxNjczOTc5NTQ2Mjk4IiBjbGFzcz0iaWNvbiIgdmlld0JveD0iMCAwIDEwMjQgMTAyNCIgdmVyc2lvbj0iMS4xIiB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHAtaWQ9IjI1MzkiIGlkPSJteF9uXzE2NzM5Nzk1NDYyOTkiIHhtbG5zOnhsaW5rPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5L3hsaW5rIiB3aWR0aD0iMjAwIiBoZWlnaHQ9IjIwMCI+PHBhdGggZD0iTTUxMiA3MTkuMzZjNy42OCAwIDEyLjgtMi41NiAxNy45Mi03LjY4bDE4MS43Ni0xODEuNzZjMTAuMjQtMTAuMjQgMTAuMjQtMjUuNiAwLTM1Ljg0bC0xODEuNzYtMTgxLjc2Yy01LjEyLTUuMTItMTIuOC03LjY4LTE3LjkyLTcuNjhzLTEyLjggMi41Ni0xNy45MiA3LjY4bC0xODEuNzYgMTgxLjc2Yy0xMC4yNCAxMC4yNC0xMC4yNCAyNS42IDAgMzUuODRsMTgxLjc2IDE4MS43NmM1LjEyIDUuMTIgMTAuMjQgNy42OCAxNy45MiA3LjY4eiBtMC0zNTMuMjhsMTQ1LjkyIDE0NS45Mi0xNDUuOTIgMTQ1LjkyLTE0NS45Mi0xNDUuOTIgMTQ1LjkyLTE0NS45MnoiIHAtaWQ9IjI1NDAiIGZpbGw9IiNlNmU2ZTYiPjwvcGF0aD48cGF0aCBkPSJNNTEyIDUxLjJjLTEwMi40IDAtMjAyLjI0IDMzLjI4LTI4MS42IDk3LjI4VjEwMi40YzAtMTIuOC0xMC4yNC0yNS42LTI1LjYtMjUuNi0xMi44IDAtMjUuNiAxMC4yNC0yNS42IDI1LjZ2OTkuODRjMCA3LjY4IDIuNTYgMTUuMzYgNy42OCAyMC40OCAyLjU2IDIuNTYgNS4xMiA1LjEyIDEwLjI0IDUuMTIgMi41NiAwIDUuMTIgMi41NiA3LjY4IDIuNTZoMTAyLjRjMTIuOCAwIDI1LjYtMTAuMjQgMjUuNi0yNS42IDAtMTIuOC0xMC4yNC0yNS42LTI1LjYtMjUuNmgtMzMuMjhjNjkuMTItNDguNjQgMTUxLjA0LTc2LjggMjM4LjA4LTc2LjggMjI1LjI4IDAgNDA5LjYgMTg0LjMyIDQwOS42IDQwOS42IDAgMTUuMzYgMTAuMjQgMjUuNiAyNS42IDI1LjZzMjUuNi0xMC4yNCAyNS42LTI1LjZjMC0yNTMuNDQtMjA3LjM2LTQ2MC44LTQ2MC44LTQ2MC44ek04MzcuMTIgODAxLjI4Yy01LjEyLTUuMTItMTIuOC03LjY4LTIwLjQ4LTcuNjhoLTEwMi40Yy0xMi44IDAtMjUuNiAxMC4yNC0yNS42IDI1LjYgMCAxMi44IDEwLjI0IDI1LjYgMjUuNiAyNS42aDMzLjI4Yy02OS4xMiA0OC42NC0xNTEuMDQgNzYuOC0yMzguMDggNzYuOC0yMjUuMjggMC00MDkuNi0xODQuMzItNDA5LjYtNDA5LjYgMC0xNS4zNi0xMC4yNC0yNS42LTI1LjYtMjUuNnMtMjUuNiAxMC4yNC0yNS42IDI1LjZjMCAyNTMuNDQgMjA3LjM2IDQ2MC44IDQ2MC44IDQ2MC44IDEwMi40IDAgMjAyLjI0LTMzLjI4IDI4MS42LTk3LjI4djQ2LjA4YzAgMTIuOCAxMC4yNCAyNS42IDI1LjYgMjUuNiAxMi44IDAgMjUuNi0xMC4yNCAyNS42LTI1LjZ2LTEwMi40YzIuNTYtNS4xMiAwLTEyLjgtNS4xMi0xNy45MnoiIHAtaWQ9IjI1NDEiIGZpbGw9IiNlNmU2ZTYiPjwvcGF0aD48L3N2Zz4=");

/***/ }),

/***/ "./img/indicator.svg":
/*!***************************!*\
  !*** ./img/indicator.svg ***!
  \***************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = ("data:image/svg+xml;base64,PHN2ZyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAyMiAyMiI+DQogICAgPHBhdGggZD0iTTE2LjExOCAzLjY2N2guMzgyYTMuNjY3IDMuNjY3IDAgMDEzLjY2NyAzLjY2N3Y3LjMzM2EzLjY2NyAzLjY2NyAwIDAxLTMuNjY3IDMuNjY3aC0xMWEzLjY2NyAzLjY2NyAwIDAxLTMuNjY3LTMuNjY3VjcuMzMzQTMuNjY3IDMuNjY3IDAgMDE1LjUgMy42NjZoLjM4Mkw0Ljk1IDIuMDUzYTEuMSAxLjEgMCAwMTEuOTA2LTEuMWwxLjU2NyAyLjcxNGg1LjE1NkwxNS4xNDYuOTUzYTEuMTAxIDEuMTAxIDAgMDExLjkwNiAxLjFsLS45MzQgMS42MTR6IiBmaWxsPSIjMzMzIj48L3BhdGg+DQogICAgPHBhdGggZD0iTTUuNTYxIDUuMTk0aDEwLjg3OGEyLjIgMi4yIDAgMDEyLjIgMi4ydjcuMjExYTIuMiAyLjIgMCAwMS0yLjIgMi4ySDUuNTYxYTIuMiAyLjIgMCAwMS0yLjItMi4yVjcuMzk0YTIuMiAyLjIgMCAwMTIuMi0yLjJ6IiBmaWxsPSIjZmZmIj48L3BhdGg+DQogICAgPHBhdGggZD0iTTYuOTY3IDguNTU2YTEuMSAxLjEgMCAwMTEuMSAxLjF2Mi42ODlhMS4xIDEuMSAwIDExLTIuMiAwVjkuNjU2YTEuMSAxLjEgMCAwMTEuMS0xLjF6TTE1LjAzMyA4LjU1NmExLjEgMS4xIDAgMDExLjEgMS4xdjIuNjg5YTEuMSAxLjEgMCAxMS0yLjIgMFY5LjY1NmExLjEgMS4xIDAgMDExLjEtMS4xeiIgZmlsbD0iIzMzMyI+PC9wYXRoPg0KPC9zdmc+");

/***/ }),

/***/ "./img/ploading.gif":
/*!**************************!*\
  !*** ./img/ploading.gif ***!
  \**************************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = ("data:image/gif;base64,R0lGODlhWgBaALMOAHR0dAICAnd3dwEBAXh4eAMDAwkJCQ0NDQsLCxwcHA4ODggICHl5eQAAAAAAAAAAACH/C05FVFNDQVBFMi4wAwEAAAAh/wtYTVAgRGF0YVhNUDw/eHBhY2tldCBiZWdpbj0i77u/IiBpZD0iVzVNME1wQ2VoaUh6cmVTek5UY3prYzlkIj8+IDx4OnhtcG1ldGEgeG1sbnM6eD0iYWRvYmU6bnM6bWV0YS8iIHg6eG1wdGs9IkFkb2JlIFhNUCBDb3JlIDUuNi1jMTMyIDc5LjE1OTI4NCwgMjAxNi8wNC8xOS0xMzoxMzo0MCAgICAgICAgIj4gPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4gPHJkZjpEZXNjcmlwdGlvbiByZGY6YWJvdXQ9IiIgeG1sbnM6eG1wTU09Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9tbS8iIHhtbG5zOnN0UmVmPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvc1R5cGUvUmVzb3VyY2VSZWYjIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtcE1NOk9yaWdpbmFsRG9jdW1lbnRJRD0ieG1wLmRpZDpiYWE1ODg5ZS1jN2RmLTRmZmUtYjkzOS0wMmVkMTZhNmNjZDIiIHhtcE1NOkRvY3VtZW50SUQ9InhtcC5kaWQ6M0I2ODI2NjA1NzhGMTFFNkEyMEVDNzhEOUY1RkQxRjgiIHhtcE1NOkluc3RhbmNlSUQ9InhtcC5paWQ6M0I2ODI2NUY1NzhGMTFFNkEyMEVDNzhEOUY1RkQxRjgiIHhtcDpDcmVhdG9yVG9vbD0iQWRvYmUgUGhvdG9zaG9wIENDIDIwMTUuNSAoTWFjaW50b3NoKSI+IDx4bXBNTTpEZXJpdmVkRnJvbSBzdFJlZjppbnN0YW5jZUlEPSJ4bXAuaWlkOjljYjgzNjY2LWYxYWUtNGMyZi1hMGEwLThhODJmYjIxM2U0MyIgc3RSZWY6ZG9jdW1lbnRJRD0iYWRvYmU6ZG9jaWQ6cGhvdG9zaG9wOmU1NDE3YzFmLTllODAtMTE3OS04NjdiLWUyN2Y3M2VkMTZkOSIvPiA8L3JkZjpEZXNjcmlwdGlvbj4gPC9yZGY6UkRGPiA8L3g6eG1wbWV0YT4gPD94cGFja2V0IGVuZD0iciI/PgH//v38+/r5+Pf29fTz8vHw7+7t7Ovq6ejn5uXk4+Lh4N/e3dzb2tnY19bV1NPS0dDPzs3My8rJyMfGxcTDwsHAv769vLu6ubi3trW0s7KxsK+urayrqqmop6alpKOioaCfnp2cm5qZmJeWlZSTkpGQj46NjIuKiYiHhoWEg4KBgH9+fXx7enl4d3Z1dHNycXBvbm1sa2ppaGdmZWRjYmFgX15dXFtaWVhXVlVUU1JRUE9OTUxLSklIR0ZFRENCQUA/Pj08Ozo5ODc2NTQzMjEwLy4tLCsqKSgnJiUkIyIhIB8eHRwbGhkYFxYVFBMSERAPDg0MCwoJCAcGBQQDAgEAACH5BAkKAA4ALAAAAABaAFoAAAT/0MlJq7046827/2AojmRpnmiqrmzrvnAsz3Ta3HW+3bjuV7wbg/H7BYXEYu7YGCaVjuDr6Hwqjy2qEzphNlTaIZfi/ZqY2zHZW0KL1RVGeRS2wiXD+ad+x8jZHXx9GX9MO2GDG3mGGG52iX5ojUFVRWWXmJmam1IknJ+goXoioqWmnHSnqquUpDxVsLGys7S1tk6Uj4dIt72+v7K5IcKQF8R7r1asPC7HHs7L0Z3Ogclr0tES1BzH2NiLSMPWUcnAsd7gTboaxLnm77e527vq2uMm8FXy98/j8z77woFoxw9Fp2pI/mUgKBDMQXrp3iATqNBeD3rMIBaqN9BfwWsZ/7kBmpTwo0aLHIF4kchupIWAKftRLHgpDYeND7skq2jMY0NyjlgqwnlRZ8mfCDlCqyO0A1E7MJueBBrTnc0RG1lGXbfQZ0w8sFLEAhmRK0khKJtWConv6lZXaKlKNWpmyk6TJxVqoWvw7iu49fQyLOrJWitx4QTzQhnX4sTAeLsmjuyO8cWcLScjFan5K9kkl9KapSuG50vDlFtlkjtaNGvEkDeDXIlprsrOts+WjkzVUZmrkmN7zsu7dzkiK3OTRl78NO7WQenK7vkc9u7pt9UJrZz0+vDMwpVPGGuBPOfwrbO/8SbNu3j1oNkvc5/+s3T5oraYhn8f/6e1zfFHQZY+BPbiWkdIFajgLMs9ZgoX+1nmn0upYOfchPK95iCG+L034HHpAAAAh6V4OOAsQYh4hAAC3EJAMO3VV55WmLBYiwAv+pKiirzoE+CGDbAoZFu4eCGiiOdYCBgPQrK4wiVHXlDJk0w4mUUZAGgAXApNDtmMkVn+0KWVLhxppojFsHBmlGm26eabcMYp55x01mnnnXjSEAEAIfkECQoADgAsAAAAAFoAWgAABP/QyUmrvTjrzbv/YCiOZGmeaKqubOu+cCxbTT3fX23jfKbXDEZvSPk1gkSiMShM8pZMZ1HngjalEiPLip1qUdZr12Hcmbhj73eETlMY5QbpxxS7JfCyqH2v5NccfH1+cR6CgxVxchqHiIR6GGGOG38/kXRIkxyVR2SKn6CFbKGkpaangDmoq6ytdh2tsbKhmaqzpHW5uru8lbWGOr3Cw8TFvJghmMbLUVvBrxvImmq/sM+Jt9mWE9K2R2La4VRZ1yDdnuLaTefWQFfp4kjsgeXo38z4xMjz0fXK+QB37asHzB03gkqu8WPkD+GQgQbNNYxIY1zBRf0MLvQx8duFVBn/zXDU6DBkrXllqo3cdokkxYsnEaZUudJiRZceJVI8N5NmTZsHcfpcGXOnpBCcRJITCo2h0afuhlISFdTjxpY57RW91/QDJ3AKSzrNyi5XCl/vwr5sB6RqtToZPZhdalXs2LRrvXwEqgEuXXl2ia4Ty7Jq1qX0mO4pt7DbTDWJ6+aNnKkx48duu2oVchVrZcKXqW6+yzmw58GTR4eie/im5NZsP6emhUsqRNiUUeO2F4cJqGanRycLvfs3kqSmb0sNLnxvzyvIlyvXzPyqMD9oqU9fHLEzLwy7TJZOTfovdRF+y3d2ThZeNsDkBRt272r87vLz6Z+Kst71L/2x8GcapHsqBWTgMOZxB9uBDIbXnDcAAricaxFKeB42FUY4IYbpHJCAAgYYsEABoQRgYgA6nIiicaNw9VUQ/zCoIhMz6tKfBcco84MAAjTIQI0/npjjfZtspQiPKdTogJI4AgcGKDwKoIKQKprok5MpCFBGlEgmWSWTPPBoBJctBCkklmFyKaULZgYwVxJqwnDmhtN4hWadJtyJ55589unnn4AGKuigG0QAACH5BAkKAA4ALAAAAABaAFoAAAT/0MlJq7046827/2AojmRpnmiqrmzrvnAsW009319t43ym1wxGb0j5NYJEojEoTPKWTOfzF5VOfi1o0yoxNlTaLdf7PXmr3CvZpE1bGORyKOx+x+Udev0CX3vaexhBcRyAgRmDXhp6U3ghfVQ+VEhDd5aXmJmajjmbnp+gOiShpKWWlCCmqqtHYn86TLGys7S1tre3kyKTuL2+v7mwrnnChxi6c8VqrFnKqc4OrNJ4yM9AYtPSy627xaLZzNHQr9fiNsC22ULVneWR6PC/k+zk3PQk8bHz48Tu/ErC7hXy9q9SwIKLCJZbgaVeooXW7CHsImpgww37ILaTqLGIIotG/zAe7NiPI7cKd1BJ8nMh48mISNidUomIUMuRLzfGdJYyloeHsG5eEyjSH8Qzsx5Byklx6MRjCl/y8olPFg2cNEvu7Jg0BVWPTkmC3MrUZ8VlYLAOG7sO4cWmnEa4zMrWHFOwJ5egmLu2qEm6dvMS1bkUcEKjd+FWOYNyh9bCfQ//7csr8Fm7hAevnAx14Z1thvECDd2Zs1DBmkgrhtwNMWBFS2ZGFs06mevIsnvOBl0bZluxcJEKWfqUt+bSZEP3nDBa9erjp5PvBtrKFS2HvQnbVW2LDxq20K+6BqeqeWuT5Mtnxx48/SdZ4RtDJO4+tb7i4nPm2/+L9/ms/AVYi3Z/tiVmTHx41Vefc/IpuOBupzmYHoMJZoPAAQ5SWEEtwkCGDgHBXOMhEwhuSEuHU+FCAIi48JIiJN/ttVxaQBjCXIxmKPKVjHHgiGOOYUAYERQGlbhRZUNc98KLSXT1woDGtLBjlFRWaeWVWGap5ZZcdumlChEAACH5BAkKAA4ALAAAAABaAFoAAAT/0MlJq7046827/2AojmRpnmiqrmzrvnAsW009319t43ym1wxGbzj5AYXEoTGITOKWTOfzx2xKHb8W1Ho1NlTba8X7PW25YnJ55I2KLQw1mxp8Y+LyHL1uv6vXGmd9GkF5gXuDHIVkPoiJHYtGF4JEf5aXmJmZJZqdnp+GIaCjpJloHKWpqlCiOlWvsLGys7S1THStR7a7vL2xuCDAj5Oup4dHwxjCHsvJRcW5DXxYqzouzajQ1NXc1hLYG8Dd4zvgx9JI5ONC5o3I2+i+sOSR6NF8e/L6tXTtytr+POyr0k+bnncBbxR8d9AevGkpsjArlnAMQIPEdmTzthFdRQri/zA+k9SBUTiKIjviE/kHojuS/458HOlwCc02Ll9KzOgx5UmGVL6pefUBjxee9XL+dPiQ3Rk3RY3upJnU2MtpuPJBFSEViEWZPs9hNQhLhVSIC5lOBErWDccSZal2tRqTqbmpQgFB2prV1b0mdzHCBKlRrFOGJQ0GRjw4r966dWY2lQvZCivKhic3HIuYMGObjtUiPSx6KWevoz3vdHSVtNLWqo0JGwovnmnXdGkI9smo09bKmteqBafqd2rJyxYP7Vost+PIYYFjvsBcF5kqiZEh3925wq3mSGQJh949881UEmJtxh2Mu7pU5EubD/1+VPzX0unX10R0e2fw+nDTX5l0qXk30IG0xOacgtThtxeCaBGoW3nO+LffhRTWheF+Dk64IYcLqjZOACQGUEABJA6g4orVdNjgPDoAIKNWENYDwA8y/pIhITA2cCNONRIUI44A2MLCj17IyIIRSCp5h4sgIBmjjAAsqYaTPTBJZZUrXInlEFt+CQYZVDoRZgwzTunMCmGWuWYLW74p55x01mnnnXjmqecNEQAAIfkEBQoADgAsAAAAAFoAWgAABP/QyUmrvTjrzbv/YCiOZGmeaKqubOu+cCzP9Nrcdb7duO5XvBuD8fsFhcSi7jhMKh3BF7P5lBxb02HVelUdG9Qtt3vKiinfRjnYdJ6h31LW/XYw0qJvuE65x0FmfBdpah1zghh+ZBmHiIl4jIGOGYo8kTx7NYSbnJ2en5t5oKOkpaEhpqmqo1ofq6+wenQ7mG22t7i5uru8TWyzGr+9w8TFub+AtZODyq7NY7AwyM5IaLHXURPTHttw2NjW1dRgSd/m5NCt3M3Cxm3nrd0c0+3u9rr0z4bs+iP37/zErRMnz0hAdOPi9TNIcGGwg+qY9dg38WFDgRQRFtSWLWNFiRr/HV4KiTGcSCCEMOQrSesiQpO1gIH8Y1IhS4skX0KLKfNCpUXeckacB9FNmjYhhhytWe4kSJsvZWUC4evZSp1EXaqb09MDLo5au1pgks5JOxZIy6pNJmTtBFuXUCCbgkqZPKSWUH4UVU1Y3b4nO4IFY2JuzL/oNroNOpRtYqcq7QbWRzZcRi2KcWJ2ygZmIWiXiWQeufnm4M8pFz8VDXk1466lVI+VbFoz69o7SzXW+xh3ZMC+g07ZNHU2cKyhX3f4aYtQcd6lkWftLd1nrtzVmSpHHH03pbRv/bak7t228HOvuot1jR6ber7U2696z72p/Pm3s5N28q9/se2OReTffYC7AJjQemeMZpx8MSjI230Q6rdghPeV9xuF8lmYyC48kcMLhhJuiIsw9RDo3IitOQbJBsawwQMAANxymA2cwEjjFzZapyFiL8Lo4404iojFET4WCSSRPgRRJABSEJKjDksyGUOUP/pg5AxUPrnMllx26eWXYIYp5phkjhkBADs=");

/***/ }),

/***/ "./img/state.png":
/*!***********************!*\
  !*** ./img/state.png ***!
  \***********************/
/***/ ((__unused_webpack_module, __webpack_exports__, __webpack_require__) => {

"use strict";
__webpack_require__.r(__webpack_exports__);
/* harmony export */ __webpack_require__.d(__webpack_exports__, {
/* harmony export */   "default": () => (__WEBPACK_DEFAULT_EXPORT__)
/* harmony export */ });
/* harmony default export */ const __WEBPACK_DEFAULT_EXPORT__ = ("data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAQsAAAELCAYAAADOVaNSAAAACXBIWXMAAC4jAAAuIwF4pT92AAAGaGlUWHRYTUw6Y29tLmFkb2JlLnhtcAAAAAAAPD94cGFja2V0IGJlZ2luPSLvu78iIGlkPSJXNU0wTXBDZWhpSHpyZVN6TlRjemtjOWQiPz4gPHg6eG1wbWV0YSB4bWxuczp4PSJhZG9iZTpuczptZXRhLyIgeDp4bXB0az0iQWRvYmUgWE1QIENvcmUgNS42LWMxNDUgNzkuMTYzNDk5LCAyMDE4LzA4LzEzLTE2OjQwOjIyICAgICAgICAiPiA8cmRmOlJERiB4bWxuczpyZGY9Imh0dHA6Ly93d3cudzMub3JnLzE5OTkvMDIvMjItcmRmLXN5bnRheC1ucyMiPiA8cmRmOkRlc2NyaXB0aW9uIHJkZjphYm91dD0iIiB4bWxuczp4bXA9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC8iIHhtbG5zOmRjPSJodHRwOi8vcHVybC5vcmcvZGMvZWxlbWVudHMvMS4xLyIgeG1sbnM6cGhvdG9zaG9wPSJodHRwOi8vbnMuYWRvYmUuY29tL3Bob3Rvc2hvcC8xLjAvIiB4bWxuczp4bXBNTT0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wL21tLyIgeG1sbnM6c3RFdnQ9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZUV2ZW50IyIgeG1wOkNyZWF0b3JUb29sPSJBZG9iZSBQaG90b3Nob3AgQ0MgMjAxOSAoTWFjaW50b3NoKSIgeG1wOkNyZWF0ZURhdGU9IjIwMTktMDUtMDZUMjE6Mzk6MzErMDg6MDAiIHhtcDpNb2RpZnlEYXRlPSIyMDE5LTA1LTA2VDIxOjQwOjU1KzA4OjAwIiB4bXA6TWV0YWRhdGFEYXRlPSIyMDE5LTA1LTA2VDIxOjQwOjU1KzA4OjAwIiBkYzpmb3JtYXQ9ImltYWdlL3BuZyIgcGhvdG9zaG9wOkNvbG9yTW9kZT0iMyIgcGhvdG9zaG9wOklDQ1Byb2ZpbGU9InNSR0IgSUVDNjE5NjYtMi4xIiB4bXBNTTpJbnN0YW5jZUlEPSJ4bXAuaWlkOjNhN2I0MGQwLTlkN2ItNDAwOS04YmMwLTY1NjZmY2I2OGQ5MyIgeG1wTU06RG9jdW1lbnRJRD0iYWRvYmU6ZG9jaWQ6cGhvdG9zaG9wOjllYTQ0NDEzLTA5YWMtNGE0YS05OGI4LTZmMjQ1ZTViYmI4NiIgeG1wTU06T3JpZ2luYWxEb2N1bWVudElEPSJ4bXAuZGlkOjlmZmM1YzJkLTA4ODEtNGU2My1hYTdhLWJmMDhiZTU3YzQ5ZSI+IDx4bXBNTTpIaXN0b3J5PiA8cmRmOlNlcT4gPHJkZjpsaSBzdEV2dDphY3Rpb249ImNyZWF0ZWQiIHN0RXZ0Omluc3RhbmNlSUQ9InhtcC5paWQ6OWZmYzVjMmQtMDg4MS00ZTYzLWFhN2EtYmYwOGJlNTdjNDllIiBzdEV2dDp3aGVuPSIyMDE5LTA1LTA2VDIxOjM5OjMxKzA4OjAwIiBzdEV2dDpzb2Z0d2FyZUFnZW50PSJBZG9iZSBQaG90b3Nob3AgQ0MgMjAxOSAoTWFjaW50b3NoKSIvPiA8cmRmOmxpIHN0RXZ0OmFjdGlvbj0iY29udmVydGVkIiBzdEV2dDpwYXJhbWV0ZXJzPSJmcm9tIGFwcGxpY2F0aW9uL3ZuZC5hZG9iZS5waG90b3Nob3AgdG8gaW1hZ2UvcG5nIi8+IDxyZGY6bGkgc3RFdnQ6YWN0aW9uPSJzYXZlZCIgc3RFdnQ6aW5zdGFuY2VJRD0ieG1wLmlpZDozYTdiNDBkMC05ZDdiLTQwMDktOGJjMC02NTY2ZmNiNjhkOTMiIHN0RXZ0OndoZW49IjIwMTktMDUtMDZUMjE6NDA6NTUrMDg6MDAiIHN0RXZ0OnNvZnR3YXJlQWdlbnQ9IkFkb2JlIFBob3Rvc2hvcCBDQyAyMDE5IChNYWNpbnRvc2gpIiBzdEV2dDpjaGFuZ2VkPSIvIi8+IDwvcmRmOlNlcT4gPC94bXBNTTpIaXN0b3J5PiA8L3JkZjpEZXNjcmlwdGlvbj4gPC9yZGY6UkRGPiA8L3g6eG1wbWV0YT4gPD94cGFja2V0IGVuZD0iciI/Ph7aCJkAADakSURBVHic7Z15mBTV1f+/VdVd3T09+8DAMMzCNjAgDPu+CQgq/lzAgBolRN8YQUTDi7yuPC4JwdckAhoNcUtIBH0TFMH4aCIGBRQBwxNBEFDRwCCLCA7M0vvvj+4abp0uuqt7umd6ps/neUo53VW3aupWnb7n3HPOlQKBABiGYaIht/QFMAzTOmBlwTCMKVhZMAxjClYWDMOYgpUFwzCmYGXBMIwpWFkwDMMwDMMwDNPMSC19ASlKS9wXDqWNjebuo7TvH/ZZMK0R/pFrAVhZMAxjClYWDMOYwtLSF9BCRBvG0u+N9o9lKGzG3vUnoI22QnP3DxB+f2O9322+f9JRWZh5EKM9jPGMyKI9jPQcRt+3+QfyAsTaH/H4NIz6J9r9Fr9v8/3DZgjDMKZgZcEwjCnS0QwxItZhbqzDXqPhKfVRmBnySkRui5gxA+k+Rv0RqylC+4N9SARWFsGHSo4gA+H3SUHkh9UI+vB5Da6D7i8+kNSGNrKRW+sDLJF/03uhGMhSBFk2aINC7xXtDx/Zh8pmaK39YQibIQzDmIKVBcMwpmBlwTCMKdLBZxHNOSZBbxPLAKzkGAeRVegVrcXgPCIBBG1eEReRG4jsht7P4SdtmLGHU9FmNuoP6jOiPgobken9txJZMWhDxEx/uKC//x5E749UvN8Jg0cWDMOYIh1GFgnn7rvvbvfDH/5wdKdOnQY4HI4Sv98vBwIB+Hy+ujNnzhx45513Nv70pz890NLXmc6sXLmyYuzYsSMLCwurFEXJ0D53u92nvvvuu/1PP/306ytWrPi2Ja+xtZEOqb4JNUN27949pbS0dJosy07tM01ZiNTU1By48cYbn3n//ffrkDwzJNrUaSoOi5NqhkyZMiXjqaeeuqF9+/ajIl6EJAW++OKL9QMHDlwf+igRZkissTOtinRUFjL0DxpVDgrCfRTZALB3795ZnTp1GmGz2TyyLDc+GBaLxSfLcuOD4ff7Zb/fL/l8vq8ef/zxhx988EE/gE5RrvMskeugfxi9CD6wGj7E/rC2xMMbLWBKgn6EKyNcOWQSOQN6haICkKuqqhyvvvrq3Hbt2nVQVbXxXsmyHBDlQCAg+f1+CQA8Hs8/HQ7HkwBKyTmosqhH5P4w+kGIFtjVqmCfhUlWr15d1alTpxGxHKMoSvmiRYv+96WXXipP0mUxIebPn1+8YcOG/87Ozi6O5Tir1XpxTU3NFcm6rrYEKwuTTJw48QfxHCfLcvurrrrq/oULF8b0EDPmueuuu4rvvffeuQ6HIz+e4x0Ox8yxY8dmRN8zvYk0vdRWaLLP4umnn+47YMCACZpssVj8kiQ1DullWQ5I0vkmA4GAFAgExHOogwYNGgZg3wcffEDNDQ03kT3Qmw1+6Ie1ZnwWqYCZ+099FtTxrhK50Udxzz33dF64cOF8i8XS+LIriuJXFKXxXkmSBFEGoOsfSZLUzMzMw3/961+PCfvQ8G4vIvcH0Dr6I27SdTYkkrKwgDycFRUVfT0eT+O9cjgcLovF0mifiv4KAJAkya8oivgwBnJzc5WHH374JzNnzvzdkCFDPgLQlVxTHpEt0OcruKB3whnZzKIMNE8yVDxJeFQ5i/fbyGeUS2QrAPmpp57qM3369OszMzNlRVHqGxsgygLk75QkKSDL+ssqLi4uB/ApOQ9VFjSRL5LPqE0pCiA9lUWsv3Twer0KGSmEKQhdg8FRhjjygMVi8VksFnXAgAHzv//++ydzcnJOksOMfj1FfNArDz+5Tj9S12EdKcmL3m8F4c8lvReWF154YcjUqVNnAoAsy3VUeYsOaMMLEkaGABByeNLr0I1GDK47Ve93UmCfRQvgdDrv+POf/zy0pa+jtfLiiy8O1hQF03ywsjDBgQMHqhPd5pQpU65ft27duES329Z57bXXxl5++eXXJbrdL7744kii22xrpMMwyqhQCrWZRRvZglBchUbv3r3L3n777XusVqsDAPLy8mpUVaX1DyKizetrNDQ02ADg6NGj23r06PEMgBJyCHWo1Yc2DRf0TlEfjJ2kussgcjwOuVh9FLTWhAX6+099FBYAWaSNbAD48MMPr+3SpcsQu93uFs1Au93uJg7NADUzolArSdJj5DMahFUD/f10Q3+//Qj3IbWpOIt08FnEU3VJN0u0d+9e344dO7YNHz58AhCc7Yj1IqiPIyMjowEAunfv3j8QCMyQJGkrOcSO8FmCSC+3jPCgoFiDhBJRdNaocA1VFhYi2wXZCsAJPeqmTZuml5WVDfT7/VBV1SMq69BsVFOuew3C/SL0ZafBfOnwQ6uDzRCTTJs2beOBAwf+laTmJ27fvv36JLXd6tm0adP0ioqKgUlqfqMkSeuj78awsoiB8ePHr02WwigvLx/y+eef/3dVVRWdNkxb+vbt6zh48OCCJCuK5Ulqu82RikOpRF9TNDOD2sxWhM/r6/I69uzZM7Fbt26ND7Cqqp5IU6lmqK2ttQPA2bNnq3/yk58seeONN7KgH557oTcr6hHMH9HwEVnbR8SMDyPa3xHNR0GH8zTpS4V+mpj6iBQAzsrKSvu6det+kp2dXayZbBoOh6NB9FHEeu8DgYBUU1OzOTc39ynh4yFkt3PQT1Wfhv7+0biX5vBZtGjsRqopi0RcTzQfBS2MYkZZ0FDtYx999NGknj17XgEATqezQZznTwCHhg0b9qft27fTYDHR1q8NbRpeBB9wkH1EaGZltKLARkQrpkvjRajvxQ59ohi933JFRUXOq6++enNeXl4RAHTo0OE7scGm+ijq6uredDqdfyEfU2XxHfQOzDNEjkdZJOJlbzGFwWZInAwbNuyNrVu3/jFJzXd54403HpkzZ060TNU2x3/9138VvfHGG7driiIJLHc6nc8nqe02TarlhjTHyMJMiroderKJfA6Ad/Xq1UfGjBlzqlu3bhc11QyhuN1u58iRIwfW19fv37lz51mD6/ZA/0vmR/SpUzr6ieeazUyVitCSg3SE1Hi/b7311k6LFy++zWaz6aZOMzMzdeaUJEkQc3FiYLkkSRtD56Rp8HT0SFPSG4gcT4mAVk1LKAspwhbte6NNNvGZTP5NZasga2aJeHy2wXWqAOyrV6/+dsKECUeKi4sHhOpYyEDsdjTF6/VaFEWxjBs3rp8kSZ9s3bq1JnRtmqmgPZiaWSUhaIoowuYV/l4Z50PERcXT1PttIW1aiWzD+dgK7SW1CbIVgBpSFD9VVdWu5XZoW0ZGhktTEKEtEIuyCAQC0pEjR1bm5ua+HzqfA0AHBKdotS2f/F31OB9CLyGoLERZUxbiMX6EPyeJft7NtJkUku2zMBPjEElORBtGNjbNAaDz/HQkQX91OpBjPrvrrruy/+d//meexWJxZGRkuKhTronUbty48Y1JkybtFz7TXjoND4J2tchpIlMHqBfhyU+RnHJGD2S04sZO6EcSTgSL1zTyq1/9qnLmzJk/sFqtDqvV6s3JydH5XkhSWKzU7ty589EhQ4aI/ZEHYDDZ7xsin4TeJ3EGsQVlaZ+JmAmCi7X6WbNVS0v2yCIVlQVNZJIRPiymQ1SqPDLJMd9u27btW5/P99nIkSMH2mw2yWq1xhThGQW1Q4cOg3v16lXz2muvHQ19RofzfoSX5qOy0cMcy8NlpCzoM2RmNqRxn5UrVw6aMWPG9YqiWIGgYrDb7TpzqgmjtFoA9xUXFx+BPirUgfDKZdQ5TCuVaSMLDTZDEkzaKAsA7m3btp31+XyfjRs3rp/NZkvoqM3j8VjKysoG9O3b97uQwmj1ymLlypWDrrrqqmt1jSVOWdQCuE+SpEOh87GyaCI8G5Jgli1bVr106dLFAA4lo/1LL730+rfffntKMtpuTlavXj2KKooEcgjAnSFFwSSIpvz6xTMKMPpVp3KkRYqpDAPZzDXQ2RBxJGFBeG5CByIXkmM6EPl0cXFx3SuvvDKtsLCwnd1ud5WUlBxHAqmurt7RuXPnNcJHDoT7Vmip+++JbFRBPNKvo+YAFqGjMKMkMHG04fnwww+v6dq162AAsFqt3ry8vAtVD4uHQ5dccsnKd955p0D4rB2AkYLsAnCCHPcfItM4i7PQB2nR2agAwhdXNuOzoHEukXxIRkWBo+X/JKwKPI8skkR1dbVr5syZr544cSIpa1MUFBSM2rdv3+xktJ1M3nvvvR9oiiIJHAJw3zvvvEOD0ZgEwMoiiWgKY//+/V8ko/2ioqIR+/btm91ais2+++6705KY57ENQR8FK4okwcoiyVRXV7suueSSPwPYmIz2i4qKRrz88ssLBg8enLIKo6Kiwr5nz565PXv2HJCkU2yUJGkJK4rkEovPwsysQ7SkIi34SZTpAsO0MA1dgIbW4KA2dDQfhpnZEBrBmWcgizkQ7aC33bOgt93PAfh81apVYwcPHtwDACorK79EAqmtra2eMmXKyq1bt4rRjn3JbjTugkYl0lW2KEY+CxpXkSMKFRUV3nXr1v04Nze3CAhGY2ZlZdF4j6awMVQLROyjAaFNww29/8aFcH/OMSLTYjc0opPWRDWaDTEj04rhkc5hFMtB831oFC+tUm7kBzHlw+CRRTMya9as99977z1aQToh2Gy20ldeeWXBqFGjUibFvUuXLva1a9feoimKJLCeU8ybj1jiLMyMLKLNwdOqSVRujtkQep20TaPRC30BHdD/bRnkGBpd6UbQu46//e1vR1RVPTdlyhQau9EkQmHmBTNnzhxrt9s/fe+992oQPotD4y5o6b5oWadmRo92AJg+fXqHZ5999qa8vLx8MTRbVVWvzWajv47xsFySpLWhf3eCvo+KQpsGTd83SuencRZ0+UKjexVrhKYZOZbZESD67IiZKFFT8MiiBXjiiScOAkjKL6KiKI758+cvWLx4cedktG+G6dOnd1i6dOmNmZmZOdH3jgstIYxpRiL5LGL9RdaSikSo7a8lEInfR6rZ0GijrVixokfv3r07O51Oe0FBQYlWPDcQCOgUHl3NPPSZRL7XyVrylyZ7vV7diMjlcqlU9vl8jce43W6bKHu9XovX6238xZVl2Wez2XS25dChQ/cpitJFkiQnEKyJEXbhMRAIBCTxGnw+X/3+/ftptqbu5lgsFh9dWS1StGQgEIB4jtB5dPfK4/FY2rVr18FqtdqAoI9CrPVhsVh8Tc318Hg8X4t9um/fvuLa2trGkcW5c+ecdXV1jbEysiz7VVV1ibLdbqcFdXT9Eypo1HidiqKErUJH1yYRR1BGRYPF748ePXpAluXArl279q9evfrI1q1b6xDuT6NrqBhlFtMREnXyGsXSUD9VtNGl/qUxIJqyMApuijZ8p8qBrobdqCzmzp3b6eqrr+7dvXv38tzc3B6NJyEdFGvx3Gj7BwKBsH1EZQIAPp9PFvfx+/0RZfoiA0C7du10zsacnJyEe/K///57XXCZ2+3WmQz0gdYyOi/UntG9iXY/c3JyahOcJwOPx6Nb9On48eMFokKnL3JoBTKdrCiKbrhOFZgsy/5I9yLavdLOY/b7urq6I8eOHTv4ySeffDF79ux/hz6mVcWMlAU1n6KZUzRM3SgQz9C8Sqnq3tdee23+LbfcMrx3795DHA5HXqhT21Q5dYYxIiMjo3OXLl1KysvLJ548ebL+m2+++feLL7645YknnqAr17UYkRycZsyQSGYJYLDsHNnHCkC+7rrr8letWjXthhtu+GFRUVF3zcSQJMloHVGqrWMNWTezf8Rfz5AsmZUB0IWSQVPY7XZ7Ipx+Oqj5RE0GWkTGZFEZM8l/jdjtdk+iFX5opNd43tra2gzxb6O/+kYyfa6MnrNI98LMvYrjewmAJMuyNTs7u/OYMWNGz549u3v79u1Pb9q06TSMw73pSMPM1Kkps8Po4oz+bSQb2VM0JoLmCWQSWTdrMHToUPXZZ5+9tqSkZBgQHPrRxWJIJ8ZU9MQI6rOIsI8oU7MkTDmIxxisoh62yBD1UdBMy0TQ0NBAlYVOmUeyqS+EgU8o4kE2m82dDGUhXse5c+ccXq+38bky8hfQ5yjaj5CJtVKjXWZczyoxZxufo1OnTn2xatWqVxYvXkxzjOgPPs2zofEhHuhjN0xnz0ZSFkZTlJECqIxWv6be8OzQcVizZs3oiRMnjlNVtVHBqKrqVVW18Rc29AvAZgiTlojKAggq+3379m2dN2/eG1u2bNGc1zRgsIbItdArB1q0xyi4zDA5TXzZY41PiFbLEjCYDendu7fz3XffvaVPnz7DrVarLGp0rYxa4wlNOJEYpq1iZN5mZmb2mD59+uDMzMwDmzZtqkH4DzSN6KTLMMZTh6P5RxYPPfRQ/1tuueVmzSdBp6d4ZMEw5zEaWdTX1zeOxHft2vWXiy++eBc5LGkji2jKgo4kIsVEKAivA9GoLP76178OGT9+/E1im1RZWCwWn9Vq1Wm9RFfNZpjWAvV9+Xw+iU5/79u375MhQ4b8UfiIvsdUWdD1TrwIzz8xjAqlykAk2urXNujngGnpMiBkT/3jH/+4om/fvqMyMzPrxAAdq9XqZWXAMPHjcrms9fX1Ox944IFnfvvb39YB6Ed2oUV76Ep2bkReLAna8ZF8FtGqWBmZIXQ2xLF58+ZrKysrhwHhy/yFouLAMEx8+Hw+xWq1dpowYUL/mpqaD3bs2JFPdqERnF7olYEZH0aYGZLwkcXmzZt/0qNHj0GazCMLhkksLpfLKvx7b05Ozutkl4SNLGhuR7RNIZuFbFZte/nll0dUVlb2t1gsPm3Tcg+0jUcVDNM0xPfJ4XBUfvDBB9fg/Dq4tQi+t1Zho+8sfaeNFu2SAEji0nL0zaWL8dA49QyELyicBQCLFi3qOnbs2NmZmZk6zyyPJBgmsdC8m4EDB/Y6fPjwgJKSEi19vxv0Ew9Gi2yJ+HCBNPiEp6hfeumlebfffvtNiW6XYRhztG/ffvr777/fO9Htin6JaD4LGt6tDWvE7+3r16+/1el05gHhab/s0GSY5KLlyBQVFQ0+derUOzt37syC3iKgSxhQh6fRkgaJH1m89tprk5JYQo1hGJPIspyxZMmS2xLZpuizoLaLSj5zQO+jyIKQKHbFFVdkjx8/fjQEz2qCyqcxDGMS8Z2z2WxVa9eu3TV9+vS9wi5OAOICTOegf8/pYtnaZ4kbWSxdunRGotpiGCYxjBs3LmHvZUKUxT333NO1sLCwayLaYhgmcdhstoKtW7dOTURbCVEWt956a7IWuGUYpon07NlzwogRI5q8RIRYuSqaz8IO/XytA4B9xYoV/W02WwEAb2Zmpq5ALMMwLYvD4XDZ7XZp3bp1/Tt06PAKgBIAYkg4jatwI3zJiAZAnzVqVL+CBm/QXBB55MiRA7XMODGUm2GYlkerD1NYWDgZwIsIKoeo7zVpRoLBhzFxww03dOzYsWN5U9pgGKZZcAYCgYlNaaBJyuLmm28e3pTjGYZpVpqsLLREEZpgouL8Mnw2BH0UTmFz9+7du6vNZnPbbDa3WOGKYZiU5KJ77rnHCeCksHmgf68zoH/vbQjphEhZpzLZdNlpzz333FC73a5qdTM5QYxhUp+rr756MM5Xy9IWIKIDBfruB5coiPeklZWV3Zp22QzDNDfdunUbGe+xcSuL8vLyhGe1MQyTXJxOZ+ehQ4fGFXNxQTMDpJhNaLMBsN1www1lqqraxcI2dO1IhmFSD0VR/Pfee28JzhfH8UPvn1AR/t4rABQxkYyOMmj1bs3hidGjR/cAEBALb7DPgmFSH1mW/YMHDy4HsCX0UT7CF1+mayDLjf+JlcrKyrJ4jmMYpuVxOp2l8RwXl7IoLCzsGM9xDMO0PA6HozKe42KZOm30azidzlxtJWqh+C6vHMYwKY622PjKlStzcX5pQ1qw94JTpxeq3G3o3Pzxj3/c3efzyZIkBVRV9WgbWUmMYZgURFEUv6qqnltvvdUO4AyCuSJ2sqlka3RwMimIz+eTamtr/3Ps2LHNhw8f/hoAunXrVpmfnz/I6XSWKorCDmWmKRTGekDMymLkyJFxOUcY83i9Xmnv3r3PVlVVvUe+2gtg7TvvvNN74MCB07Kzs3ux0mDiJC5lcaF1Q2iKugRApis7M4nF5/NpiuJTAMXCVw4E4/YxadIkAHhl8eLFXW688cabOnbs6FAUxe9wOMTVsSFJEisSJh6MdAFijbNQfT6fha7uzCSOmpqaz0IjiquhVxY9AVQI8uePPPLI54888sjHixcvvuiaa64pveiii74V2+L6IkwcaEmlIvHHWTDJIRAI4LPPPns71uMeeeSRPZdddtmar7/++kOv18tKnEkKrCxSCLfbLY8cOXJHPMceO3asrnv37s+vXbv28ZMnTx4MBNgCYRILK4sUwuVy7WtqGzfeeONnnTp1emzLli1/qKuro7UUGSZuIvkstGQycV+rz+dTQj6L5ri+dKYfADGz1wXgP4KsAOguyF0BDNWECRMmoFOnThsXLVrUf+rUqUMtFkugvLz8aFKvmGmNqBAWC0NQD7DPIt04evSo66677vpoxowZfzxw4MBRn8/H/gwmblhZpAG7du2qmTJlyqubNm1acu7cue94VMjEAyuLNGLSpEl7s7Ky7jx48ODq+vp6XuOFiYlIQVlaEolO5hiLZuMsgNOCXAdAfME9oS0SdEnJWQDQs2dPFBYW/usPf/hD+bBhwyqsVmsAAFRV9fBi1mkHDb6k7732WewjC20Iywoj6XgRXB1K3BrIVi9sDdAXYnVBX7XZCaBM206cONH+8ssv/9uiRYueOnTo0GGPx6P4/X4eaaYnNOPcEH440pznn3/+aFVV1e/+7//+76+nT58+09LXw6QurCwYAMCcOXN2zpgx4xdffvnlOo/Hw6NGJoxoa51KRGafRfNRC+B7QT4L4Jwga2aHCJ3moFWcs4ncTzxm69attd26dTvXt2/fLffff//o8ePH5+Tl5dVq30uS5Oe6JW0S+p4bvt+xjCwu5AhlkoMX552Y2kZ9EpF8GEYbbS8XwVRlbSsC0Hn37t226667bsfdd9/9+unTp8/4fD4plG3MI9G2h+n3mTufuSB/+tOfvuzYseMDW7ZsebW+vp5Dx9McVhZMVCZPnrzxyiuvfPDLL7/cxv6M9IXL6qUuHgSnSzU0M0ODxl0EEKynKEJ9GvRF9yNYX1X8XtynBKG8gS1btqCysnLPrFmz9t5+++1Xde/evViSJOTk5NQK+0OWZS7c3EZhZZG6aD4KjQYEFYRGHYJOUI0AwoO0aP9SB6gXemWhraKtUQKglyDXrFq1as+qVavemjt3brebb755XFVVlU4ByTIPVtsq3LNMXDz99NNfXHXVVS99+umnGxoaGjh0PA1gZcHETXV1tat///7rlixZ8lB1dfW/Oau1bcPKInUJIOhT0DYf2bxko9Oims9D3GKdeq1D+NRrWAj5L3/5S2d5efnr99xzzyuHDx8+1dDQYHW5XFaXy2X1+/2sQNoI7LNIXaiDU3u5NeqgD9Lyk++NoN/XQv8MOKAP5KI+DBXB3BKRPto/li1bhmXLlv1lxYoVwy677LIxTqfTlp+fX8PJaW0DHlkwCWf+/PkfTZ48+aktW7bs4KnWtgMrCyYpHDp0qGHGjBl/X7Zs2S9Onz69n/0ZrR9WFqlLIMrmR2Sfhhm/RlN9GvUGx+SK24MPPujKz89/8cknn/zL559/HmhoaFC9Xq+ibT6fj5/BVgL7LFIXTQFoaC97JJlOYXqJ7CayC8HCvxo2BP0SGg4EF8rVsELvJwHCl8EbQOQKAN6f/exnWLJkyb8feOCBdtOnT++akZFhA4ILIWVlZdWBSXlYqzPNxsmTJ9133nnnlltvvXX57t27d/FMSeuClQXT7Lz55pvfjRs37pXHH3/898ePHz/GBYRbB6wsmBbjscce+6JHjx6P7d69+zleECn1YZ9F6qI5MkVZTNLSHJga1IcBhCeSUbkB+sQxK/TPBPVhWKAvIgwEi/KInCFyLvT5J30BDBfk+qqqquoOHTpsf/DBBwdedtll3UtLS0+KDfACz6kBjyyYlOD48ePuefPmbZsxY8aLJ06cOMhTrakHKwsmpfj444/PFBcXL3333Xefqqmp4QWRUghWFkxKcvnll/+roKDg7r17925wuVyc1ZoCsM+i9WLk06A/w7QQjdH34nDfC33chRt6f4OM8PwSGstB/SZ10D9nNN9ERtCvoZEJoKMm9O/fHyNGjHhz0aJFY4cMGdJLURTk5+fXqKrK+SbNDCuLtgVVBtGUBxBePSvSuN9otSojp6kIdZrK5JwqggpEww6gndjAhx9++O0111yz6brrrvt83rx547KzsyVVFf2uTHPAZgjTanj55ZePjB49+qW33nprNU+1Nj+sLJhWx/Tp07cuXLjwrq+++ooXRGpG2Axp20QzS4z8HJFkox8X+gwZ1f0UP6O1Ra3Qm0cWhIoEC9D8k8HPPPNM/TPPPPPt0KFD1y1fvnxCnz59CmVZDgCA1Wr1qqpKfSlME2Fl0baI5cUHjB2gEpHFfSSEVxCP9svug95pSov4WEmbVHkA4cqiSDvv9u3bMWLEiHXz58/Pnjt37jV5eXnZTqezgZVF4mEzhGkTrFix4lCvXr1+s3Hjxvdra2up05VJAKwsmDbFDTfc8O7UqVMfOnbs2Fav18v+jATCZkh6Ey080k/2MXr5aI0Mox8g0QyhOSwWcg4xruNCnzlJm6UQYjN27tx5pqio6NMf/vCHgQULFkwuLS215ufnizksAc2/wZiHlUV6YeYFiaQcJIM2zPhFRAVCCxFboPdRGPks6HMqQZ/gVkZkF4CGl156CW+99dYnL7300uCJEyc2/i2SJBldJxMFNkOYNs2pU6c8s2bN+tvp06ePtPS1tHZYWTBtnhMnTtStWbPmBY7JaBqsLBgRs3EZTSkkHGsRYaNCwnSj+wcQNGe0zXnnnXdi+/bt+8+ePZtRV1dni/8WpS/ss2Ao1GdBFQaNs4imYHzQ+z4s0DsnFYT7MOjUp5EfRPRRSAgW6tHIgT45LQtAxxdeeKHkZz/7WeeMjIy6rKysr8HEBI8smLTh8OHDXEW8CbCyYNKGkpKSjOh7MReClQWTNkyePLlzS19Da4Z9FkwkzMQixFpgxwf9j5QMvR9EQbhfhAZlWaD3Wdig91ko0D/brqlTp8oDBgzIkGW5wWaz0UAyxgSsLJg2T35+vrp48eLRsswD6abAd49p0wwcODD7+eefH5WdnW0URs7EAI8smDbLnDlzulx33XXdc3JyZIRPxzIxwsqCSTRG+SRSjDId8RrJ1O/RKF977bXeO+64o2PXrl0LJEk6raqqp3379mfM/wmMEawsmDZDYWGhumTJkqHjxo0rcTgc7MRMMKwsmDbBz3/+8/5TpkypzMrKsiqKwssEJAFWFkyr5o477uh+/fXXX5aXl5cdSj1nkgQrCyYaRj6ISLJiIMtEpuuI2Mn34joiQDC3Q8Q3fPhw29KlS/9f165dO2dkZLgURflW+9JisfisVmtjDU5Jkrh2RQJgZcFEwsxPdTTlQRcVos5JBXoFY0GEiuGlpaX2xx57bNTIkSMHWyyWAICAzWbziMpBURQ/r7yeeFhZMK2GX//610OuuOKK0ZmZmVZwpatmh5UFk/LMnj27bO7cuROLiorac0m8loOVBRPN1IgUEwGE+yhopKQVerNDJfsoCF8o2QYA48ePz3n44Ycvrays7BEK1XYBQZ+E6IdQVdUjFuCVZdloTVemibCySG/MKIpoAVPRlIWNHEOTvizQOzSlLl26OB988MFhkyZNGmOxWAJOp/OcrkGbzaMoCiuEZoaVBZNSLF26tP/UqVPHFxQUZIPNjZSClQWTEsyZM6f8Rz/60fjS0tJSmJuFYZoZVhbpRawxE3TaU0L4M6MSmVajskNvqtjEY7p3725/+umnL+vXr1+VoigBWZa94jQoAFitVt00KC8Q1DKwskhvoiV9GSVs0WeGVsp2EJmuHqaGNjz33HNjx48fP7CgoEBCsNI3LBaLLzMzs970X8A0G6wsmGbn7rvv7nnTTTdNys/Pzw45KjnpqxXAyoJpNiZMmJD385///NqSkpJyrlrV+mBl0bZpqo+C5nUYmSHUZxG2gE+PHj1sv/rVryYOGTJkeChEu9EnIcuyLjSbp0RTF1YWbQszykBEQWTlYIX+GVEQrgyyiZwrCitXrux36aWXjnU6nTYAPrvd7rbb7Y1mhyRJHETVSmBlwSSFW265pey2226bVFpams+jhbYBKwsmoYwcOTL3oYceuqSysrK7LMuQJMkb/SimNcDKgkkIFRUVjkcffXTU8OHDJ4f8Ekwbg5VF6yJaDATN06D9S30WKiL7KHQBVKHvMkkbgaVLl/a/8sorx+Xn52fn5eWd1l2AxeLjIKq2ASsLJm5++tOflv/4xz8eW1ZWVsol7do+rCyYmBkzZkzuvffeO66qqqqfoig8akgTWFkwMfHHP/5x3OjRo4eGpkJZUaQRrCxSl4gL6SC88K2E8BgI2r+01kQG9H4OO8JrTViAYIj27NmzJ5SUlKiSJPkB1AMA53GkD6wsmIhcffXVHRYsWDC5S5cuJVar1SdJUkNLXxPTMrCyYAzp1q2b/bHHHhNDtJk0h5UFE8bvfve7EZMnTx7ncDgcfr+fFQUDgJVFKkPjJqzQxzyo0C/OY7ROBp3PpLUmHOI55s2bV3LHHXdcVVBQkC1JEqxWa72qqo1LAUqSFODQ7fSFlUXqEq1YLk360j4TMQrCorKihWj37du3q81m8yI0y2GxWHw2m43XDWUAsLJIa7p06WJ/4IEHRmhVtGVZ5lW8mAvCyiJNeeKJJwZOnTr1CqfTaQfHSzAmYGWRuligj4tQoY+B8EAoIgPAj+hrePjnz5/f5aabbrq4tLS0zGq1emVZboyTUBTFL/okaOFcJr1hZZG6KAhXFqLPwQ69UzMA4wV+AADDhw/PefTRRwdfdNFF/bUQbYfD4aarjXPSF3MhWFmkAb///e/HjB8/flBubq4MNjmYOGFl0YZZsGBBxY9+9KOJ+fn52aGsUJ7ZYOImFmURIP9nEowkSeICPQrC19sQv6dTqT4AtUAwRHvhwoWTevToUaQoChAqtW+3213i+egCw5xmnpbQ9/mC73ckZREwakiSJIgPGJM47HZ76Zw5c5zPPPNMLcKDsmwIOjE1aCKZu7S0VPrlL385etSoUYNCCwqfFdt3Op2c18EYYUph8OINKYTVag3cdtttg+M59qGHHhrw97///bZx48YN5FwOJhmwzyLF6Nq16zQA75ndf9q0acULFy4cX1hYmCGW2GeYRBOzz+Lo0aM1SboWBoDT6cyvrq6+rbi4+BsAZ4SvdGZhnz59Mh599NHRVVVV5bIsBxRF8Yp5HAAv2MNE5MQFPo/os7iQ45L6LAIA/F9//TUriyQiSRKKiopGbd68+dh99923cfPmzWdCX7UH0H7MmDEFs2fP7n3xxRcXZ2dn1wM4BwSdlTk5ObUtdd1Mq0NTFgHofWGGvkogDjNk//79rCySjCRJ6N27d8WaNWv6HD169Gu/34+6uroMVVULCwoKnLIsw2KxsLOSaQrHYz0gZmWxa9eus9H3YhKBxWIJlJaWlgJAQ0ODzeVy0axRhokLSZIuZIZckLhmQ/7zn/9Ux3McwzApwZ54DhJ9FtQZ5oc+98CnyUeOHDlRVVWV6fP5GpUNL3CbWPLz89ncYxKO3++X6uvrD+N84SQZ+oREL8ILKfm1HWPm4MGD38ZzHMMwLU9NTc2n8RwXl7J466232AxhmFbK66+/3nzK4l//+lfNiRMneHTBMK0Mv9//1Zw5c+KaYo/ks2j0UYTwQsha3L9//9GSkpJejQ1ZLP6MjAyezmOYFOazzz7bAaBI+CgD4T4LWvQozGcRMLH5tW3t2rU7fD6fImycZ8IwKc769es/RtC5qW0WCO81Irz/Es6nPeeSdguhL+NWACBfkL/55ptv5thstgIgGEGYlZVVl7g/i2GYBHNIkqSfA2gnfNYJ+pFGPYDD5LgjQBOzTnfv3v3PphzPMEyzsr4pB2tl1nQmRmjzkM0NwCVsGbfffvvh+vp6e319vd3lctH6jwzDpA61ALYhaEF0EbYs6N9rF8LffR8AXzSfBVUg4qYcPHgwsHv37j0+n0/2+/3ss2CY1GW9JEm1CFZccwob9VkYbQEAgSa/4L/4xS/eb2obDMMknY1NbaDJymLz5s1n9u7d+0lT22EYJmmsiSdxjBLJZ+Elm+a30DYrgDwAeQ888MAnZ8+e9Xu9XkXcmnpxDMM0Dbfb7Vq4cOEmANmhLZNsCvTvtRvh774fgF+sIK0FaGmbDcEK0pryUBFUEFqwVhaCU6n2o0ePShUVFcqAAQNyA4GApG0Wi4UTyximBdm7d+/rs2bNOoTg+2wDUBzatEWrGgCcxfn32hWSxYFDLUJxFloshZOcJw/6Fa5yAeQIcjsEYy80jtfW1k5VFKW99gGvwM0wLcoeSZJ+A328VGVo0zgFQEzdaEB4YZyTQIKre7/99ttPJLI9hmHiphbAs4lsMNLUKfVh+MhG7Rp52rRptdu3b3+3oaFBbWho4KpODNNyrJEk6RCC1kCBsNkQ/u7Sd9tw6lSbYwXCC164he+AYBioWIbPBr2ZYgEwauzYsd9u3LixvrS0tJgLyDJMi7BRkiQtWrMPgiHdGm4AYmGlcwitZCd8r1u5DqHEsqQEUs2aNeuNmpoarvTEMM3PIQDPJaPhpCiL6upq1/333/836DUWwzDJpRbAfaFIzYRDp05FrAhOnWo+DGto02Rt6kWzaxCSJQDS559/Xl9cXLy+e/fuYwKBgN3r9VoURQnwOqkMkxQ0RQHow7m7Qr+gtmZmiNOi9Tj/XnsA1EHvv2wAgopCUwhiTIUfwVGH6NjURiFacokma8knQHB61aJtGzZssH/66acnhg4dOqi+vj7b4XC4rFYr9Y0wDNM0NEVxCMAV0CeK5eL8ey4h6KM4g/Pv7dnQpr3XrlB7osPThUTkhkRjw4YNJxYtWvSy2+2mThOGYZrOIZxXFEmlWTJFN2zYcGL27Nl/rKmpOdoc52OYNKHZFAXQTMoCAD766KPv58+f/wgSkP3GMAzWS5J0Z7KcmUZIOK8wqINThV6Z2KEPG9UcKBo2BEvxiZQTuQ6A9/bbb+962223XVJWVnbW4XA0mieyLAd4oSKGMSYQCEgej+fUl19++dvKykqtnP9Msls2kY8h5KAMUQN9nEUtQotrh/AifBazAWjGkYXIb3/72y8vvfTSPxw9evTDljg/w7RGvF7vhmXLlt0tKIpmJeaFkRNFdXW1q1evXs/v27dvfVlZ2Q+sVmuflroWhklx9gBYpqrqKehH981KiykLjZCW/HTfvn19unTpMtZms41r6WtimBRhI4K+Cc2B2aI1YrS5V6MLsUBvplgR9GNoOHB+cVWEvssjbXQmcib0+SQN0OefnJoyZYoyc+bMfiNGjOjXvn37dllZWfWNFytJfo7TYNoqPp9Prq+vP3L8+PH3ly9fvunJJ5/MANBe2CUDQAk5zEHk00T+FvpcD5oLUh/aGi8Deh8HEFpcTFQW1H9hEb4DzkdwamiLlIjf55I2OhG5HfTDqEzoRzfHEXTIAAAmT55sefzxx6UOHTpUOhyOUrvdXqqqKtfIYNoShwB8CWDP66+//sXVV18tOhs7Q/+DmwmgOzn+DJGPEPkUglGbGrXQK4sG6JWDH+GJZKmvLBD02orOHE8gELAB6BCSuyC8aE80ro9x/6awphnPxZwnVfu4FkHlAADHDepiZkAfmp1SyqLFfRaxErrB2k3eHevxgUBgIsKneJOCJEmsLFqAQCDQXMriRDr1sbgwMk3w8kE/spAM9hH9DUbOF7rwkBd6v0c22UeBXpMGAPQk10DPYyeyDfpRkk085h//+EfxoEGDGu08q9XqTfSyi36/X/L5fN8K11YOYDbZjZ6TanMf9Pc7gPD731aRiGyU5Cj2sQqhj3fs2FFQUVHR+GtpsVh8TqczoYt2+/1+yePxnMKF+9jI9qd9TmvG+IgsxkEBwf6no5FTRP6WyGcgLGiO8JGEVqRXPAf1CwYAfSfQB9FIMYidqFXLEqH+BPoCNJALsUF/czKgd9hYoDczFIQ7dKgZ4oBeoWSI8qZNmyz9+vVL9ohKcrvdXwvnzQbQl+xzlsj1RNaqKmsYKYu2qDxE01iD/ujQHwRdn2/fvt3TtWvXZM8cSLW1tftw4T42Cm6ifU7fB1p4xgX9i+2BPqAKCH9ujEwI8b2kq6RTBXXB5yztVhFbt27dzuY4z5kzZ3Y0x3mYcD744IMvm+M8Bw4cSKs+TjtlsXfv3tMnT55M6sPk9/tPdO7ceVMyz8FcmNWrV+9uaGj4Lpnn8Hq9n44YMeKrZJ4j1Yg0HI9mhlATxI9wfwIdhsnQDykDRK6H3sxQoLfzjMwQagdSM0Rbz7Gxjfvvv3/H8uXLewOA3+9PeOr8Rx999DaAUuGjPAD04f2eyPTv8CDy8LAtmiCAsRlCiz9TM8QJfZ+ry5cv3zFv3rxrAMBms7kTXQ/2zTfffAvhfSz6D7zQ51wAxiaE2Mcu6P0HVPaYaJP+nfUIN3ViNUMAhHeKCB11yGR/hewjIzwUNZPItFPpi0yT1ahyUBDu0KQ+C52PAsHFkMRzfA/gzG9+85vBl1122SCHw9FQVlb2DRLHHkmSngfQTfisAMAwst8ZItOHwEhZ0CS7tqIwqCM9mrKwQ//s0Sn47wGc+fOf/zx+0KBBFUno4/WSJP0T+j7Oh76PvQj3UdAfCKMX2R1BNvKD0DapTP0i2uqCGj6EOzTpc+YH0tAM0ViwYMHOjz/+eH+Cmz0E4BcJbpOJk7vvvvuD48eP09mCpnJIkqSkFMRNddJWWQDAjTfeuCmBCkMrRMJFilOEb775xn399ddv+OqrrxJVdGkbgPsS1FarI5IZQoeDVJahH+7LCB8uZhDZAf1wkZoM2nqMGtTsUBBu6sRqhtBY+OPbtm1z9urV60pVVe2yLPtVVdX5Y6IVGT5+/Pg7jz/++F9+/etfa4qia2gTr6EnOewMkaOZIdpiLxptxQQBwp9D+iNm5LMQ+5iaIWF9/O9//7tdRUXFVCBYN8VqtcbUx4cPH95QWlr6F+Ej2seZAHoIcjxmiLb4uIYb4dOe1LdFz0F9GFpxXrGNWH0WfiC6sogmU2ViVECHyrSgDk1WE9ugfhAzCimasqAhtQ0AviktLbU99dRTE/v27VvWuXNnXYdYLJYLJa9tQ3DlJ3voPBrlAMoE2YfwTqYJP+nssxChgUiAsc8ikrKgfVwH4KsxY8bk3nfffWP79+9f2a5duzNigxH6eCOCfVwE/Vq/ZdD3sQf6HwAvwl9cKkfzJ1DZKBSbPlc0EIzG62hFuDXMxO8EAFYWQEhZaEKPHj0aNm/e7M3MzOyjKEp7RVHKrVarNrrZg2D+yiEA24TY/kqwskgUSVMWmtCrVy/3pk2bZIM+1nI3jiPY19sEs7IKaa4sWl1uSLI5ePBgbceOHT8A8E/hYxolx7RiPvvss9qOHTvuBPdxTMSiLLS1RUTZaB8R+ktI8x08CI/NEL+n07VG+Sn0V4iOcCxEpiMiD/neAaCYtOlGZHKh/yWTof8F8CL8YaQy/UVwQ/+3Up8FDOS2gNHIwqguq7iPgsh97IZ+dOKEuT4W728O9P4xCfo+dBPZi/A+NZIj+ROo7Ed4fJOZac9Iz5HpZyqSGRKNaGYJEK4IFLIPTYOnsRv0xTcydajDkw5RaSwHzT+xQm9C0HwUIHwYTG94HfSdRh1sRjH9Z4hMZ1E4N+Q8ZoKyqMKnfSwWsrUg3Hyl56A/bLH2sVFQFpWNEsnEcxiZDNHysYyCJaMlJJp6jtJ66pRhGPOwsmAYxhSsLBiGMUVTfBZm2osmGzknIwWCGTm/aK0Do+nZSIFfRjUzaP4JPQe1A2ngi7hYNGAc00/tV6MpLy5+E8TMAljR+lj0UcTbxz4iR+pjH6I7sY38UtGmy6PJZvwRcT1HPLJgGMYUrCwYhjEFKwuGYUyRaJ+FmfZjCTE3+iyaPUsXR6L2LV3SQIbefjWqy0HPQe1ZD5FpApBR8VYq0zBeI1s0XX0WRgtgiftE62OagBhPH9MgOdrHfkQvqU8Dv8wUZY41JiJpsTjJDvemF2oUgRlPGyI0gk1C5ArJtE0aJSobnNMoylTchzqqjGL6owXPGFVUZgenuX3M9DGIHGsf01ydaH1sNtoy1j5u6vdxw2YIwzCmYGXBMIwpku2ziIdYr8lMrVAqi0NOCXr71ij/JNo10ZgIGuMfT0x/OuSBmKUl+pjeb+pfMOpjD5HNmJqRZEqLPgOpmKLeVJ+Gkd0X6UEwcqBS52I0ZUEfJKMEoGjZgUkLpmkDtEQfG2VQx9LHZuqPtKo+ZjOEYRhTsLJgGMYUqWiGNBW6GBKVtc80jKZz6TSamSEqlRNdEi9lh6ctQDL62Mw5qRwtr6NN9XEqOjhjJdrfYJSsRuVoi9xEUxZmbOpY7deUfnCamebo42ikfR+zGcIwjClYWTAMY4pY7bZUJNbQ4GjQwsTaZ+JmdIzR/kxiaI4+NnMMlSP1cZvr/7bgs4hGPH9jIu5LSgfYtDG4j5sBNkMYhjEFKwuGYUzByoJhGIZhGIZhmGbm/wMD91GhpxHALQAAAABJRU5ErkJggg==");

/***/ }),

/***/ "./node_modules/url/url.js":
/*!*********************************!*\
  !*** ./node_modules/url/url.js ***!
  \*********************************/
/***/ ((__unused_webpack_module, exports, __webpack_require__) => {

"use strict";
/*
 * Copyright Joyent, Inc. and other Node contributors.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a
 * copy of this software and associated documentation files (the
 * "Software"), to deal in the Software without restriction, including
 * without limitation the rights to use, copy, modify, merge, publish,
 * distribute, sublicense, and/or sell copies of the Software, and to permit
 * persons to whom the Software is furnished to do so, subject to the
 * following conditions:
 *
 * The above copyright notice and this permission notice shall be included
 * in all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS
 * OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 * MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN
 * NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
 * DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
 * OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE
 * USE OR OTHER DEALINGS IN THE SOFTWARE.
 */



var punycode = __webpack_require__(/*! punycode */ "./node_modules/punycode/punycode.js");

function Url() {
  this.protocol = null;
  this.slashes = null;
  this.auth = null;
  this.host = null;
  this.port = null;
  this.hostname = null;
  this.hash = null;
  this.search = null;
  this.query = null;
  this.pathname = null;
  this.path = null;
  this.href = null;
}

// Reference: RFC 3986, RFC 1808, RFC 2396

/*
 * define these here so at least they only have to be
 * compiled once on the first module load.
 */
var protocolPattern = /^([a-z0-9.+-]+:)/i,
  portPattern = /:[0-9]*$/,

  // Special case for a simple path URL
  simplePathPattern = /^(\/\/?(?!\/)[^?\s]*)(\?[^\s]*)?$/,

  /*
   * RFC 2396: characters reserved for delimiting URLs.
   * We actually just auto-escape these.
   */
  delims = [
    '<', '>', '"', '`', ' ', '\r', '\n', '\t'
  ],

  // RFC 2396: characters not allowed for various reasons.
  unwise = [
    '{', '}', '|', '\\', '^', '`'
  ].concat(delims),

  // Allowed by RFCs, but cause of XSS attacks.  Always escape these.
  autoEscape = ['\''].concat(unwise),
  /*
   * Characters that are never ever allowed in a hostname.
   * Note that any invalid chars are also handled, but these
   * are the ones that are *expected* to be seen, so we fast-path
   * them.
   */
  nonHostChars = [
    '%', '/', '?', ';', '#'
  ].concat(autoEscape),
  hostEndingChars = [
    '/', '?', '#'
  ],
  hostnameMaxLen = 255,
  hostnamePartPattern = /^[+a-z0-9A-Z_-]{0,63}$/,
  hostnamePartStart = /^([+a-z0-9A-Z_-]{0,63})(.*)$/,
  // protocols that can allow "unsafe" and "unwise" chars.
  unsafeProtocol = {
    javascript: true,
    'javascript:': true
  },
  // protocols that never have a hostname.
  hostlessProtocol = {
    javascript: true,
    'javascript:': true
  },
  // protocols that always contain a // bit.
  slashedProtocol = {
    http: true,
    https: true,
    ftp: true,
    gopher: true,
    file: true,
    'http:': true,
    'https:': true,
    'ftp:': true,
    'gopher:': true,
    'file:': true
  },
  querystring = __webpack_require__(/*! qs */ "./node_modules/qs/lib/index.js");

function urlParse(url, parseQueryString, slashesDenoteHost) {
  if (url && typeof url === 'object' && url instanceof Url) { return url; }

  var u = new Url();
  u.parse(url, parseQueryString, slashesDenoteHost);
  return u;
}

Url.prototype.parse = function (url, parseQueryString, slashesDenoteHost) {
  if (typeof url !== 'string') {
    throw new TypeError("Parameter 'url' must be a string, not " + typeof url);
  }

  /*
   * Copy chrome, IE, opera backslash-handling behavior.
   * Back slashes before the query string get converted to forward slashes
   * See: https://code.google.com/p/chromium/issues/detail?id=25916
   */
  var queryIndex = url.indexOf('?'),
    splitter = queryIndex !== -1 && queryIndex < url.indexOf('#') ? '?' : '#',
    uSplit = url.split(splitter),
    slashRegex = /\\/g;
  uSplit[0] = uSplit[0].replace(slashRegex, '/');
  url = uSplit.join(splitter);

  var rest = url;

  /*
   * trim before proceeding.
   * This is to support parse stuff like "  http://foo.com  \n"
   */
  rest = rest.trim();

  if (!slashesDenoteHost && url.split('#').length === 1) {
    // Try fast path regexp
    var simplePath = simplePathPattern.exec(rest);
    if (simplePath) {
      this.path = rest;
      this.href = rest;
      this.pathname = simplePath[1];
      if (simplePath[2]) {
        this.search = simplePath[2];
        if (parseQueryString) {
          this.query = querystring.parse(this.search.substr(1));
        } else {
          this.query = this.search.substr(1);
        }
      } else if (parseQueryString) {
        this.search = '';
        this.query = {};
      }
      return this;
    }
  }

  var proto = protocolPattern.exec(rest);
  if (proto) {
    proto = proto[0];
    var lowerProto = proto.toLowerCase();
    this.protocol = lowerProto;
    rest = rest.substr(proto.length);
  }

  /*
   * figure out if it's got a host
   * user@server is *always* interpreted as a hostname, and url
   * resolution will treat //foo/bar as host=foo,path=bar because that's
   * how the browser resolves relative URLs.
   */
  if (slashesDenoteHost || proto || rest.match(/^\/\/[^@/]+@[^@/]+/)) {
    var slashes = rest.substr(0, 2) === '//';
    if (slashes && !(proto && hostlessProtocol[proto])) {
      rest = rest.substr(2);
      this.slashes = true;
    }
  }

  if (!hostlessProtocol[proto] && (slashes || (proto && !slashedProtocol[proto]))) {

    /*
     * there's a hostname.
     * the first instance of /, ?, ;, or # ends the host.
     *
     * If there is an @ in the hostname, then non-host chars *are* allowed
     * to the left of the last @ sign, unless some host-ending character
     * comes *before* the @-sign.
     * URLs are obnoxious.
     *
     * ex:
     * http://a@b@c/ => user:a@b host:c
     * http://a@b?@c => user:a host:c path:/?@c
     */

    /*
     * v0.12 TODO(isaacs): This is not quite how Chrome does things.
     * Review our test case against browsers more comprehensively.
     */

    // find the first instance of any hostEndingChars
    var hostEnd = -1;
    for (var i = 0; i < hostEndingChars.length; i++) {
      var hec = rest.indexOf(hostEndingChars[i]);
      if (hec !== -1 && (hostEnd === -1 || hec < hostEnd)) { hostEnd = hec; }
    }

    /*
     * at this point, either we have an explicit point where the
     * auth portion cannot go past, or the last @ char is the decider.
     */
    var auth, atSign;
    if (hostEnd === -1) {
      // atSign can be anywhere.
      atSign = rest.lastIndexOf('@');
    } else {
      /*
       * atSign must be in auth portion.
       * http://a@b/c@d => host:b auth:a path:/c@d
       */
      atSign = rest.lastIndexOf('@', hostEnd);
    }

    /*
     * Now we have a portion which is definitely the auth.
     * Pull that off.
     */
    if (atSign !== -1) {
      auth = rest.slice(0, atSign);
      rest = rest.slice(atSign + 1);
      this.auth = decodeURIComponent(auth);
    }

    // the host is the remaining to the left of the first non-host char
    hostEnd = -1;
    for (var i = 0; i < nonHostChars.length; i++) {
      var hec = rest.indexOf(nonHostChars[i]);
      if (hec !== -1 && (hostEnd === -1 || hec < hostEnd)) { hostEnd = hec; }
    }
    // if we still have not hit it, then the entire thing is a host.
    if (hostEnd === -1) { hostEnd = rest.length; }

    this.host = rest.slice(0, hostEnd);
    rest = rest.slice(hostEnd);

    // pull out port.
    this.parseHost();

    /*
     * we've indicated that there is a hostname,
     * so even if it's empty, it has to be present.
     */
    this.hostname = this.hostname || '';

    /*
     * if hostname begins with [ and ends with ]
     * assume that it's an IPv6 address.
     */
    var ipv6Hostname = this.hostname[0] === '[' && this.hostname[this.hostname.length - 1] === ']';

    // validate a little.
    if (!ipv6Hostname) {
      var hostparts = this.hostname.split(/\./);
      for (var i = 0, l = hostparts.length; i < l; i++) {
        var part = hostparts[i];
        if (!part) { continue; }
        if (!part.match(hostnamePartPattern)) {
          var newpart = '';
          for (var j = 0, k = part.length; j < k; j++) {
            if (part.charCodeAt(j) > 127) {
              /*
               * we replace non-ASCII char with a temporary placeholder
               * we need this to make sure size of hostname is not
               * broken by replacing non-ASCII by nothing
               */
              newpart += 'x';
            } else {
              newpart += part[j];
            }
          }
          // we test again with ASCII char only
          if (!newpart.match(hostnamePartPattern)) {
            var validParts = hostparts.slice(0, i);
            var notHost = hostparts.slice(i + 1);
            var bit = part.match(hostnamePartStart);
            if (bit) {
              validParts.push(bit[1]);
              notHost.unshift(bit[2]);
            }
            if (notHost.length) {
              rest = '/' + notHost.join('.') + rest;
            }
            this.hostname = validParts.join('.');
            break;
          }
        }
      }
    }

    if (this.hostname.length > hostnameMaxLen) {
      this.hostname = '';
    } else {
      // hostnames are always lower case.
      this.hostname = this.hostname.toLowerCase();
    }

    if (!ipv6Hostname) {
      /*
       * IDNA Support: Returns a punycoded representation of "domain".
       * It only converts parts of the domain name that
       * have non-ASCII characters, i.e. it doesn't matter if
       * you call it with a domain that already is ASCII-only.
       */
      this.hostname = punycode.toASCII(this.hostname);
    }

    var p = this.port ? ':' + this.port : '';
    var h = this.hostname || '';
    this.host = h + p;
    this.href += this.host;

    /*
     * strip [ and ] from the hostname
     * the host field still retains them, though
     */
    if (ipv6Hostname) {
      this.hostname = this.hostname.substr(1, this.hostname.length - 2);
      if (rest[0] !== '/') {
        rest = '/' + rest;
      }
    }
  }

  /*
   * now rest is set to the post-host stuff.
   * chop off any delim chars.
   */
  if (!unsafeProtocol[lowerProto]) {

    /*
     * First, make 100% sure that any "autoEscape" chars get
     * escaped, even if encodeURIComponent doesn't think they
     * need to be.
     */
    for (var i = 0, l = autoEscape.length; i < l; i++) {
      var ae = autoEscape[i];
      if (rest.indexOf(ae) === -1) { continue; }
      var esc = encodeURIComponent(ae);
      if (esc === ae) {
        esc = escape(ae);
      }
      rest = rest.split(ae).join(esc);
    }
  }

  // chop off from the tail first.
  var hash = rest.indexOf('#');
  if (hash !== -1) {
    // got a fragment string.
    this.hash = rest.substr(hash);
    rest = rest.slice(0, hash);
  }
  var qm = rest.indexOf('?');
  if (qm !== -1) {
    this.search = rest.substr(qm);
    this.query = rest.substr(qm + 1);
    if (parseQueryString) {
      this.query = querystring.parse(this.query);
    }
    rest = rest.slice(0, qm);
  } else if (parseQueryString) {
    // no query string, but parseQueryString still requested
    this.search = '';
    this.query = {};
  }
  if (rest) { this.pathname = rest; }
  if (slashedProtocol[lowerProto] && this.hostname && !this.pathname) {
    this.pathname = '/';
  }

  // to support http.request
  if (this.pathname || this.search) {
    var p = this.pathname || '';
    var s = this.search || '';
    this.path = p + s;
  }

  // finally, reconstruct the href based on what has been validated.
  this.href = this.format();
  return this;
};

// format a parsed object into a url string
function urlFormat(obj) {
  /*
   * ensure it's an object, and not a string url.
   * If it's an obj, this is a no-op.
   * this way, you can call url_format() on strings
   * to clean up potentially wonky urls.
   */
  if (typeof obj === 'string') { obj = urlParse(obj); }
  if (!(obj instanceof Url)) { return Url.prototype.format.call(obj); }
  return obj.format();
}

Url.prototype.format = function () {
  var auth = this.auth || '';
  if (auth) {
    auth = encodeURIComponent(auth);
    auth = auth.replace(/%3A/i, ':');
    auth += '@';
  }

  var protocol = this.protocol || '',
    pathname = this.pathname || '',
    hash = this.hash || '',
    host = false,
    query = '';

  if (this.host) {
    host = auth + this.host;
  } else if (this.hostname) {
    host = auth + (this.hostname.indexOf(':') === -1 ? this.hostname : '[' + this.hostname + ']');
    if (this.port) {
      host += ':' + this.port;
    }
  }

  if (this.query && typeof this.query === 'object' && Object.keys(this.query).length) {
    query = querystring.stringify(this.query, {
      arrayFormat: 'repeat',
      addQueryPrefix: false
    });
  }

  var search = this.search || (query && ('?' + query)) || '';

  if (protocol && protocol.substr(-1) !== ':') { protocol += ':'; }

  /*
   * only the slashedProtocols get the //.  Not mailto:, xmpp:, etc.
   * unless they had them to begin with.
   */
  if (this.slashes || (!protocol || slashedProtocol[protocol]) && host !== false) {
    host = '//' + (host || '');
    if (pathname && pathname.charAt(0) !== '/') { pathname = '/' + pathname; }
  } else if (!host) {
    host = '';
  }

  if (hash && hash.charAt(0) !== '#') { hash = '#' + hash; }
  if (search && search.charAt(0) !== '?') { search = '?' + search; }

  pathname = pathname.replace(/[?#]/g, function (match) {
    return encodeURIComponent(match);
  });
  search = search.replace('#', '%23');

  return protocol + host + pathname + search + hash;
};

function urlResolve(source, relative) {
  return urlParse(source, false, true).resolve(relative);
}

Url.prototype.resolve = function (relative) {
  return this.resolveObject(urlParse(relative, false, true)).format();
};

function urlResolveObject(source, relative) {
  if (!source) { return relative; }
  return urlParse(source, false, true).resolveObject(relative);
}

Url.prototype.resolveObject = function (relative) {
  if (typeof relative === 'string') {
    var rel = new Url();
    rel.parse(relative, false, true);
    relative = rel;
  }

  var result = new Url();
  var tkeys = Object.keys(this);
  for (var tk = 0; tk < tkeys.length; tk++) {
    var tkey = tkeys[tk];
    result[tkey] = this[tkey];
  }

  /*
   * hash is always overridden, no matter what.
   * even href="" will remove it.
   */
  result.hash = relative.hash;

  // if the relative url is empty, then there's nothing left to do here.
  if (relative.href === '') {
    result.href = result.format();
    return result;
  }

  // hrefs like //foo/bar always cut to the protocol.
  if (relative.slashes && !relative.protocol) {
    // take everything except the protocol from relative
    var rkeys = Object.keys(relative);
    for (var rk = 0; rk < rkeys.length; rk++) {
      var rkey = rkeys[rk];
      if (rkey !== 'protocol') { result[rkey] = relative[rkey]; }
    }

    // urlParse appends trailing / to urls like http://www.example.com
    if (slashedProtocol[result.protocol] && result.hostname && !result.pathname) {
      result.pathname = '/';
      result.path = result.pathname;
    }

    result.href = result.format();
    return result;
  }

  if (relative.protocol && relative.protocol !== result.protocol) {
    /*
     * if it's a known url protocol, then changing
     * the protocol does weird things
     * first, if it's not file:, then we MUST have a host,
     * and if there was a path
     * to begin with, then we MUST have a path.
     * if it is file:, then the host is dropped,
     * because that's known to be hostless.
     * anything else is assumed to be absolute.
     */
    if (!slashedProtocol[relative.protocol]) {
      var keys = Object.keys(relative);
      for (var v = 0; v < keys.length; v++) {
        var k = keys[v];
        result[k] = relative[k];
      }
      result.href = result.format();
      return result;
    }

    result.protocol = relative.protocol;
    if (!relative.host && !hostlessProtocol[relative.protocol]) {
      var relPath = (relative.pathname || '').split('/');
      while (relPath.length && !(relative.host = relPath.shift())) { }
      if (!relative.host) { relative.host = ''; }
      if (!relative.hostname) { relative.hostname = ''; }
      if (relPath[0] !== '') { relPath.unshift(''); }
      if (relPath.length < 2) { relPath.unshift(''); }
      result.pathname = relPath.join('/');
    } else {
      result.pathname = relative.pathname;
    }
    result.search = relative.search;
    result.query = relative.query;
    result.host = relative.host || '';
    result.auth = relative.auth;
    result.hostname = relative.hostname || relative.host;
    result.port = relative.port;
    // to support http.request
    if (result.pathname || result.search) {
      var p = result.pathname || '';
      var s = result.search || '';
      result.path = p + s;
    }
    result.slashes = result.slashes || relative.slashes;
    result.href = result.format();
    return result;
  }

  var isSourceAbs = result.pathname && result.pathname.charAt(0) === '/',
    isRelAbs = relative.host || relative.pathname && relative.pathname.charAt(0) === '/',
    mustEndAbs = isRelAbs || isSourceAbs || (result.host && relative.pathname),
    removeAllDots = mustEndAbs,
    srcPath = result.pathname && result.pathname.split('/') || [],
    relPath = relative.pathname && relative.pathname.split('/') || [],
    psychotic = result.protocol && !slashedProtocol[result.protocol];

  /*
   * if the url is a non-slashed url, then relative
   * links like ../.. should be able
   * to crawl up to the hostname, as well.  This is strange.
   * result.protocol has already been set by now.
   * Later on, put the first path part into the host field.
   */
  if (psychotic) {
    result.hostname = '';
    result.port = null;
    if (result.host) {
      if (srcPath[0] === '') { srcPath[0] = result.host; } else { srcPath.unshift(result.host); }
    }
    result.host = '';
    if (relative.protocol) {
      relative.hostname = null;
      relative.port = null;
      if (relative.host) {
        if (relPath[0] === '') { relPath[0] = relative.host; } else { relPath.unshift(relative.host); }
      }
      relative.host = null;
    }
    mustEndAbs = mustEndAbs && (relPath[0] === '' || srcPath[0] === '');
  }

  if (isRelAbs) {
    // it's absolute.
    result.host = relative.host || relative.host === '' ? relative.host : result.host;
    result.hostname = relative.hostname || relative.hostname === '' ? relative.hostname : result.hostname;
    result.search = relative.search;
    result.query = relative.query;
    srcPath = relPath;
    // fall through to the dot-handling below.
  } else if (relPath.length) {
    /*
     * it's relative
     * throw away the existing file, and take the new path instead.
     */
    if (!srcPath) { srcPath = []; }
    srcPath.pop();
    srcPath = srcPath.concat(relPath);
    result.search = relative.search;
    result.query = relative.query;
  } else if (relative.search != null) {
    /*
     * just pull out the search.
     * like href='?foo'.
     * Put this after the other two cases because it simplifies the booleans
     */
    if (psychotic) {
      result.host = srcPath.shift();
      result.hostname = result.host;
      /*
       * occationaly the auth can get stuck only in host
       * this especially happens in cases like
       * url.resolveObject('mailto:local1@domain1', 'local2@domain2')
       */
      var authInHost = result.host && result.host.indexOf('@') > 0 ? result.host.split('@') : false;
      if (authInHost) {
        result.auth = authInHost.shift();
        result.hostname = authInHost.shift();
        result.host = result.hostname;
      }
    }
    result.search = relative.search;
    result.query = relative.query;
    // to support http.request
    if (result.pathname !== null || result.search !== null) {
      result.path = (result.pathname ? result.pathname : '') + (result.search ? result.search : '');
    }
    result.href = result.format();
    return result;
  }

  if (!srcPath.length) {
    /*
     * no path at all.  easy.
     * we've already handled the other stuff above.
     */
    result.pathname = null;
    // to support http.request
    if (result.search) {
      result.path = '/' + result.search;
    } else {
      result.path = null;
    }
    result.href = result.format();
    return result;
  }

  /*
   * if a url ENDs in . or .., then it must get a trailing slash.
   * however, if it ends in anything else non-slashy,
   * then it must NOT get a trailing slash.
   */
  var last = srcPath.slice(-1)[0];
  var hasTrailingSlash = (result.host || relative.host || srcPath.length > 1) && (last === '.' || last === '..') || last === '';

  /*
   * strip single dots, resolve double dots to parent dir
   * if the path tries to go above the root, `up` ends up > 0
   */
  var up = 0;
  for (var i = srcPath.length; i >= 0; i--) {
    last = srcPath[i];
    if (last === '.') {
      srcPath.splice(i, 1);
    } else if (last === '..') {
      srcPath.splice(i, 1);
      up++;
    } else if (up) {
      srcPath.splice(i, 1);
      up--;
    }
  }

  // if the path is allowed to go above the root, restore leading ..s
  if (!mustEndAbs && !removeAllDots) {
    for (; up--; up) {
      srcPath.unshift('..');
    }
  }

  if (mustEndAbs && srcPath[0] !== '' && (!srcPath[0] || srcPath[0].charAt(0) !== '/')) {
    srcPath.unshift('');
  }

  if (hasTrailingSlash && (srcPath.join('/').substr(-1) !== '/')) {
    srcPath.push('');
  }

  var isAbsolute = srcPath[0] === '' || (srcPath[0] && srcPath[0].charAt(0) === '/');

  // put the host back
  if (psychotic) {
    result.hostname = isAbsolute ? '' : srcPath.length ? srcPath.shift() : '';
    result.host = result.hostname;
    /*
     * occationaly the auth can get stuck only in host
     * this especially happens in cases like
     * url.resolveObject('mailto:local1@domain1', 'local2@domain2')
     */
    var authInHost = result.host && result.host.indexOf('@') > 0 ? result.host.split('@') : false;
    if (authInHost) {
      result.auth = authInHost.shift();
      result.hostname = authInHost.shift();
      result.host = result.hostname;
    }
  }

  mustEndAbs = mustEndAbs || (result.host && srcPath.length);

  if (mustEndAbs && !isAbsolute) {
    srcPath.unshift('');
  }

  if (srcPath.length > 0) {
    result.pathname = srcPath.join('/');
  } else {
    result.pathname = null;
    result.path = null;
  }

  // to support request.http
  if (result.pathname !== null || result.search !== null) {
    result.path = (result.pathname ? result.pathname : '') + (result.search ? result.search : '');
  }
  result.auth = relative.auth || result.auth;
  result.slashes = result.slashes || relative.slashes;
  result.href = result.format();
  return result;
};

Url.prototype.parseHost = function () {
  var host = this.host;
  var port = portPattern.exec(host);
  if (port) {
    port = port[0];
    if (port !== ':') {
      this.port = port.substr(1);
    }
    host = host.substr(0, host.length - port.length);
  }
  if (host) { this.hostname = host; }
};

exports.parse = urlParse;
exports.resolve = urlResolve;
exports.resolveObject = urlResolveObject;
exports.format = urlFormat;

exports.Url = Url;


/***/ }),

/***/ "./node_modules/util-deprecate/browser.js":
/*!************************************************!*\
  !*** ./node_modules/util-deprecate/browser.js ***!
  \************************************************/
/***/ ((module, __unused_webpack_exports, __webpack_require__) => {


/**
 * Module exports.
 */

module.exports = deprecate;

/**
 * Mark that a method should not be used.
 * Returns a modified function which warns once by default.
 *
 * If `localStorage.noDeprecation = true` is set, then it is a no-op.
 *
 * If `localStorage.throwDeprecation = true` is set, then deprecated functions
 * will throw an Error when invoked.
 *
 * If `localStorage.traceDeprecation = true` is set, then deprecated functions
 * will invoke `console.trace()` instead of `console.error()`.
 *
 * @param {Function} fn - the function to deprecate
 * @param {String} msg - the string to print to the console when `fn` is invoked
 * @returns {Function} a new "deprecated" version of `fn`
 * @api public
 */

function deprecate (fn, msg) {
  if (config('noDeprecation')) {
    return fn;
  }

  var warned = false;
  function deprecated() {
    if (!warned) {
      if (config('throwDeprecation')) {
        throw new Error(msg);
      } else if (config('traceDeprecation')) {
        console.trace(msg);
      } else {
        console.warn(msg);
      }
      warned = true;
    }
    return fn.apply(this, arguments);
  }

  return deprecated;
}

/**
 * Checks `localStorage` for boolean values for the given `name`.
 *
 * @param {String} name
 * @returns {Boolean}
 * @api private
 */

function config (name) {
  // accessing global.localStorage can trigger a DOMException in sandboxed iframes
  try {
    if (!__webpack_require__.g.localStorage) return false;
  } catch (_) {
    return false;
  }
  var val = __webpack_require__.g.localStorage[name];
  if (null == val) return false;
  return String(val).toLowerCase() === 'true';
}


/***/ }),

/***/ "./node_modules/xtend/immutable.js":
/*!*****************************************!*\
  !*** ./node_modules/xtend/immutable.js ***!
  \*****************************************/
/***/ ((module) => {

module.exports = extend

var hasOwnProperty = Object.prototype.hasOwnProperty;

function extend() {
    var target = {}

    for (var i = 0; i < arguments.length; i++) {
        var source = arguments[i]

        for (var key in source) {
            if (hasOwnProperty.call(source, key)) {
                target[key] = source[key]
            }
        }
    }

    return target
}


/***/ }),

/***/ "?4f7e":
/*!********************************!*\
  !*** ./util.inspect (ignored) ***!
  \********************************/
/***/ (() => {

/* (ignored) */

/***/ }),

/***/ "?ed1b":
/*!**********************!*\
  !*** util (ignored) ***!
  \**********************/
/***/ (() => {

/* (ignored) */

/***/ }),

/***/ "?d17e":
/*!**********************!*\
  !*** util (ignored) ***!
  \**********************/
/***/ (() => {

/* (ignored) */

/***/ })

/******/ 	});
/************************************************************************/
/******/ 	// The module cache
/******/ 	var __webpack_module_cache__ = {};
/******/ 	
/******/ 	// The require function
/******/ 	function __webpack_require__(moduleId) {
/******/ 		// Check if module is in cache
/******/ 		var cachedModule = __webpack_module_cache__[moduleId];
/******/ 		if (cachedModule !== undefined) {
/******/ 			return cachedModule.exports;
/******/ 		}
/******/ 		// Create a new module (and put it into the cache)
/******/ 		var module = __webpack_module_cache__[moduleId] = {
/******/ 			id: moduleId,
/******/ 			loaded: false,
/******/ 			exports: {}
/******/ 		};
/******/ 	
/******/ 		// Execute the module function
/******/ 		__webpack_modules__[moduleId].call(module.exports, module, module.exports, __webpack_require__);
/******/ 	
/******/ 		// Flag the module as loaded
/******/ 		module.loaded = true;
/******/ 	
/******/ 		// Return the exports of the module
/******/ 		return module.exports;
/******/ 	}
/******/ 	
/************************************************************************/
/******/ 	/* webpack/runtime/compat get default export */
/******/ 	(() => {
/******/ 		// getDefaultExport function for compatibility with non-harmony modules
/******/ 		__webpack_require__.n = (module) => {
/******/ 			var getter = module && module.__esModule ?
/******/ 				() => (module['default']) :
/******/ 				() => (module);
/******/ 			__webpack_require__.d(getter, { a: getter });
/******/ 			return getter;
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/define property getters */
/******/ 	(() => {
/******/ 		// define getter functions for harmony exports
/******/ 		__webpack_require__.d = (exports, definition) => {
/******/ 			for(var key in definition) {
/******/ 				if(__webpack_require__.o(definition, key) && !__webpack_require__.o(exports, key)) {
/******/ 					Object.defineProperty(exports, key, { enumerable: true, get: definition[key] });
/******/ 				}
/******/ 			}
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/global */
/******/ 	(() => {
/******/ 		__webpack_require__.g = (function() {
/******/ 			if (typeof globalThis === 'object') return globalThis;
/******/ 			try {
/******/ 				return this || new Function('return this')();
/******/ 			} catch (e) {
/******/ 				if (typeof window === 'object') return window;
/******/ 			}
/******/ 		})();
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/hasOwnProperty shorthand */
/******/ 	(() => {
/******/ 		__webpack_require__.o = (obj, prop) => (Object.prototype.hasOwnProperty.call(obj, prop))
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/make namespace object */
/******/ 	(() => {
/******/ 		// define __esModule on exports
/******/ 		__webpack_require__.r = (exports) => {
/******/ 			if(typeof Symbol !== 'undefined' && Symbol.toStringTag) {
/******/ 				Object.defineProperty(exports, Symbol.toStringTag, { value: 'Module' });
/******/ 			}
/******/ 			Object.defineProperty(exports, '__esModule', { value: true });
/******/ 		};
/******/ 	})();
/******/ 	
/******/ 	/* webpack/runtime/node module decorator */
/******/ 	(() => {
/******/ 		__webpack_require__.nmd = (module) => {
/******/ 			module.paths = [];
/******/ 			if (!module.children) module.children = [];
/******/ 			return module;
/******/ 		};
/******/ 	})();
/******/ 	
/************************************************************************/
var __webpack_exports__ = {};
// This entry need to be wrapped in an IIFE because it need to be in strict mode.
(() => {
"use strict";
/*!****************!*\
  !*** ./app.js ***!
  \****************/
__webpack_require__.r(__webpack_exports__);
/* harmony import */ var artplayer__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! artplayer */ "./node_modules/artplayer/dist/artplayer.js");
/* harmony import */ var artplayer__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(artplayer__WEBPACK_IMPORTED_MODULE_0__);
/* harmony import */ var mpegts_js__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! mpegts.js */ "./node_modules/mpegts.js/dist/mpegts.js");
/* harmony import */ var mpegts_js__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(mpegts_js__WEBPACK_IMPORTED_MODULE_1__);
/* harmony import */ var artplayer_plugin_danmuku__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! artplayer-plugin-danmuku */ "./node_modules/artplayer-plugin-danmuku/dist/artplayer-plugin-danmuku.js");
/* harmony import */ var artplayer_plugin_danmuku__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(artplayer_plugin_danmuku__WEBPACK_IMPORTED_MODULE_2__);
/* harmony import */ var _img_ploading_gif__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./img/ploading.gif */ "./img/ploading.gif");
/* harmony import */ var _img_state_png__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! ./img/state.png */ "./img/state.png");
/* harmony import */ var _img_indicator_svg__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! ./img/indicator.svg */ "./img/indicator.svg");
/* harmony import */ var _img_filp_svg__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! ./img/filp.svg */ "./img/filp.svg");
/* harmony import */ var stream_http__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! stream-http */ "./node_modules/stream-http/index.js");
/* harmony import */ var stream_http__WEBPACK_IMPORTED_MODULE_7___default = /*#__PURE__*/__webpack_require__.n(stream_http__WEBPACK_IMPORTED_MODULE_7__);










(() => {
    class FIFO {
        #indexedDB;
        #ok=false;
        #db;
        #size=0;
        #cu=1;
        #dbN="FIFO"+new Date().getTime();
        #objN="fifo"+new Date().getTime();

        constructor(okf = (_)=>{}) {
            const that = this;
            this.#indexedDB = window.indexedDB;
            if (!this.#indexedDB) {
                console.error("IndexedDB could not be found in this browser.");
            }

            this.close().catch();

            const request = this.#indexedDB.open(this.#dbN, 1);

            request.onerror = function (event) {
                console.error("An error occurred with IndexedDB");
                console.error(event);
            };
            
            request.onupgradeneeded = function () {
                that.#db = request.result;
                that.#db.createObjectStore(that.#objN, { keyPath: "id", autoIncrement: true });
            };
            
            request.onsuccess = function () {
                console.log("Database opened successfully");
                that.#db = request.result;
                that.#ok = true;
                if(okf)okf(that);
            };
        }

        #getTx(mode,func) {
            if(!this.#ok)return;
            const transaction = this.#db.transaction(this.#objN, mode);
            transaction.onerror = (event) => {
                console.error("An error occurred with put");
                console.error(event);
            };
            transaction.oncomplete = function () {};
            return func(transaction, transaction.objectStore(this.#objN));
        }

        #stillTx(transaction,func) {
            return func(transaction, transaction.objectStore(this.#objN));
        }

        size(){
            return new Promise((resolve) => resolve(this.#size));
        }

        showSize(){
            return this.#getTx("readonly",  (transaction, store)=>{
                const idQuery = store.count();
                idQuery.onsuccess = function () {
                    console.log(this.#size);
                };
            });
        }

        put(data){
            const that = this;
            return this.#getTx("readwrite",  (transaction, store)=>{
                return new Promise((resolve) => {
                    store.put({ data: data });
                    that.#size += 1;
                    resolve(that.#size);
                });
            });
        }

        get(){
            const that = this;
            return this.#getTx("readwrite", (transaction, store)=>{
                return new Promise((resolve, reject) => {
                    const idQuery = store.get(that.#cu);
                    idQuery.onsuccess = async function () {
                        if(idQuery.result){
                            that.#size -= 1;
                            that.#cu += 1;
                            await that.#stillTx(transaction,  (transaction, store)=>{
                                return new Promise((resolve) => {
                                    transaction.oncomplete = function () {
                                        resolve();
                                    };
                                    store.delete(idQuery.result.id)
                                });
                            });
                            resolve({size: that.#size, data: idQuery.result.data});
                        } else reject();
                    };
                });
            });
        }

        /**
         * @returns .then(e=>{}).catch(e=>{});
         */
        close(){
            if(this.#ok)this.#db.close();
            return new Promise((resolve, reject) => {
                const DBDeleteRequest = this.#indexedDB.deleteDatabase(this.#dbN);
                DBDeleteRequest.onerror = (event) => {
                    reject("Error deleting database.");
                };

                DBDeleteRequest.onsuccess = (event) => {
                    if(event.result===undefined)resolve("Database deleted successfully.");
                    else reject("Error deleting fail.");
                };
            });
        }

        deleteOnExit() {
            let that = this;
            window.addEventListener('beforeunload', function (e) {
                that.close().catch(()=>{});
            });
        }

        static test() {
            new FIFO(async fifo=>{
                fifo.put(1).then(size=>size!=1?console.error("size:1 ",size):console.log("1ok"));
                fifo.put(2).then(size=>size!=2?console.error("size:2 ",size):console.log("2ok"));
                fifo.put(3).then(size=>size!=3?console.error("size:3 ",size):console.log("3ok"));
                fifo.put(4).then(size=>size!=4?console.error("size:4 ",size):console.log("4ok"));
                fifo.size().then(size=>size!=4?console.error("size:4 ",size):console.log("5ok"));
                console.log('1!')
                await fifo.get().then(result=>result.id!=1?console.error(result):console.log("6ok")).catch(()=>{});
                console.log('2!')
                await fifo.get().then(result=>result.id!=2?console.error(result):console.log("7ok")).catch(()=>{});
                console.log('3!')
                fifo.close().then(r=>console.log(r)).catch(result=>console.error(result));
                console.log("fin");
            });
        }
    }

    class EventPromise {
        #eventEL = document.createElement("_");
        
        eventCall(name, data = undefined, el = this.#eventEL){
            let e = new Event(name, {bubbles: true, cancelable: false})
            e.detail = data;
            el.dispatchEvent(e);
        }

        promise(name, bootFunc = ({event: event})=>{}){
            return EventPromise.toPromise(this, name, bootFunc);
        }

        /**
         * cover event listener to promise
         * @param {*} object 
         * @param {*} event name 
         * @param {*} bootFunc {event: event} => {}
         * @returns .then(({event: event, data: data}) => {}).catch(({event: event, error: error}) => {})
         */
        static toPromise(object, name, bootFunc = ({event: event})=>{}){
            return new Promise((resolve, reject) => {
                let event = object.addEventListener(name, data =>{
                    object.removeEventListener(name, event);
                    resolve({object:object, name:name, event: event, data: data});
                });
                try {
                    bootFunc({event: event});
                } catch (error) {
                    object.removeEventListener(name, event);
                    reject({object:object, name:name, event: event, error: error});
                }
            });
        }

        addEventListener(name, func, el = this.#eventEL){
            let eventFunc = e=>func(e.detail);
            el.addEventListener(name, eventFunc);
            return eventFunc;
        }

        removeEventListener(name, eventFunc, el = this.#eventEL){
            el.addEventListener(name, eventFunc);
        }

        constructor(name){
            this.#eventEL = document.createElement(name);
        }

        static test(){
            let ep = new EventPromise();
            ep.addEventListener("test", data=>{
                if (data=="ss")console.log("event ok");
                else console.error(data);
            });
            ep.promise("test").then(data=>{
                if (data=="ss")console.log("promise ok");
                else console.error(data);
            });
            ep.eventCall('test','ss');
        }
    }

    class MSC extends EventPromise {
        #fetchDone = false;
        #forceExit = false;
        #exit = () => this.#forceExit || this.#bufLen <= 1 && this.#fifoL == 0 && this.#fetchDone;
        #fifo;

        #id = new Date().getTime();
        #url = "";
        #loadedRange = 0;
        #video;
        #fifoL = 0;
        #bufLen = 0;
        #sourceBuffer;
        #mediaSource;

        #mp4LoadFromDB = 20;
        #mp4StopFromDB = 30;
        #mp4LoadFromWeb = 1000;
        #mp4StopFromWeb = 2000;

        #loopIfFalse(f, miliSec = 1000, rejectFail = false){
            return new Promise((reslove, reject)=>{
                if(f())return reslove();
                let l = () => setTimeout(()=>{
                    if(f())return reslove();
                    else if(rejectFail)return reject();
                    else return l();
                },miliSec);
                l();
            });
        }

        #fetchLoop = () => {
            let that = this;
            var reqHeaders = new Headers();
            reqHeaders.append("Range", "bytes="+that.#loadedRange+"-");

            fetch(new Request(that.#url,{
                method: "GET",
                headers: reqHeaders,
                mode: "cors",
                cache: "default",
            }))
            .then((response) => {
                const reader = response.body.getReader();
                reader.read().then(function pump({ done, value }) {
                    if(done)return that.eventCall("fetch.done", "ok");
                    if(that.#exit())return;
                    
                    that.#loadedRange += value.length;
                    that.#fifo.put(value).then(tfifoL=>{that.#fifoL = tfifoL;});

                    if(that.#fifoL>that.#mp4StopFromWeb){
                        reader.cancel();
                        return that.#loopIfFalse(()=>that.#exit() || that.#fifoL<that.#mp4LoadFromWeb).then(()=>that.#fetchLoop());
                    }
                    return reader.read().then(pump);
                });
            })
            .catch(({event: event, error: error}) => that.eventCall("error", {altmsg: error}));
        }

        #sourceBufferLoop = () => {
            let that = this;
            let deal = () => {

                if(that.#mediaSource.sourceBuffers.length != 0 && that.#sourceBuffer.buffered.length != 0)
                    that.#bufLen = that.#sourceBuffer.buffered.end(that.#sourceBuffer.buffered.length-1) - that.#video.currentTime;
                else that.#bufLen = 0;

                if(that.#exit()){
                    try {
                        that.eventCall("mediaSource.sourceended");
                        that.#mediaSource.endOfStream();
                    } catch {}
                    return;
                }

                if(that.#bufLen<that.#mp4StopFromDB){
                    return that.#fifo.get()
                    .then(({size: size, data: data})=>{
                        that.#fifoL = size;
                        that.#sourceBuffer.appendBuffer(data);
                    })
                    .catch(()=>setTimeout(deal, 1000));
                } else {
                    return that.#loopIfFalse(()=>{
                        if(that.#mediaSource.sourceBuffers.length != 0 && that.#sourceBuffer.buffered.length != 0)
                            that.#bufLen = that.#sourceBuffer.buffered.end(that.#sourceBuffer.buffered.length-1) - that.#video.currentTime;
                        else that.#bufLen = 0;
                        return that.#exit() || that.#bufLen<that.#mp4LoadFromDB;
                    }).then(deal);
                }
            };

            that.#sourceBuffer.addEventListener("updateend", deal);
            
            deal();
        }

        #stateLoop(){
            let that = this;
            setTimeout(()=>{
                if(that.#exit())return;
                console.log("[%s] fifo: %d buf: %d", that.#id, that.#fifoL, that.#bufLen);
                that.#stateLoop();
            }, 2000);
        }

        #watchExit(){
            let exitf = (o) => {
                this.#forceExit = true;
                this.removeEventListener("mediaSource.sourceended", exitf);
                this.removeEventListener("beforeunload", exitf, window);
                this.removeEventListener("mediaSource.error", exitf);
                this.removeEventListener("error", exitf, this.#video);
                this.removeEventListener("error", exitf, this.#sourceBuffer);
                if(o.event && o.event.name && o.event.name.indexOf("error") != -1)console.error(o);
                else console.log(o);
                if(o.event && o.event.altmsg)alert(o.altmsg);
            }
            this.promise("mediaSource.sourceended").then(exitf).catch(()=>{});
            this.promise("mediaSource.error").then(exitf).catch(()=>{});
            EventPromise.toPromise(window, "beforeunload").then(exitf).catch(()=>{});
            EventPromise.toPromise(this.#video, "error").then(exitf).catch(()=>{});
            EventPromise.toPromise(this.#sourceBuffer, "error").then(exitf).catch(()=>{});
        }

        constructor({
            video: video, 
            url: url, 
            mimeType: mimeType = 'video/mp4; codecs="avc1.640032,mp4a.40.2"', 
            mode: mode = "sequence",
            mp4LoadFromDB = 20,
            mp4StopFromDB = 30,
            mp4LoadFromWeb = 1000,
            mp4StopFromWeb = 2000
        }){
            super();

            let that = this;
            that.#url = url;
            that.#video = video;
            that.#mp4LoadFromDB = mp4LoadFromDB;
            that.#mp4StopFromDB = mp4StopFromDB;
            that.#mp4LoadFromWeb = mp4LoadFromWeb;
            that.#mp4StopFromWeb = mp4StopFromWeb;

            if (!MediaSource.isTypeSupported(mimeType)) {
                that.eventCall("mediaSource.error", {altmsg: mimeType+" not Supported"});
                return;
            }

            this.#mediaSource = new MediaSource();
            this.#mediaSource.addEventListener('sourceopen', () => {

                that.eventCall("mediaSource.sourceopen");

                that.#sourceBuffer = that.#mediaSource.addSourceBuffer(mimeType);
                that.#sourceBuffer.mode = mode;

                if(that.#mediaSource.sourceBuffers.length == 0){
                    that.eventCall("mediaSource.error", {altmsg: "addSourceBuffer error"});
                    return;
                }

                this.promise("fetch.done").then(()=>{
                    that.#fetchDone = true;
                    console.log("[%s] fetch.done", that.#id);
                });

                that.#watchExit();

                that.#stateLoop();

                that.#sourceBufferLoop();

                that.#fetchLoop();
            });

            new FIFO(fifo => {
                console.log(that);
                fifo.deleteOnExit();
                that.#fifo = fifo;
                that.#video.src = URL.createObjectURL(that.#mediaSource);
            });
        }
    }

    console.log("init 31");
        let player,
        flvPlayer,
        danmuEmit = document.createElement("div"),
        config = {
            conn: undefined,
            container: '.artplayer-app',
            url: "../stream?_=" + new Date().getTime()+"&ref="+new URL(window.location.href).searchParams.get("ref"),
            title: "" + new Date().getTime(),
            type: new URL(window.location.href).searchParams.get("format")||"flv",
            volume: 0.5,
            hotkey: true,
            isLive: true,
            muted: false,
            autoplay: true,
            autoMini: true,
            screenshot: true,
            setting: true,
            loop: false,
            flip: true,
            playbackRate: true,
            aspectRatio: true,
            fullscreen: true,
            fullscreenWeb: true,
            subtitleOffset: true,
            miniProgressBar: true,
            mutex: true,
            backdrop: true,
            playsInline: true,
            autoPlayback: true,
            theme: '#23ade5',
            lang: navigator.language.toLowerCase(),
            whitelist: ['*'],
            moreVideoAttr: {
                crossOrigin: 'anonymous',
            },
            settings: [],
            contextmenu: [],
            layers: [],
            quality: [],
            thumbnails: {},
            subtitle: {},
            highlight: [],
            controls: [
                {
                    name: '翻转',
                    index: 10,
                    position: 'right',
                    html: '<img width="22" heigth="22" src="'+ _img_filp_svg__WEBPACK_IMPORTED_MODULE_6__["default"] +'">',
                    click: function (...args) {
                        let f = function(...e){
                            // if(e)alert(e);
                            rotate(document.querySelector('.art-video'));
                            rotate(document.querySelector('.art-danmuku'));
                        }, rotate = function(element) {
                            if(element.style.transform == 'rotateZ(0deg)' || element.style.transform == ''){
                                element.style.transform = 'rotateZ(180deg)';
                            }
                            else {
                                element.style.transform = 'rotateZ(0deg)';
                            }
                        };

                        switch (screen.orientation.type) {
                            case "landscape-primary":
                                screen.orientation.lock("landscape-secondary").catch(e=>{f(e);});
                                break;
                            case "landscape-secondary":
                                screen.orientation.lock("landscape-primary").catch(e=>{f(e);});
                                break;
                            case "portrait-secondary":
                                screen.orientation.lock("portrait-primary").catch(e=>{f(e);});
                                break;
                            case "portrait-primary":
                                screen.orientation.lock("portrait-secondary").catch(e=>{f(e);});
                                break;
                            default:
                                f();
                        }
                    },
                }
            ],
            plugins: [
                artplayer_plugin_danmuku__WEBPACK_IMPORTED_MODULE_2___default()({
                    danmuku: [],
                    speed: 7,
                    opacity: 0.7,
                    mount: danmuEmit,
                }),
            ],
            icons: {
                loading: '<img src=' + _img_ploading_gif__WEBPACK_IMPORTED_MODULE_3__["default"] + '>',
                state: '<img width="150" heigth="150" src=' + _img_state_png__WEBPACK_IMPORTED_MODULE_4__["default"] + '>',
                indicator: '<img width="16" heigth="16" src=' + _img_indicator_svg__WEBPACK_IMPORTED_MODULE_5__["default"] + '>',
            },
            customType: {
                mp4: (video, url) => {
                    if(url.indexOf("now")!=-1)new MSC({video: video, url: url});
                    else video.src = url;
                },
                flv: function (video, url) {
                    var needUnload = true;
                    if(flvPlayer){
                        needUnload = false;
                        flvPlayer.destroy();
                    }
                    if (mpegts_js__WEBPACK_IMPORTED_MODULE_1___default().getFeatureList().mseLivePlayback) {
                        flvPlayer = mpegts_js__WEBPACK_IMPORTED_MODULE_1___default().createPlayer({
                            type: 'flv',  // could also be mpegts, m2ts, flv
                            isLive: true,
                            url: url
                        });
                        flvPlayer.attachMediaElement(video);
                        flvPlayer.load();
                        flvPlayer.on("error", function(){
                            flvPlayer.destroy();
                            var c = config;
                            c.type="mp4";
                            initPlay(c);
                        })
                        if(needUnload){
                            setTimeout(function(){
                                if(flvPlayer.paused)flvPlayer.unload();
                            },1000);
                        }
                    }
                },
            },
        };
    
    /**
     * ws 收发
     */
     function ws(player) {
        if (window["WebSocket"]) {
            var conn = new WebSocket("ws://" + window.location.host + window.location.pathname+"ws?ref="+new URL(window.location.href).searchParams.get("ref"));
            let interval_handle = undefined;
            player.on('video:error', conn.close);
            conn.onclose = function (evt) {
                clearInterval(interval_handle)
            };
            conn.onmessage = function (evt) {
                try {
                    let data = JSON.parse(evt.data)
                    player.plugins.artplayerPluginDanmuku.emit({
                        text: data.text,
                        color: data.style.color,
                        border: data.style.border,
                        mode: data.style.mode,
                    });
                    if(!interval_handle)interval_handle = setInterval(()=>{
                        if(conn && player && player.currentTime)conn.send(player.currentTime);
                    },3000);
                } catch (e) {
                    console.log(e)
                    console.log(evt.data)
                }
            };
            conn.onerror = () => {
                clearInterval(interval_handle)
            };
            conn.onopen = function () {
                conn.send(`pause`);
                config.conn = conn;
            };
        }
    }

    function initPlay(config) {
        if(player != undefined && player.destroy != undefined)player.destroy();
        player = new (artplayer__WEBPACK_IMPORTED_MODULE_0___default())(config);
        ws(player);
        player.on('ready', () => {
            player.autoHeight();
        });
        player.on('resize', () => {
            player.autoHeight();
        });
        player.on('video:play', (...args) => {
            if(config.conn != undefined)config.conn.send(`play`);
        });
        player.on('pause', (...args) => {
            if(config.conn != undefined)config.conn.send(`pause`);
        });
        player.on('video:error', (...args) => {
            console.log("clear danmu");
            player.plugins.artplayerPluginDanmuku.config({
                danmuku: [],
                speed: 7,
                opacity: 0.7,
                mount: danmuEmit,
            });
            player.plugins.artplayerPluginDanmuku.load();
            if(config.conn != undefined){
                config.conn.close();
                config.conn = undefined;
            }
            ws(player);
        });
        player.on('video:ended', (...args) => {
            if(config.conn != undefined){
                config.conn.close();
                config.conn = undefined;
            }
            if(flvPlayer)flvPlayer.unload();
        });
        player.on('artplayerPluginDanmuku:visible', (danmu) => {
            danmu.$ref.innerHTML = danmu.$ref.innerText.replace(/(\[.*\])/,"<img src=\"emots/$1.png\" alt=\"$1\" style=\"height: 1.125em;\"></img>");
        });
        player.on('artplayerPluginDanmuku:emit', (danmu) => {
            if(config.conn != undefined)config.conn.send("%S"+danmu.text);
        });
        document.addEventListener("resize", player.autoSize);
        // window.addEventListener('beforeunload', function (e) {
        //     tabUnload = true;
        // });
    }

    stream_http__WEBPACK_IMPORTED_MODULE_7___default().get('../keepAlive', function (res) {
        res.on('data', function (buf) {
            config.url += "&key="+buf;
            initPlay(config);
            let i = setInterval(function () {
                stream_http__WEBPACK_IMPORTED_MODULE_7___default().get('../keepAlive?key='+buf, function (res) {
                    if (res.statusCode>=300)clearInterval(i);
                })
            },15000);
        });
    })
})();
})();

/******/ })()
;
//# sourceMappingURL=bundle.js.map