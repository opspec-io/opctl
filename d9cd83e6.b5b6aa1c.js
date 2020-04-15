(window.webpackJsonp=window.webpackJsonp||[]).push([[44],{138:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return o})),r.d(t,"metadata",(function(){return i})),r.d(t,"rightToc",(function(){return p})),r.d(t,"default",(function(){return u}));var n=r(1),a=r(6),c=(r(0),r(147)),o={title:"String Parameter [object]"},i={id:"opspec/reference/structure/op-directory/op/parameter/string",title:"String Parameter [object]",description:"An object defining a parameter which accepts a [string typed value](../../../../types/string.md).",source:"@site/docs/opspec/reference/structure/op-directory/op/parameter/string.md",permalink:"/docs/opspec/reference/structure/op-directory/op/parameter/string",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/opspec/reference/structure/op-directory/op/parameter/string.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1585210706,sidebar:"docs",previous:{title:"Socket Parameter [object]",permalink:"/docs/opspec/reference/structure/op-directory/op/parameter/socket"},next:{title:"Identifier [string]",permalink:"/docs/opspec/reference/structure/op-directory/op/identifier"}},p=[{value:"Properties:",id:"properties",children:[{value:"constraints",id:"constraints",children:[]}]}],s={rightToc:p},l="wrapper";function u(e){var t=e.components,r=Object(a.a)(e,["components"]);return Object(c.b)(l,Object(n.a)({},s,r,{components:t,mdxType:"MDXLayout"}),Object(c.b)("p",null,"An object defining a parameter which accepts a ",Object(c.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/opspec/reference/types/string"}),"string typed value"),"."),Object(c.b)("h2",{id:"properties"},"Properties:"),Object(c.b)("ul",null,Object(c.b)("li",{parentName:"ul"},"must have:",Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#description"}),"description")))),Object(c.b)("li",{parentName:"ul"},"may have:",Object(c.b)("ul",{parentName:"li"},Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#constraints"}),"constraints")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#default"}),"default")),Object(c.b)("li",{parentName:"ul"},Object(c.b)("a",Object(n.a)({parentName:"li"},{href:"#issecret"}),"isSecret"))))),Object(c.b)("h3",{id:"constraints"},"constraints"),Object(c.b)("p",null,"A ",Object(c.b)("a",Object(n.a)({parentName:"p"},{href:"https://tools.ietf.org/html/draft-wright-json-schema-00"}),"JSON Schema v4 [object]")," defining constraints on the parameters value."),Object(c.b)("h4",{id:"default"},"default"),Object(c.b)("p",null,"A string to use as the value of the parameter when no argument is provided."),Object(c.b)("h4",{id:"description"},"description"),Object(c.b)("p",null,"A ",Object(c.b)("a",Object(n.a)({parentName:"p"},{href:"/docs/opspec/reference/structure/op-directory/op/markdown"}),"markdown [string]")," defining a human friendly description of the parameter."),Object(c.b)("h4",{id:"issecret"},"isSecret"),Object(c.b)("p",null,"A boolean indicating if the value of the parameter is secret. This will cause it to be hidden in UI's for example."))}u.isMDXComponent=!0},147:function(e,t,r){"use strict";r.d(t,"a",(function(){return u})),r.d(t,"b",(function(){return m}));var n=r(0),a=r.n(n);function c(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function o(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?o(Object(r),!0).forEach((function(t){c(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function p(e,t){if(null==e)return{};var r,n,a=function(e,t){if(null==e)return{};var r,n,a={},c=Object.keys(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||(a[r]=e[r]);return a}(e,t);if(Object.getOwnPropertySymbols){var c=Object.getOwnPropertySymbols(e);for(n=0;n<c.length;n++)r=c[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(a[r]=e[r])}return a}var s=a.a.createContext({}),l=function(e){var t=a.a.useContext(s),r=t;return e&&(r="function"==typeof e?e(t):i({},t,{},e)),r},u=function(e){var t=l(e.components);return a.a.createElement(s.Provider,{value:t},e.children)},b="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return a.a.createElement(a.a.Fragment,{},t)}},f=Object(n.forwardRef)((function(e,t){var r=e.components,n=e.mdxType,c=e.originalType,o=e.parentName,s=p(e,["components","mdxType","originalType","parentName"]),u=l(r),b=n,f=u["".concat(o,".").concat(b)]||u[b]||d[b]||c;return r?a.a.createElement(f,i({ref:t},s,{components:r})):a.a.createElement(f,i({ref:t},s))}));function m(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var c=r.length,o=new Array(c);o[0]=f;var i={};for(var p in t)hasOwnProperty.call(t,p)&&(i[p]=t[p]);i.originalType=e,i[b]="string"==typeof e?e:n,o[1]=i;for(var s=2;s<c;s++)o[s]=r[s];return a.a.createElement.apply(null,o)}return a.a.createElement.apply(null,r)}f.displayName="MDXCreateElement"}}]);