package org.bbelovic.kata.exportlocation;

import java.util.Objects;

import static java.lang.String.format;

public class ExportLocation {
    private long sharedId;
    private boolean admin;
    private final boolean preferred;

    public ExportLocation(long sharedId, boolean isAdmin, boolean preferred) {
        this.sharedId = sharedId;
        this.admin = isAdmin;
        this.preferred = preferred;
    }

    public long getSharedId() {
        return sharedId;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        ExportLocation that = (ExportLocation) o;
        return sharedId == that.sharedId && admin == that.admin && preferred == that.preferred;
    }

    @Override
    public int hashCode() {
        return Objects.hash(sharedId, admin, preferred);
    }

    @Override
    public String toString() {
        return format("ExportLocation[sharedId=%d, admin=%s, preferred=%s]", sharedId, admin, preferred);
    }

    public boolean isAdmin() {
        return admin;
    }

    public boolean isPreferred() {
        return preferred;
    }
}
