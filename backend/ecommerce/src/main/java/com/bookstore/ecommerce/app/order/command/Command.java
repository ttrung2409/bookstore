package com.bookstore.ecommerce.app.order.command;

import com.bookstore.ecommerce.app.order.command.model.CreateOrderRequest;

public interface Command {
  String createOrder(CreateOrderRequest request) throws Exception;
}
