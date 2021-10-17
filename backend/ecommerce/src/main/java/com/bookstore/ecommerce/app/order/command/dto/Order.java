package com.bookstore.ecommerce.app.order.command.dto;

import lombok.Data;

@Data
public class Order {
  private Customer customer;
  private OrderItem[] items;
}
