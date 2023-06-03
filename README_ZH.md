# task_bot

一个发送日常任务，帮助养成日常习惯的机器人，由 `GitHub Actions` 驱动。

[English](./README.md)

## 工作原理

1. 在 `daily.toml` 配置日常任务

```toml
[task.reading]
describe = "今天的回顾旧书是 %v, 享受阅读的乐趣吧！"
tasks = [
     "《这些人那些事》",
     "《The Minimalist Entrepreneur》",
     "《侠隐》",
]
```

2. 把邮件配置添加到 `Actions secrets` 

去仓库页面 -> `Settings` -> `Security` -- `Secrets and variables` -> `Actions` -> `New repository secret`

| Name         | Secret(示例)   |
| :----------- | :---------------- |
| EMAIL_SERVER | smtp.mail.com:587 |
| EMAIL_USER   | who_send_task     |
| EMAIL_PASS   | sender_password   |
| EMAIL_FROM   | sender@mail.com   |
| EMAIL_TO     | task_to@mail.com  |

3. 在 `send_task.yml` 编排任务时间

```yml
on:
  schedule:
    - cron: '* 2 * * *' # 每天 UTC 时间早上两点发送日常任务细节的邮件给你
```

更多配置方法: [cron config](https://en.wikipedia.org/wiki/Cron)

## 还有一件事

如果你想要更多新功能，开一个 issue 告诉我。

## 许可证

MIT
