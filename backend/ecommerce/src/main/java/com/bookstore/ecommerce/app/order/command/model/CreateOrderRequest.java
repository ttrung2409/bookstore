package com.bookstore.ecommerce.app.order.command.model;

import lombok.Data;

@Data
public class CreateOrderRequest {
  private Customer customer;
  private Book[] books;
}
