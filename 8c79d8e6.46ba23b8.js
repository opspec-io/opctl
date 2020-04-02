(window.webpackJsonp=window.webpackJsonp||[]).push([[33],{159:function(e,t,n){"use strict";n.r(t),n.d(t,"frontMatter",(function(){return c})),n.d(t,"metadata",(function(){return i})),n.d(t,"rightToc",(function(){return p})),n.d(t,"default",(function(){return u}));var r=n(1),o=n(9),a=(n(0),n(181)),c={title:"Kubernetes",sidebar_label:"Kubernetes"},i={id:"setup/kubernetes",title:"Kubernetes",description:"## Examples",source:"@site/docs/setup/kubernetes.md",permalink:"/docs/setup/kubernetes",editUrl:"https://github.com/opctl/opctl/edit/master/docs/docs/setup/kubernetes.md",lastUpdatedBy:"Chris Dostert",lastUpdatedAt:1585855263,sidebar_label:"Kubernetes",sidebar:"docs",previous:{title:"Gitlab",permalink:"/docs/setup/gitlab"},next:{title:"Travis",permalink:"/docs/setup/travis"}},p=[{value:"Examples",id:"examples",children:[]}],l={rightToc:p},s="wrapper";function u(e){var t=e.components,n=Object(o.a)(e,["components"]);return Object(a.b)(s,Object(r.a)({},l,n,{components:t,mdxType:"MDXLayout"}),Object(a.b)("h2",{id:"examples"},"Examples"),Object(a.b)("p",null,"Deploy opctl in kubernetes"),Object(a.b)("pre",null,Object(a.b)("code",Object(r.a)({parentName:"pre"},{className:"language-yaml"}),"apiVersion: apps/v1\nkind: Deployment\nmetadata:\n  name: opctl-in-kubernetes\nspec:\n  replicas: 1\n  template:\n    spec:\n      containers:\n        - name: opctl\n          image: opctl/opctl:0.1.29-dind\n          ports:\n            # expose to other containers\n            - name: http\n              containerPort: 42224\n              protocol: TCP\n          securityContext:\n            privileged: true\n")))}u.isMDXComponent=!0},181:function(e,t,n){"use strict";n.d(t,"a",(function(){return u})),n.d(t,"b",(function(){return f}));var r=n(0),o=n.n(r);function a(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function c(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?c(Object(n),!0).forEach((function(t){a(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):c(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function p(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var l=o.a.createContext({}),s=function(e){var t=o.a.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):i({},t,{},e)),n},u=function(e){var t=s(e.components);return o.a.createElement(l.Provider,{value:t},e.children)},b="mdxType",d={inlineCode:"code",wrapper:function(e){var t=e.children;return o.a.createElement(o.a.Fragment,{},t)}},m=Object(r.forwardRef)((function(e,t){var n=e.components,r=e.mdxType,a=e.originalType,c=e.parentName,l=p(e,["components","mdxType","originalType","parentName"]),u=s(n),b=r,m=u["".concat(c,".").concat(b)]||u[b]||d[b]||a;return n?o.a.createElement(m,i({ref:t},l,{components:n})):o.a.createElement(m,i({ref:t},l))}));function f(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var a=n.length,c=new Array(a);c[0]=m;var i={};for(var p in t)hasOwnProperty.call(t,p)&&(i[p]=t[p]);i.originalType=e,i[b]="string"==typeof e?e:r,c[1]=i;for(var l=2;l<a;l++)c[l]=n[l];return o.a.createElement.apply(null,c)}return o.a.createElement.apply(null,n)}m.displayName="MDXCreateElement"}}]);