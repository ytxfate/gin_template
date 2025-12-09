module.exports = {
  apps: [{
    name: "gin_template-api",           // 应用名称
    script: "./gin_template-api",       // 编译后的Go可执行文件
    cwd: "./",                    // 应用启动目录
    watch: false,                 // 是否监听文件变化重启
    max_memory_restart: "1024M",  // 内存限制
    log_date_format: "YYYY-MM-DD HH:mm:ss.SSS",                // 日志格式
    error_file: "./logs/gin_template-api-error.log",            // 错误日志
    out_file: "./logs/gin_template-api-out.log",                // 标准输出日志
    autorestart: true,  // 自动重启
    max_restarts: 10,   // 最大重启次数
    restart_delay: 5000,// 重启延迟（毫秒）
    exec_mode: "fork",
  }]
};
