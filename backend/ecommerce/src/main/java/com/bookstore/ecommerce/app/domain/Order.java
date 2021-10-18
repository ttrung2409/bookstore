package com.bookstore.ecommerce.app.domain;

import java.time.Instant;
import java.util.ArrayList;
import java.util.UUID;

import com.bookstore.ecommerce.app.domain.data.Book;
import com.bookstore.ecommerce.app.domain.data.Customer;
import com.bookstore.ecommerce.app.domain.data.OrderItem;
import com.bookstore.ecommerce.app.domain.data.OrderStatus;

import lombok.Getter;

public class Order {
  @Getter
  private final com.bookstore.ecommerce.app.domain.data.Order state;

  public Order(Customer customer, Book[] books) {
    final var id = UUID.randomUUID().toString();

    final var items = new ArrayList<OrderItem>();

    for (final var book : books) {
      items.add(OrderItem.builder()
        .key(new OrderItem.Key(id, book.getId()))
        .bookTitle(book.getTitle())
        .bookSubtitle(book.getSubtitle())
        .bookDescription(book.getDescription())
        .build());
    }

    this.state = com.bookstore.ecommerce.app.domain.data.Order.builder()
      .id(id)
      .number(id)
      .createdAt(Instant.now())
      .status(OrderStatus.Queued.toString())
      .customerName(customer.getName())
      .customerPhone(customer.getPhone())
      .customerDeliveryAddress(customer.getDeliveryAddress())
      .items(items)
      .build();
  }
}
