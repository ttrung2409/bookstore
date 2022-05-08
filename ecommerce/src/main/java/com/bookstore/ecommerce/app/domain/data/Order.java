package com.bookstore.ecommerce.app.domain.data;

import java.time.Instant;
import java.util.ArrayList;
import java.util.List;
import javax.persistence.CascadeType;
import javax.persistence.Column;
import javax.persistence.Entity;
import javax.persistence.FetchType;
import javax.persistence.Id;
import javax.persistence.JoinColumn;
import javax.persistence.OneToMany;
import javax.persistence.Table;
import javax.persistence.Transient;
import org.modelmapper.ModelMapper;
import lombok.AccessLevel;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.Setter;
import lombok.var;

@Entity
@Table(name = "order")
@AllArgsConstructor
@Data
@Builder
public class Order implements Cloneable {
  @Id
  private String id;
  @Column(name = "created_at")
  private Instant createdAt;
  private String status;
  @Column(name = "customer_id")
  private String customerId;
  @Column(name = "customer_name")
  private String customerName;
  @Column(name = "customer_phone")
  private String customerPhone;
  @Column(name = "customer_delivery_address")
  private String customerDeliveryAddress;

  @OneToMany(cascade = CascadeType.ALL, fetch = FetchType.EAGER)
  @JoinColumn(name = "order_id")
  private List<OrderItem> items;

  public Order clone() {
    var order = new ModelMapper().map(this, Order.class);
    var items = new ArrayList<OrderItem>();
    for (OrderItem item : order.items) {
      items.add(item.clone());
    }

    return order;
  }
}
