import csv
import json

jewels_base = [
    ("Antidote Jewel 1", 727, 0),
	("Antipara Jewel 1", 728, 1),
	("Pep Jewel 1", 729, 2),
	("Steadfast Jewel 1", 730, 3),
	("Antiblast Jewel 1", 731, 4),
	("Suture Jewel 1", 732, 5),
	("Def Lock Jewel 1", 733, 6),
	("Earplug Jewel 3", 734, 7),
	("Wind Resist Jewel 2", 735, 8),
	("Footing Jewel 2", 736, 9),
	("Fertilizer Jewel 1", 737, 0xFFFFFFFF),
	("Heat Resist Jewel 2", 738, 0xFFFFFFFF),
	("Attack Jewel 1", 739, 10),
	("Defense Jewel 1", 740, 11),
	("Vitality Jewel 1", 741, 12),
	("Recovery Jewel 1", 742, 13),
	("Fire Res Jewel 1", 743, 14),
	("Water Res Jewel 1", 744, 15),
	("Ice Res Jewel 1", 745, 16),
	("Thunder Res Jewel 1", 746, 17),
	("Dragon Res Jewel 1", 747, 18),
	("Resistor Jewel 1", 748, 19),
	("Blaze Jewel 1", 749, 20),
	("Stream Jewel 1", 750, 21),
	("Frost Jewel 1", 751, 22),
	("Bolt Jewel 1", 752, 23),
	("Dragon Jewel 1", 753, 24),
	("Venom Jewel 1", 754, 25),
	("Paralyzer Jewel 1", 755, 26),
	("Sleep Jewel 1", 756, 27),
	("Blast Jewel 1", 757, 28),
	("Poisoncoat Jewel 3", 758, 29),
	("Paracoat Jewel 3", 759, 30),
	("Sleepcoat Jewel 3", 760, 31),
	("Blastcoat Jewel 3", 761, 32),
	("Powercoat Jewel 3", 762, 0xFFFFFFFF),
	("Release Jewel 3", 763, 33),
	("Expert Jewel 1", 764, 34),
	("Critical Jewel 2", 765, 35),
	("Tenderizer Jewel 2", 766, 36),
	("Charger Jewel 2", 767, 37),
	("Handicraft Jewel 3", 768, 38),
	("Draw Jewel 2", 769, 39),
	("Destroyer Jewel 2", 770, 40),
	("KO Jewel 2", 771, 41),
	("Drain Jewel 1", 772, 42),
	("Rodeo Jewel 2", 773, 0xFFFFFFFF),
	("Flight Jewel 2", 774, 43),
	("Throttle Jewel 2", 775, 44),
	("Challenger Jewel 2", 776, 45),
	("Flawless Jewel 2", 777, 46),
	("Potential Jewel 2", 778, 47),
	("Fortitude Jewel 1", 779, 48),
	("Furor Jewel 2", 780, 49),
	("Sonorous Jewel 1", 781, 50),
	("Magazine Jewel 2", 782, 51),
	("Trueshot Jewel 1", 783, 52),
	("Artillery Jewel 1", 784, 53),
	("Heavy Artillery Jewel 1", 785, 54),
	("Sprinter Jewel 2", 786, 55),
	("Physique Jewel 2", 787, 56),
	("Flying Leap Jewel 1", 788, 0xFFFFFFFF),
	("Refresh Jewel 2", 789, 57),
	("Hungerless Jewel 1", 790, 58),
	("Evasion Jewel 2", 791, 59),
	("Jumping Jewel 2", 792, 60),
	("Ironwall Jewel 1", 793, 61),
	("Sheath Jewel 1", 794, 62),
	("Friendship Jewel 1", 795, 63),
	("Enduring Jewel 1", 796, 64),
	("Satiated Jewel 1", 797, 65),
	("Gobbler Jewel 1", 798, 66),
	("Grinder Jewel 1", 799, 67),
	("Bomber Jewel 1", 800, 68),
	("Fungiform Jewel 1", 801, 69),
	("Angler Jewel 1", 802, 0xFFFFFFFF),
	("Chef Jewel 1", 803, 0xFFFFFFFF),
	("Transporter Jewel 1", 804, 0xFFFFFFFF),
	("Gathering Jewel 1", 805, 0xFFFFFFFF),
	("Honeybee Jewel 1", 806, 0xFFFFFFFF),
	("Carver Jewel 1", 807, 0xFFFFFFFF),
	("Protection Jewel 1", 808, 70),
	("Meowster Jewel 1", 809, 71),
	("Botany Jewel 1", 810, 72),
	("Geology Jewel 1", 811, 73),
	("Mighty Jewel 2", 812, 74),
	("Stonethrower Jewel 1", 813, 75),
	("Tip Toe Jewel 1", 814, 76),
	("Brace Jewel 3", 815, 77),
	("Scoutfly Jewel 1", 816, 0xFFFFFFFF),
	("Crouching Jewel 1", 817, 0xFFFFFFFF),
	("Longjump Jewel 1", 818, 0xFFFFFFFF),
	("Smoke Jewel 1", 819, 78),
	("Mirewalker Jewel 1", 820, 79),
	("Climber Jewel 1", 821, 0xFFFFFFFF),
	("Radiosity Jewel 1", 822, 0xFFFFFFFF),
	("Research Jewel 1", 823, 0xFFFFFFFF),
	("Specimen Jewel 1", 824, 80),
	("Miasma Jewel 1", 825, 97),
	("Scent Jewel 1", 826, 81),
	("Slider Jewel 1", 827, 0xFFFFFFFF),
	("Intimidator Jewel 1", 828, 82),
	("Hazmat Jewel 1", 829, 0xFFFFFFFF),
	("Mudshield Jewel 1", 830, 0xFFFFFFFF),
	("Element Resist Jewel 1", 831, 0xFFFFFFFF),
	("Slider Jewel 2", 832, 83),
	("Medicine Jewel 1", 833, 84),
	("Forceshot Jewel 3", 834, 85),
	("Pierce Jewel 3", 835, 86),
	("Spread Jewel 3", 836, 87),
	("Enhancer Jewel 2", 837, 88),
	("Crisis Jewel 1", 838, 89),
	("Dragonseal Jewel 3", 839, 90),
	("Discovery Jewel 2", 840, 0xFFFFFFFF),
	("Detector Jewel 1", 841, 0xFFFFFFFF),
	("Maintenance Jewel 1", 842, 91),
	("Mighty Bow Jewel 2", 874, 92),
	("Mind's Eye Jewel 2", 875, 93),
	("Shield Jewel 2", 876, 94),
	("Sharp Jewel 2", 877, 95),
	("Elementless Jewel 2", 878, 96),
]

if __name__ == '__main__':
    with open('decoration_base_translations.csv', 'r') as f:
        reader = csv.reader(f)
        translation = []
        for row in reader:
            translation.append({
                'name_en': row[0],
                'name_ja': row[1],
                'name_zh': row[10],
                'name_ko': row[9]
            })

    with open('decoration_base.csv', 'r') as f:
        reader = csv.reader(f)
        base = []
        for row in reader:
            base.append({
                'name_en': row[0],
                'skill_en': row[2]
            })

    with open('Monster Hunter_ World Jewels - Print.csv', 'r') as f:
        reader = csv.reader(f)
        max_list = []
        for row in reader:
            max_list.append({
                'skill_en': row[4],
                'max': row[3]
            })

    final_list = []

    for jb in jewels_base:
        jewel = {}
        jewel['name'] = jb[0]
        jewel['itemId'] = jb[1]
        jewel['equippedItemId'] = jb[2]
        jewel['locales'] = {}
        skill_en = ''
        for b in base:
            if jb[0] == b['name_en']:
                skill_en = b['skill_en']
        for m in max_list:
            if m['skill_en'] == skill_en:
                jewel['max'] = int(m['max'])
        for t in translation:
            if t['name_en'] == jewel['name']:
                jewel['locales']['en'] = t['name_en']
                jewel['locales']['ja'] = t['name_ja']
                jewel['locales']['zh'] = t['name_zh']
                jewel['locales']['ko'] = t['name_ko']
        final_list.append(jewel)

    with open('dist/decorations.json', 'w') as f:
        json.dump(final_list, f, ensure_ascii=False)
