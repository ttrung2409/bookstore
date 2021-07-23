package com.bookstore.ecommerce.app.operation.command;

public interface CreateOrderCommand {
  String execute(Order order);
}
