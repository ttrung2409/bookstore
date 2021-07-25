package com.bookstore.ecommerce.app.operation.command;

import lombok.Data;

@Data
public class Order {
  private Customer customer;
  private OrderItem[] items;
}
