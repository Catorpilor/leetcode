package stock

type spanner struct {
    price int
    span  int
}

type StockSpanner struct {
    store []spanner
}

func Construct() *StockSpanner {
    return &StockSpanner{}
}

// Next worst case time complexity O(N), average O(1)
func (ss *StockSpanner) Next(price int) int {
    res := 1
    nl := len(ss.store)
    for nl > 0 && ss.store[nl-1].price <= price {
        top := ss.store[nl-1]
        ss.store = ss.store[:nl-1]
        nl--
        res += top.span
    }
    ss.store = append(ss.store, spanner{price: price, span: res})
    return res
}
