# Go-Learning
A daily record of learning the Go language.


memcache/
│
├── cmd/                 # 应用程序的入口点
│   └── main.go      # 示例服务的主程序
│
├── pkg/                 # 可以被外部应用引用的库代码
│   ├── cache/           # 缓存逻辑的核心代码
│   │   ├── cache.go     # 定义缓存的接口和基础逻辑
│   │   └── memory.go    # 具体的内存缓存实现
│   └── config/          # 配置加载和解析
│       └── config.go    # 加载和解析配置文件的逻辑
│
├── internal/            # 私有的应用程序和库代码
│   └── util/            # 内部使用的工具和辅助函数
│       └── util.go      # 工具和辅助函数的实现
│
├── .gitignore           # Git忽略文件列表
├── go.mod               # Go模块描述文件
├── go.sum               # Go模块的依赖树
└── README.md            # 项目说明文件
