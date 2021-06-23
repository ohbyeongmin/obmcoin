https://github.com/ohbyeongmin/obmcoin/commit/49ef0465c588e047c29ae830ec14f5af91049c61
// // 보내는 사람의 balance 가 amount(보내는 코인) 보다 적으면 에러
// if Blockchain().BalanceByAddress(from) < amount {
// return nil, errors.New("NOT ENOUGH MONEY")
// }
// var txIns []*TxIn
// var txOuts []*TxOut
// total := 0
// // 보내는 사람의 txOut 을 모두 가져옴
// oldTxOuts := Blockchain().TxOutsByAddress(from)
// for \_, txOut := range oldTxOuts {
// // out에서 코인의 양을 차례로 더하는 과정중 amount 보다 많아지면 loop 를 멈춤
// if total > amount {
// break;
// }
// txIn := &TxIn{txOut.Owner, txOut.Amount}
// txIns = append(txIns, txIn)
// total += txIn.Amount
// }
// change := total - amount
// // total 에서 남는 코인은 보낸사람의 것으로 out을 내어준다.
// if change != 0 {
// changeTxOut := &TxOut{from, change}
// txOuts = append(txOuts, changeTxOut)
// }
// txOut := &TxOut{to, amount}
// txOuts = append(txOuts, txOut)
// tx := &Tx{
// Id: "",
// Timestamp: int(time.Now().Unix()),
// TxIns: txIns,
// TxOuts: txOuts,
// }
// tx.getId()
// return tx, nil
