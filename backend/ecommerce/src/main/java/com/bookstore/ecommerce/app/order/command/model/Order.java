package com.bookstore.ecommerce.app.order.command.model;

import lombok.Data;

@Data
public class Order {
  private Customer customer;
  private OrderItem[] items;
}
