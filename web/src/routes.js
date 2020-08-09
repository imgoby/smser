export default [
    {
        path: '/user',
        component: resolve=>require(['@/components/common/UserLayout'], resolve),
        children: [
            {
                path: 'login',
                component: resolve=>require(['@/components/auth/Login'], resolve),
                parent: "",
                name: "登录"
            }
        ]
    },
    {
        path: '/',
        component: resolve=>require(['@/components/common/Layout'], resolve),
        children: [
            {
                path: '/',
                component: resolve=>require(['@/components/index/Index'], resolve),
                parent: "",
                name: "首页"
            },
            {
                path: '/course/list',
                component: resolve=>require(['@/components/course/List'], resolve),
                parent: "课程管理",
                name: "课程列表"
            },
            {
                path: '/course/create',
                component: resolve=>require(['@/components/course/Operate'], resolve),
                parent: "课程管理",
                name: "课程添加"
            },
            {
                path: '/message/list',
                component: resolve=>require(['@/components/message/Message'], resolve),
                name: "消息列表"
            },
            {
                path: '/logs/worker',
                component: resolve=>require(['@/components/log/WorkerLog'], resolve),
                name: "消费日志"
            },
            {
                path: '/logs/message',
                component: resolve=>require(['@/components/log/MessageLog'], resolve),
                name: "发送日志"
            },
            {
                path: '/logs/request',
                component: resolve=>require(['@/components/log/RequestLog'], resolve),
                name: "请求日志"
            },
            {
                path: '/logs/log',
                component: resolve=>require(['@/components/log/Log'], resolve),
                name: "运行日志"
            },
        ]
    },
];
