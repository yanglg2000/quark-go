"use strict";(self.webpackChunk=self.webpackChunk||[]).push([[717],{52756:function(Pe,b,e){e.r(b);var Z=e(97857),L=e.n(Z),$=e(15009),c=e.n($),G=e(19632),F=e.n(G),H=e(99289),C=e.n(H),q=e(5574),p=e.n(q),d=e(67294),r=e(80854),z=e(64793),J=e(94785),P=e(17331),N=e(45104),Q=e(66949),M=e(43083),D=e(91373),V=e(42952),X=e(92443),Y=e(87547),a=e(85893),w=function(ee){var ne=(0,r.useModel)("pageLoading"),te=ne.pageLoading,ae=(0,d.useState)({}),A=p()(ae,2),u=A[0],_e=A[1],ue=(0,d.useState)({}),K=p()(ue,2),U=K[0],y=K[1],se=(0,r.useModel)("pageLoading"),O=se.setPageLoading,oe=(0,r.useModel)("@@initialState"),R=oe.initialState,s=R==null?void 0:R.accountInfo,j=(0,r.useLocation)(),m=J.Z.parse(j.search),re=(0,d.useState)(""),I=p()(re,2),le=I[0],T=I[1],ie=(0,d.useState)([]),S=p()(ie,2),h=S[0],B=S[1],de=(0,d.useState)([""]),W=p()(de,2),me=W[0],k=W[1],Ee=(0,d.useState)(!1),x=p()(Ee,2),ce=x[0],pe=x[1],ve=function(){var _=C()(c()().mark(function n(){var o,E,t,g,l,f;return c()().wrap(function(i){for(;;)switch(i.prev=i.next){case 0:return i.next=2,(0,r.request)("./config.json");case 2:if(o=i.sent,E=o.api.layout,E){i.next=7;break}return y("\u8BF7\u8BBE\u7F6E\u5E03\u5C40\u63A5\u53E3\uFF01"),i.abrupt("return");case 7:return O(!0),i.next=10,(0,D.U)({url:E});case 10:t=i.sent,_e(t),t.menu&&m.api&&(g=(0,M.hA)(t.menu,decodeURIComponent(v)),l=(0,M.Kg)(t.menu,g),h.push.apply(h,F()(l)),B(h),k([g])),f=(0,M.Ah)(t.menu,j.pathname,decodeURIComponent(v)),T(f),O(!1);case 16:case"end":return i.stop()}},n)}));return function(){return _.apply(this,arguments)}}(),fe=function(){var _=C()(c()().mark(function n(){var o,E,t;return c()().wrap(function(l){for(;;)switch(l.prev=l.next){case 0:if(o=(0,M.Ah)(u.menu,j.pathname,decodeURIComponent(v)),T(o),v){l.next=5;break}return y(null),l.abrupt("return");case 5:return E={},Object.keys(m).forEach(function(f){f!=="api"&&(E[f]=m[f])}),O(!0),l.next=10,(0,D.U)({url:v,data:E});case 10:t=l.sent,y(t),O(!1);case 13:case"end":return l.stop()}},n)}));return function(){return _.apply(this,arguments)}}(),v="";m!=null&&m.api&&(v=m.api),(0,d.useEffect)(function(){ve()},[]),(0,d.useEffect)(function(){fe()},[m.api]);var Me=function(n){k([n.key]);var o=(0,M.A9)(u.menu,n.key);if(o.indexOf("http")===0)return window.open(o,"_blank"),!1;r.history.push(o)},Oe=function(n){B(n)},he=function(n){switch(n.key){case"logout":ge();break;case"setting":r.history.push({pathname:"/layout/index",search:"api=/api/admin/account/setting/form"});break}},ge=function(){var _=C()(c()().mark(function n(){var o;return c()().wrap(function(t){for(;;)switch(t.prev=t.next){case 0:return t.next=2,(0,D.U)({url:"/api/admin/logout/index/handle"});case 2:o=t.sent,o.status==="success"&&sessionStorage.removeItem("token"),r.history.push("/");case 5:case"end":return t.stop()}},n)}));return function(){return _.apply(this,arguments)}}(),Ce=[{key:"setting",icon:(0,a.jsx)(V.Z,{}),label:"\u4E2A\u4EBA\u8BBE\u7F6E"},{key:"logout",icon:(0,a.jsx)(X.Z,{}),label:"\u9000\u51FA\u767B\u5F55"}];return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsxs)(r.Helmet,{children:[(0,a.jsx)("meta",{charSet:"utf-8"}),(0,a.jsx)("title",{children:le})]}),(u==null?void 0:u.title)&&(0,a.jsx)(z.f,L()(L()({},u),{},{loading:te,logo:u.logo?u.logo:Q,iconfontUrl:u.iconfontUrl,openKeys:h,selectedKeys:me,menuProps:{onOpenChange:Oe,onClick:Me},onCollapse:function(n){pe(n)},menuDataRender:function(){return u.menu},actionsRender:function(){return[(0,a.jsx)(P.Z,{body:u.actions},"action")]},rightContentRender:function(){return(0,a.jsx)(N.Z,{menu:{items:Ce,onClick:he},avatar:s!=null&&s.avatar?s==null?void 0:s.avatar:(0,a.jsx)(Y.Z,{}),name:ee.layout==="side"?ce||s==null?void 0:s.nickname:s==null?void 0:s.nickname})},footerRender:function(){return(0,a.jsx)(P.Z,{body:u.footer})},children:U?(0,a.jsx)(P.Z,{body:U}):(0,a.jsx)(r.Outlet,{})}))]})};b.default=w}}]);
