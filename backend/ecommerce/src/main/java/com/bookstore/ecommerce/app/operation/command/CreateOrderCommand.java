package com.bookstore.ecommerce.app.operation.command;

public interface CreateOrderCommand {
  void execute(CreateOrderRequest request) throws Exception;
}
