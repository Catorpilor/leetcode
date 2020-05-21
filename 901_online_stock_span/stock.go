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

// Next amotized O(1)
// One price will be pushed once and popped once.
// So 2 * N times stack operations and N times calls.
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
