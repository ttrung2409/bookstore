package com.bookstore.ecommerce.app.domain;

public class Order {
  private com.bookstore.ecommerce.app.data.Order state;

  public Order(com.bookstore.ecommerce.app.data.Order order) {
    this.state = order;
  }
}
