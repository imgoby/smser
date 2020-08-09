export default [
    {
        menu: "控制面板",
        icon: "",
        url: "/",
        children: [
            {
                menu: "数据中心",
                icon: "",
                url: "/",
            },
        ]
    },
    {
        menu: "课程管理",
        icon: "",
        url: "/course",
        children: [
            {
                menu: "课程列表",
                icon: "",
                url: "/course/list",
            },
            {
                menu: "添加课程",
                icon: "",
                url: "/course/create",
            }
        ]
    },
    {
        menu: "消息管理",
        icon: "",
        url: "/message",
        children: [
            {
                menu: "消息列表",
                icon: "",
                url: "/message/list",
            },
        ]
    },
    {
        menu: "日志管理",
        icon: "",
        url: "/log",
        children: [
            {
                menu: "消费日志",
                icon: "",
                url: "/logs/worker",
            },
            {
                menu: "请求日志",
                icon: "",
                url: "/logs/request",
            },
            {
                menu: "发送日志",
                icon: "",
                url: "/logs/message",
            },
            {
                menu: "运行日志",
                icon: "",
                url: "/logs/log",
            },
        ]
    },
]
