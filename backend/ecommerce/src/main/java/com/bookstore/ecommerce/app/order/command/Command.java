package com.bookstore.ecommerce.app.order.command;

import java.util.concurrent.CompletableFuture;
import com.bookstore.ecommerce.app.order.command.dto.CreateOrderRequest;

public interface Command {
  CompletableFuture<String> createOrder(CreateOrderRequest request) throws Exception;
}
