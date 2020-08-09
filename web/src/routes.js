const CourseList = resolve=>require(['@/components/course/List'], resolve);
const CourseOperate = resolve=>require(['@/components/course/Operate'], resolve);
const Index = resolve=>require(['@/components/index/Index'], resolve);
const Login = resolve=>require(['@/components/auth/Login'], resolve);

export default [
    {
        path: '/',
        component: Index,
        name: "首页",
        parent: "控制面板"
    },
    {
        path: '/index',
        component: Index,
        name: "首页",
        parent: "控制面板"
    },
    {
        path: '/course/list',
        component: CourseList,
        parent: "课程管理",
        name: "课程列表"
    },
    {
        path: '/course/create',
        component: CourseOperate,
        parent: "课程管理",
        name: "添加课程"
    },
    {
        path: '/login',
        component: Login,
        parent: "",
        name: "登陆"
    },
];