(window.webpackJsonp=window.webpackJsonp||[]).push([[11],{147:function(e,t,n){"use strict";n.d(t,"a",(function(){return b})),n.d(t,"b",(function(){return f}));var r=n(0),a=n.n(r);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function i(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function l(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?i(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):i(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function c(e,t){if(null==e)return{};var n,r,a=function(e,t){if(null==e)return{};var n,r,a={},o=Object.keys(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||(a[n]=e[n]);return a}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(r=0;r<o.length;r++)n=o[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(a[n]=e[n])}return a}var p=a.a.createContext({}),s=function(e){var t=a.a.useContext(p),n=t;return e&&(n="function"==typeof e?e(t):l({},t,{},e)),n},b=function(e){var t=s(e.components);return a.a.createElement(p.Provider,{value:t},e.children)},u="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return a.a.createElement(a.a.Fragment,{},t)}},m=Object(r.forwardRef)((function(e,t){var n=e.components,r=e.mdxType,o=e.originalType,i=e.parentName,p=c(e,["components","mdxType","originalType","parentName"]),b=s(n),u=r,m=b["".concat(i,".").concat(u)]||b[u]||d[u]||o;return n?a.a.createElement(m,l({ref:t},p,{components:n})):a.a.createElement(m,l({ref:t},p))}));function f(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var o=n.length,i=new Array(o);i[0]=m;var l={};for(var c in t)hasOwnProperty.call(t,c)&&(l[c]=t[c]);l.originalType=e,l[u]="string"==typeof e?e:r,i[1]=l;for(var p=2;p<o;p++)i[p]=n[p];return a.a.createElement.apply(null,i)}return a.a.createElement.apply(null,n)}m.displayName="MDXCreateElement"},98:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return i})),n.d(t,"metadata",(function(){return l})),n.d(t,"rightToc",(function(){return c})),n.d(t,"default",(function(){return b}));var r=n(1),a=n(6),o=(n(0),n(147)),i={title:"What & Why",sidebar_label:"What & Why"},l={id:"what-why",title:"What & Why",description:"## What are ops",source:"@site/docs/what-why.md",permalink:"/docs/what-why",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/what-why.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1578700982,sidebar_label:"What & Why",sidebar:"docs",next:{title:"Bare Metal",permalink:"/docs/setup/bare-metal"}},c=[{value:"What are ops",id:"what-are-ops",children:[]},{value:"What problems do ops solve?",id:"what-problems-do-ops-solve",children:[]},{value:"Implementation Goals",id:"implementation-goals",children:[]}],p={rightToc:c},s="wrapper";function b(e){var t=e.components,n=Object(a.a)(e,["components"]);return Object(o.b)(s,Object(r.a)({},p,n,{components:t,mdxType:"MDXLayout"}),Object(o.b)("h2",{id:"what-are-ops"},"What are ops"),Object(o.b)("ul",null,Object(o.b)("li",{parentName:"ul"},"we think of ops as operations as code"),Object(o.b)("li",{parentName:"ul"},"an op accepts inputs, produces outputs, and may have side effects"),Object(o.b)("li",{parentName:"ul"},"ops are designed to be:",Object(o.b)("ol",{parentName:"li"},Object(o.b)("li",{parentName:"ol"},"Composable: ops can be composed of smaller ops that are defined to run in serial or parallel"),Object(o.b)("li",{parentName:"ol"},"Portable: an op's definition contains everything it needs to run and what inputs it expects, and ops leverage docker containers to run anywhere"),Object(o.b)("li",{parentName:"ol"},"Distributable: ops can be referenced remotely, and can be remotely invoked"),Object(o.b)("li",{parentName:"ol"},"Versionable: an op is defined in a simple ",Object(o.b)("inlineCode",{parentName:"li"},"yaml")," file which makes versioning easy using standard source control")))),Object(o.b)("h2",{id:"what-problems-do-ops-solve"},"What problems do ops solve?"),Object(o.b)("ul",null,Object(o.b)("li",{parentName:"ul"},"automating manual technical operations"),Object(o.b)("li",{parentName:"ul"},"reliable and easy local development for software services"),Object(o.b)("li",{parentName:"ul"},"portable pipelines that live and change with the code they build and deploy"),Object(o.b)("li",{parentName:"ul"},"turning tacit operational knowledge into executable documentation"),Object(o.b)("li",{parentName:"ul"},"providing microservice development teams with an easy to understand, standard interface for operations")),Object(o.b)("h2",{id:"implementation-goals"},"Implementation Goals"),Object(o.b)("ul",null,Object(o.b)("li",{parentName:"ul"},"decentralized"),Object(o.b)("li",{parentName:"ul"},"vendor & platform agnostic"),Object(o.b)("li",{parentName:"ul"},"single executable")))}b.isMDXComponent=!0}}]);