class Solution:
    def JobSequencing(self, id, deadline, profit):
        tuple_items = [(a, b, c) for (a, b, c) in zip(profit, deadline, id)]
        tuple_items.sort(reverse=True)

        n = len(id)

        res = [0] * n

        for item_profit, item_deadline, _ in tuple_items:
            ind = item_deadline - 1

            while ind >= 0 and res[ind] != 0:
                ind -= 1

            if ind >= 0:
                res[ind] = item_profit

        return len(res) - res.count(0), sum(res)
