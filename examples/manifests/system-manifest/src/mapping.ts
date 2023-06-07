import { BigInt, cosmos, store } from "@graphprotocol/graph-ts";
import { Rule, Binding, Relation, Proposal, Block } from "../generated/schema";

export function handleRuleProposal(data: cosmos.EventData): void {

  const name = data.event.getAttributeValue("rule_name");
  const content = data.event.getAttributeValue("rule_content");
  const hash = data.event.getAttributeValue("rule_hash");
  const operationType = data.event.getAttributeValue("operation_type");

  if (operationType == "insert") {
    let rule = new Rule(name);
    rule.content = content;
    rule.name = name;
    rule.hash = hash;
    rule.save();
  } else if (operationType == "update") {
    let rule = Rule.load(name);
    if (rule != null) {
      rule.content = content;
      rule.hash = hash;
      rule.save();
    }
  } else if (operationType == "delete") {
    let rule = Rule.load(name);
    if (rule != null) {
      store.remove("Rule", name);
    }
  }
}

export function handleBindingProposal(data: cosmos.EventData): void {

  const name = data.event.getAttributeValue("binding_name");
  const content = data.event.getAttributeValue("binding_content");
  const hash = data.event.getAttributeValue("binding_hash");
  const rulefilesNames = data.event.getAttributeValue("binding_rule_files_names");
  const operationType = data.event.getAttributeValue("operation_type");

  if (operationType == "insert") {
    let binding = new Binding(name)
    binding.content = content;
    binding.hash = hash;
    binding.name = name;
    binding.rules = rulefilesNames.split(",");
    binding.save();
  } else if (operationType == "update") {
    let binding = Binding.load(name);
    if (binding != null) {
      binding.content = content;
      binding.hash = hash;
      binding.rules = rulefilesNames.split(",");
      binding.save();
    }
  } else if (operationType == "delete") {
    let binding = Binding.load(name);
    if (binding != null) {
      store.remove("Binding", name);
    }
  }
}
export function handleRelationProposal(data: cosmos.EventData): void {

  const bindingName = data.event.getAttributeValue("binding_name");
  const contractAddress = data.event.getAttributeValue("contract_address");
  const operationType = data.event.getAttributeValue("operation_type");

  if (operationType == "insert") {
    let relation = new Relation(contractAddress);
    relation.binding = bindingName;
    relation.contract_address = contractAddress;
    relation.save();
  } else if (operationType == "update") {
    let relation = Relation.load(contractAddress);
    if (relation != null) {
      relation.binding = bindingName;
      relation.save();
    }
  } else if (operationType == "delete") {
    let relation = Relation.load(contractAddress);
    if (relation != null) {
      store.remove("Relation", contractAddress);
    }
  }
}

export function handleProposal(data: cosmos.EventData): void {

  const id = data.event.getAttributeValue("proposal_id");
  const result = data.event.getAttributeValue("proposal_result");

  let proposal = new Proposal(id);
  proposal.result = result;
  proposal.save();
}
export function handleBlock(data: cosmos.Block): void {
  let newBlock = new Block(data.header.hash.toHexString());
  newBlock.hash = data.header.hash.toHexString();
  newBlock.height = BigInt.fromString(data.header.height.toString());
  newBlock.app_hash = data.header.appHash.toHexString();
  newBlock.data_hash = data.header.dataHash.toHexString();

  newBlock.proposer_address = data.header.proposerAddress.toHexString();
  newBlock.timestamp = BigInt.fromString(data.header.time.nanos.toString());
  newBlock.save();
}