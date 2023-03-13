const Subscription = require('egg').Subscription;

class UpdateCache extends Subscription {
    // 通过 schedule 属性来设置定时任务的执行间隔等配置
    static get schedule() {
        return {
            interval: '1440m',
            type: 'all', // 指定所有的 worker 都需要执行
        };
    }

    // subscribe 是真正定时任务执行时被运行的函数
    async subscribe() {
        console.log("schedule start")
        const { ctx } = this;
        const ret = await ctx.service.opensea.saveCollection();
        console.log(ret)
    }
}

module.exports = UpdateCache;