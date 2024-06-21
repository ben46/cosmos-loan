
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


## 添加expected_keepers代码
添加接口函数, 以及返回错误

## 添加keeper/msg_server_xxx代码
(处理CRUD)

## 添加types/message validateBasic代码

## 什么时候添加panic错误?

在提供的代码示例中，`panic` 错误主要用于处理 **不可恢复的错误**，例如：

* **解析地址失败：** 当尝试从 Bech32 字符串解析地址时，如果解析失败，就会使用 `panic` 错误。例如：

```go
borrower, err := sdk.AccAddressFromBech32(msg.Creator)
if err != nil {
    panic(err)
}
```

* **解析 Coins 失败：** 当尝试解析 Coins 对象时，如果解析失败，就会使用 `panic` 错误。例如：

```go
collateral, err := sdk.ParseCoinsNormalized(loan.Collateral)
if err != nil {
    panic(err)
}
```

* **解析 Deadline 失败：** 当尝试解析 Deadline 时，如果解析失败，就会使用 `panic` 错误。例如：

```go
deadline, err := strconv.ParseInt(loan.Deadline, 10, 64)
if err != nil {
    panic(err)
}
```

**使用 `panic` 错误的理由：**

* 这些错误通常是 **不可恢复的**，例如地址解析失败意味着代码无法继续执行。
* 使用 `panic` 错误可以 **快速终止程序**，避免程序继续运行并可能导致更严重的问题。

**需要注意的是：**

* 在生产环境中，使用 `panic` 错误应该谨慎，因为它们会导致程序崩溃。
* 对于可恢复的错误，应该使用 `error` 对象进行处理，而不是 `panic` 错误。

**建议：**

* 在开发环境中，可以使用 `panic` 错误来帮助快速定位问题。
* 在生产环境中，应该使用 `error` 对象进行处理，并记录错误信息，以便进行调试和修复。

## 什么时候返回types自定义错误?

在提供的代码示例中，自定义错误主要用于处理 **业务逻辑相关的错误**，例如：

* **贷款状态错误：** 当贷款状态不符合预期时，就会返回自定义错误。例如：

```go
if loan.State != "requested" {
    return nil, errorsmod.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
}
```

* **操作时间错误：** 当操作时间不符合要求时，就会返回自定义错误。例如：

```go
if ctx.BlockHeight() < deadline {
    return nil, errorsmod.Wrap(types.ErrDeadline, "Cannot liquidate before deadline")
}
```

**使用自定义错误的理由：**

* 自定义错误可以 **更清晰地描述错误原因**，方便开发者进行调试和修复。
* 自定义错误可以 **方便地进行错误处理**，例如根据错误类型进行不同的处理逻辑。

**建议：**

* 对于业务逻辑相关的错误，应该使用自定义错误进行处理。
* 自定义错误应该包含清晰的错误信息，方便开发者理解错误原因。

## 为什么有的验证代码写在types里面, 有的写在keeper里面?

* `types` 文件夹中的验证主要用于确保消息的格式和数据类型正确，属于基本验证, **是和数据库没有交互的**。
* `keeper` 文件夹中的验证主要用于确保消息在当前状态下是合法的，属于业务逻辑验证。

* **`keeper` 文件夹中的验证：** 这些验证主要用于 **业务逻辑验证**，确保消息在当前状态下是合法的。例如，`x/loan/keeper/msg_server_approve_loan.go` 中的 `ApproveLoan` 函数检查了贷款是否处于 "requested" 状态，以及贷款的创建者是否与当前用户一致。

