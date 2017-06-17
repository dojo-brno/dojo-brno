package org.bbelovic.kata.exportlocation;

import org.junit.Assert;
import org.junit.Test;

import java.util.List;

import static java.util.Arrays.asList;
import static org.junit.Assert.assertEquals;

public class ExportLocationSelectorShould {

    private static final List<ExportLocation> EXPORT_LOCATIONS = createExportLocations();
    private final ExportLocationSelector selector = new ExportLocationSelector();

    @Test
    public void
    choose_only_export_locations_whose_shared_instance_id_matches_provided_share_id() {
        ExportLocation actual = selector.chooseExportLocation(EXPORT_LOCATIONS, 1L);
        ExportLocation expected = new ExportLocation(1L, false, true);
        assertEquals(expected, actual);
    }

    @Test(expected = RuntimeException.class)
    public void
    signal_error_when_no_export_location_matches_shared_id_criteria() {
        selector.chooseExportLocation(EXPORT_LOCATIONS, 10L);
    }

    @Test
    public void
    choose_only_export_location_whose_isAdmin_property_is_false() {
        ExportLocation actual = selector.chooseExportLocation(EXPORT_LOCATIONS, 3L);
        ExportLocation expected = new ExportLocation(3L, false, true);
        assertEquals(expected, actual);
    }

    @Test(expected = RuntimeException.class)
    public void
    signal_error_when_no_export_location_matches_isAdmin_criteria() {
        selector.chooseExportLocation(EXPORT_LOCATIONS, 2L);
    }

    @Test
    public void
    favour_export_location_whose_preferred_property_is_set_to_true() {
        ExportLocation actual = selector.chooseExportLocation(EXPORT_LOCATIONS, 1L);
        ExportLocation expected = new ExportLocation(1L, false, true);
        Assert.assertEquals(expected, actual);
    }



    private static List<ExportLocation> createExportLocations() {
        return asList(new ExportLocation(1L, false, false),
                new ExportLocation(1L, false, true),
                new ExportLocation(2L, true, true),
                new ExportLocation(3L, true, true),
                new ExportLocation(3L, false, true));
    }
}
