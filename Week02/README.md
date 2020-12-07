# 第二周学习笔记

问题：
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

答：
需要Wrap这个error，并抛给上层。
dao层中遇到的sql.ErrNoRows是由kit库返回，所以需要对error进行wrap封装成本服务的error。
另外由于是查询数据的业务，所以需要将error上报到业务层让业务层对返回数据进行处理。
