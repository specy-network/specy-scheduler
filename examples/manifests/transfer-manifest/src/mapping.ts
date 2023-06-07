import { BigInt, cosmos, log, store } from "@graphprotocol/graph-ts";
import { Transfer } from "../generated/schema";

/*
处理transfer类event
*/
export function handleTransferEvent(data: cosmos.EventData, txHash: string): void {
  const sender = data.event.getAttributeValue("sender");
  const receiver = data.event.getAttributeValue("recipient");
  let transfer = new Transfer(txHash);
  transfer.id = txHash;
  transfer.hash = txHash;
  transfer.sender = sender;
  transfer.receiver = receiver;
  //{"type":"transfer","sender":"Address1","recipient":"Address2","amount":1000,"denom":"stake","module":"bank"}
  transfer.value = BigInt.fromString(data.event.getAttributeValue("amount"));
  transfer.tokenname = data.event.getAttributeValue("denom");
  transfer.timestamp = BigInt.fromString(
    data.block.header.time.seconds.toString()
  );
  transfer.contract_address = data.event.getAttributeValue("contract_address");
  transfer.save();
}

/*
交易处理逻辑，由于cosmos中event无法获取合约地址和交易hash等信息，故此处由txhandle统一处理所有event
之后根据event类型交由具体handle处理
*/
export function handleTx(data: cosmos.TransactionData): void {
  const txHash = data.tx.hash.toHexString();
  const events = data.tx.result.events;
  let found = false;
  let transfer = Transfer.load(txHash);
  for (let index = 0; index < events.length; index++) {
    const event: cosmos.Event = events[index];
    let hob = new cosmos.HeaderOnlyBlock(data.block.header);
    let ed = new cosmos.EventData(event, hob);
    if (event.eventType == "transfer") {
      if (transfer == null) {
        handleTransferEvent(ed, txHash);
        log.info("处理transfer完成:{}", [txHash])
      }
      found = true;
    }
  }
  if (found == false && transfer != null) {
    store.remove("Transfer", txHash);
    log.info("监管未通过,删除transfer记录:{}", [txHash])
  }
}
