package com.bookstore.ecommerce.app.operation.command;

import lombok.Data;

@Data
public class CreateOrderRequest {
  private Customer customer;
  private Book[] books;
}
