package com.bookstore.ecommerce.app.domain.data;

import java.time.Instant;
import java.util.List;

import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.Id;
import javax.persistence.Table;
import javax.persistence.Transient;

import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Getter;

@Entity
@Table(name = "order")
@AllArgsConstructor
@Builder
public class Order {
  @Getter
  @Id
  private String id;
  @Getter
  private String number;
  @Getter
  @Column(name = "created_at")
  private Instant createdAt;
  @Getter
  private String status;
  @Getter
  @Column(name = "customer_name")
  private String customerName;
  @Getter
  @Column(name = "customer_phone")
  private String customerPhone;
  @Getter
  @Column(name = "customer_delivery_address")
  private String customerDeliveryAddress;
  @Getter
  @Transient
  private OrderItem[] items;

  public Order() {
  }
}
