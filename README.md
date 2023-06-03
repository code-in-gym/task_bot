# task_bot
A robot that sends a daily task and motivates to cultivate a daily habit, power by `GitHub Actions`.

[中文](./README_ZH.md)

## How it works

1. config daily task in `daily.toml`

```toml
[task.reading]
describe = "Today's review of an old book is %v. Enjoy reading for fun!"
tasks = [
    "《animal farm》",
    "《The Minimalist Entrepreneur》",
    "《Lemo》",
]
```

2. add email config to Actions secrets

Go to repository page -> `Settings` -> `Security` -- `Secrets and variables` -> `Actions` -> `New repository secret`

| Name         | Secret(example)   |
| :----------- | :---------------- |
| EMAIL_SERVER | smtp.mail.com:587 |
| EMAIL_USER   | who_send_task     |
| EMAIL_PASS   | sender_password   |
| EMAIL_FROM   | sender@mail.com   |
| EMAIL_TO     | task_to@mail.com  |

3. schedule task in `send_task.yml`

```yml
on:
  schedule:
    - cron: '* 2 * * *' # send you email with task detail at 2:00 AM UTC every day
```

more: [cron config](https://en.wikipedia.org/wiki/Cron)

## One more thing

Tell me with an issue if you want more feature about this project.

## License

MIT
