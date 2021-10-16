package com.bookstore.ecommerce.app.order.command;

public interface CreateOrderCommand {
  void execute(CreateOrderRequest request) throws Exception;
}
