package com.bookstore.ecommerce.app.order.query.dto;

import java.time.Instant;
import java.util.List;
import java.util.stream.Collectors;
import lombok.Builder;
import lombok.Data;

@Data
@Builder
public class Order {
  private String id;
  private String number;
  private String status;
  private Instant createdAt;
  private Customer customer;
  private List<OrderItem> items;

  public static Order fromDataObject(com.bookstore.ecommerce.app.domain.data.Order order) {
    return Order.builder()
      .id(order.getId())
      .number(order.getNumber())
      .status(order.getStatus())
      .createdAt(order.getCreatedAt())
      .customer(Customer.builder()
        .id(order.getCustomerId())
        .name(order.getCustomerName())
        .phone(order.getCustomerPhone())
        .deliveryAddress(order.getCustomerDeliveryAddress())
        .build())
      .items(order.getItems().stream()
        .map(item -> OrderItem.fromDataObject(item))
        .collect(Collectors.toList()))
      .build();
  }
}
