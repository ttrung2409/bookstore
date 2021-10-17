package com.bookstore.ecommerce.app.order.command.dto;

import lombok.Data;

@Data
public class CreateOrderRequest {
  private Customer customer;
  private Book[] books;
}
