package com.bookstore.ecommerce.app.domain;

import java.time.Instant;
import java.util.ArrayList;
import java.util.List;
import java.util.UUID;

import com.bookstore.ecommerce.app.domain.data.Book;
import com.bookstore.ecommerce.app.domain.data.Customer;
import com.bookstore.ecommerce.app.domain.data.OrderItem;
import com.bookstore.ecommerce.app.domain.data.OrderStatus;
import com.bookstore.ecommerce.app.domain.events.OrderCancelled;
import com.bookstore.ecommerce.app.domain.events.OrderCreated;
import com.google.common.base.Strings;
import lombok.var;

public class Order extends EventSource {
  private final com.bookstore.ecommerce.app.domain.data.Order state;

  public Order(Customer customer, List<Book> books) {
    final var orderId = UUID.randomUUID().toString();

    final var items = new ArrayList<OrderItem>();

    for (final var book : books) {
      items.add(OrderItem.builder()
        .key(new OrderItem.Key(orderId, book.getId()))
        .bookTitle(book.getTitle())
        .bookSubtitle(book.getSubtitle())
        .bookDescription(book.getDescription())
        .build());
    }

    this.state = com.bookstore.ecommerce.app.domain.data.Order.builder()
      .id(orderId)
      .createdAt(Instant.now())
      .status(OrderStatus.Pending.toString())
      .customerId(customer.getId())
      .customerName(customer.getName())
      .customerPhone(customer.getPhone())
      .customerDeliveryAddress(customer.getDeliveryAddress())
      .items(items)
      .build();

    this.pendingEvents.add(new OrderCreated(this.state));
  }


  public Order(com.bookstore.ecommerce.app.domain.data.Order order) {
    var cloned = order.clone();
    if (Strings.isNullOrEmpty(cloned.getId())) {
      cloned.setId(UUID.randomUUID().toString());
    }

    this.state = cloned;
  }

  public com.bookstore.ecommerce.app.domain.data.Order getState() {
    return this.state.clone();
  }

  public void cancel() throws Exception {
    if (this.state.getStatus() != OrderStatus.Pending.toString()
      && this.state.getStatus() != OrderStatus.Accepted.toString()) {
      throw new Exception(
        String.format("order status is %s, no cancellation allowed", this.state.getStatus()));
    }

    this.state.setStatus(OrderStatus.Cancelled.toString());

    this.pendingEvents.add(new OrderCancelled(this.state.getId()));
  }
}
