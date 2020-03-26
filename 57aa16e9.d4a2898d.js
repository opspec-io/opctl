(window.webpackJsonp=window.webpackJsonp||[]).push([[23],{149:function(e,t,r){"use strict";r.r(t),r.d(t,"frontMatter",(function(){return i})),r.d(t,"metadata",(function(){return o})),r.d(t,"rightToc",(function(){return l})),r.d(t,"default",(function(){return s}));var n=r(1),c=r(9),a=(r(0),r(181)),i={title:"Dir"},o={id:"opspec/reference/types/dir",title:"Dir",description:"Dir typed values are a filesystem directory entry.",source:"@site/docs/opspec/reference/types/dir.md",permalink:"/docs/opspec/reference/types/dir",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/opspec/reference/types/dir.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1585204905,sidebar:"docs",previous:{title:"Boolean",permalink:"/docs/opspec/reference/types/boolean"},next:{title:"File",permalink:"/docs/opspec/reference/types/file"}},l=[{value:"Entry Referencing",id:"entry-referencing",children:[]}],p={rightToc:l},b="wrapper";function s(e){var t=e.components,r=Object(c.a)(e,["components"]);return Object(a.b)(b,Object(n.a)({},p,r,{components:t,mdxType:"MDXLayout"}),Object(a.b)("p",null,"Dir typed values are a filesystem directory entry."),Object(a.b)("p",null,"Dirs..."),Object(a.b)("ul",null,Object(a.b)("li",{parentName:"ul"},"are mutable, i.e. making changes to a directory results in the directory being changed everywhere it's referenced."),Object(a.b)("li",{parentName:"ul"},"can be passed in/out of ops via ",Object(a.b)("a",Object(n.a)({parentName:"li"},{href:"/docs/opspec/reference/structure/op-directory/op/parameter/dir"}),"dir parameters"),"."),Object(a.b)("li",{parentName:"ul"},"are not coercible to any other type.")),Object(a.b)("h3",{id:"entry-referencing"},"Entry Referencing"),Object(a.b)("p",null,"Dir entries (child files/directories) can be referenced via ",Object(a.b)("inlineCode",{parentName:"p"},"$(ROOT/ENTRY)")," syntax."),Object(a.b)("h4",{id:"entry-referencing-example-embedded"},"Entry Referencing Example (embedded)"),Object(a.b)("p",null,"given:"),Object(a.b)("ul",null,Object(a.b)("li",{parentName:"ul"},Object(a.b)("inlineCode",{parentName:"li"},"/file1.json")," is embedded in op")),Object(a.b)("pre",null,Object(a.b)("code",Object(n.a)({parentName:"pre"},{className:"language-yaml"}),"$(/file1.json)\n")),Object(a.b)("h4",{id:"entry-referencing-example-scope"},"Entry Referencing Example (scope)"),Object(a.b)("p",null,"given:"),Object(a.b)("ul",null,Object(a.b)("li",{parentName:"ul"},Object(a.b)("inlineCode",{parentName:"li"},"someDir"),Object(a.b)("ul",{parentName:"li"},Object(a.b)("li",{parentName:"ul"},"is in scope dir"),Object(a.b)("li",{parentName:"ul"},"contains ",Object(a.b)("inlineCode",{parentName:"li"},"file2.txt"))))),Object(a.b)("pre",null,Object(a.b)("code",Object(n.a)({parentName:"pre"},{className:"language-yaml"}),"$(someDir/file2.txt)\n")))}s.isMDXComponent=!0},181:function(e,t,r){"use strict";r.d(t,"a",(function(){return s})),r.d(t,"b",(function(){return m}));var n=r(0),c=r.n(n);function a(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function i(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(e);t&&(n=n.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,n)}return r}function o(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?i(Object(r),!0).forEach((function(t){a(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):i(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function l(e,t){if(null==e)return{};var r,n,c=function(e,t){if(null==e)return{};var r,n,c={},a=Object.keys(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||(c[r]=e[r]);return c}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(n=0;n<a.length;n++)r=a[n],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(c[r]=e[r])}return c}var p=c.a.createContext({}),b=function(e){var t=c.a.useContext(p),r=t;return e&&(r="function"==typeof e?e(t):o({},t,{},e)),r},s=function(e){var t=b(e.components);return c.a.createElement(p.Provider,{value:t},e.children)},u="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return c.a.createElement(c.a.Fragment,{},t)}},f=Object(n.forwardRef)((function(e,t){var r=e.components,n=e.mdxType,a=e.originalType,i=e.parentName,p=l(e,["components","mdxType","originalType","parentName"]),s=b(r),u=n,f=s["".concat(i,".").concat(u)]||s[u]||d[u]||a;return r?c.a.createElement(f,o({ref:t},p,{components:r})):c.a.createElement(f,o({ref:t},p))}));function m(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var a=r.length,i=new Array(a);i[0]=f;var o={};for(var l in t)hasOwnProperty.call(t,l)&&(o[l]=t[l]);o.originalType=e,o[u]="string"==typeof e?e:n,i[1]=o;for(var p=2;p<a;p++)i[p]=r[p];return c.a.createElement.apply(null,i)}return c.a.createElement.apply(null,r)}f.displayName="MDXCreateElement"}}]);