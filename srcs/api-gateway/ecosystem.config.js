module.exports = {
    apps: [
        {
            name: "api-gateway",
            script: "tmp/main",
            watch: false,
            autorestart: true,
            restart_delay: 2000,
            kill_timeout: 5000,
            max_memory_restart: "200M",
            env_file: ".env",           
            env: {
                PATH: process.env.PATH
            },
            out_file: "~/.pm2/logs/api-gateway-out.log",
            error_file: "~/.pm2/logs/api-gateway-error.log",
            merge_logs: true,
            time: true
        }
    ]
};