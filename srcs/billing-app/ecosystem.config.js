module.exports = {
    apps: [
        {
            name: "billing-app",
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
            out_file: "~/.pm2/logs/billing-app-out.log",
            error_file: "~/.pm2/logs/billing-app-error.log",
            merge_logs: true,
            time: true
        }
    ]
};