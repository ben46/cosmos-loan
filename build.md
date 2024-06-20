
**增加loan模块, 依赖bank模块**

```
ignite scaffold module loan --dep bank
```

**list是什么命令?**

“list”脚手架命令用于生成文件，实现在区块链状态中存储和与存储为列表的数据交互的逻辑。

```
ignite scaffold list loan amount fee collateral deadline state borrower lender --no-message
```

## 消息处理: 请求,批准,偿还,清算,取消贷款

申请贷款, 输入参数: 金额, 手续费, 保证金, 到期时间

`ignite scaffold message request-loan amount fee collateral deadline`

批准贷款, 贷款id

`ignite scaffold message approve-loan id:uint`

取消贷款, 贷款id

`ignite scaffold message cancel-loan id:uint`

偿还贷款, 贷款id

`ignite scaffold message repay-loan id:uint`

清算贷款, 贷款id

`ignite scaffold message liquidate-loan id:uint`